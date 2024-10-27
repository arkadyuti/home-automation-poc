package gpioHandler

import (
	"encoding/json"
	"fmt"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/helpers"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

type TogglePinPayload struct {
	PinId  gpio.PinIO
	Event  string
	BulbId string
}

var (
	LightPintIds = map[string]gpio.PinIO{
		"fairyLights": rpi.P1_40,
		"warmLight":   rpi.P1_38,
	}
)

var TogglePinChan = make(chan TogglePinPayload)

func init() {
	go TogglePin()
}

func TogglePin() {
	if _, err := host.Init(); err != nil {
		fmt.Println("Failed to initialize periph.io:", err)
		return
	}
	for event := range TogglePinChan {
		pinStatusRaw, err := helpers.ReadFile("pin-status.json")
		if err != nil {
			pinStatusRaw = []byte("{}")
		}

		var pinStatus map[string]bool
		json.Unmarshal(pinStatusRaw, &pinStatus)

		switch event.Event {
		case "on":
			fmt.Println("Light on for: ", event.BulbId)
			pinStatus[event.BulbId] = true
			event.PinId.Out(gpio.Low)
		case "off":
			fmt.Println("Light off for: ", event.BulbId)
			pinStatus[event.BulbId] = false
			event.PinId.Out(gpio.High)
		}
		data, _ := json.Marshal(pinStatus)
		err = helpers.WriteStringToFile("pin-status.json", string(data))
		if err != nil {
			fmt.Println("TogglePin: write error:", err)
			return
		}
		fmt.Println("pin status in file after: ", string(data))
	}
}
