package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("print (1) to display menu")

	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)
	fmt.Println(choice)

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{
			name: "Coffe",
			prices: map[string]float64{
				"regular": 12.5,
				"large":   20,
			},
		},
		{
			name: "tea",
			prices: map[string]float64{
				"single": 10,
				"double": 15,
				"triple": 20,
			},
		},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for key, value := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", key, value)
		}
	}
}
