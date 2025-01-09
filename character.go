package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"xosc.org/btwscrolls/rolls"
)

type Character struct {
	Name            string
	Dead            bool
	Id              int
	Class           int
	Hp              int
	HitDie          int
	InitiativeBonus int
	Armor           int
	Exp             int
	Level           int
	BaseAttackBonus int
	FortunePoints   int
	Coppers         int
	Str             int
	Con             int
	Dex             int
	Wis             int
	Intel           int
	Cha             int
	Alignment       int
	next            *Character
}

const (
	Warrior = 1
	Rouge   = 2
	Mage    = 3
)

func (c Character) getBonus(ability int) int {
	abilityBonuses := map[int]int{
		1: -4, 2: -3, 3: -3,
		4: -2, 5: -2, 6: -1, 7: -1, 8: -1,
		9: 0, 10: 0, 11: 0, 12: 0,
		13: +1, 14: +1, 15: +1,
		16: +2, 17: +2, 18: +3, 19: +3,
	}

	if bonus, exists := abilityBonuses[ability]; exists {
		return bonus
	}
	return 0
}

func (c Character) getAlignment() string {
	if c.Alignment == 1 {
		return "Lawful"
	} else if c.Alignment == 2 {
		return "Neutral"
	} else {
		return "Chaotic"
	}
}

func (c Character) toString() string {
	return fmt.Sprintf("Name: %s Level: %d XP: %d Alignment: %s\n\nSTR: %d (%d) DEX: %d (%d) WIS: %d (%d) INT: %d (%d) CHA: %d (%d)\n\nAC: %d BAB: %d, Initiative: %d Fortune Points: %d",
		c.Name, c.Level, c.Exp, c.getAlignment(),
		c.Str, c.getBonus(c.Str), c.Dex, c.getBonus(c.Dex), c.Wis, c.getBonus(c.Wis),
		c.Intel, c.getBonus(c.Intel), c.Cha, c.getBonus(c.Cha),
		c.Armor, c.BaseAttackBonus, c.InitiativeBonus, c.FortunePoints)
}

func CreateNewCharacter(name []string) (Character, error) {
	var n string
	if len(name) == 0 || len(name[0]) == 0 {
		fmt.Print("Enter a name for your character: ")
		if _, err := fmt.Scanln(&n); err != nil {
			log.Fatalf("Error reading name: %v", err.Error())
			os.Exit(1)
		}
	} else {
		n = name[0]
	}

	if GlobalList.characterExists(n) {
		return Character{}, fmt.Errorf("A character with that name already exists")
	}
	fmt.Println("Creating a character named", n)

	fmt.Println("What class shall your character be?")
	fmt.Println("")
	fmt.Println("1: Warrior")
	fmt.Println("2: Rouge")
	fmt.Println("3: Mage")
	fmt.Println("")
	fmt.Println("Enter a number between 1 and 3")

	c := &Character{Name: n}
	c.Id = rand.Intn(9999999)
	c.Class = AskForInt("Class [1-3]: ", 1, 3)

newagain:
	abilityRolls := make([]int, 6)

	fmt.Print("Rolled 6x 4d6 and added the highest results: ")
	for i := range abilityRolls {
		temp := rolls.RollDice(4, 6)
		sort.Ints(temp)
		sum := 0
		for _, value := range temp[1:] {
			sum += value
		}
		abilityRolls[i] = sum
	}

	fmt.Println(abilityRolls)
	fmt.Printf("\nDo you want to roll again? [Yes == 1/No == 0]: ")
	if AskForInt("", 0, 1) == 1 {
		goto newagain
	}

	fmt.Printf("Now distribute the values to your abilities\n\n")

	c.Str = AskForInt("STR: ", 1, 19)
	c.Dex = AskForInt("DEX: ", 1, 19)
	c.Con = AskForInt("CON: ", 1, 19)
	c.Intel = AskForInt("INT: ", 1, 19)
	c.Wis = AskForInt("WIS: ", 1, 19)
	c.Cha = AskForInt("CHA: ", 1, 19)

	fmt.Println("Choose an alignment for your character?")
	fmt.Println("")
	fmt.Println("1: Lawful")
	fmt.Println("2: Neutral")
	fmt.Println("3: Chaotic")
	fmt.Println("")

	c.Alignment = AskForInt("Alignment [1-3]: ", 1, 3)

	fmt.Println(c.toString())

	return *c, nil
}

func (c Character) skillCheck(cmds []string) error {
	var t int
	switch(strings.ToLower(cmds[0])) {
		case "str":
			t = CurChar.Str
		case "dex":
			t = CurChar.Dex
		case "con":
			t = CurChar.Con
		case "int":
			t = CurChar.Intel
		case "wis":
			t = CurChar.Wis
		case "cha":
			t = CurChar.Cha
		default:
			return fmt.Errorf("Unknown ability, please choose from STR, DEX, CON, INT, WIS, CHA")
	}

	res := rolls.RollDice(1, 20)

	var modifier int
	// An additional bonus/penalty was provided
	if len(cmds) > 1 {
		x, err := strconv.Atoi(cmds[1])
		if err != nil {
			return fmt.Errorf("Cannot convert bonus/penalty to number")
		}
		modifier = x

		if err := validateIntRange(modifier, -20, 20); err != nil {
			return fmt.Errorf("%v", err.Error())
		}
		var sign string
		if modifier >= 0 {
			sign = "+"
		}
		fmt.Printf("<%d> vs %d%s%d", res[0], t, sign, modifier)
	} else {
		fmt.Printf("<%d> vs %d", res[0], t)
	}

	if res[0] > (t + modifier) {
		return fmt.Errorf(" failed")
	}

	return nil
}

func (w Character) getXpMaxPerLevel() int {
	if w.Class == Warrior {
		switch w.Level {
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
	} else if w.Class == Rouge {
		switch w.Level {
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
	} else { // Mage
		switch w.Level {
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
}

