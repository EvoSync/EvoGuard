package functions

import (
	"math/rand"
	"strings"
	"time"
)

// ExportLast will just ignore every value but the last one
func ExportLast(i int, e error) error {
	return e
}

// RemovePortFromIP will take the string and manipulate the string to remove the port from it
func RemovePortFromIP(addr string) string {
	return strings.Join(strings.Split(addr, ":")[:strings.Count(addr, ":")], ":")
}

func Shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}
