# Signature Service 

### Project Setup

- Go project containing the setup
- Basic API structure and functionality
- Encoding / decoding of different key types (only needed to serialize keys to a persistent storage)
- Key generation algorithms (ECC, RSA)
- Library to generate UUIDs, included in `go.mod`

### Prerequisites & Tooling

- Golang (v1.20+)

# APP

## RUN

    ```
    go run main.go
    ```

## TEST

    ```
    go test ./...
    ```