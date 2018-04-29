package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:123456@tcp(golangdock_db_1)/golang")

type Product struct {
    Id int `db:"id" json:"id"`
    Name string `db:"name" json:"title"`
    Price int `db:"price" json:"completed"`
}

func getAll(c *gin.Context){
    var (
        product Product
        products []Product
    )

    rows, err := db.Query("select id, name, price from product;")

    if err != nil {
        c.JSON(422, gin.H{"error": err})
    }

    for rows.Next() {
        rows.Scan(&product.Id, &product.Name, &product.Price)
        products = append(products, product)
    }

    defer rows.Close()

    c.JSON(http.StatusOK, products)
}

func getById(c *gin.Context) {
    var product Product

    id := c.Param("id")
    row := db.QueryRow("select id, name, price from product where id = ?;", id)

    err = row.Scan(&product.Id, &product.Name, &product.Price)
    if err != nil {
        c.JSON(http.StatusOK, "Fallo en la BD")
    } else {
        c.JSON(http.StatusOK, product)
    }
}

func main() {

    if err != nil {
        fmt.Print(err.Error())
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Print(err.Error())
    }

    router := gin.Default()

    //Test
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Tohure")
    })

    router.GET("/api/product/:id", getById)
    router.GET("/api/products", getAll)
    //router.POST("/api/product", add)
    //router.PUT("/api/product/:id", update)
    //router.DELETE("/api/product/:id", delete)

    router.Run(":8080")
}