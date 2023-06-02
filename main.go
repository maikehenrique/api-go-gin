package main

import (
	"github.com/maikehenrique/api-go-gin/database"
	"github.com/maikehenrique/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
