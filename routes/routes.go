package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maikehenrique/api-go-gin/controllers"
	docs "github.com/maikehenrique/api-go-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.PaginaIndex)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
