# lightshow

RGB LED light effects and things for the Raspberry Pi.

<img src="images/pi3_gpio.png" />

## Links

* [Different Styles of Pixels](http://www.doityourselfchristmas.com/wiki/index.php?title=Different_Styles_of_Pixels)
* [Pixel Wiring Colors](http://www.doityourselfchristmas.com/wiki/index.php?title=Pixel_Wiring_Colors)
* [GE Color Effects Bulbs](http://www.geholidaylighting.com/color-effects/us/items/72425)
  * [Hacking the remote](https://lukecyca.com/2013/g35-rf-remote.html)
* [Arduino GECE library](https://github.com/sowbug/G35Arduino)
  * [Related forum thread on using it](http://doityourselfchristmas.com/forums/showthread.php?37502-E1-31-Arduino-GECE-controller)
* [Another blog post on hacking the GE lights](http://culverlabs.co/post/38529142159/christmaslights)

## Build for Pi

```
export GOOS=linux
export GOARCH=arm
go build *.go
```
