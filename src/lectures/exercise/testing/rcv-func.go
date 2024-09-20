//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
	Name      string
}

func (player *Player) modifyHealth(num int) {
	health := player.Health + num
	if health <= 0 {
		player.Health = 0
		player.Energy = 0 // 命都没了肯定也没能量了哇
	} else if health >= player.MaxHealth {
		player.Health = player.MaxHealth
	} else {
		player.Health = health
	}
	fmt.Println("The player stat(modified Health):", player)
}

func (player *Player) modifyEnergy(num int) {
	energy := player.Energy + num
	if energy <= 0 {
		player.Energy = 0
	} else if energy >= player.MaxEnergy {
		player.Energy = player.MaxEnergy
	} else {
		player.Energy = energy
	}
	fmt.Println("The player stat(modifed Energy):", player)
}

func main() {
	player := Player{
		Health:    100,
		MaxHealth: 100,
		Energy:    0,
		MaxEnergy: 100,
		Name:      "Gonzo",
	}
	fmt.Println("Initial Player:", player)
	// 这个receiver function其实就是class里的method吧
	player.modifyHealth(-10)
	player.modifyEnergy(50)

	// player.modifyHealth(-100)
	player.modifyHealth(80)
	player.modifyEnergy(80)

	var test1 uint = 10
	var test2 uint = 15
	result := test1 - test2
	fmt.Println(result) // super big!18446744073709551611

}
