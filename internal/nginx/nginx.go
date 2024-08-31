package nginx

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/AkhileshThykkat/fastapi-hoster/internal/config"
	"github.com/AkhileshThykkat/fastapi-hoster/internal/ssl"
)

const nginxTemplate = `server {
    listen 80;
    server_name {{.Domain}} www.{{.Domain}};

    location / {
        return 301 https://$host$request_uri;
    }
}

# HTTPS server
server {
    listen 443 ssl;
    server_name {{.Domain}} www.{{.Domain}};

    ssl_certificate /etc/letsencrypt/live/{{.Domain}}/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/{{.Domain}}/privkey.pem;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;
    ssl_ciphers "ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256";
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 1d;
    ssl_session_tickets off;
    ssl_stapling on;
    ssl_stapling_verify on;
    resolver 8.8.8.8 8.8.4.4 valid=300s;
    resolver_timeout 5s;

    location / {
        proxy_pass http://localhost:{{.Port}};
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
`

func ConfigureNginx(reader *bufio.Reader, cfg *config.ServiceConfig) {
	cfg.Domain = config.GetInput(reader, "Enter the domain name: ")

	for {
		err := ssl.CreateSSLCertificate(cfg.Domain)
		if err == nil {
			break
		}
		if !config.ConfirmInput(reader, "Error creating SSL certificate. Retry?") {
			return
		}
	}

	nginxConfigName := config.GetInput(reader, "Enter a name for the Nginx configuration: ")
	createNginxConfig(cfg, nginxConfigName)

	enableNginxConfig(nginxConfigName)
	testNginxConfig()
	reloadNginx()

	fmt.Printf("Your application is now hosted at https://%s\n", cfg.Domain)
}

func createNginxConfig(cfg *config.ServiceConfig, configName string) {
	tmpl, err := template.New("nginx").Parse(nginxTemplate)
	if err != nil {
		fmt.Println("Error creating Nginx template:", err)
		return
	}

	fileName := filepath.Join("/etc/nginx/sites-available", configName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating Nginx config file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, cfg)
	if err != nil {
		fmt.Println("Error writing Nginx config file:", err)
		return
	}

	fmt.Printf("Nginx configuration created: %s\n", fileName)
}

func enableNginxConfig(configName string) {
	source := filepath.Join("/etc/nginx/sites-available", configName)
	target := filepath.Join("/etc/nginx/sites-enabled", configName)
	err := os.Symlink(source, target)
	if err != nil {
		fmt.Printf("Error enabling Nginx config: %v\n", err)
	} else {
		fmt.Println("Nginx configuration enabled")
	}
}

func testNginxConfig() {
	cmd := exec.Command("nginx", "-t")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx configuration test failed: %v\n%s\n", err, output)
	} else {
		fmt.Println("Nginx configuration test passed")
	}
}

func reloadNginx() {
	cmd := exec.Command("systemctl", "reload", "nginx")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error reloading Nginx: %v\n", err)
	} else {
		fmt.Println("Nginx reloaded successfully")
	}
}
