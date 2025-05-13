# PulseLedger

A Go Fiber-based microservice to manage accounts and transactions, using SQLite and Viper for configuration.

---

## Build & Run

> This project uses **Go 1.24** and builds using a multi-stage **Alpine Linux-based Docker image** (`golang:1.24.2-alpine`).  
> The runtime container is based on `alpine:latest`.

### Run Locally (Without Docker)

```bash
go mod tidy
go run main.go
```

### Change Port Locally

Update `config.yaml`:

```yaml
server:
  port: "8080"

database:
  path: "./pulseledger.db"
  reset_on_start: true
```

---

### Run with Docker

```bash
./run
```

> Make sure `.env` is present with:

```env
SERVER_PORT=8080
DATABASE_PATH=./pulseledger.db
DATABASE_RESET_ON_START=true
```

---

## Run Tests

```bash
go test ./...
```

---

## Eample cURL Commands

### Create Account

```bash
curl -X POST http://localhost:8080/api/v1/accounts \
  -H "Content-Type: application/json" \
  -d '{"document_number": 1234567890}'
```

Invalid example:

```bash
curl -X POST http://localhost:8080/api/v1/accounts \
  -H "Content-Type: application/json" \
  -d '{}'
```

### Create Transaction

```bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "account_id": 1,
    "operation_type_id": 1,
    "amount": 100.50
  }'
```

### Get Account by ID

```bash
curl http://localhost:8080/api/v1/accounts/1
```

> Replace `1` with the actual account ID you want to retrieve.

---

## Folder Structure

```
pulseledger/
├── config/         # Configuration loader (Viper)
├── db/             # DB init and migrations
├── handlers/       # HTTP handlers
├── repositories/   # DB layer
├── services/       # Business logic
├── dto/            # Request/response types
├── main.go         # Entry point
```
