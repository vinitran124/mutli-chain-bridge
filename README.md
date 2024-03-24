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

Run the migration
```
cd backend
go run cmd/*.go migration --action up   // action up to update database
go run cmd/*.go migration --action down // action down to revert database
```

## Token Information

| Symbol | ChainId     | Address                       |
| :-------- | :------- | :-------------------------------- |
| VINI      | 11155111 | 0x15f8253779428d9ea5b054deef3e454d539ddf7e |
| VINI      | 97 | 0x6b08b796b4b43d565c34cf4b57d8c871db410ebe |