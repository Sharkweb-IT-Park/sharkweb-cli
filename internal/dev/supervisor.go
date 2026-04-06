package dev

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Supervisor struct {
	Processes []*Process
}

func (s *Supervisor) Add(p *Process) {
	s.Processes = append(s.Processes, p)
}

func (s *Supervisor) StartAll() {

	// Start all processes
	for _, p := range s.Processes {
		go p.Start()
	}

	// Listen for Ctrl + C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan

	fmt.Println("\n🛑 Shutting down all services...")

	for _, p := range s.Processes {
		if p.Cmd != nil && p.Cmd.Process != nil {
			p.Cmd.Process.Kill()
		}
	}

	os.Exit(0)
}
