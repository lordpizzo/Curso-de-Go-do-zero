package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lógica de log
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lógica de autenticação
		if erro := auth.ValidarToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
