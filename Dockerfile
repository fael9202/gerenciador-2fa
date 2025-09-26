FROM golang:1.21-alpine

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN go build -o main ./cmd/api

# Expor a porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"] 