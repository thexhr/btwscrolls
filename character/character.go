package character

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"xosc.org/btwscrolls/clog"
)

type Character struct {
	name            string
	id              int
	dead            bool
	hp              int
	hitDie          int
	initiativeBonus int
	armor           int
	exp             int
	level           int
	baseAttackBonus int
	fortunePoints   int
	coppers         int
	str             int
	strBonus        int
	dex             int
	dexBonus        int
	wis             int
	wisBonus        int
	intel           int
	intelBonus      int
	cha             int
	chaBonus        int
}

type Warrior struct {
	Character
	knacks int
}

type Rouge struct {
	Character
}

type Mage struct {
	Character
}

func CreateNewCharacter(name []string) {

	if len(name) == 0  || len(name[0]) == 0 {
		fmt.Print("Enter a name for your character: ")
		var n string
		if _, err := fmt.Scanln(&n); err != nil {
			log.Fatalf("Error reading name: %v", err.Error())
			os.Exit(1)
		}
		fmt.Println("Creating a character named", n)
	}

	fmt.Println("What class shall your character be?")
	fmt.Println("")
	fmt.Println("1: Warrior")
	fmt.Println("2: Rouge")
	fmt.Println("3: Mage")
	fmt.Println("")
	fmt.Println("Enter a number between 1 and 3")

	switch AskForInt("Class [1-3]: ", 3) {
		case 1:
			var c Warrior
		case 2:
			var c Rouge
		case 3:
			var c Mage
		default:
			log.Println("Invalid character class")
			return
	}

}

func AskForInt(desc string, maximum int) (ret int) {
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

	if err = validateIntRange(ret, 1, maximum); err != nil {
		log.Printf("%v", err.Error())
		goto again
	}

	clog.Debug("Read %d", ret)

	return ret

}

func validateIntRange(cur int, minimum int, maximum int) error {
	if cur < minimum || cur > maximum {
		return fmt.Errorf("Value out of range. It has to be between %d and %d", minimum, maximum)
	}

	return nil
}
