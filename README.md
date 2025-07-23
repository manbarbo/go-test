# Stock Recommendation Dashboard

Application that displays the best stock recommendations for investing, based on broker analysis and score calculations. 
The frontend is a SPA built with Vue 3 + Tailwind CSS, and the backend is developed in Go with a CockroachDB database.

## Architecture
backedng (Go)
- api: Routes and HTTP controlers
- db: Conection and DB queries
- load: App to load Stock Information into DB
- middlewares: Midlewares as cors
- models: Data structs
- migrations: App to run the DB migrations
- services: Business logic
- utils: Tools

dashboard (Vue + Tailwind)

- src
- - assets: Statics files 
- - componets: Vue omponents
- - models: Data structs
- - pages: Main views
- - router: App routing

## Requirements
- Go 1.24+
- Node 22+ 
- npm
- CockroachDB (Postgresql)
- [Vite](https://vitejs.dev/)
- `.env` configuration

## Backend - Go

### Environment variables (`.env`)

```env
API_TOKEN=
API_URL=

DB_USER=
BD_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_OPTIONS=
BD_MIGRATION_CONNECTOR=

CORS_ORIGINS=
```

### Run
Migrations
```bash
go run migrations/main.go
```

Load Data
```bash
go run load/main.go
```

Backend
```bash
go run main.go
```

### Main Endpoints

GET /stocks: List all the stock information

GET /recommendation?top=5: Top 5 recomatation by score

GET /health: heat test

## Frontend - Vue + Tailwind

### Environment variables (`.env`)
```env
VITE_API_URL=
```

### Install and run
```bash
cd dashboard
npm install
npm run dev
```


## Author
Manuel Barona
Fullstack developer

## Licence
MIT