package main

import (
	"fmt"
	"net"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	HyperionFrameBufferLimit = 490 * 3 // max 490 leds, 3 bytes for each led
)

type wsEngine interface {
	Init() error
	Render() error
	Wait() error
	Fini()
	Leds(channel int) []uint32
}

type hyperionFeed struct {
	ws wsEngine
}

func (hf *hyperionFeed) setup() error {
	return hf.ws.Init()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = 255
	opt.Channels[0].LedCount = 113

	dev, err := ws2811.MakeWS2811(&opt)
	checkError(err)

	checkError(dev.Init())
	defer dev.Fini()

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, HyperionFrameBufferLimit)
		rlen, _, err := conn.ReadFromUDP(message)
		if err != nil {
			panic(err)
		}

		ledNo := 0
		for x := 0; x < rlen; x += 3 {
			r := uint32(message[x])
			g := uint32(message[x+1])
			b := uint32(message[x+2])

			rgb := r
			rgb = (rgb << 8) + g
			rgb = (rgb << 8) + b

			dev.Leds(0)[ledNo] = rgb

			// logrus.WithFields(logrus.Fields{"rgb": rgb, "ledNo": ledNo}).Info("Read LED state")

			ledNo++
		}
		checkError(dev.Render())

		// logrus.WithFields(logrus.Fields{"frameLength": rlen}).Info("Read frame")
	}
}
