# lightshow

RGB LED light effects and things for the Raspberry Pi.

Arduino's are tricky for anything fancy regarding services, limited in RAM and plus I wanted to use Golang. Getting software PWM to work on a RPI is a bit of a challenge! All dependencies are vendored for convenience.

This project not possible without the original Arduino [github.com/sowbug/G35Arduino](https://github.com/sowbug/G35Arduino) project for inspiration and [WiringPi](http://wiringpi.com/)

## Hardware

This project uses a [Raspberry Pi 3](https://www.amazon.com/Raspberry-Pi-RASP-PI-3-Model-Motherboard/dp/B01CD5VC92/ref=sr_1_1?s=pc&ie=UTF8&qid=1480741328&sr=1-1&keywords=Raspberry+Pi+3) and GE Color Effects Lights that you can buy at [Lowe's](https://www.lowes.com/pd/GE-Color-Effects-36-Count-29-17-ft-Multi-Function-Color-Changing-G28-LED-Plug-in-Indoor-Outdoor-Christmas-String-Lights-ENERGY-STAR/1000110759). You should be able to support other LED pixels if you know what you are doing around timings and such.

Pi Model 3 Pinout Reference:

<img src="images/pi3_gpio.png" width="300px"/>

Raspberry PWM Pinout:

| Header Pin    | Name          | BCM GPIO |
| ------------- | ------------- | -------- |
| 12            | PWM0          | 18       |

## Linux Setup

* Used [Minibian](https://minibianpi.wordpress.com/download/) distro. Setup your SSH access as you prefer, default user is `root` / `raspberry`.
* Used ethernet for internet access, you can do wifi if you please

### PWM Support via Pi Blaster

1. Disable Raspberry Pi audio (it uses PWM). You will need to blacklist the Broadcom audio kernel module by creating a file `/etc/modprobe.d/snd-blacklist.conf` with

    ```
    blacklist snd_bcm2835
    ```

    If the audio device is still loading after blacklisting, you may also need to comment it out in the `/etc/modules` file. Some distributions use audio by default, even if nothing is being played. If audio is needed, you can use a USB audio device instead.

2. Install latest Pi firmware
    ```
    wget https://raw.githubusercontent.com/Hexxeh/rpi-update/master/rpi-update
    bash rpi-update
    ```
    Reboot the pi

3. Install Pi Blaster
  ```
  sudo apt-get install pi-blaster
  ```

  pi-blaster creates a special file (FIFO) in `/dev/pi-blaster`. Any application on your Raspberry Pi can write to it (this means that only pi-blaster needs to be root, your application can run as a normal user).

  Important: when using pi-blaster, the GPIO pins you send to it are configured as output.

  To set the value of a PIN, you write a command to /dev/pi-blaster in the form &lt;GPIOPinName&gt;=&lt;value&gt; where <value> must be a number between 0 and 1 (included).

  You must use the GPIO number (BCM xx in the diagram below).

## Running

```
go build
./lightshow
```

## Links

To investigate

* [go bindings for Raspberry Pi PWM library for WS281X LEDs](https://github.com/mcuadros/go-rpi-ws281x)
  * [Userspace Raspberry Pi PWM library for WS281X LEDs](https://github.com/jgarff/rpi_ws281x/tree/27229b24028082b2ef9314c6247f621af296486d)


* [Different Styles of Pixels](http://www.doityourselfchristmas.com/wiki/index.php?title=Different_Styles_of_Pixels)
* [Pixel Wiring Colors](http://www.doityourselfchristmas.com/wiki/index.php?title=Pixel_Wiring_Colors)
* [GE Color Effects Bulbs](http://www.geholidaylighting.com/color-effects/us/items/72425)
  * [Hacking the remote](https://lukecyca.com/2013/g35-rf-remote.html)
* [Arduino GECE library](https://github.com/sowbug/G35Arduino)
  * [Related forum thread on using it](http://doityourselfchristmas.com/forums/showthread.php?37502-E1-31-Arduino-GECE-controller)
* [Another blog post on hacking the GE lights](http://culverlabs.co/post/38529142159/christmaslights)

PWM and GPIO things

* [Python and RPi.GPIO](http://raspi.tv/2013/rpi-gpio-0-5-2a-now-has-software-pwm-how-to-use-it)
* [Customize GPIO and PWM clock on linux](https://www.raspberrypi.org/documentation/configuration/pin-configuration.md)
* [How to use GPIOs on raspberry pi (Simple I/O, PWM and UART)](https://sites.google.com/site/semilleroadt/raspberry-pi-tutorials/gpio)
* [Raspberry Pi GPIO via the Shell](https://luketopia.net/2013/07/28/raspberry-pi-gpio-via-the-shell/)
* [Wiring Pi](http://wiringpi.com/)
  * [WiringPi-Go: Golang wrapped version](https://github.com/hugozhu/rpi)
* [pi golang gpio module](https://github.com/stianeikeland/go-rpio/blob/master/rpio.go)

Lights and LED

* [Hacking Christmas Lights -  GE Color Effects](http://www.deepdarc.com/2010/11/27/hacking-christmas-lights/)
