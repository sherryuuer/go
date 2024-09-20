// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func RightValue(num int) bool {
	player := Player{
		Health:    100,
		MaxHealth: 100,
		Energy:    0,
		MaxEnergy: 100,
		Name:      "Gonzo",
	}
	player.modifyHealth(num)
	return player.Health >= 0 && player.Health <= player.MaxHealth
}

func TestModifyHealth(t *testing.T) {
	table := []struct {
		num  int
		want bool
	}{
		{-100, true},
		{-25, true},
		{0, true},
		{25, true},
		{50, true},
		{100, true},
	}

	for _, data := range table {
		result := RightValue(data.num)
		if result != data.want {
			t.Errorf("Result is wrong with num %v", data.num)
		}
	}

}

// A good way
func newPlay() Player {
	return Player{
		Health:    100,
		MaxHealth: 100,
		Energy:    0,
		MaxEnergy: 100,
		Name:      "Gonzo",
	}
}

func TestHealth(t *testing.T) {
	player := newPlay()
	player.modifyHealth(999)
	if player.Health > player.MaxHealth {
		t.Fatalf("Health (%v) went over limit (%v)", player.Health, player.MaxHealth)
	}
	player.modifyEnergy(player.MaxHealth + 1)
	if player.Health < 0 {
		t.Fatalf("Health (%v) can not go blow 0", player.Health)
	}
}
