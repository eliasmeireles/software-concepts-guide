package main

import (
	statelessvsstateful "github.com/eliasmeireles/software-concepts-guide/software-architecture/stateless-vs-stateful"
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/server"
	"strconv"
)

// Estrutura para armazenar o estado da sessão
type Session struct {
	VisitCount int
}

// Mapa para armazenar as sessões dos usuários
var sessions = make(map[string]*Session)

// Exemplo 2: Stateful (Sessão de Usuário)
//
// Neste exemplo, vamos criar um sistema stateful que mantém o estado de uma sessão de usuário.
// O servidor armazena informações sobre o cliente (como um contador de visitas) entre as requisições.
func main() {
	server.Default().
		RegisterMiddleware(sectionHandler, "SECTION/HANDLER").
		Get(statelessvsstateful.GetUsers, "users").
		Get(statelessvsstateful.GetUser, "users/{id}").
		NotFoundHandler().
		StartServer()
}

// Handler para uma requisição stateful
func sectionHandler(ctx *api_context.ApiRequestContext[*api_context.DefaultContext]) (doNext bool) {
	sessionsIds := ctx.Headers["X-Session-Id"]
	sessionID := ""

	if len(sessionsIds) > 0 {
		sessionID = sessionsIds[0]
		if sessionID == "" {
			sessionID = "sessao-" + strconv.Itoa(len(sessions)+1)
		}
	} else {
		sessionID = "sessao-" + strconv.Itoa(len(sessions)+1)
	}

	(*ctx.Writer).Header().Set("X-Session-ID", sessionID)

	// Verifica se a sessão já existe
	session, exists := sessions[sessionID]
	if !exists {
		// Cria uma nova sessão
		session = &Session{VisitCount: 0}
		sessions[sessionID] = session
	}

	// Incrementa o contador de visitas
	session.VisitCount++
	return true

}
