package PowerStrip

import "github.com/Iot-hack-labs/gpio2pwr/pkg/power"

const RotatingLight = "RotatingLight"
const Lamp = "Lamp"
const Fan = "Fan"
const Speakers = "Speakers"

func New() *power.PowerStrip {
	Instance := power.NewPowerStrip()
	Instance.AddOutlet(RotatingLight, 5)
	Instance.AddOutlet(Lamp, 6)
	Instance.AddOutlet(Fan, 19)
	Instance.AddOutlet(Speakers, 26)
	return Instance
}
