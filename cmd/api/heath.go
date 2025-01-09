package main

import (
	"net/http"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//
// @produce		json
//
//	@Success		200	{object}	string "ok"
//
// @Router			/health [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//Headers for the response in jason format
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"status": "ok"}`))

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
