package main

import (
    "log"
    "gerenciador-2fa/internal/database"
    "gerenciador-2fa/internal/handlers"
    "gerenciador-2fa/internal/middleware"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
    // Configurar o Gin para confiar apenas em proxies locais
    gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = log.Writer()
    
    r := gin.New()
    
    // Configurar os IPs confiáveis
    // Aqui estamos confiando apenas em localhost e redes privadas comuns
    r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16", "10.0.0.0/8"})
    
    // Middleware padrão
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(corsMiddleware())

    // Rotas públicas
    public := r.Group("/api/v1")
    {
        public.GET("/health", handlers.HealthCheck)
        public.POST("/register", handlers.Register)
        public.POST("/login", handlers.Login)
    }

    // Rotas protegidas
    protected := r.Group("/api/v1")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/profile", handlers.GetProfile)
        protected.POST("/totp", handlers.AddTOTPAccount)
        protected.GET("/totp", handlers.ListTOTPAccounts)
        protected.DELETE("/totp/:id", handlers.DeleteTOTPAccount)
    }

    return r
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    // Carregar variáveis de ambiente
    if err := godotenv.Load(); err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    // Conectar ao MongoDB
    if err := database.ConnectDB(); err != nil {
        log.Fatal("Não foi possível conectar ao MongoDB:", err)
    }

    // Inicializar o router
    r := setupRouter()

    // Iniciar o servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Erro ao iniciar o servidor:", err)
    }
}