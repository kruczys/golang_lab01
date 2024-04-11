package main

import (
    "fmt"
    "math/rand"
)

func main() {
    game_rounds := 1000000
    player_wins_staying := 0
    player_wins_switching := 0
    // is_staying := false
    is_staying := true
   
    for i := 0; i < game_rounds; i++ {
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

    if is_staying {
        fmt.Println("Player tactic was to stay")
        fmt.Println("wins: ", player_wins_staying)
        fmt.Println("loses: ", game_rounds - player_wins_staying)
    } else {
        fmt.Println("Player tactic was to switch")
        fmt.Println("wins: ", player_wins_switching)
        fmt.Println("loses: ", game_rounds - player_wins_switching)
    }
}
