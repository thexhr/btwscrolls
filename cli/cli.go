package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/readline"
)

type command struct {
	cmd      string
	callback func(cmds []string)
	desc     string
}

var commands = []command{
	{cmd: "quit", callback: cmd_quit, desc: "quit"},
	{cmd: "q", callback: cmd_quit, desc: "quit"},
	{cmd: "cd", callback: cmd_cd, desc: "Switch to something else"},
	{cmd: "ls", callback: cmd_ls, desc: "Show something else"},
}

func cmd_quit(cmds []string) {
	os.Exit(0)
}

func cmd_cd(cmds []string) {
	fmt.Println("Called cd with args:", cmds)
}

func cmd_ls(cmds []string) {
	fmt.Println("Called ls")
}

func FindCommand(cmd string) (command, error) {
	for _, v := range commands {
		if cmd == v.cmd {
			return v, nil
		}
	}
	return command{}, fmt.Errorf("command not found")
}

func ExecuteCommand(commandLine string) {
	cmdArray := strings.Split(commandLine, " ")
	if len(cmdArray) == 0 {
		return
	}

	cmd, err := FindCommand(string(cmdArray[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	f := cmd.callback
	f(cmdArray[1:])
}

func Tab(line []rune, pos int, dtx readline.DelayedTabContext) (string, []string, map[string]string, readline.TabDisplayType) {
	var suggestions []string

	for i := range commands {
		if strings.HasPrefix(commands[i].cmd, string(line)) {
			suggestions = append(suggestions, commands[i].cmd[pos:])
		}
	}

	return string(line[:pos]), suggestions, nil, readline.TabDisplayGrid
}

