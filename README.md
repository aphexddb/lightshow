# lightshow

RGB LED light effects and things for the Raspberry Pi.

Arduino's are tricky for anything fancy regarding services, limited in RAM and plus I wanted to use Golang. Getting PWM to work on a RPI is a challenge! All dependencies are vendored for convenience.

## Hardware

This project uses a Raspberry Pi 3.

<img src="images/pi3_gpio.png" width="480px"/>

## Linux

Use the [Minibian](https://minibianpi.wordpress.com/download/) distro. Setup your SSH access as you prefer, default user is `root` / `raspberry`.

### PWM Support

We need the PWM Linux kernel module, I used [this guide](https://learn.adafruit.com/adafruit-raspberry-pi-lesson-9-controlling-a-dc-motor/the-pwm-kernel-module):

1. Install Hexxeh firmware
  ```
  wget https://raw.githubusercontent.com/Hexxeh/rpi-update/master/rpi-update
  bash rpi-update
  ```
2. asd

## Links

* [Different Styles of Pixels](http://www.doityourselfchristmas.com/wiki/index.php?title=Different_Styles_of_Pixels)
* [Pixel Wiring Colors](http://www.doityourselfchristmas.com/wiki/index.php?title=Pixel_Wiring_Colors)
* [GE Color Effects Bulbs](http://www.geholidaylighting.com/color-effects/us/items/72425)
  * [Hacking the remote](https://lukecyca.com/2013/g35-rf-remote.html)
* [Arduino GECE library](https://github.com/sowbug/G35Arduino)
  * [Related forum thread on using it](http://doityourselfchristmas.com/forums/showthread.php?37502-E1-31-Arduino-GECE-controller)
* [Another blog post on hacking the GE lights](http://culverlabs.co/post/38529142159/christmaslights)

PWM and GPIO things

* [How to use GPIOs on raspberry pi (Simple I/O, PWM and UART)](https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio)
* [Raspberry Pi GPIO via the Shell](https://luketopia.net/2013/07/28/raspberry-pi-gpio-via-the-shell/)
* [Wiring Pi](http://wiringpi.com/)
  * [WiringPi-Go: Golang wrapped version](https://github.com/hugozhu/rpi)
* [pi golang gpio module](https://github.com/stianeikeland/go-rpio/blob/master/rpio.go)


## Build for Pi

```
export GOOS=linux
export GOARCH=arm
go build *.go
```
