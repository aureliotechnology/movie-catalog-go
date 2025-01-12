# Etapa 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

# Copiar arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do projeto
COPY . .

# Compilar o binário de forma estática
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o movie-catalog-go -ldflags '-extldflags "-static"' ./cmd/main.go

# Etapa 2: Runtime
FROM scratch

WORKDIR /root/

# Copiar o binário gerado para a imagem final
COPY --from=builder /app/movie-catalog-go .

# Expor a porta usada pela aplicação
EXPOSE 8080

# Comando para rodar o binário
CMD ["./movie-catalog-go"]
