package character

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"xosc.org/btwscrolls/clog"
	"xosc.org/btwscrolls/rolls"
)

type Character struct {
	name            string
	dead            bool
	id              int
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
	xpPerLevel int
}

type Rouge struct {
	Character
	xpPerLevel int
}

type Mage struct {
	Character
	xpPerLevel int
}

type CharacterInterface interface {
	toString() string
}

func (c Character) toString() string {
	return fmt.Sprintf("Name: %s Level: %d XP: %d/%d", c.name, c.level, c.exp, 1000)
}

func (c Warrior) toString() string {
	return fmt.Sprintf("Name: %s Level: %d XP: %d/%d", c.name, c.level, c.exp, c.getXpMaxPerLevel())
}

func CreateNewCharacter(name []string) CharacterInterface {

	var n string
	if len(name) == 0  || len(name[0]) == 0 {
		fmt.Print("Enter a name for your character: ")
		if _, err := fmt.Scanln(&n); err != nil {
			log.Fatalf("Error reading name: %v", err.Error())
			os.Exit(1)
		}
		fmt.Println("Creating a character named", n)
	} else {
		n = name[0]
	}

	fmt.Println("What class shall your character be?")
	fmt.Println("")
	fmt.Println("1: Warrior")
	fmt.Println("2: Rouge")
	fmt.Println("3: Mage")
	fmt.Println("")
	fmt.Println("Enter a number between 1 and 3")

	var c CharacterInterface
	switch AskForInt("Class [1-3]: ", 3) {
		case 1:
			c = Warrior{
				Character: Character {name: n},
				knacks: 0,
			}
		case 2:
			c = Rouge {
				Character: Character {name: n},
			}
		case 3:
			c = Mage {
				Character: Character {name: n},
			}
		default:
			log.Println("Invalid character class")
			return c
	}

	clog.Debug(c.toString())

	return c
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

func (w Warrior) getXpMaxPerLevel() int {
	switch w.level {
		case 1:
			return 0
		case 2:
			return 2000
		case 3:
			return 4000
		case 4:
			return 8000
		case 5:
			return 16000
		case 6:
			return 32000
		case 7:
			return 64000
		case 8:
			return 12000
		case 9:
			return 240000
		case 10:
			return 360000
		default:
			return 0
	}
}

func (w Rouge) getXpMaxPerLevel() int {
	switch w.level {
		case 1:
			return 0
		case 2:
			return 1500
		case 3:
			return 3000
		case 4:
			return 6000
		case 5:
			return 12000
		case 6:
			return 25000
		case 7:
			return 50000
		case 8:
			return 100000
		case 9:
			return 200000
		case 10:
			return 300000
		default:
			return 0
	}
}

func (w Mage) getXpMaxPerLevel() int {
	switch w.level {
		case 1:
			return 0
		case 2:
			return 2500
		case 3:
			return 5000
		case 4:
			return 10000
		case 5:
			return 20000
		case 6:
			return 40000
		case 7:
			return 80000
		case 8:
			return 150000
		case 9:
			return 300000
		case 10:
			return 400000
		default:
			return 0
	}
}


