package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestRandomInt(t *testing.T) {
	dataset := [][]int64{
		{100, 200},
		{300, 400},
		{500, 900},
	}

	for _, data := range dataset {
		randomNumber := RandomInt(data[0], data[1])
		if randomNumber > data[1] || randomNumber < data[0] {
			s := fmt.Sprintf("Expected number between %d to %d, but randomNumber : %d", 
						data[0], data[1], randomNumber)
			t.Error(s)
		}
	}
}

func TestRandomString(t *testing.T) {
	Len := 10
	randomString := RandomString(Len)
	require.Equal(t, Len, len(randomString))
}

func TestRandomOwner(t *testing.T) {
	ownerName := RandomOwner()
	require.Equal(t, 10, len(ownerName))
}

func TestRandomMoney(t *testing.T) {
	randonMoney := RandomMoney()

	if randonMoney < 200 || randonMoney > 2000 {
		t.Errorf("Expected number between 200 to 2000, but randomNumber : %d", randonMoney)
	}
}

// RandomCurrency
func TestRandomCurrency(t *testing.T) {
	currencies := []string{
		"USD", "EUR", "CAD",
	}
	randomCurrency := RandomCurrency()
	for _, c := range currencies {
		if randomCurrency == c {return}
	}
	t.Error("see currencies slices")
}
