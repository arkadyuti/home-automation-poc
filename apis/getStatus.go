package apis

import (
	"errors"
	"fmt"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/gpioHandler"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/helpers"
	"net/http"
	"periph.io/x/periph/conn/gpio"
)

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ReturnErrorResponse(w, http.StatusBadRequest, errors.New("invalid request method"))
		return
	}

	resp := map[string]string{
		"fairyLights": "off",
		"warmLight":   "off",
	}

	for itemId, itemPin := range gpioHandler.LightPintIds {
		fmt.Println("item", itemId, itemPin)
		if itemPin.Read() == gpio.Low {
			resp[itemId] = "on"
		}
		if itemPin.Read() == gpio.High {
			resp[itemId] = "off"
		}
	}
	w.Header().Set("Content-Type", "application/json")
	helpers.ReturnResponse(w, http.StatusOK, resp)
}
