# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9

# Copy the local package files to the container's workspace.
ADD . /go/src/gawk

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet

# Document that the service listens on port 8080.
EXPOSE 8080