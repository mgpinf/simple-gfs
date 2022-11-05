package main

// Initialises all components

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type master struct {
	chunkServers map[int][]int
}

var (
	m master
)

func (m *master) initMaster() {
	m.chunkServers = make(map[int][]int)
	for i := 0; i < fileCount; i++ {
		m.chunkServers[i] = make([]int, replicationFactor)
	}
	for i := 0; i < fileCount; i++ {
		for j := 0; j < serverCount; j++ {
			m.chunkServers[i][j] = (i + j) % serverCount
			// add chunks and their replicas to different servers
			ss[j].chs = append(ss[j].chs, chunk{id: i, data: fmt.Sprintf("file_%v", i)})
		}
	}
}

func (m *master) initServers() {
	ss = make([]server, serverCount)
	for i := 0; i < serverCount; i++ {
		ss[i].id = i
		ss[i].chs = make([]chunk, 0)
	}
}

func (m *master) initClients() {
	cs = make([]client, clientCount)
	for i := 0; i < clientCount; i++ {
		cs[i].id = i
	}
}

func (m *master) serialExec() {
	for i := 0; i < 20; i++ {
		opHandler()
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *master) concurrentExec() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			opHandler()
			time.Sleep(500 * time.Millisecond)
		}()
	}
	wg.Wait()
}

func (m *master) obtainKeys() []int {
	keys := make([]int, 0, len(m.chunkServers))
	for k := range m.chunkServers {
		keys = append(keys, k)
	}
	return keys
}

func (m *master) generateSnapshot() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Seconds().Do(func() {
		snapshotContent := fmt.Sprintf("%v\n%v\n%v\n", replicationFactor, fileCount, serverCount)
		keys := m.obtainKeys()
		for i := range keys {
			snapshotContent += fmt.Sprintf("%v ", keys[i])
			for j := range m.chunkServers[keys[i]] {
				snapshotContent += fmt.Sprintf("%v ", m.chunkServers[i][j])
			}
			snapshotContent += "\n"
			cmd := exec.Command("touch", ".snapshot/files/"+fmt.Sprintf("%d", ss[m.chunkServers[keys[i]][0]].id))
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			fileContent := fmt.Sprintf("%v\n", ss[m.chunkServers[keys[i]][0]].chs[i].data)
			f1, err11 := os.Create(".snapshot/files/" + fmt.Sprintf("%v", ss[m.chunkServers[keys[i]][0]].chs[i].id))
			if err11 != nil {
				log.Fatal(err)
			}
			_, err12 := f1.WriteString(fileContent)
			if err12 != nil {
				log.Fatal(err12)
			}
			f1.Close()
		}
		f, err := os.Create(".snapshot/SNAPSHOT")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err2 := f.WriteString(snapshotContent)
		if err2 != nil {
			log.Fatal(err2)
		}
		magenta("Snapshot generated\n")
	})
	s.StartAsync()
}

func (m *master) createSnapshotDirectoryIfNotExists() {
	cmd := exec.Command("mkdir", "-p", ".snapshot/files")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (m *master) initAll() {
	m.initServers()
	m.initMaster()
	m.initClients()
	m.createSnapshotDirectoryIfNotExists()
	m.generateSnapshot()
	if serial {
		m.serialExec()
		return
	}
	m.concurrentExec()
}
