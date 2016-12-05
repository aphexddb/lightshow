package main // import "github.com/aphexddb/lightshow"

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

var (
	config  *Configuration
	cfgFile = "config.json"
)

func init() {
	// load config on init
	cfg, cfgErr := LoadConfig(cfgFile)
	if cfgErr != nil {
		log.Errorf("Unable to load config file: %s", cfgErr.Error())
		os.Exit(1)
	}
	log.Infof("Loaded config from %s", cfgFile)
	config = cfg

	// show all logs
	log.SetLevel(log.DebugLevel)
}

// catch signals issued by OS
func handleSignals() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		sig := <-sigs
		if sig == syscall.SIGINT {
			log.Error("Interrupt signal (SIGINT) received, exiting")
			os.Exit(0)
		} else if sig == syscall.SIGTERM {
			log.Error("Software termination signal (SIGTERM) received, exiting")
			os.Exit(0)
		} else if sig == syscall.SIGQUIT {
			log.Error("Quit signal (SIGQUIT) received, exiting")
			os.Exit(0)
		} else if sig == syscall.SIGKILL {
			// Kill signal, exit immediately
			os.Exit(0)
		} else {
			log.Errorf("Signal (%s) received, ignoring", sig.String())
		}
		done <- true
	}()
}

func main() {
	handleSignals()
	go ServeHTTP()

	ls := &LightString{
		Pin:   8,
		Count: 10,
	}
	ls.Init()
	ls.FillString(2)
	ls.SetColor(0, 2)
	ls.SetColor(100, 2)
	ls.Render()

	fmt.Printf("ColorWhite: [%v] %v\n", ColorWhite, reflect.TypeOf(ColorWhite))
	fmt.Printf("ColorRed: [%v] %v\n", ColorRed, reflect.TypeOf(ColorRed))
	fmt.Printf("ColorGreen: [%v] %v\n", ColorGreen, reflect.TypeOf(ColorGreen))
	fmt.Printf("ColorBlue: [%v] %v\n", ColorBlue, reflect.TypeOf(ColorBlue))
	fmt.Printf("ColorHue(RBRed): [%v] %v\n", ColorHue(uint8(RBRed)), reflect.TypeOf(ColorHue(uint8(RBRed))))
	fmt.Printf("RainbowColor(RBRed): [%v] %v\n", RainbowColor(RBRed), reflect.TypeOf(RainbowColor(RBRed)))
	fmt.Printf("ChannelMax: [%v] %v\n", ChannelMax, reflect.TypeOf(ChannelMax))
	fmt.Printf("HueMax: [%v] %v\n", HueMax, reflect.TypeOf(HueMax))

	LedBLinkLoop()

}
