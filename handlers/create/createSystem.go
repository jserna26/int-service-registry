package create

import (
	"encoding/json"
	"net/http"

	"github.com/jserna26/SystemPOC/handlers"
	"github.com/jserna26/SystemPOC/repo"
	"github.com/jserna26/SystemPOC/types"
)

func NewCreateSystemHandler(sysRepo repo.SystemRepo) http.Handler {
	return handlers.RepoHandler{
		Repo: sysRepo,
		Hf:   createSystemHandler,
	}
}

func createSystemHandler(w http.ResponseWriter, r *http.Request, h handlers.RepoHandler) {
	var sysIn types.NewSystemType
	sysIn.Status = types.StatusEnabled
	err := json.NewDecoder(r.Body).Decode(&sysIn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.ErrorType{Message: "CreateSystemHandler - Decode JSON: " + err.Error()})
		return
	}
	sysOut, err := h.Repo.CreateSystem(sysIn)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.ErrorType{Message: "CreateSystemHandler - Create System: " + err.Error()})
		return
	}

	json.NewEncoder(w).Encode(sysOut)
	w.WriteHeader(http.StatusOK)
}
