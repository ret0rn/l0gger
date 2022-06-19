package main

import (
	"fmt"

	"github.com/ret0rn/l0gger"
)

func main() {
	var log, err = l0gger.New(nil, l0gger.JsonFormater, l0gger.DebugLvl) // log to console
	// or
	// var log, err = l0gger.New(os.Stdout, l0gger.JsonFormater, l0gger.DebugLvl)

	if err != nil {
		fmt.Printf("Error: [%s]\n", err.Error())
		panic(err)
	}

	log.Debug("Debug test")
	log.Info("Info test")
	log.Warn("Warning test")
	log.Error("Error test")
	log.Fatal("Fatal test")
}
