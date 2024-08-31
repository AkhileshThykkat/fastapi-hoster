package launch

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/AkhileshThykkat/fastapi-hoster/internal/config"
)

func LaunchApplication(reader *bufio.Reader) {
	appDir := config.GetInput(reader, "Enter the FastAPI application root directory: ")
	venvPath := config.GetInput(reader, "Enter the virtual environment path: ")
	port := config.GetInput(reader, "Enter the port number to run the application: ")

	activateCmd := fmt.Sprintf("source %s/bin/activate", venvPath)
	uvicornCmd := fmt.Sprintf("uvicorn main:app --reload --port %s", port)
	fullCmd := fmt.Sprintf("%s && cd %s && %s", activateCmd, appDir, uvicornCmd)

	cmd := exec.Command("bash", "-c", fullCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Launching FastAPI application on port %s...\n", port)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error launching application: %v\n", err)
	}
}
