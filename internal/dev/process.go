package dev

import (
	"context"
	"fmt"
	"os"
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

		// 🔥 Force colors (THIS is key)
		cmd.Env = append(os.Environ(),
			"FORCE_COLOR=1",
			"TERM=xterm-256color",
		)

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

		// 🔥 Stream raw (preserves ANSI)
		go StreamOutput(p.Name, stdout)
		go StreamOutput(p.Name, stderr)

		err = cmd.Wait()

		if ctx.Err() == context.Canceled {
			fmt.Printf("🛑 %s stopped\n", p.Name)
			return
		}

		if err != nil {
			fmt.Printf("⚠️ %s crashed. Restarting...\n", p.Name)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("⚠️ %s exited normally\n", p.Name)
		return
	}
}
