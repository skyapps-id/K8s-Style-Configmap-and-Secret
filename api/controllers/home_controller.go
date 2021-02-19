package controllers

import (
	"net/http"

	"github.com/skyapps-id/K8s-Style-Configmap-and-Secret/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
