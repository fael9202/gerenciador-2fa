package database

import (
	"context"
	"fmt"
	"log"
	"gerenciador-2fa/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *models.User) error {
	collection := GetDatabase().Collection("users")
	
	// Verificar se já existe um usuário com este email
	existingUser, err := GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("usuário com este email já existe")
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Erro ao criar usuário no MongoDB: %v", err)
		return err
	}
	
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	collection := GetDatabase().Collection("users")
	var user models.User
	
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Erro ao buscar usuário por email: %v", err)
		return nil, err
	}
	
	return &user, nil
}

func UpdateTOTPStatus(userID primitive.ObjectID, enabled bool) error {
	collection := GetDatabase().Collection("users")
	
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"totp_enabled": enabled}}
	
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// UserExists verifica se um usuário com o email fornecido já existe
func UserExists(email string) (bool, error) {
	collection := GetDatabase().Collection("users")
	
	count, err := collection.CountDocuments(context.Background(), bson.M{"email": email})
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
} 