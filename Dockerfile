# Source
FROM golang

# Working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY cmd /app/cmd
COPY pkg/cache_manager /app/pkg/cache_manager
COPY pkg/generator /app/pkg/generator
COPY internal/server /app/internal/server
COPY config /app/config
COPY database /app/database
COPY go.mod /app/
COPY go.sum /app/

# Inform Docker that the container listens on the specified network ports
EXPOSE 4000

# Provide defaults for an executing container
CMD ["go", "run", "cmd/app/main.go"]