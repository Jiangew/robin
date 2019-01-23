package main

import (
	"fmt"	
)

func main()  {
	fmt.Println("------ Composite Types ------")

	// Arrays
	// var days []string
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days[0])
	fmt.Println(days[5])
	
	// Slices
	weekdays:=days[0:5]
	fmt.Println(weekdays)

	// Maps
	subscribers:=map[string]int{
		"TutorialEdge":     2240,
  		"MKBHD":            6580350,
  		"Fun Fun Function": 171220,
	}
	fmt.Println(subscribers["MKBHD"])

	// Structs
	type Persion struct {
		name string
		age int
	}
	// var ganluo Persion
	jiangew:=Persion{name:"Jiangew", age:27}
	jiangew.age=18
	fmt.Println(jiangew)

	// Nested Structs
	type Team struct {
		name string
		committers [2]Persion
	}

	var newTeam Team
	fmt.Println(newTeam)

	committers:=[...]Persion{Persion{name:"Jiangew"}, Persion{name:"Daxu"}}
	monkeyBang:=Team{name:"MonkeyBang",committers:committers}
	fmt.Println(monkeyBang)

	fmt.Println("------ Composite Types End ------")
}
