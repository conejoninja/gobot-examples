package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewTCPAdaptor("192.168.1.9:3030")
	led1 := gpio.NewLedDriver(firmataAdaptor, "15")
	led2 := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led1.Toggle()
		})

		gobot.Every(3*time.Second, func() {
			led2.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2},
		work,
	)

	robot.Start()
}
