package main

import (
	"fmt"
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
)

func randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func generateRandomString(length int) string {
	randomize()
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func opHandler() {
	randomize()
	op := rand.Intn(opCount)
	randomize()
	i := rand.Intn(clientCount)
	c := cs[i]
	cyan("Client %v ", c.id)
	file := c.selectFile()
	if op == 1 {
		cyan("reading file %v", file)
		c.read(file)
	} else {
		cyan("appending to file %v", file)
		randomString := generateRandomString(4)
		c.append(file, randomString)
	}
}

// used to simulate success or failure
func result() bool {
	randomize()
	chooser, _ := wr.NewChooser(
		wr.Choice{Item: true, Weight: 8},
		wr.Choice{Item: false, Weight: 2},
	)
	res := chooser.Pick().(bool)
	return res
}
