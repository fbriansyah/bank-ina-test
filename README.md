# Tasks Project
Application for creating task.

## Project Structure
1. `/cmd` main package
2. `/db` migration and query files for SQLC
3. `/internal` the meat of application
    - `/internal/adapter` all adapter i use in this project, like chi and postgres
    - `/internal/application` this folder contain domain logic in this project
    - `/internal/application/port` abstraction for adapter
4. `/util` utility function and config

## How to Run This Project

### Manual With Docker
```bash
make postgres # create postgres db instance
make createdb # create database for application
make re-db # re-structure database
make run # run application
```

### Using Docker Compose
```bash
make compose-up
```

## How To Access
We can import postman collection from file `Task.postman_collection.json`