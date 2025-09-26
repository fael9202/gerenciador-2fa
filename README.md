# Gerenciador 2FA

Um gerenciador de contas TOTP (Time-based One-Time Password) desenvolvido em Go com API REST para armazenar e gerenciar códigos de autenticação de dois fatores.

## 🚀 Funcionalidades

- **Autenticação de Usuários**: Sistema de registro e login com JWT
- **Gerenciamento de Contas TOTP**: Adicionar, listar e remover contas 2FA
- **API REST**: Endpoints bem estruturados com validação
- **Banco de Dados**: MongoDB para persistência de dados
- **Docker**: Containerização completa da aplicação
- **Segurança**: Senhas criptografadas com bcrypt e tokens JWT

## 🛠️ Tecnologias

- **Backend**: Go 1.23.6
- **Framework Web**: Gin
- **Banco de Dados**: MongoDB
- **Autenticação**: JWT (JSON Web Tokens)
- **Criptografia**: bcrypt para senhas
- **Containerização**: Docker & Docker Compose

## 📋 Pré-requisitos

- Go 1.23.6 ou superior
- MongoDB
- Docker (opcional)

## 🚀 Instalação e Execução

### Opção 1: Execução Local

1. **Clone o repositório**
```bash
git clone https://github.com/seu-usuario/gerenciador-2fa.git
cd gerenciador-2fa
```

2. **Instale as dependências**
```bash
go mod download
```

3. **Configure as variáveis de ambiente**
Crie um arquivo `.env` na raiz do projeto:
```env
MONGODB_URI=mongodb://localhost:27017
JWT_SECRET=seu_jwt_secret_aqui
```

4. **Inicie o MongoDB**
```bash
# Com Docker
docker run -d -p 27017:27017 --name mongodb mongo:latest

# Ou use o docker-compose
docker-compose up -d mongodb
```

5. **Execute a aplicação**
```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

### Opção 2: Docker Compose

1. **Clone e configure**
```bash
git clone https://github.com/seu-usuario/gerenciador-2fa.git
cd gerenciador-2fa
```

2. **Configure o .env**
```env
MONGODB_URI=mongodb://mongodb:27017
JWT_SECRET=seu_jwt_secret_aqui
```

3. **Execute com Docker Compose**
```bash
docker-compose up --build
```

## 📚 API Endpoints

### Rotas Públicas

#### Health Check
```http
GET /api/v1/health
```
Verifica o status da aplicação e conexão com o banco.

#### Registro de Usuário
```http
POST /api/v1/register
Content-Type: application/json

{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

#### Login
```http
POST /api/v1/login
Content-Type: application/json

{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

**Resposta:**
```json
{
  "status": "success",
  "message": "Login realizado com sucesso",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Rotas Protegidas (Requer JWT)

Todas as rotas abaixo requerem o header `Authorization: Bearer <token>`

#### Perfil do Usuário
```http
GET /api/v1/profile
Authorization: Bearer <token>
```

#### Adicionar Conta TOTP
```http
POST /api/v1/totp
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Google",
  "issuer": "Google LLC",
  "secret": "JBSWY3DPEHPK3PXP"
}
```

#### Listar Contas TOTP
```http
GET /api/v1/totp
Authorization: Bearer <token>
```

#### Remover Conta TOTP
```http
DELETE /api/v1/totp/{id}
Authorization: Bearer <token>
```

## 🏗️ Estrutura do Projeto

```
gerenciador-2fa/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   └── config.go           # Configurações da aplicação
│   ├── database/
│   │   ├── mongodb.go          # Conexão com MongoDB
│   │   ├── user.go             # Operações de usuário
│   │   └── totp.go             # Operações TOTP
│   ├── handlers/
│   │   ├── auth.go             # Handlers de autenticação
│   │   ├── user.go             # Handlers de usuário
│   │   ├── totp.go             # Handlers TOTP
│   │   └── health.go           # Health check
│   ├── middleware/
│   │   └── auth.go             # Middleware de autenticação
│   ├── models/
│   │   ├── user.go             # Modelo de usuário
│   │   └── totp.go             # Modelo TOTP
│   ├── response/
│   │   └── response.go         # Utilitários de resposta
│   └── errors/
│       └── errors.go           # Tratamento de erros
├── docker-compose.yml          # Configuração Docker Compose
├── Dockerfile                  # Imagem Docker
├── go.mod                      # Dependências Go
└── README.md                   # Este arquivo
```

## 🔧 Configuração

### Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|---------|
| `MONGODB_URI` | URI de conexão com MongoDB | `mongodb://localhost:27017` |
| `JWT_SECRET` | Chave secreta para JWT | **Obrigatório** |

### Configurações da Aplicação

- **Porta**: 8080
- **Expiração do Token**: 24 horas
- **Timeout MongoDB**: 10 segundos

## 🧪 Testando a API

### Exemplo com cURL

1. **Registrar usuário**
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"email":"teste@exemplo.com","password":"senha123"}'
```

2. **Fazer login**
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email":"teste@exemplo.com","password":"senha123"}'
```

3. **Adicionar conta TOTP**
```bash
curl -X POST http://localhost:8080/api/v1/totp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  -d '{"name":"Google","issuer":"Google LLC","secret":"JBSWY3DPEHPK3PXP"}'
```

## 🔒 Segurança

- Senhas são criptografadas usando bcrypt
- Tokens JWT com expiração de 24 horas
- Middleware de autenticação em rotas protegidas
- Validação de entrada em todos os endpoints
- CORS configurado para desenvolvimento

## 🐳 Docker

### Build da Imagem
```bash
docker build -t gerenciador-2fa .
```

### Executar Container
```bash
docker run -p 8080:8080 \
  -e MONGODB_URI=mongodb://host.docker.internal:27017 \
  -e JWT_SECRET=seu_jwt_secret \
  gerenciador-2fa
```

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 📞 Suporte

Se você encontrar algum problema ou tiver dúvidas, por favor abra uma issue no repositório.

---

**Desenvolvido com ❤️ em Go**