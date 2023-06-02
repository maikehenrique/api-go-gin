package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maikehenrique/api-go-gin/controllers"
	"github.com/maikehenrique/api-go-gin/database"
	"github.com/maikehenrique/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {

	aluno := models.Aluno{
		Nome: "Nome do Aluno Teste",
		CPF:  "12364578901",
		RG:   "123456789",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/mai", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E ai mai, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody), "Deveriam ser iguais")
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12364578901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12364578901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{
		Nome: "Nome do Aluno Teste",
		CPF:  "47364578901",
		RG:   "123456700",
	}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "47364578901", alunoMock.CPF)
	assert.Equal(t, "123456700", alunoMock.RG)

}
