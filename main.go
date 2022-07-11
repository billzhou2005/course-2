package main

import (
	"fmt"
	"time"
	"zhouyan/course2/game"
)

func init() {

}

func main() {
	var player game.Player
	player.Name = "Jerry"
	player.Age = 17
	player.UserID = "13200001011"
	player.Weapon = append(player.Weapon, "knife", "pencil")
	player.Score = 0

	fmt.Println(player)
	count := 0
	timer := time.NewTimer(1 * time.Second)
	for {
		timer.Reset(1 * time.Second) // the timer is reused here
		select {
		case <-timer.C: //如果时间到了
			player.Score += 5
			fmt.Println("every 1 seconds", player)
			count++
		}
	}
}
