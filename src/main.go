package main

//constants used:
//	fileCount - Number of files system can support
//	replicationFactor - 3 including the primary (1 + 2)
//	opCount - To choose a random operation between read and write/append to facilitate automated selection
const (
	fileCount         int = 3
	replicationFactor int = 3
	opCount           int = 2
)

var (
	files []int = []int{0, 1, 2}
)

func main() {
	m.initAll()
}
