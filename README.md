# Project Sothea Backend

## Overview

This is the backend for the patient management system for Project Sothea, and is to be set up in conjunction with the frontend.  
The backend is written in Go, and uses a PostgreSQL database to store patient data. The backend provides a RESTful API for the frontend to interact with, and is responsible for handling requests to create, read, update, and delete patient data.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Golang](https://golang.org/) - The Go programming language.
- [PostgreSQL](https://www.postgresql.org/) - An open-source relational database system.
- [Docker](https://www.docker.com/) - A platform for building, shipping, and running applications in containers.
- [pgAdmin](https://www.pgadmin.org/) - A comprehensive database management tool for PostgreSQL. Good to have for database management.

### Installation and Setup
1. Clone the repository to your local machine: `git clone https://github.com/Project-Sothea/Sothea-Backend.git`
 
2. In the project folder, build the project with `go build -o bin/sothea-backend` 
 (this ensures the binary is located in the `/bin` folder, so it can be gitignored)

3. Set up the required docker containers for the database (see below).

4. Run the Go binary with `./bin/sothea-backend --mode=dev`, starting it up in development mode.
 
5. The server should now be accessible on `http://localhost:9090`

6. You can now make requests to the server using a tool like Postman or curl.
 
7. To stop the server, enter `Ctrl + C` in the terminal, then run `docker stop sothea-db` to stop the database container.

### Setting Up Docker
To facilitate easy setup of the patients database with preloaded data, we've opted to use Docker with a PostgreSQL image. To set up the database, follow the steps below:
1. Make sure Docker is running in the background.

2. Build the Docker image for the Postgres database: `docker build -t sothea-db .`

3. Run the Postgres database container with `docker run --rm --name sothea-db -d -p 5432:5432 sothea-db`

4. To stop the container, run `docker stop sothea-db`

### Modes of Operation
When running the Go binary, you can specify the mode of operation using the `--mode` flag. The available modes are:  
- `dev` - Development mode, using config.json for configuration. This mode will run the server on port 9090.
- `prod` - Production mode, using prod.json for configuration. This mode will run the server on "192.168.0.100:9090", a static IP address that we use on our production server on the deployed network.   
Do ensure that the frontend is appropriately configured to make requests to the correct server address.

### Common Issues
- Database role not found / Authentication Failed
This usually happens if there are already pre-existing Postgres instances running on port 5432. To resolve this, stop check the processes running on port 5432, and stop the existing Postgres processes.