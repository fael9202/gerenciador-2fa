package handlers

import (
    "gerenciador-2fa/internal/database"
    "gerenciador-2fa/internal/models"
    "gerenciador-2fa/internal/response"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "os"
    "time"
)

// Estruturas de requisição
type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// Register registra um novo usuário
func Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "Dados inválidos")
        return
    }

    exists, err := database.UserExists(req.Email)
    if err != nil {
        response.InternalError(c, "Erro ao verificar usuário")
        return
    }
    if exists {
        response.Error(c, http.StatusConflict, "Email já está em uso")
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        response.InternalError(c, "Erro ao processar senha")
        return
    }

    user := models.User{
        Email:     req.Email,
        Password:  string(hashedPassword),
        CreatedAt: time.Now(),
    }

    if err := database.CreateUser(&user); err != nil {
        response.InternalError(c, "Erro ao criar usuário")
        return
    }

    response.Created(c, "Usuário criado com sucesso", nil)
}

// Login realiza o login do usuário
func Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "Dados inválidos")
        return
    }

    user, err := database.GetUserByEmail(req.Email)
    if err != nil {
        response.Error(c, http.StatusUnauthorized, "Credenciais inválidas")
        return
    }

    // Verificar senha
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        response.Error(c, http.StatusUnauthorized, "Credenciais inválidas")
        return
    }

    // Gerar JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID.Hex(),
        "email":   user.Email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        response.InternalError(c, "Erro ao gerar token")
        return
    }

    response.OK(c, "Login realizado com sucesso", gin.H{"token": tokenString})
} 