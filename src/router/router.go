package router

import (
	"github.com/gorilla/mux"
	"github.com/jailcomfranssa/api-golang/src/router/rotas"
)

//Gerar vai retornar um router com as totas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
