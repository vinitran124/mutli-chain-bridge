## Requirements

- go version: 1.22
- docker compose version: 2.23

## Deployment

Initialize postgres and redis database:

```
docker compose up redis
docker compose up postgres
```

Run the api
```
cd backend
go run cmd/*.go api           //--cfg flag for path of configuration file
```