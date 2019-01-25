package redisc

import (
	"testing"
)

func TestDelKeys(t *testing.T) {
	Delkey("hackernewsset")
	Delkey("hackernews")
	Delkey("hackernewshash")
	Delkey("story")
}
