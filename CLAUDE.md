# CLAUDE.md - goswucards

## Project Overview
**goswucards** is a Go library that provides comprehensive data about Star Wars Unlimited trading card game cards. It serves as an importable library for projects that need access to card information, including metadata and images.

## Project Structure
```
goswucards/
├── CLAUDE.md              # Project documentation and AI instructions
├── LICENSE                # License file
├── Makefile               # Build automation
├── README.md              # Project readme
├── go.mod                 # Go module definition
├── images/                # Downloaded card images
│   └── SOR001.png        # Sample card image
└── internal/              # Internal tooling
    ├── main.go            # Scraper and code generator
    └── template           # Go template for generating package
```

## Core Components

### 1. `internal/main.go`
- Scrapes card data from swudb web application and generates the library package

### 2. `internal/template`
- Go template file used to generate the goswucards.go file structure

### 3. `Makefile`
- Build automation file that provides commands for generating the package

## Data Source
- Primary source: swudb web application
- Scrapes comprehensive card information including metadata
- Downloads high-quality card images