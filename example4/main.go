package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"fmt"
)

func main() {
	master := gobot.NewMaster()

	esp8266 := firmata.NewTCPAdaptor("192.168.1.9:3030")
	arduino := firmata.NewAdaptor("/dev/ttyUSB0")

	led := gpio.NewLedDriver(arduino, "13")
	button := gpio.NewButtonDriver(esp8266, "4")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("ON")
			led.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("OFF")
			led.Off()
		})
	}

	ledrobot := gobot.NewRobot("ledbot",
		[]gobot.Connection{arduino},
		[]gobot.Device{led},
		work,
	)

	buttonrobot := gobot.NewRobot("buttonbot",
		[]gobot.Connection{esp8266},
		[]gobot.Device{button},
		nil,
	)

	master.AddRobot(ledrobot)
	master.AddRobot(buttonrobot)

	master.Start()
	
}
