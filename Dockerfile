# Etapa de build
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# Dependências importantes
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila com flags de segurança (build the `cmd` package to produce the web server)
RUN go build -ldflags="-s -w" -o server ./cmd

# -------------------------------------------------------------------

# Etapa final (run)
FROM alpine:3.19

# Cria usuário não-root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Diretório da aplicação
WORKDIR /app

# Copia binário da etapa de build
COPY --from=builder /app/server /app/server

# Permissões
RUN chown -R appuser:appgroup /app
RUN chmod +x /app/server

# Troca para o usuário não-root
USER appuser

# EntryPoint
ENTRYPOINT ["/app/server"]
