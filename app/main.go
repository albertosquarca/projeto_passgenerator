package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// Obtém o diretório de trabalho atual
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Constrói o caminho absoluto para a pasta de modelos
	templatesDir := filepath.Join(currentDir, "templates")

	r := gin.Default()

	// Carrega modelos HTML do caminho absoluto
	r.LoadHTMLGlob(filepath.Join(templatesDir, "*.html"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
