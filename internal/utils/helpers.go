package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomName() string {

	rand.Seed(time.Now().UnixNano())

	firstNames := []string{
		"Emma", "Liam", "Olivia", "Noah", "Ava",
		"William", "Sophia", "James", "Isabella", "Oliver",
		"Charlotte", "Elijah", "Amelia", "Benjamin", "Mia",
		"Lucas", "Harper", "Henry", "Evelyn", "Alexander",
	}

	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
		"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
		"Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris",
	}

	return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
}

func GenerateFDSTrxID() string {
	// Format current time to yyyyMMddHHmmss
	timestamp := time.Now().Format("20060102150405")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a 12-digit random number as a string
	randomDigits := ""
	for i := 0; i < 12; i++ {
		randomDigits += fmt.Sprintf("%d", rand.Intn(10))
	}

	return timestamp + randomDigits
}
