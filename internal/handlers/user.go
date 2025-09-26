package handlers

import (
    "gerenciador-2fa/internal/database"
    "gerenciador-2fa/internal/response"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProfile retorna todas as contas TOTP do usuário
func GetProfile(c *gin.Context) {
    userID, _ := c.Get("user_id")
    userObjID, err := primitive.ObjectIDFromHex(userID.(string))
    if err != nil {
        response.BadRequest(c, "ID de usuário inválido")
        return
    }

    accounts, err := database.GetUserTOTPAccounts(userObjID)
    if err != nil {
        response.InternalError(c, "Erro ao buscar contas TOTP")
        return
    }

    response.OK(c, "Contas TOTP recuperadas com sucesso", accounts)
}