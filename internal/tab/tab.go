package tab

import (
	"fmt"
	"os/exec"
)

func OpenTab(n int) {
	app := "gnome-terminal"

	args := [][]string{
		{"--tab", "--", "tmux"},
		{"--tab", "--working-directory", "/home/akshbansal/d/temp"},
		{"--tab", "--working-directory", "/home/akshbansal/d/temp", "--", "tmux"},
	}

	cmd := exec.Command(app, args[n]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
