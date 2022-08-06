package tab

import (
	"log"
	"os"
	"os/exec"
)

func OpenTab(n int) {
	app := "gnome-terminal"
	homeDir, _ := os.UserHomeDir()
	args := [][]string{
		{"--tab", "--", "tmux"},
		{"--tab", "--working-directory", homeDir + "/d/code/web/github/gymkhana"},
		{"--tab", "--working-directory", homeDir + "/d/temp", "--", "tmux"},
	}

	cmd := exec.Command(app, args[n]...)
	_, err := cmd.Output()

	if err != nil {
		log.Fatal("invalid command")
		return
	}
}
