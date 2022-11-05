package main

import (
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
	defer ss[server].mu.Unlock()
	underlineWhite("File %v contents:\n", chunk)
	ss[server].mu.Lock()
	yellow("%v\n", ss[server].chs[chunk].data)
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

func (c *client) append(file int, data string) bool {
	servers := c.queryMaster(file)
	primary, secondaries := servers[0], servers[1:]
	ss[primary].appendPrimary(file, data)
	res := ss[primary].appendSecondaries(secondaries, file, data)
	return res
}
