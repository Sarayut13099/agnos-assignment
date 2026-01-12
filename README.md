# Agnos Assignment


## 🧱 Tech Stack

-   Golang (Gin)
-   GORM
-   PostgreSQL 18
-   Docker & Docker Compose
-   Nginx (Reverse Proxy)

------------------------------------------------------------------------

## 📁 Project Structure

    .
    ├── docker-compose.yml
    ├── nginx/
    │   └── nginx.conf
    ├── postgres/
    │   ├── health-app/
    │   │   └── init.sql
    │   └── his/
    │       └── init.sql
    └── services/
        ├── health-app/
        │   └── cmd/api/main.go
        └── his/
            └── cmd/api/main.go

------------------------------------------------------------------------

## 🚀 Run the System

``` bash
docker compose down -v
docker compose up --build

OR
docker compose up -d
```

Access via: http://localhost

------------------------------------------------------------------------

## 🌐 API Endpoints

### Health App

Base URL: `http://localhost/health-app`

------------------------------------------------------------------------

### HIS

Base URL: `http://localhost/his`


------------------------------------------------------------------------

## 🛢 Database Access

### Health App DB

-   Host: localhost
-   Port: 5432
-   DB: health_app
-   User: admin
-   Pass: adminpass

### HIS DB

-   Host: localhost
-   Port: 5433
-   DB: his_db
-   User: admin
-   Pass: adminpass

------------------------------------------------------------------------
