package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
    // Configura la cadena de conexión a la base de datos
    db, err := sql.Open("mysql", "root:Maniac321.@tcp(192.168.36.90:3306)/bifrost_base")
    if err != nil {
        panic(err.Error())
    }

    // Verifica si la conexión a la base de datos es exitosa
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Database connected successfully!")

    // Retorna la conexión a la base de datos
    return db
}
