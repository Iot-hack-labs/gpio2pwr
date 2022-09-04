package power

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
)

type PowerStrip struct {
	PowerOutlets map[string]rpio.Pin
}

func NewPowerStrip() *PowerStrip {
	if err := rpio.Open(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return &PowerStrip{PowerOutlets: make(map[string]rpio.Pin)}
}

func (p *PowerStrip) AddOutlet(name string, GPIOPin uint8) {
	p.PowerOutlets[name] = rpio.Pin(GPIOPin)
	p.PowerOutlets[name].Output()
	p.PowerOutlets[name].High()
}

func (p *PowerStrip) On(name string) {
	if outlet, ok := p.PowerOutlets[name]; ok {
		outlet.Low()
	}
}

func (p *PowerStrip) Off(name string) {
	if outlet, ok := p.PowerOutlets[name]; ok {
		outlet.High()
	}
}

func (p *PowerStrip) Toggle(name string) {
	if outlet, ok := p.PowerOutlets[name]; ok {
		outlet.Toggle()
	}
}

func (p *PowerStrip) AllOn() {
	for _, outlet := range p.PowerOutlets {
		outlet.Low()
	}
}

func (p *PowerStrip) AllOff() {
	for _, outlet := range p.PowerOutlets {
		outlet.High()
	}
}

func (p *PowerStrip) ToggleAll() {
	for _, outlet := range p.PowerOutlets {
		outlet.Toggle()
	}
}

func (p *PowerStrip) Close() {
	rpio.Close()
}
