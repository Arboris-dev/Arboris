package httpwriters

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	jsonData, marshalErr := json.Marshal(payload)

	if marshalErr != nil {
		slog.Error("Unable to marshal payload", "ERROR", marshalErr)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	_, writeErr := w.Write(jsonData)

	if writeErr != nil {
		slog.Error("Unable to write the payload", "ERROR", writeErr)
	}
	return
}

func RespondWithErr(w http.ResponseWriter, errMsg string, statusCode int) {
	if statusCode > 499 {
		slog.Error("Client side error", "ERROR", errMsg)
	}
	type errResponse struct {
		msg string
	}
	RespondWithJSON(w, statusCode, errResponse{
		msg: errMsg,
	})
	return
}
