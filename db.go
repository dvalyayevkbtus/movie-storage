package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type MovieDb struct {
	db *sql.DB
}

type Movie struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func GetDatabase(conf Config) (*MovieDb, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &MovieDb{db}, nil
}

func (m MovieDb) MigrateDatabase() error {
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS movie (
			id bigserial not null,
			name character varying(256) not null,
			category character varying(256) not null,

			constraint movie_pk primary key (id)
		)
	`)
	if err != nil {
		return err
	}
	log.Info("Database migrated successfully!")
	return nil
}

func (m MovieDb) SelectMovies() ([]Movie, error) {
	rows, err := m.db.Query("SELECT id, name, category FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Movie{}
	for rows.Next() {
		m := Movie{}
		rErr := rows.Scan(&m.Id, &m.Name, &m.Category)
		if rErr != nil {
			return nil, rErr
		}
		result = append(result, m)
	}

	return result, nil
}

func (m MovieDb) CreateMovie(movie Movie) error {
	_, err := m.db.Exec("INSERT INTO movie (name, category) VALUES ($1, $2)", movie.Name, movie.Category)
	return err
}

func (m MovieDb) CloseDatabase() error {
	return m.db.Close()
}
