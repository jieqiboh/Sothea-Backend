# Project Sothea Backend

## Overview

TO ADD

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

3. Make sure Docker is running in the background.

4. Start up the Postgres database container, pre-loaded with patient data: `docker compose up -d`

5. Run the Go binary with `./bin/sothea-backend`
 
6. The server should now be running on `localhost:9090`

7. You can now make requests to the server using a tool like Postman or curl.
 
8. To stop the server, press `Ctrl + C` in the terminal, then run `docker compose down` to stop the database container.