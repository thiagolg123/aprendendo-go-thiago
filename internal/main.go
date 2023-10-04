package main

import (
	"fmt"
	"net/http"

	sqlConn "github.com/thiagolg123/aprendendo-go-thiago/internal/databasePostgres"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

func main() {
	conn, e := sqlConn.CreateConn()

	if e != nil {
		panic(e)
	}
	defer conn.Close()
	fmt.Println("CONEXAO COM POSTGRES OK!")

	router := gin.Default()

	router.GET("/getAlbums", getAlbuns)
	router.Run("127.0.0.1:8080")
}

func createAlbum() []album {
	var albums = []album{
		{ID: "1", Title: "legal", Artist: "Thiago", Price: 10.25},
		{ID: "2", Title: "cool", Artist: "Thayna", Price: 20.25},
		{ID: "3", Title: "thasSoCool", Artist: "MAMAE", Price: 30.25},
	}

	return albums
}

func getAlbuns(ginCon *gin.Context) {
	//find in db or something

	ginCon.IndentedJSON(http.StatusOK, createAlbum())
}
