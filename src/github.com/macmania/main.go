package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-sphero"
)

func main() {

  spheroAdaptor := new(gobotSphero.SpheroAdaptor)
  spheroAdaptor.Name = "Sphero"
  spheroAdaptor.Port = "/dev/rfcomm0"

  sphero := gobotSphero.NewSphero(spheroAdaptor)
  sphero.Name = "Sphero"

  work := func() {
    gobot.Every("2s", func() {
      sphero.Roll(100, uint16(gobot.Rand(360)))
    })
  }

  robot := gobot.Robot{
    Connections: []gobot.Connection{spheroAdaptor},
    Devices:     []gobot.Device{sphero},
    Work:        work,
  }

  robot.Start()
}
