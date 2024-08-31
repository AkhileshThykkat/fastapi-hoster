# FastAPI Hoster

FastAPI Hoster is a straightforward, Go-based tool for those who’d rather spend time creating than worrying about deployment. It covers the essentials—Nginx, SSL, firewall rules—so you don’t have to. It’s simple and does what it needs to, making it easy to host and manage your FastAPI applications. Let FastAPI Hoster take care of the setup while you focus on what you do best.

_P.S.: This is just the initial version—more updates are on the way!_

## Features

- Launch FastAPI applications
- Configure and manage Nginx for reverse proxying
- Set up SSL certificates using Certbot
- Configure UFW (Uncomplicated Firewall) rules
- Create and manage systemd services for your FastAPI apps

## Installation

To use FastAPI Hoster, you need to have Go installed on your system. If you haven't installed Go, you can download it from [the official Go website](https://golang.org/dl/).

Clone the repository:

```bash
git clone https://github.com/AkhileshThykkat/fastapi-hoster.git
cd fastapi-hoster
```

Build the project:

```bash
go build -o fastapi-hoster ./cmd/fastapi-hoster
```

## Usage

Run the FastAPI Hoster:

```bash
./fastapi-hoster
```

Follow the interactive prompts to:

1. Launch your FastAPI application
2. Host your application with Nginx
3. Configure SSL certificates
4. Set up UFW rules

## Contributing

Contributions are welcome to the FastAPI Hoster project! Here's how you can contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please make sure to update tests as appropriate and adhere to the existing coding style.

## License

Distributed under the MIT License. See `LICENSE` file for more information.

## Contact

Your Name - [@AkhileshThykkat](https://x.com/AkhileshThykkat) - akhileshthykkat843@gmail.com

Project Link: [https://github.com/AkhileshThykkat/fastapi-hoster](https://github.com/AkhileshThykkat/fastapi-hoster)

## Acknowledgements

- [FastAPI](https://fastapi.tiangolo.com/)
- [Nginx](https://nginx.org/)
- [Certbot](https://certbot.eff.org/)
- [UFW](https://help.ubuntu.com/community/UFW)
