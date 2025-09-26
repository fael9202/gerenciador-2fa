package handlers

import (
    "gerenciador-2fa/internal/database"
    "gerenciador-2fa/internal/models"
    "gerenciador-2fa/internal/response"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type AddTOTPRequest struct {
    Name      string `json:"name" binding:"required"`
    Issuer    string `json:"issuer" binding:"required"`
    Secret    string `json:"secret" binding:"required"`
}

// AddTOTPAccount adiciona uma nova conta TOTP
func AddTOTPAccount(c *gin.Context) {
    var req AddTOTPRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "Dados inválidos")
        return
    }

    userID, _ := c.Get("user_id")
    userObjID, _ := primitive.ObjectIDFromHex(userID.(string))

    account := models.TOTPAccount{
        UserID:    userObjID,
        Name:      req.Name,
        Issuer:    req.Issuer,
        Secret:    req.Secret,
        CreatedAt: time.Now(),
    }

    if err := database.CreateTOTPAccount(&account); err != nil {
        response.InternalError(c, "Erro ao salvar conta TOTP")
        return
    }

    response.Created(c, "Conta TOTP adicionada com sucesso", account)
}

// ListTOTPAccounts lista todas as contas TOTP do usuário
func ListTOTPAccounts(c *gin.Context) {
    userID, _ := c.Get("user_id")
    userObjID, _ := primitive.ObjectIDFromHex(userID.(string))

    accounts, err := database.GetUserTOTPAccounts(userObjID)
    if err != nil {
        response.InternalError(c, "Erro ao buscar contas TOTP")
        return
    }

    response.OK(c, "Contas TOTP recuperadas com sucesso", accounts)
}

// DeleteTOTPAccount remove uma conta TOTP
func DeleteTOTPAccount(c *gin.Context) {
    accountID, err := primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
        response.BadRequest(c, "ID inválido")
        return
    }

    userID, _ := c.Get("user_id")
    userObjID, _ := primitive.ObjectIDFromHex(userID.(string))

    if err := database.DeleteTOTPAccount(userObjID, accountID); err != nil {
        response.InternalError(c, "Erro ao deletar conta TOTP")
        return
    }

    response.OK(c, "Conta TOTP removida com sucesso", nil)
} 