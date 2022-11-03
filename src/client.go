package main

import (
	"fmt"
	"math/rand"
)

type client struct {
	id int
}

const (
	clientCount int = 3
)

var (
	cs []client
)

// read / append part-1
func (c *client) queryMaster(file int) []int {
	return m.chunkServers[file]
}

// read part-2
func (c *client) queryServer(chunk int, servers []int) {
	randomize()
	i := rand.Intn(replicationFactor)
	server := servers[i]
	underlineWhite("\nFile contents:\n")
	yellow(ss[server].chs[chunk].data)
	fmt.Println()
}

func (c *client) read(file int) {
	servers := c.queryMaster(file)
	c.queryServer(file, servers)
}

// select random file
func (c *client) selectFile() int {
	randomize()
	i := rand.Intn(fileCount)
	return files[i]
}

func (c *client) append(file int, data string) {
	servers := c.queryMaster(file)
	primary, secondary1, secondary2 := servers[0], servers[1], servers[2]
	fmt.Println()
	ss[primary].appendPrimary(primary, file, data)
	res := ss[primary].appendSecondaries(secondary1, secondary2, file, data)
	for !res {
		boldRed("Failure! Attempting again\n")
		ss[primary].appendPrimary(primary, file, data)
		res = ss[primary].appendSecondaries(secondary1, secondary2, file, data)
	}
	boldGreen("Success!\n")
}
