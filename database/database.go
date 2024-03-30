package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
    // Configura la cadena de conexión a la base de datos
    // db, err := sql.Open("mysql", "root:xxx.@tcp(1x.x.x.x:3306)/bifrost_base")
    db, err := sql.Open("mysql", "__here_database_connection_string__")
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
