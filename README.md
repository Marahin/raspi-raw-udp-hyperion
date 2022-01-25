## raspi-raw-udp-hyperion

This is a simple UDP server that gets RAW UDP input from [hyperion.ng](https://github.com/hyperion-project/hyperion.ng) and feeds the WS2812B LED strip with it. 

It uses [rpi-ws281x-go](https://github.com/rpi-ws281x/rpi-ws281x-go) to talk to the WS2812B LED strip.

### how to

* take any steps necessary to have the [builder image ready](https://github.com/rpi-ws281x/rpi-ws281x-go#cross-compiling)

* clone this repo
* change LED count to the amount of leds you're using
* `make xbuild`
* put it on your raspberry
* run it

#### systemd.service

i also use a systemd service so its easier to run and forget:


```sh
# /lib/systemd/system/rruh.service
[Unit]
Description=Start rruh
After=multi-user.target

[Service]
ExecStart=/home/pi/rruh-armv6

[Install]
WantedBy=multi-user.target
```