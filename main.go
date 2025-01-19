package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"xosc.org/btwscrolls/clog"

	"github.com/lmorg/readline"
)

type CharList struct {
	head *Character
	next *Character
	len  int
}

var GlobalList CharList
var CurChar *Character
var Prompt string
var rl *readline.Instance

func main() {
	rl = readline.NewInstance()
	rl.TabCompleter = Tab

	log.SetFlags(0)

	clog.DebugLog = true

	SetupBaseDir()

	GlobalList = CharList{}

	CurChar = nil
	if err := GlobalList.loadCharacters(); err != nil {
		log.Fatalf("%v", err.Error())
	}

	setPrompt("")
	GlobalList.loadLastActiveChar()

	var history []string
	history, err := loadHistoryFromDisk()
	if err != nil {
		history = make([]string, 1)
	}

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(line) == 0 {
			continue
		}
		if line == "q" || line == "quit" {
			break
		}

		cmd := strings.TrimSpace(line)

		if len(history) > 0 {
			history = append(history, cmd)
		}
		ExecuteCommand(cmd)
	}

	fmt.Println(rl.History.Dump())

	GlobalList.saveCharacter()
	writeHistoryToDisk(history)
}

func loadHistoryFromDisk() ([]string, error) {
	var hist []string
	f, err := os.Open(btwscrollsHome + "/history")
	if err != nil {
		return hist, fmt.Errorf("No history file present")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		hist = append(hist, line)
	}

	return hist, nil
}

func writeHistoryToDisk(arr []string) error {
	f, err := os.OpenFile(btwscrollsHome + "/history", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range arr {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (l *CharList) returnChar(name string) (*Character, error) {
	ptr := l.head
	if ptr == nil {
		return &Character{}, fmt.Errorf("No characters found")
	}

	for ptr != nil {
		if ptr.Name == name {
			return ptr, nil
		}
		ptr = ptr.next
	}

	return &Character{}, fmt.Errorf("Character not found")
}

func (l *CharList) addCharToList(c Character) error {
	if l.head == nil {
		l.head = &c
		l.len++
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = &c
		l.len++
	}

	return nil
}

func (l *CharList) showAllCharacters() error {
	ptr := l.head
	if ptr == nil {
		return fmt.Errorf("No characters yet created")
	}

	for ptr != nil {
		fmt.Println(ptr.Name)
		ptr = ptr.next
	}

	return nil
}

func (l *CharList) characterExists(name string) bool {
	ptr := l.head
	if ptr == nil {
		return false
	}

	for ptr != nil {
		if ptr.Name == name {
			return true
		}
		ptr = ptr.next
	}

	return false
}

func (l *CharList) toSlice() []Character {
	var characters []Character
	ptr := l.head
	for ptr != nil {
		characters = append(characters, *ptr)
		ptr = ptr.next
	}
	return characters
}

func (l *CharList) saveCharacter() error {
	x := l.toSlice()
	jsonData, err := json.Marshal(x)
	if err != nil {
		return fmt.Errorf("Error marshaling to JSON: %v", err)
	}

	err = os.WriteFile(GetHomeDir()+"/characters.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	clog.Debug("Character saved to character.json")

	return nil
}

func (l *CharList) loadCharacters() error {
	jsonData, err := os.ReadFile(GetHomeDir() + "/characters.json")
	if err != nil {
		// Although this is an error, we ignore it here since there will be no
		// JSON file at the first startup
		return nil
	}

	var characters []Character
	err = json.Unmarshal(jsonData, &characters)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	if len(characters) > 0 {
		for i := 0; i < len(characters); i++ {
			l.addCharToList(characters[i])
		}
	}

	return nil
}

func (l *CharList) listEmpty() bool {
	if l.len > 0 {
		return false
	}

	return true
}

func (l *CharList) loadLastActiveChar() {
	if l.listEmpty() {
		return
	}

	ptr := l.head
	for ptr != nil {
		if ptr.LastActive == 1 {
			CurChar = ptr
			setPrompt(CurChar.Name)
			return
		}
		ptr = ptr.next
	}

}
