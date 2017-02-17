package handlers

import (
	"net/http"

	"github.com/jserna26/SystemPOC/repo"
)

type HandlerFunction func(w http.ResponseWriter, r *http.Request, h RepoHandler)

type RepoHandler struct {
	Repo repo.SystemRepo
	Hf   HandlerFunction
}

func (h RepoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Hf(w, r, h)
}
