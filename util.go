package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func AskForInt(desc string, minimum int, maximum int) (ret int) {
again:
	fmt.Print(desc)
	var n string
	if _, err := fmt.Scanln(&n); err != nil {
		log.Printf("Error reading %s: %v", desc, err.Error())
		goto again
	}

	ret, err := strconv.Atoi(n)
	if err != nil {
		log.Printf("Invalid input")
		goto again
	}

	if err = validateIntRange(ret, minimum, maximum); err != nil {
		log.Printf("%v", err.Error())
		goto again
	}

	return ret

}

func validateIntRange(cur int, minimum int, maximum int) error {
	if cur < minimum || cur > maximum {
		return fmt.Errorf("Value out of range. It has to be between %d and %d", minimum, maximum)
	}

	return nil
}

func setPrompt(prompt string) {
	var p string
	if len(prompt) == 0 {
		p = fmt.Sprintf("> ")
	} else {
		p = fmt.Sprintf("%s > ", prompt)
	}

	rl.SetPrompt(p)
}
