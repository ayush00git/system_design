package helpers

import (
	"time"
	"math/rand"	
)

func GenerateCode(n int) string {
	
	charset := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

	// making a new seed source (so that the server always generate a random string)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	b := make([]byte, n)	// a byte slice of n size ["0", "0", "0"...n]

	for i := range b {
		randomInt := r.Intn(len(charset))		// any number b/w 0 and len(charset)
		b[i] = charset[randomInt]
	}

	return string(b)
}
