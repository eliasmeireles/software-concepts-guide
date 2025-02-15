package main

import (
	statelessvsstateful "github.com/eliasmeireles/software-concepts-guide/software-architecture/stateless-vs-stateful"
	"github.com/softwareplace/http-utils/server"
)

// ### Exemplo 1: Stateless (API RESTful)
//
// Neste exemplo, vamos criar uma API RESTful simples que não mantém nenhum estado entre as requisições.
// Cada requisição é independente e o servidor não armazena nenhuma informação sobre o cliente.

func main() {
	server.Default().
		Get(statelessvsstateful.GetUsers, "users").
		Get(statelessvsstateful.GetUser, "users/{id}").
		NotFoundHandler().
		StartServer()
}
