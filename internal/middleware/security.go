package middleware

import (
	"github.com/gin-gonic/gin"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Segurança contra XSS
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		// Segurança contra sniffing de MIME
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		// Segurança de Frame (previne clickjacking)
		c.Writer.Header().Set("X-Frame-Options", "DENY")

		// Política de Conteúdo (CSP)
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")

		// Segurança para Referer
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")

		// HSTS (somente HTTPS)
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// Continuação da requisição
		c.Next()
	}
}
