package main

import (
	"log"
	"os"
)

var btwscrollsHome string

func SetupBaseDir() {
	if xdgh := os.Getenv("XDG_CONFIG_HOME"); len(xdgh) > 0 {
		btwscrollsHome = xdgh + "/btwscrolls"
	} else if home := os.Getenv("HOME"); len(home) > 0 {
		btwscrollsHome = home + "/.config/btwscrolls"
	} else {
		log.Fatal("Neither $XDG_CONFIG_HOME nor $HOME set")
	}

	if _, err := os.Stat(btwscrollsHome); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(btwscrollsHome, 0755); err != nil {
				log.Fatalf("Cannot create home dir: %v", err.Error())
			}
		}
	}
}

func GetHomeDir() string {
	return btwscrollsHome
}
