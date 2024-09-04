package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAccountNumber() string {
	// Create a new random number generator with a seed based on the current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random number for the last 7 digits
	lastSevenDigits := r.Intn(10000000) // 7-digit number from 0000000 to 9999999

	// Combine the fixed "077" prefix with the random 7-digit number
	return fmt.Sprintf("001%07d", lastSevenDigits)
}
