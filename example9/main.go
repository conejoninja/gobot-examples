package main

import (
	"github.com/conejoninja/gobot-examples/example9/dash"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	master := gobot.NewMaster()


	arduino := firmata.NewAdaptor("/dev/ttyUSB0")

	rgbled := gpio.NewRgbLedDriver(arduino, "9", "10", "11")
	servo := gpio.NewServoDriver(arduino, "4")


	// my custom dashboard
	dash.Start(rgbled)


	work := func() {
		rgbled.SetRGB(255, 255, 255)
	}

	rgbledrobot := gobot.NewRobot("RGB-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{rgbled, servo},
		work,
	)

	master.AddRobot(rgbledrobot)

	master.Start()

}
