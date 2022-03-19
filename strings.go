package testament

import (
	"fmt"
	"math/rand"
	"strings"
)

// RandomString returns a randomly generates string with the length of count.
func RandomString(count int) string {
	const runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, count)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// RandomLowerString returns a randomly generates lower-cased string with the
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

// RandomStringSlice return a random string slice with maximum length of max.
func RandomStringSlice(max int) []string {
	ret := make([]string, rand.Intn(max))
	for i := range ret {
		ret[i] = RandomString(20)
	}
	return ret
}

// RandomEmailAddress returns a random email address.
func RandomEmailAddress() string {
	return fmt.Sprintf("%s@%s.%s", RandomLowerString(10), RandomLowerString(10), RandomLowerString(3))
}
