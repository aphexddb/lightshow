package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

// FoobarPWM tests things
func FoobarPWM(pinNumber int) {

	//
	// TODO: Try this! https://github.com/hugozhu/rpi
	//

	flag.Parse()

	embd.InitGPIO()
	defer embd.CloseGPIO()

	pwm, err := embd.NewPWMPin(fmt.Sprintf("%v", pinNumber))
	if err != nil {
		log.Error("Unable to create PWM pin %v: %v", pinNumber, err)
		os.Exit(1)
	}
	defer pwm.Close()

	// SetPeriod sets the period of a pwm pin.
	SetPeriod(1000000000) // 1 sec in nanoseconds

	// SetDuty sets the duty of a pwm pin.
	SetDuty(1000) // nanoseconds

	// SetMicroseconds sends a command to the PWM driver to generate a us wide pulse.
	us := 10000
	pwm.SetMicroseconds(us)

}
