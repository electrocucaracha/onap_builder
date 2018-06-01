package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Dispatcher struct {
	Id       int
	CmdQueue chan []*Cmd
	ErrQueue chan error
	quit     chan bool
}

func NewDispatcher(id int, cmdQueue chan []*Cmd) *Dispatcher {
	d := &Dispatcher{
		Id:       id,
		CmdQueue: cmdQueue,
		ErrQueue: make(chan error),
		quit:     make(chan bool),
	}
	go d.run()
	return d
}

func (dispatcher *Dispatcher) run() {
	log.Println("starting dispatcher...")
	for {
		select {
		case <-dispatcher.quit:
			log.Printf("stopping dispatcher id: %v...", dispatcher.Id)
			dispatcher.quit <- true
			return
		case cmds := <-dispatcher.CmdQueue:
			if cmds == nil {
				return
			}
			log.Printf("running dispatcher id: %v - %v commmand(s)...", dispatcher.Id, len(cmds))
			for _, cmd := range cmds {
				command := exec.Command(cmd.Name, cmd.Args...)
				command.Dir = cmd.Dir

				log.Printf("executing: %s\n", strings.Join(cmd.Args, " "))
				stdout, err := command.StdoutPipe()
				if err != nil {
					log.Fatal(err)
				}
				if err := command.Start(); err != nil {
					os.Stderr.WriteString(fmt.Sprintf("error: %s\n", err.Error()))
					dispatcher.ErrQueue <- err
				}
				scanner := bufio.NewScanner(stdout)
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					m := scanner.Text()
					fmt.Println(m)
				}
				command.Wait()
			}
		}
	}
}

func (dispatcher *Dispatcher) Stop() {
	dispatcher.quit <- true
	<-dispatcher.quit
	log.Printf("dispatcher id: %v stopped", dispatcher.Id)
}
