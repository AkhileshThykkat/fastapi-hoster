package host

import (
	"bufio"
	"fmt"

	"github.com/AkhileshThykkat/fastapi-hoster/internal/config"
	"github.com/AkhileshThykkat/fastapi-hoster/internal/nginx"
	"github.com/AkhileshThykkat/fastapi-hoster/internal/service"
	"github.com/AkhileshThykkat/fastapi-hoster/internal/ufw"
)

func HostApplication(reader *bufio.Reader) {
	cfg := &config.ServiceConfig{}

	cfg.AppDir = config.GetInput(reader, "Enter the FastAPI application root directory: ")
	cfg.VenvPath = config.GetInput(reader, "Enter the virtual environment path: ")
	cfg.Port = config.GetInput(reader, "Enter the port number: ")
	cfg.Workers = config.GetInput(reader, "Enter the number of worker processes: ")
	cfg.ServiceName = config.GetInput(reader, "Enter the name for the systemd service: ")

	service.CreateSystemdService(cfg)

	if config.ConfirmInput(reader, "Enable this service?") {
		service.EnableService(cfg.ServiceName)
	}

	if config.ConfirmInput(reader, "Start the service?") {
		service.StartService(cfg.ServiceName)
	}

	if config.ConfirmInput(reader, "Allow UFW?") {
		ufw.ConfigureUFW(cfg.Port)
	}

	if config.ConfirmInput(reader, "Proceed with Nginx domain config?") {
		nginx.ConfigureNginx(reader, cfg)
	} else {
		fmt.Printf("Your application is now hosted at http://your_ip_address:%s\n", cfg.Port)
	}
}
