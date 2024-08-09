# Use the official PostgreSQL image from the Docker Hub
FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=qd
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=cardadb

# Copy initialization scripts (if any) to the Docker image
COPY ./init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432