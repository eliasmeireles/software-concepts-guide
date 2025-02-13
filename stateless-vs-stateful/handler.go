package stateless_vs_stateful

import (
	"github.com/softwareplace/http-utils/api_context"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(ctx *api_context.ApiRequestContext[*api_context.DefaultContext]) {
	log.Printf("GET /users")
	ctx.Response(UsersData, http.StatusOK)

}

func GetUser(ctx *api_context.ApiRequestContext[*api_context.DefaultContext]) {
	log.Printf("GET /users/{id}")
	id := ctx.PathValues["id"]
	if id != "" {
		for _, user := range UsersData {
			requestId, _ := strconv.Atoi(id)

			if user.ID == requestId {
				ctx.Response(user, http.StatusOK)
				return
			}
		}
	}

	ctx.Response(map[string]string{"error": "User not found"}, http.StatusNotFound)

}
