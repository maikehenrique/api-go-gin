package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikehenrique/api-go-gin/database"
	"github.com/maikehenrique/api-go-gin/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// Saudacao
//
//	@Summary		Exige saudação
//	@Description	Rota para exibir saudação
//	@Accept			json
//	@Produce		json
//	@Param			nome	path 		string				true	"Seu nome"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

// ExibeTodosAlunos
//
//	@Summary		Exibe todos Alunos
//	@Description	Rota para exibir todos os produtos
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// CriaNovoAluno
//
//	@Summary		Cria novo aluno
//	@Description	Rota para criar novo alunos
//	@ID				alunos
//	@Accept			json
//	@Produce		json
//	@Param			aluno	body		models.Aluno				true	"Modelo do Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID
//
//	@Summary		Busca um aluno por seu ID
//	@Description	Rota para criar buscar um aluno por sua chave de identificação única ID
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"ID do Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno
//
//	@Summary		Deleta um aluno por seu ID
//	@Description	Rota para criar deletar um aluno por sua chave de identificação única ID
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"ID do Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno
//
//	@Summary		Edita um aluno
//	@Description	Rota para alterar os dados de um aluno
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"ID do Aluno"
//	@Param			aluno	body		models.Aluno				true	"Modelo do Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaDadosAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF
//
//	@Summary		Busca um aluno por seu CPF
//	@Description	Rota para busca um aluno por seu CPF
//	@Accept			json
//	@Produce		json
//	@Param			cpf		path		string				true	"CPF do Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Router			/alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func PaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
