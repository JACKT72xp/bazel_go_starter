package routes

import (
    "github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
    // Configura tus rutas aquí, por ejemplo:
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    // Define tus rutas HTTP
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
}

// Funciones de controlador para las rutas
func getAlbums(c *gin.Context) {
    // Lógica para obtener todos los álbumes
}

func postAlbums(c *gin.Context) {
    // Lógica para agregar un nuevo álbum
}
