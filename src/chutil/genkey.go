package chutil

import (
	"math/rand"
	"time"
)

const keyLength = 8

var syms = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// init sets up the random seed for GenKey
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenKey generates a model id; ideally this should check that the key
// is not already taken, but we're going to assume no collisions for simplicity
// based on https://stackoverflow.com/a/22892986/2397327
func GenKey() string {
	key := make([]rune, keyLength)
	for i := range key {
		key[i] = syms[rand.Intn(len(syms))]
	}
	return string(key)
}
