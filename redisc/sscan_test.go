package redisc

import (
	"testing"
)

func TestSscan(t *testing.T) {
	newId := make(chan float64, 100)
	Sscan("favoritesettest", newId)
}
