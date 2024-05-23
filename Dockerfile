# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=jieqiboh
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=patients

# Copy the initialization scripts into the Docker entrypoint directory
COPY ./sql /docker-entrypoint-initdb.d
#COPY sql/patients_setup.sql /docker-entrypoint-initdb.d/init.sql

# Expose the PostgreSQL port
EXPOSE 5432

# Set the entrypoint to the default PostgreSQL entrypoint
ENTRYPOINT ["docker-entrypoint.sh"]

# Run the PostgreSQL server
CMD ["postgres"]