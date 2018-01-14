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

	rgbled := gpio.NewRgbLedDriver(arduino, "9", "10", "11")
	accel := i2c.NewADXL345Driver(esp8266)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			x, y, z := accel.XYZ()
			r := (x + 1) * 128
			g := (y + 1) * 128
			b := (z + 1) * 128

			if r < 0 {
				r = 0
			}
			if g < 0 {
				g = 0
			}
			if b < 0 {
				b = 0
			}

			if r > 255 {
				r = 255
			}
			if g > 255 {
				g = 255
			}
			if b > 255 {
				b = 255
			}

			fmt.Println(x, y, z)
			rgbled.SetRGB(byte(b), byte(g), byte(r))

		})
	}

	rgbledrobot := gobot.NewRobot("RGB-bot",
		[]gobot.Connection{arduino},
		[]gobot.Device{rgbled},
		work,
	)

	accelrobot := gobot.NewRobot("accel-bot",
		[]gobot.Connection{esp8266},
		[]gobot.Device{accel},
		nil,
	)

	master.AddRobot(rgbledrobot)
	master.AddRobot(accelrobot)

	master.Start()

}
