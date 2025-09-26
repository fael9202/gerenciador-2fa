package database

import (
    "context"
    "log"
    "os"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() error {
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        uri = "mongodb://localhost:27017"
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Printf("Erro ao conectar ao MongoDB: %v", err)
        return err
    }

    // Ping no banco para verificar a conexão
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Printf("Erro ao fazer ping no MongoDB: %v", err)
        return err
    }

    log.Println("Conectado ao MongoDB com sucesso!")
    return nil
}

func GetClient() *mongo.Client {
    return client
}

func GetDatabase() *mongo.Database {
    dbName := os.Getenv("MONGODB_DATABASE")
    if dbName == "" {
        dbName = "gerenciador2fa"
    }
    return client.Database(dbName)
}