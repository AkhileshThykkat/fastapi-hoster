package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/AkhileshThykkat/fastapi-hoster/internal/config"
)

const serviceTemplate = `[Unit]
Description=Uvicorn service for FastAPI application

[Service]
ExecStart={{.VenvPath}}/bin/uvicorn main:app --host 0.0.0.0 --port {{.Port}} --workers {{.Workers}}
WorkingDirectory={{.AppDir}}
Restart=always
Environment="PATH={{.VenvPath}}/bin"

[Install]
WantedBy=multi-user.target
`

func CreateSystemdService(config *config.ServiceConfig) {
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	fileName := filepath.Join("/etc/systemd/system", config.ServiceName+".service")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating service file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, config)
	if err != nil {
		fmt.Println("Error writing service file:", err)
		return
	}

	fmt.Printf("Systemd service created: %s\n", fileName)
}

func EnableService(serviceName string) {
	cmd := exec.Command("systemctl", "enable", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error enabling service: %v\n", err)
	} else {
		fmt.Println("Service enabled successfully")
	}
}

func StartService(serviceName string) {
	cmd := exec.Command("systemctl", "start", serviceName)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error starting service: %v\n", err)
	} else {
		fmt.Println("Service started successfully")
	}
}
