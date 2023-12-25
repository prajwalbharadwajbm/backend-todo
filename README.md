# A foundational backend ToDo application implemented in GoLang, incorporating CRUD functionality

## Overview
This project was solely to implement CRUD in GoLang, and also implement a test workflow automated using github actions

## Features
- CRUD functionalities
- Github workflow
- dockerfile

## Technologies
- GoLang
- Docker

## Getting Started
To get started with the ToDo application, follow these steps:

1. **Setting up the Development Environment:**
   - Ensure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
   - Clone this repository to your local machine.

2. **Installing Dependencies:**
   - Open a terminal and navigate to the project directory.
   - Run the following command to download and install project dependencies:
     ```bash
     go mod download
     ```

3. **Running the Project:**
   - Execute the following command to start the ToDo application:
     ```bash
     go run main.go
     ```
   - Open your web browser and visit [http://localhost:8080](http://localhost:8080) to interact with the ToDo application.

## Project Structure
The project is structured as follows:

- **`main.go`:**
  - The main entry point of the ToDo application.

- **`main_test.go`:**
  - Unit tests for the main application.

- **`go.mod` and `go.sum`:**
  - Go module files specifying project dependencies.
 
- **`Dockerfile`:**
  - Dockerfile for building a Docker image of the ToDo application.

- **`.github/workflows/test.yml`:**
  - GitHub Actions workflow for running tests on every push to the repository.

Feel free to explore the directories for more details on the project's structure and functionality.

## Acknowledgments
- [Learn to build and deploy your distributed applications easily to the cloud with Docker](https://docker-curriculum.com)
- [Using workflows:Creating and managing GitHub Actions workflows.](https://docs.github.com/en/actions/using-workflows)
