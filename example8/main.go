package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	master := gobot.NewMaster()

	a := api.NewAPI(master)
	a.Start()

	arduino := firmata.NewAdaptor("/dev/ttyUSB0")
	red := gpio.NewLedDriver(arduino, "9")
	red.SetName("LED-ROJO")

	green := gpio.NewLedDriver(arduino, "10")
	green.SetName("LED-VERDE")

	blue := gpio.NewLedDriver(arduino, "11")
	blue.SetName("LED-AZUL")

	work := func() {
		red.On()
		green.On()
		blue.On()
	}

	rgbledrobot := gobot.NewRobot("RGB-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{red, green , blue},
		work,
	)

	master.AddRobot(rgbledrobot)

	master.Start()

}
