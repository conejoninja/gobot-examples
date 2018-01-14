package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"fmt"
	"gobot.io/x/gobot/drivers/i2c"
	"time"
)

func main() {

	master := gobot.NewMaster()

	esp8266 := firmata.NewTCPAdaptor("192.168.1.9:3030")
	arduino := firmata.NewAdaptor("/dev/ttyUSB0")

	rgbled := gpio.NewRgbLedDriver(arduino, "9", "10", "11")
	light := i2c.NewBH1750Driver(esp8266)


	work := func() {
		gobot.Every(1*time.Second, func() {
			lux, _ := light.Lux()
			fmt.Println("lux:", lux)
			if lux > 112 {
				lux = 112
			}
			rgbled.SetRGB(255, byte(255-2*lux), byte(255-lux/2))
		})
	}

	rgbledrobot := gobot.NewRobot("RGB-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{rgbled},
		work,
	)

	lightrobot := gobot.NewRobot("light-bot",
		[]gobot.Connection{esp8266},
		[]gobot.Device{light},
		nil,
	)

	master.AddRobot(rgbledrobot)
	master.AddRobot(lightrobot)

	master.Start()

}
