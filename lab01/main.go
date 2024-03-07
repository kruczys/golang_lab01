package main

import (
    "fmt"
    "math/rand"
)

func main() {
    game_rounds := 100
    player_wins_staying := 0
    player_wins_switching := 0
   
    for i := 0; i <= game_rounds; i++ {
        var single_round [3]int
        winning_box := rand.Intn(3)
        player_choice := rand.Intn(3)
        single_round[winning_box] = 1

        var showed_box int
        for showed_box = 0; showed_box == winning_box || showed_box == player_choice; {
            showed_box = rand.Intn(3)
        }

        if single_round[player_choice] == 1 {
            player_wins_staying += 1
        } else {
            player_wins_switching += 1
        }
    }

    fmt.Println("wins by staying", player_wins_staying)
    fmt.Println("wins by switching", player_wins_switching)
}
