package main

import (
	"log"
	"time"
)

// Accepts Adafruit.com, Vilros.com, Pishop.us, Sparkfun.com
// Vilros links cannot have www. in front of them

var pi8links []string = []string{"https://www.adafruit.com/product/4564", "https://www.sparkfun.com/products/16811", "https://vilros.com/products/raspberry-pi-4-model-b-8gb-ram", "https://www.pishop.us/product/raspberry-pi-4-model-b-8gb"}
var pi4links []string = []string{"https://www.adafruit.com/product/4296", "https://vilros.com/products/raspberry-pi-4-4gb-ram", "https://www.pishop.us/product/raspberry-pi-4-model-b-4gb", "https://www.sparkfun.com/products/15447"}

func main() {
	log.Println("                                                                                                             ")
	log.Println("     _/_/_/            _/_/_/    _/  _/_/_/_/_/                                _/                            ")
	log.Println("  _/          _/_/    _/    _/          _/      _/  _/_/    _/_/_/    _/_/_/  _/  _/      _/_/    _/  _/_/   ")
	log.Println(" _/  _/_/  _/    _/  _/_/_/    _/      _/      _/_/      _/    _/  _/        _/_/      _/_/_/_/  _/_/        ")
	log.Println("_/    _/  _/    _/  _/        _/      _/      _/        _/    _/  _/        _/  _/    _/        _/           ")
	log.Println(" _/_/_/    _/_/    _/        _/      _/      _/          _/_/_/    _/_/_/  _/    _/    _/_/_/  _/            ")
	log.Println("                                                                                                             ")
	log.Println("GoPiTracker v1.2.1 - Written by Garrett")
	log.Println("Tracking Raspberry Pi 4 - 4gb & 8gb stock. Running every 10 minutes.")
	log.Println("---------------------------------------")
	// schedule the checkAllStock function to run every 10 minutes
	scheduleEvery(600000, checkAllStock)
}
func scheduleEvery(duration int, f func()) {
	f()
	time.Sleep(time.Duration(duration) * time.Millisecond)
	scheduleEvery(duration, f)
}
