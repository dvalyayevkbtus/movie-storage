# movie-storage

This is a sample Golang HTTP service.

## Endpoints

**App is running on port 8080**

There are 3 endpoints:

- GET `/health` - To check liveness (`curl http://<host>:<port>/health`)
- GET `/movie` - To get all movies (`curl http://<host>:<port>/movie`)
- POST `/movie` - To create movies: (`curl -X POST -d '{"name":"my-movie","catergory":"horror"} http://<host>:<port>/movie'`)
    body example:
    ```json
        {"name": "my-movie", "category": "horror"}
    ```

## How to build

### Host

1. Check if you have go `go version` or install it
2. Run `go build .`. The file will be located in this dir `movie-storage`

### Dockerfile

1. Use `golang:1.24.2` as a base image for build
2. Copy `go.mod` and `go.sum` and run `go mod download`
3. Copy all files inside and run `go build .`
4. Prepare a running image. You can use `FROM alpine`.
5. Copy binary from build to run and configure entrypoint to just run your binary

## How to run

### Prepare a config

This app needs a PostgreSQL running.
Create a file `movie-conf.json`:
    
```json
{
    "db_host": "<host>",
    "db_port": "<port>",
    "db_username": "<user>",
    "db_password": "<password>",
    "db_name": "<database name>"
}
```

### Configure environment

You need to setup `MOVIE_CONF_PATH` environment variable with a full path to your config file.

### Running

You need just run
