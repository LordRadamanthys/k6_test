package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/test/client"
	"github.com/test/model"
	"github.com/test/repository"
	"github.com/test/utils/config"
)

type Teste struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Ola")
	config.CreateMongoConnection()
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "qualquer coisa")
	})

	r.POST("/test", func(ctx *gin.Context) {
		body := Teste{}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, body)

	})

	r.GET("/anime/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)
		a, err := client.GetAnime(idInt)

		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		ctx.JSON(200, a)

	})

	r.POST("/anime", func(ctx *gin.Context) {

		animeBody := model.Anime{}

		if err := ctx.ShouldBindJSON(&animeBody); err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		if err := repository.SaveAnime(animeBody); err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		ctx.JSON(201, nil)

	})

	r.Run(":8080")

}
