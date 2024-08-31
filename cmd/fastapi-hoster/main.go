package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AkhileshThykkat/fastapi-hoster/internal/host"
	"github.com/AkhileshThykkat/fastapi-hoster/internal/launch"
)

func main() {
	fmt.Println("Welcome to the FastAPI Systemd Service Setup!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Launch application")
		fmt.Println("2. Host application")
		fmt.Println("3. Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			launch.LaunchApplication(reader)
		case "2":
			host.HostApplication(reader)
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
