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
1. Navigate to the project directory and build the Go binary with `go build`

2. Make sure Docker is running in the background.

3. Start up the Postgres database container, pre-loaded with patient data: `docker-compose up -d`

4. Run the Go binary with `./sothea_backend`
 
5. The server should now be running on `localhost:9090`

6. You can now make requests to the server using a tool like Postman or curl.
 
7. To stop the server, press `Ctrl + C` in the terminal, then run `docker-compose down` to stop the database container.