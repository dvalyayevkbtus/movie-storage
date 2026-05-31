package main

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type MoviesHttp struct {
	db *MovieDb
}

func CreateMovieHttp(db *MovieDb) *MoviesHttp {
	return &MoviesHttp{db}
}

func (m *MoviesHttp) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		m.getAll(rw, req)
	} else if req.Method == http.MethodPost {
		m.create(rw, req)
	} else {
		MethodNotAllowed(rw)
	}
}

func (m *MoviesHttp) getAll(rw http.ResponseWriter, req *http.Request) {
	movies, err := m.db.SelectMovies()
	if err != nil {
		log.Errorf("Cannot select movies: %v", err)
		InternalServerError(rw)
		return
	}

	body, mErr := json.Marshal(movies)
	if mErr != nil {
		log.Errorf("Cannot select movies: %v", mErr)
		InternalServerError(rw)
		return
	}

	SuccessString(rw, string(body))
}

func (m *MoviesHttp) create(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Errorf("Cannot read body! %v", err)
		InternalServerError(rw)
		return
	}

	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		log.Errorf("Cannot read body! %v", err)
		InternalServerError(rw)
		return
	}

	err = m.db.CreateMovie(movie)
	if err != nil {
		log.Errorf("Cannot create new movie in database! %v", err)
	}

	SuccessString(rw, "Success!")
}


TASK 1

sudo groupadd sysadm
sudo groupadd operations

sudo vi /etc/sudoers.d/sysadm

%sysadm ALL=(ALL) NOPASSWD: /bin/systemctl

sudo chmod 440 /etc/sudoers.d/sysadm
sudo cp /etc/sudoers.d/sysadm /home/kbtu/sysadm.bak

sudo vi /etc/sudoers.d/operations

%operations ALL=(ALL) NOPASSWD: /usr/local/bin/build-pipeline.sh

sudo chmod 440 /etc/sudoers.d/operations
sudo cp /etc/sudoers.d/operations /home/kbtu/operations.bak
TASK 2 & 3

cd /home/kbtu/
git clone https://github.com/dvalyayevkbtu/movie-storage.git
cd movie-storage
sudo podman build -t movie-storage:latest .
sudo podman run -d 
--name movie-app 
--memory 256m 
--memory-swap 256m 
--cpus 0.5 
-p 8080:8080 movie-storage:latest curl http://localhost:8080
sudo podman push movie-storage:latest docker://docker.io/yourusername/movie-storage:latest
TASK 4 
Bash
sudo vi /usr/local/bin/build-pipeline.sh

#!/bin/bash
cd /tmp
rm -rf /tmp/movie-storage
git clone https://github.com/dvalyayevkbtu/movie-storage.git
cd movie-storage
podman build -t movie-storage:latest .
logger -t movie-pipeline "Audit: Container image built successfully"
podman run -d --name movie-test --memory 256m --memory-swap 256m --cpus 0.5 -p 8080:8080 movie-storage:latest
logger -t movie-pipeline "Audit: Container started"
sleep 3
if curl -s http://localhost:8080 > /dev/null; then
    logger -t movie-pipeline "Audit: Service check PASSED"
else
    logger -t movie-pipeline "Audit: Service check FAILED"
fi
podman stop movie-test
podman rm movie-test
logger -t movie-pipeline "Audit: Container cleaned up and deleted"
cd /tmp
rm -rf /tmp/movie-storage
logger -t movie-pipeline "Audit: Local Git repository purged"

sudo chmod 750 /usr/local/bin/build-pipeline.sh
sudo chown root:operations /usr/local/bin/build-pipeline.sh
sudo cp /usr/local/bin/build-pipeline.sh /home/kbtu/build-pipeline.sh

sudo vi /etc/audit/rules.d/movie-audit.rules

-w /usr/local/bin/build-pipeline.sh -p wa -k movie-pipeline-tamper
Bash
sudo cp /etc/audit/rules.d/movie-audit.rules /home/kbtu/movie-audit.rules
sudo augenrules --load
sudo systemctl enable --now auditd

TASK 5 

sudo vi /etc/nftables.conf

Plaintext
#!/usr/sbin/nft -f
flush ruleset
table inet filter {
  chain input {
    type filter hook input priority 0;
    ct state established,related accept
    iif "lo" accept
    tcp dport { 22, 4200 } accept
    drop
  }
  chain forward { type filter hook forward priority 0; }
  chain output { type filter hook output priority 0; }
}
Bash
sudo cp /etc/nftables.conf /home/kbtu/nftables.conf
sudo nft -f /etc/nftables.conf
sudo systemctl enable --now nftables
