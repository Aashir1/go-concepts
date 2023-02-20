package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int){ //same as below just a syntactic sugar of below function
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}

func takeDamageFromExplosion(player *Player, dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}

func main() {

	player := &Player{
		health: 100,
	}
	// by defining player with '&' sign, the value of palyer is 8 byte long integer
	// without '&' it will be simple struct

	fmt.Println("before explosion: ", player)
	player.takeDamageFromExplosion(45)
	fmt.Println("after explosion: ", player)

}

// why pointers?
// 1. constantly change in state
// 2. If we are working with big amount of data we need to copy that data again and again so its better to use pointers.
