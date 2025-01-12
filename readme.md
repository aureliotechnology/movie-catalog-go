# Movie Catalog Go

**Movie Catalog Go** é um projeto de API desenvolvido em Go utilizando arquitetura hexagonal (Ports and Adapters). O objetivo principal é fornecer uma base sólida e escalável para gerenciamento de catálogos de filmes, seguindo boas práticas de desenvolvimento de software.

---

## **Estrutura do Projeto**

```plaintext
movie-catalog-go/
├── api/
│   ├── http/
│   │   ├── health_handler.go      # Manipuladores de rotas relacionadas à saúde
│   │   ├── routes.go              # Configuração centralizada de rotas
│   │   ├── health_handler_test.go # Testes do manipulador de saúde
│   │   └── routes_test.go         # Testes para configuração de rotas
├── cmd/
│   └── main.go                    # Ponto de entrada da aplicação
├── internal/
│   ├── adapter/
│   │   └── repository/            # Implementação de repositórios e conexão com banco de dados
│   ├── app/                       # Serviços e casos de uso
│   │   └── health_service.go      # Serviço responsável pela lógica de saúde
│   ├── core/
│   │   └── port/
│   │       └── in/                # Interfaces de entrada (ex.: HealthPort)
├── docker-compose.yml             # Configuração para rodar com Docker Compose
├── Dockerfile                     # Dockerfile para construção da imagem da aplicação
├── go.mod                         # Gerenciamento de dependências Go
├── go.sum                         # Checksums das dependências Go
└── README.md                      # Documentação do projeto
```

---

## **Funcionalidades**

### 1. **Rota `/health`**
- Verifica se a aplicação está rodando corretamente.
- **Método:** `GET`
- **Resposta:** `{ "message": "ok" }`

### 2. **Rota `/health-db`**
- Verifica a conexão com o banco de dados PostgreSQL.
- **Método:** `GET`
- **Resposta:**
  - Sucesso: `{ "message": "ok" }`
  - Falha: `{ "message": "database connection error" }`

---

## **Como Rodar o Projeto**

### **Pré-requisitos**
- Go 1.21 ou superior
- Docker e Docker Compose

### **Rodando Localmente**

1. Clone o repositório:
   ```bash
   git clone <repository-url>
   cd movie-catalog-go
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute o projeto:
   ```bash
   go run cmd/main.go
   ```

4. Acesse os endpoints:
   - [http://localhost:8080/health](http://localhost:8080/health)
   - [http://localhost:8080/health-db](http://localhost:8080/health-db)

### **Rodando com Docker Compose**

1. Suba o ambiente:
   ```bash
   docker-compose up --build
   ```

2. Acesse os endpoints:
   - [http://localhost:8080/health](http://localhost:8080/health)
   - [http://localhost:8080/health-db](http://localhost:8080/health-db)

---

## **Testes**

### **Rodar Todos os Testes**

Execute o comando abaixo para rodar os testes unitários:
```bash
go test ./... -v
```

### **Cobertura de Código**

Gere um relatório de cobertura de código:
```bash
go test ./... -coverprofile=coverage.out
```
Abra o relatório:
```bash
go tool cover -html=coverage.out
```

---

## **Tecnologias Utilizadas**

- **Linguagem:** Go (Golang)
- **Banco de Dados:** PostgreSQL
- **Arquitetura:** Hexagonal (Ports and Adapters)
- **Containerização:** Docker, Docker Compose

---

## **Contribuindo**

Contribuições são bem-vindas! Siga os passos abaixo para contribuir:

1. Faça um fork do repositório.
2. Crie uma branch para suas alterações:
   ```bash
   git checkout -b minha-feature
   ```
3. Envie um Pull Request com suas mudanças.

---

## **Licença**

Este projeto está licenciado sob a Licença MIT.
