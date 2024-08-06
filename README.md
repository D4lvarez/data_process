# Requirements
1. Install Go.
2. Install MySQL or MariaDB.
3. Execute:
```bash
go mod tidy
```


# Initialize Database
1. Rename file *.env.example* to *.env* and set values for variables.

2. Execute on terminal this command:
```bash
go run cmd/database/main.go
```