package router

import "github.com/gin-gonic/gin"

// CORSMiddleware intercepta e autoriza as requisições do seu frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Responde ao Preflight (OPTIONS) do navegador
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Initialize() {
	router := gin.Default()
	
	// ATENÇÃO: O middleware precisa ser ativado ANTES de iniciar as rotas
	router.Use(CORSMiddleware())
	
	InitializeRoutes(router)
	router.Run(":8080")
}