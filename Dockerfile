FROM golang

# Set working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY cache_manager /app/cache_manager
COPY generator /app/generator
COPY server /app/server

# Install dependencies
RUN apt-get update && apt-get install -y golang-go
RUN go mod init my_project/urlgen
# TODO:  RUN go get ...

EXPOSE 4000

CMD ["go", "run", "server/server.go"]