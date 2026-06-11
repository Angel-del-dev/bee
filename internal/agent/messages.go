package agent

import (
	"math/rand"
	"time"
)

var messages = []string{
	"Buzzing online!",
	"Ready to buzz!",
	"No time to pollen around!",
	"Bee awake!",
}

func RandomMessage() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return messages[r.Intn(len(messages))]
}
