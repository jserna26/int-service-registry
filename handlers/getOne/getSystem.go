package getOne

import (
	"encoding/json"
	"net/http"

	"github.com/jserna26/SystemPOC/handlers"
	"github.com/jserna26/SystemPOC/repo"
	"github.com/jserna26/SystemPOC/types"
)

func NewGetSystemHandler(sysRepo repo.SystemRepo) http.Handler {
	return handlers.RepoHandler{
		Repo: sysRepo,
		Hf:   getSystemHandlerFunc,
	}
}

func getSystemHandlerFunc(w http.ResponseWriter, r *http.Request, h handlers.RepoHandler) {

	vars := r.URL.Query()
	system := vars["Name"][0]

	sysOut, err := h.Repo.GetSystem(system)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.ErrorType{Message: "GetSystemsHandler - Get System: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(sysOut)
	w.WriteHeader(http.StatusOK)

}
