package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	conf, err := GetConfig()
	if err != nil {
		log.Fatal("Cannot upload config of the app!")
		panic(err)
	}

	db, dbErr := GetDatabase(conf)
	if dbErr != nil {
		log.Fatal("Cannot connect to database!")
		panic(dbErr)
	}
	defer db.CloseDatabase()

	migErr := db.MigrateDatabase()
	if migErr != nil {
		log.Fatal("Cannot migrate database!")
		panic(migErr)
	}

	movies := CreateMovieHttp(db)

	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/movie", movies.Handler)

	log.Info("Http has started.")
	httpErr := http.ListenAndServe(":8080", nil)
	if httpErr != nil {
		log.Fatal("Error on http!")
		panic(httpErr)
	}

	log.Info("App is shutted down.")
}
