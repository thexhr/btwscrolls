package main

import (
	"fmt"
	"github.com/lmorg/readline"
	"strings"
)

type command struct {
	cmd      string
	callback func(cmds []string)
	desc     string
}

var commands = []command{
	{cmd: "cd", callback: cmd_cd, desc: "Switch to a character"},
	{cmd: "cds", callback: cmd_cds, desc: "Switch to a character and show details"},
	{cmd: "ls", callback: cmd_ls, desc: "Show all characters"},
	{cmd: "create", callback: cmd_create, desc: "Create a new character"},
	{cmd: "skillcheck", callback: cmd_skillCheck, desc: "Roll a skill check"},
}

// cd loads the chracter given by cmds[0] or unloads the active character.
func cmd_cd(cmds []string) {
	/* We got no argument and there is no character loaded */
	if len(cmds) == 0 && CurChar == nil {
		fmt.Println("Provide the name of a character as argument")
		if !GlobalList.listEmpty() {
			fmt.Println("")
			fmt.Println("You currently have the following characters:")
			GlobalList.showAllCharacters()
		}
		return
		/* We got no argument and there is a character loaded */
	} else if len(cmds) == 0 && CurChar != nil {
		CurChar.LastActive = 0
		CurChar = nil
		rl.SetPrompt("> ")
		return
		/* We got an argument and there is no/a character loaded */
	} else if len(cmds) > 0 {
		temp, err := GlobalList.returnChar(cmds[0])
		if err != nil {
			fmt.Printf("%v", err.Error())
			return
		}
		if CurChar != nil {
			CurChar.LastActive = 0
		}

		CurChar = temp
		setPrompt(cmds[0])
		CurChar.LastActive = 1
	}
}

// cds is cmd_cd() plus showing the Character struct as pretty print.
func cmd_cds(cmds []string) {
	cmd_cd(cmds)
	if CurChar != nil {
		fmt.Println(CurChar.toString())
	}
}

// ls lists the names of all esisting characters or nothing if none exist.
func cmd_ls(_ []string) {
	if err := GlobalList.showAllCharacters(); err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}

func cmd_create(cmds []string) {
	c, err := CreateNewCharacter(cmds)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	GlobalList.addCharToList(c)
	CurChar = &c

	p := fmt.Sprintf("%s > ", c.Name)
	rl.SetPrompt(p)
}

func cmd_skillCheck(cmds []string) {
	if len(cmds) < 1 {
		fmt.Println("You need to specify the ability and an optional bonus malus")
		fmt.Println("")
		fmt.Println("Example: skillcheck str 2")
		return
	}

	if CurChar == nil {
		return
	}

	if str, err := CurChar.skillCheck(cmds); err != nil {
		fmt.Printf("%v\n", err.Error())
	} else {
		fmt.Printf("%v\n", str)
	}
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
