# Use the official PostgreSQL image from the Docker Hub
FROM postgres:latest

# Copy initialization scripts (if any) to the Docker image
COPY ./init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432