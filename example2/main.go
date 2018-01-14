package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")
	button := gpio.NewButtonDriver(firmataAdaptor, "2")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("led-button-bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led, button},
		work,
	)

	robot.Start()
}
