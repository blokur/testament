package testament

import (
	"math/rand"
	"strings"
)

// RandomString returns a randomly generated string with the length of count.
func RandomString(count int) string {
	const runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// RandomLowerString returns a randomly generated lowercased string with the
// length of count.
func RandomLowerString(count int) string {
	return strings.ToLower(RandomString(count))
}


// StringSlice return a string slice with the provided length.
func StringSlice(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		ret[i] = RandomString(20)
	}
	return ret
}

// RandomStringSlice return a random string slice with maximum length max.
func RandomStringSlice(max int) []string {
	ret := make([]string, rand.Intn(max))
	for i := range ret {
		ret[i] = RandomString(20)
	}
	return ret
}
