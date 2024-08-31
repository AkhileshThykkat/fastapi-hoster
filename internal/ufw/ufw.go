package ufw

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConfigureUFW(port string) {
	cmds := [][]string{
		{"ufw", "allow", port + "/tcp"},
		{"ufw", "allow", port + "/udp"},
	}

	for _, cmd := range cmds {
		err := exec.Command(cmd[0], cmd[1:]...).Run()
		if err != nil {
			fmt.Printf("Error configuring UFW: %v\n", err)
		} else {
			fmt.Printf("UFW rule added: %s\n", strings.Join(cmd, " "))
		}
	}
}
