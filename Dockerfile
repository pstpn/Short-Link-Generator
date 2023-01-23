FROM golang

# Set working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY cache_manager /app/cache_manager
COPY generator /app/generator
COPY server /app/server
COPY config /app/config
COPY database /app/database
COPY go.mod /app/
COPY go.sum /app/

EXPOSE 4000

CMD ["go", "run", "server/server.go"]