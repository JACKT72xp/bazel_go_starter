package main

import (
    "example/bazel-go-gin-starter/database"
    "example/bazel-go-gin-starter/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Inicializa la conexión a la base de datos
    db := database.InitDB()
    defer db.Close()

    // Inicializa las rutas HTTP
	router := gin.Default()
    routes.InitRoutes(router) // Modifica InitRoutes para aceptar *gin.Engine como argumento

    router.Run(":8080")

}
