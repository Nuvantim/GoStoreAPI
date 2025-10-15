# GoStoreAPI
## Descriptions
GoStoreAPI is a RESTful API designed to power e-commerce platforms or store management systems. It provides structured endpoints to handle core operations such as product management, categories, transactions, user authentication, and other essential features for online stores or inventory-based applications.
## Prerequisites
- Go (version 1.23+)
- MySQL / Postgres
- Docker

## Installation :cd:

### 1. Clone Repository
```bash
git clone https://github.com/Nuvantim/GoStoreAPI.git
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
APP_NAME=STORE_API
URL=example.com

DB_DRIVER=mysql
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=
DB_NAME=api_store
DB_PORT=3306

MAIL_MAILER=smtp.xxx.com
MAIL_PORT=25
MAIL_USERNAME=user@mail.com
MAIL_PASSWORD=password123
MAIL_FROM_ADDRESS="person@mail.com"

PORT=8080
```

### 4. Install Dependencies
```bash
go mod tidy

```
### 5. Create an RSA key
```bash
# Generate a 4096-bit RSA private key
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:4096

# Generate a public key from the private key
openssl rsa -in private.pem -pubout -out public.pem

```

## Running the Application ⚙️

### Development Mode
```bash
go run cmd/main.go
```

## Deployment :rocket:
For the deployment process using [Docker](https://www.docker.com/), make sure docker is installed on your server
### 1. Environment Variables
The database configuration is adjusted in the file [docker-compose.yml](https://github.com/Kalveir/GoStoreAPI/blob/main/docker-compose.yml)
```
APP_NAME=STORE_API
URL=example.com

DB_DRIVER=pgsql
DB_HOST=mydb
DB_USER=postgres
DB_PASSWORD=your_database_password
DB_NAME=api_store
DB_PORT=5432

MAIL_MAILER=smtp.xxx.com
MAIL_PORT=25
MAIL_USERNAME=user@mail.com
MAIL_PASSWORD=password123
MAIL_FROM_ADDRESS="person@mail.com"

PORT=7373
```
## 2. Compile project
```bash
make build
```
## 3. Run Docker Compose
```bash
docker compose up -d
```
## 4. Check Logs App
```bash
docker logs myapp
```
