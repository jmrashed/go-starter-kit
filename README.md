# Go Starter Kit

## Overview

This is a starter kit for building a Go-based portfolio website. It uses the [Gin](https://github.com/gin-gonic/gin) web framework and [Viper](https://github.com/spf13/viper) for configuration management. The project is structured to promote clean code and scalability.

## Project Structure

```
go-starter-kit/
├── cmd/
│   └── myapp/
│       └── main.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── logger/
│   │   └── logger.go
│   ├── server/
│   │   └── server.go
├── templates/
│   ├── base.html
│   ├── home.html
│   ├── about.html
│   ├── education.html
│   ├── skills.html
│   ├── experience.html
│   ├── contact.html
├── assets/
│   ├── css/
│   │   └── style.css
│   └── js/
│       └── script.js
├── .gitignore
├── README.md
├── LICENSE
├── go.mod
├── go.sum
└── config.yaml

```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository:**

```sh
    git clone https://github.com/jmrashed/go-starter-kit.git
    cd go-starter-kit
```

2. **Install dependencies:**

```sh
    go mod tidy
```

### Configuration

Ensure you have a `config.yaml` file in the root directory of the project with the following content:

```yaml
server_address: ":8080"
```

### Running the Application

Navigate to the project root directory and run:

```sh
go run cmd/myapp/main.go
```

The application will start, and you can access it at `http://localhost:8080`.

### Project Details

#### cmd/myapp/main.go

The entry point of the application. It initializes the configuration, logger, and server.

#### pkg/config/config.go

Handles the configuration using Viper. Loads settings from `config.yaml`.

#### pkg/logger/logger.go

Sets up logging for the application.

#### pkg/server/server.go

Configures the Gin router, sets up routes, and serves static files and templates.

#### templates/

Contains HTML templates for the application. The default template is `index.html`.

#### assets/

Contains static assets like CSS and JavaScript files.

### Example `index.html`

An example HTML file located in the `templates` directory:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
    <link rel="stylesheet" href="/assets/css/tailwind.css">
</head>
<body>
    <h1>{{ .title }}</h1>
    <p>Welcome to my portfolio!</p>
</body>
</html>
```

### Example `config.yaml`

An example configuration file located in the root directory:

```yaml
server_address: ":8080"
```

## Development

To run the application in development mode:

1. Ensure Go is installed and your environment is set up.
2. Navigate to the root directory of the project.
3. Run the application using `go run cmd/myapp/main.go`.

## Deployment

To build the application for production:

1. **Build the binary:**

```sh
    go build -o bin/myapp cmd/myapp/main.go
```

2. **Run the binary:**

```sh
    ./bin/myapp
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any changes or additions.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Viper Configuration Management](https://github.com/spf13/viper)
 

### Explanation

- **Overview:** A brief introduction to the project.
- **Project Structure:** A visual representation of the directory structure.
- **Getting Started:** Instructions for cloning the repository, installing dependencies, and running the application.
- **Configuration:** Details on how to set up the configuration file.
- **Running the Application:** Step-by-step instructions on running the project.
- **Project Details:** Information about the key directories and files in the project.
- **Development:** Instructions for running the application in development mode.
- **Deployment:** Steps to build and run the application for production.
- **Contributing:** Guidelines for contributing to the project.
- **License:** Licensing information.
- **Acknowledgements:** Credits to the libraries and frameworks used in the project.
 