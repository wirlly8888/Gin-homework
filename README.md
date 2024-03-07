# Gin-homework
This a simple CRUD homework implemented by Gin.

## Before Start
You should install [Docker](https://www.docker.com/) before you run this service.
If You want to develop, please download Go 1.20+.

### Build docker image
If you have installed [Docker](https://www.docker.com/), you can execute this command to build Docker image
```bash
make build-image
```

### Install Go modules (develop)
```bash
go mod download
go mod tidy
```

## Run Server & stop

### Use docker container
- Run server
```bash
make run-server 
```

- Stop server
```bash
make stop-server
```

### Use loacl environment (develop)
- Run server
```bash
go run main.go
```

- Stop server: Just control + C

### Run test (develop)
```
go test ./...
```

### API document
If you want to learn more detail about API in this service, you can read [api.yaml](/APIs/api.yaml)