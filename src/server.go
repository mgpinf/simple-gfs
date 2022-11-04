package main

type server struct {
	id  int
	chs []chunk
}

const (
	serverCount    int = 3
	primaryCount   int = 1
	secondaryCount int = 2
)

var (
	ss []server
)

// append part-2
func (s *server) appendPrimary(file int, data string) {
	underlineWhite("File %v data in primary (server %v) before appending:\n", file, s.id)
	yellow("%v\n", ss[s.id].chs[file].data)
	ss[s.id].chs[file].data += data
	green("Primary (server %v) appended data to file %v in itself\n", s.id, file)
	underlineWhite("File %v data in primary (server %v) after appending:\n", file, s.id)
	yellow("%v\n", ss[s.id].chs[file].data)
}

// part of append part-3
func (s *server) appendSecondary(sno, secondary, file int, data string) bool {
	res := result()
	underlineWhite("File %v data in secondary %v (server %v) before appending:\n", file, sno, secondary)
	yellow("%v\n", ss[secondary].chs[file].data)
	if res {
		ss[secondary].chs[file].data += data
		green("Primary (server %v) appended data to file %v in secondary %v (server %v)\n", s.id, file, sno, secondary)
		underlineWhite("File %v data in secondary %v (server %v) after appending:\n", file, sno, secondary)
		yellow("%v\n", ss[secondary].chs[file].data)
	} else {
		red("Primary (server %v) failed to append data to file %v in secondary %v (server %v)\n", s.id, file, sno, secondary)
	}
	return res
}

// append part-3 (serial)
func (s *server) appendSecondaries(secondary1, secondary2, file int, data string) bool {
	res1, res2 := s.appendSecondary(1, secondary1, file, data), s.appendSecondary(2, secondary2, file, data)
	return res1 && res2
}
