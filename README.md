# Hackathon

Abaixo esta a estrutura do projeto realizado para desafiar nossas habilidades com ferramentas ainda não utilizadas na minha carreira

---

## Estrutura da API

API CRUD feita em GO



Banco de dados: Postgres

## Estrutura do projeto

O repositório está organizado da seguinte forma:

- `cmd/` : contém o executável principal da aplicação (`main.go`) que inicializa o servidor HTTP (Gin), middleware e insere dependências (repository → usecase → controller).
- `internal/` : pacotes internos da aplicação. Ex.: `internal/cert` contém o gerador de certificados TLS usado em desenvolvimento e `internal/middleware` contém middlewares HTTP.
- `repository/` : implementação da camada de acesso a dados (SQL). Aqui ficam os repositórios que conversam com o Postgres.
- `usecase/` : camada de regras de negócio que orquestra chamadas aos repositórios.
- `controler/` : handlers/controllers HTTP (rotas) que expõem os endpoints da API.
- `db/` : código de conexão e configuração do banco de dados (ex.: `ConnectDB`).
- `model/` : definições das structs usadas na aplicação (ex.: `Product`, `Response`).
- `Dockerfile` : Docker multi-stage que compila a aplicação e gera a imagem final mínima.
- `docker-compose.yml` : orquestra a aplicação e um container Postgres para desenvolvimento.
- `generate_cert.go` : utilitário para gerar certificados (mantido para referência; o servidor também gera certs em `internal/cert`).

## Executando localmente (resumo)

1. Inicie o Docker Desktop e rode:

```powershell
cd C:\Users\higor\Desktop\Hackathon
docker compose up -d --build
```

2. A API ficará disponível em `https://localhost:8080` (usa um certificado autoassinado gerado em tempo de execução).

3. Endpoints principais:

- `GET /ping` — healthcheck
- `GET /products` — lista produtos
- `POST /product` — cria produto
- `GET /product/:productId` — obtém produto por id
- `PUT /product/:productId` — atualiza produto
- `DELETE /product/:productId` — remove produto

## Observações

- Em desenvolvimento o servidor gera `cert.pem` e `key.pem` automaticamente e inicia em HTTPS. Para produção, use certificados válidos e gerenciados externamente.
- Nomes exportados em `main.go` podem ser renomeados para seguir convenções idiomáticas do Go (ex.: `productRepo`).


