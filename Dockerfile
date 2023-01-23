# Source
FROM golang

# Working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY cache_manager /app/cache_manager
COPY generator /app/generator
COPY server /app/server
COPY config /app/config
COPY database /app/database
COPY go.mod /app/
COPY go.sum /app/

# Inform Docker that the container listens on the specified network ports
EXPOSE 4000

# Provide defaults for an executing container
CMD ["go", "run", "server/server.go"]