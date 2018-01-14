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

	servo := gpio.NewServoDriver(arduino, "4")
	light := i2c.NewBH1750Driver(esp8266)


	work := func() {
		gobot.Every(1*time.Second, func() {
			lux, _ := light.Lux()
			fmt.Println("lux:", lux)
			if lux > 180 {
				lux = 180
			}
			servo.Move(uint8(lux))
		})
	}

	servorobot := gobot.NewRobot("servo-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{servo},
		work,
	)

	lightrobot := gobot.NewRobot("light-bot",
		[]gobot.Connection{esp8266},
		[]gobot.Device{light},
		nil,
	)

	master.AddRobot(servorobot)
	master.AddRobot(lightrobot)

	master.Start()

}
