package toggle

import (
	"fmt"
	"gitlab.com/arkadooti.sarkar/relay-raspberry-pi/gpioHandler"
	"time"

	"github.com/go-co-op/gocron"
)

var DayTimeSchedulerChannel = make(chan string)

func DayTimeScheduler() {
	sunriseTime := "02:00:00"
	sunsetTime := "18:00:00"

	// Scheduler to trigger events at sunrise and sunset
	scheduler := gocron.NewScheduler(time.Local)

	// Schedule the sunrise event
	scheduler.Every(1).Day().At(sunriseTime).Do(triggerSunriseEvent)

	// Schedule the sunset event
	scheduler.Every(1).Day().At(sunsetTime).Do(triggerSunsetEvent)

	fmt.Printf("Scheduled sunrise event at %s and sunset event at %s\n", sunriseTime, sunsetTime)
	fmt.Printf("Time now: %s", time.Now().String())

	scheduler.StartBlocking()
}

func triggerSunriseEvent() {
	for _, pinId := range gpioHandler.LightPintIds {
		gpioHandler.TogglePinChan <- gpioHandler.TogglePinPayload{
			PinId: pinId,
			Event: "off",
		}
	}
}

// Function to trigger at sunset
func triggerSunsetEvent() {
	gpioHandler.TogglePinChan <- gpioHandler.TogglePinPayload{
		PinId: gpioHandler.LightPintIds["fairyLights"],
		Event: "on",
	}
}
