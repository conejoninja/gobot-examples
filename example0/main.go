package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	pin := gpio.NewDirectPinDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			pin.On()
			time.Sleep(1 * time.Second)
			pin.Off()
			time.Sleep(1 * time.Second)
		})

	}

	robot := gobot.NewRobot("pin-bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{pin},
		work,
	)

	robot.Start()
}