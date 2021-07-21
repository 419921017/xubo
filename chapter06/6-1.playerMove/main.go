package main

import (
	"fmt"
	player "xubo/chapter06/6-1.playerMove/palyer"
)

func main() {
	p := player.NewPlayer(0.5)
	p.MoveTo(player.Vec2{3, 1})

	for !p.IsArrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
}
