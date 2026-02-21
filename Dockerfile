# ESTÁGIO 1: Compilação (O Construtor)
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Cache de dependências
COPY go.mod ./
RUN go mod download

# Copia o código e compila
COPY . .
RUN go build -o main .

# ESTÁGIO 2: Execução (O que vai para o ar)
FROM alpine:latest
WORKDIR /root/

# Copiamos APENAS o binário do estágio anterior
COPY --from=builder /app/main .

# Expõe a porta da sua API (ex: 8080)
EXPOSE 8080

# Roda o binário direto (Rápido e leve)
CMD ["./main"]
