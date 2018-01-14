package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led1 := gpio.NewLedDriver(firmataAdaptor, "13")
	led2 := gpio.NewLedDriver(firmataAdaptor, "12")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led1.Toggle()
		})

		gobot.Every(3*time.Second, func() {
			led2.Toggle()
		})
	}

	robot := gobot.NewRobot("led-bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2},
		work,
	)

	robot.Start()
}