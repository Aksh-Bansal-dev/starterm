package runner

import (
	"log"
	"os/exec"
)

func Run(command []string) {
	// app := "gnome-terminal"
	// homeDir, _ := os.UserHomeDir()
	// args := [][]string{
	// 	{"--tab", "--", "tmux"},
	// 	{"--tab", "--working-directory", homeDir + "/d/code/web/github/gymkhana"},
	// 	{"--tab", "--working-directory", homeDir + "/d/temp", "--", "tmux"},
	// }
	cmd := exec.Command(command[0], command[1:]...)
	_, err := cmd.Output()

	if err != nil {
		log.Fatal("invalid command")
		return
	}
}
