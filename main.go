package main

import (
    // Importing necessary packages: the database package for database operations,
    // the routes package for HTTP route definitions, and the gin package for the web framework.
    "example/bazel-go-gin-starter/database"
    "example/bazel-go-gin-starter/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database connection using the InitDB function from the database package.
    // The InitDB function is expected to return a database connection object.
    db := database.InitDB()

    // Defer the closure of the database connection until the end of the main function.
    // This ensures that the database connection will be closed properly before the program exits.
    defer db.Close()

    // Initialize the HTTP router using Gin's Default method which provides some default middleware.
    router := gin.Default()

    // Initialize the HTTP routes. The InitRoutes function from the routes package is called,
    // passing the router object. This function is expected to set up all the route handlers.
    routes.InitRoutes(router) // Modify InitRoutes to accept *gin.Engine as an argument.

    // Start the HTTP server on port 8080 and begin listening for requests.
    router.Run(":8080")
}
