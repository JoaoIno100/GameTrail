package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return rotas.Configurar(router)
}
