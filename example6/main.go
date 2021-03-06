package main

import (
	"time"

	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	master := gobot.NewMaster()

	esp8266 := firmata.NewTCPAdaptor("192.168.1.9:3030")
	arduino := firmata.NewAdaptor("/dev/ttyUSB0")

	servo := gpio.NewServoDriver(arduino, "4")
	accel := i2c.NewADXL345Driver(esp8266)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			x, _, _ := accel.XYZ()
			angle := uint8((x + 1) * 90)

			if angle < 0 {
				angle = 0
			}
			if angle > 180 {
				angle = 180
			}

			fmt.Println("X:", x)
			fmt.Println("Servo's angle:", angle)
			servo.Move(angle)

		})
	}

	servorobot := gobot.NewRobot("servo-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{servo},
		work,
	)

	accelrobot := gobot.NewRobot("accel-bot",
		[]gobot.Connection{esp8266},
		[]gobot.Device{accel},
		nil,
	)

	master.AddRobot(servorobot)
	master.AddRobot(accelrobot)

	master.Start()

}
