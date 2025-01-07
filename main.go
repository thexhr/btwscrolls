package main

import (
	"fmt"
	"log"
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
