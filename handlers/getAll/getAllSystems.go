package getAll

import (
	"encoding/json"
	"net/http"

	"github.com/jserna26/SystemPOC/handlers"
	"github.com/jserna26/SystemPOC/repo"
	"github.com/jserna26/SystemPOC/types"
)

func NewGetAllSystemsHandler(sysRepo repo.SystemRepo) http.Handler {
	return handlers.RepoHandler{
		Repo: sysRepo,
		Hf:   getAllSystemsHandlerFunc,
	}
}

func getAllSystemsHandlerFunc(w http.ResponseWriter, r *http.Request, h handlers.RepoHandler) {
	sysOut, err := h.Repo.GetAllSystems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(types.ErrorType{Message: "GetAllSystemsHandler - Get All Systems: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(sysOut)
	w.WriteHeader(http.StatusOK)
}

/*func GetAllSystemsHandler(sysRepo repo.SystemRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.WriteHeader(http.StatusOK)
		sysOut, err := sysRepo.GetAllSystems()
		fmt.Printf("Test: %s", sysOut)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(types.ErrorType{Message: "GetAllSystemsHandler - Get All Systems: " + err.Error()})
			return
		}
		json.NewEncoder(w).Encode(sysOut)
		w.WriteHeader(http.StatusOK)
	}
}*/
