# Source
FROM golang

# Working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY init/init_db.sql /docker-entrypoint-initdb.d/
COPY cmd /app/cmd
COPY pkg/ /app/pkg/
COPY internal/ /app/internal/
COPY config /app/config
COPY database /app/database
COPY go.mod /app/
COPY go.sum /app/

# Inform Docker that the container listens on the specified network ports
EXPOSE 4000

# Provide defaults for an executing container
CMD ["go", "run", "cmd/app/main.go"]