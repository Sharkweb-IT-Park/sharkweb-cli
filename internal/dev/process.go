package dev

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Process struct {
	Name    string
	Command string
	Args    []string
	Dir     string

	Cmd    *exec.Cmd
	cancel context.CancelFunc
}

func (p *Process) Start() {

	for {
		ctx, cancel := context.WithCancel(context.Background())
		p.cancel = cancel

		cmd := exec.CommandContext(ctx, p.Command, p.Args...)
		cmd.Dir = p.Dir

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Printf("❌ %s stdout error: %v\n", p.Name, err)
			return
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Printf("❌ %s stderr error: %v\n", p.Name, err)
			return
		}

		p.Cmd = cmd

		err = cmd.Start()
		if err != nil {
			fmt.Printf("❌ Failed to start %s: %v\n", p.Name, err)
			return
		}

		fmt.Println("🚀 Started:", p.Name)

		// 🔹 Stream logs
		go StreamOutput(p.Name, stdout)
		go StreamOutput(p.Name, stderr)

		// 🔹 Wait for exit
		err = cmd.Wait()

		// 🔴 If stopped manually → exit loop
		if ctx.Err() == context.Canceled {
			fmt.Printf("🛑 %s stopped\n", p.Name)
			return
		}

		// 🔁 Restart on crash
		if err != nil {
			fmt.Printf("⚠️ %s crashed. Restarting...\n", p.Name)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("⚠️ %s exited normally\n", p.Name)
		return
	}
}

func (p *Process) Stop() {
	if p.cancel != nil {
		p.cancel()
	}
}
