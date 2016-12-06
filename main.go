package main // import "github.com/aphexddb/lightshow"

import (
	"os"
	"os/signal"
	"syscall"
	"time"

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
	// go ServeHTTP()

	ls := NewG35String(12, 36)
	for {
		ls.FillColor(0, 36, MaxIntensity, ColorGreen)
		time.Sleep(1000 * time.Millisecond)
		ls.FillColor(0, 36, MaxIntensity, ColorRed)
		time.Sleep(1000 * time.Millisecond)
	}
}
