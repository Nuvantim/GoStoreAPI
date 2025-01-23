# GoStoreAPI

## Prerequisites
- Go (version 1.20+)
- MySQL / Postgres

## Installation :cd:

### 1. Clone Repository
```bash
git clone https://github.com/Kalveir/GoStoreAPI.git
cd GoStoreAPI
```

### 2. Configure Environment
```bash
# Copy environment template
cp .env.example .env

# Edit configuration
nano .env
```

### 3. Environment Variables
```
API_KEY=your_secure_api_key
REFRESH_KEY=your_secure_refresh_key

DB_DRIVER=mysql
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=
DB_NAME=api_store
DB_PORT=3306
PORT=8080
```

### 4. Install Dependencies
```bash
go mod tidy
```

## Running the Application ⚙️

### Development Mode
```bash
go run cli/main.go
```

## Deployment :rocket:
For the deployment process using [Docker](https://www.docker.com/), make sure docker is installed on your server
### 1. Environment Variables
The database configuration is adjusted in the file [docker-compose.yml](https://github.com/Kalveir/GoStoreAPI/blob/main/docker-compose.yml)
```
API_KEY=your_secure_api_key
REFRESH_KEY=your_secure_refresh_key

DB_DRIVER=pgsql
DB_HOST=golang-db
DB_USER=postgres
DB_PASSWORD=your_database_password
DB_NAME=api_store
DB_PORT=5432
PORT=7373
```
## 2. Run Docker Compose
```bash
docker compose up -d --build
```
## 3. Check Logs App
```bash
docker logs golang-app
```
