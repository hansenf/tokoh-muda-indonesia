package controller

import (
	"context"
	"encoding/json"
	"net/http"
	client "tmi-native/db"
	db "tmi-native/db/sqlc"
	"tmi-native/utils"

	"github.com/julienschmidt/httprouter"
)

//POST
func PostSilabus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var silabus db.CreateSilabusParams

		if err := json.NewDecoder(r.Body).Decode(&silabus); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.CreateSilabus(ctx, silabus); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}
