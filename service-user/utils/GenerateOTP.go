package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateOTP generates a 6-digit OTP code
func GenerateOTP() (string, error) {

	// Seed the random number generator with the current timestamp
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number
	otp := rand.Intn(1000000)

	// Ensure the OTP is exactly 6 digits long
	otpString := strconv.Itoa(otp)

	if len(otpString) < 6 {
		// Pad the OTP with leading zeros if necessary
		otpString = fmt.Sprintf("%06d", otp)
	}

	return otpString, nil
}
