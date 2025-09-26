package database

import (
	"context"
	"gerenciador-2fa/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTOTPAccount(account *models.TOTPAccount) error {
	collection := GetDatabase().Collection("totp_accounts")
	_, err := collection.InsertOne(context.Background(), account)
	return err
}

func GetUserTOTPAccounts(userID primitive.ObjectID) ([]models.TOTPAccount, error) {
	collection := GetDatabase().Collection("totp_accounts")
	
	cursor, err := collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var accounts []models.TOTPAccount
	if err = cursor.All(context.Background(), &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func DeleteTOTPAccount(userID, accountID primitive.ObjectID) error {
	collection := GetDatabase().Collection("totp_accounts")
	
	_, err := collection.DeleteOne(context.Background(), bson.M{
		"_id": accountID,
		"user_id": userID,
	})
	return err
} 