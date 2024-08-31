package config

import (
	"bufio"
	"fmt"
	"strings"
)
type ServiceConfig struct {
	AppDir      string
	VenvPath    string
	Port        string
	Workers     string
	ServiceName string
	Domain      string
}

func GetInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ConfirmInput(reader *bufio.Reader, message string) bool {
	fmt.Println(message)
	fmt.Print("Confirm? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(confirm)) == "y"
}
