package main

import "fmt"

type server struct {
	id  int
	chs []chunk
}

const (
	serverCount int = 3
)

var (
	ss []server
)

// append part-2
func (s *server) appendPrimary(primary, file int, data string) {
	underlineWhite("Primary %v data before appending:\n", primary)
	yellow(ss[primary].chs[file].data)
	fmt.Println()
	ss[primary].chs[file].data += data
	green("Primary %v appended data to itself\n", primary)
	underlineWhite("Primary %v data after appending:\n", primary)
	yellow(ss[primary].chs[file].data)
	fmt.Println()
}

// part of append part-3
func (s *server) appendSecondary(secondary, file int, data string) bool {
	res := result()
	underlineWhite("Secondary %v data before appending:\n", secondary)
	yellow(ss[secondary].chs[file].data)
	fmt.Println()
	if res {
		ss[secondary].chs[file].data += data
		green("Primary %v appended data to secondary %v\n", s.id, secondary)
		underlineWhite("Secondary %v data after appending:\n", secondary)
		yellow(ss[secondary].chs[file].data)
		fmt.Println()
	} else {
		red("Primary %v failed to append data to secondary %v\n", s.id, secondary)
	}
	return res
}

// append part-3
func (s *server) appendSecondaries(secondary1, secondary2, file int, data string) bool {
	res1, res2 := s.appendSecondary(secondary1, file, data), s.appendSecondary(secondary2, file, data)
	return res1 && res2
}
