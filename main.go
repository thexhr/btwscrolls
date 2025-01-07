package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"xosc.org/btwscrolls/cli"
	"xosc.org/btwscrolls/clog"

	"github.com/lmorg/readline"
)

func main() {
	rl := readline.NewInstance()
	rl.TabCompleter = cli.Tab

	log.SetFlags(0)

	clog.DebugLog = true

	setupBaseDir()

	for {
		rl.SetPrompt("> ")
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(line) == 0 {
			continue
		}

		cmd := strings.TrimSpace(line)

		// XXX add history
		cli.ExecuteCommand(cmd)
	}
}

func setupBaseDir() {
	var btwscrollsHome string
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
