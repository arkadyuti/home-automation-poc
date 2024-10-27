package apis

import (
	"encoding/json"
	"errors"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/gpioHandler"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/helpers"
	"net/http"
)

type ToggleRequest struct {
	BulbID string `json:"bulbId"`
	State  string `json:"state"`
}

type Response struct {
	Message string `json:"message"`
}

func ToggleLightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ReturnErrorResponse(w, http.StatusBadRequest, errors.New("invalid request method"))
		return
	}

	var req ToggleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ReturnErrorResponse(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if req.BulbID == "" || req.State == "" {
		helpers.ReturnErrorResponse(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if req.State == "on" {
		gpioHandler.TogglePinChan <- gpioHandler.TogglePinPayload{
			PinId:  gpioHandler.LightPintIds[req.BulbID],
			Event:  "on",
			BulbId: req.BulbID,
		}
	}

	if req.State == "off" {
		gpioHandler.TogglePinChan <- gpioHandler.TogglePinPayload{
			PinId:  gpioHandler.LightPintIds[req.BulbID],
			Event:  "off",
			BulbId: req.BulbID,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.ReturnResponse(w, http.StatusOK, Response{Message: "done"})
}
