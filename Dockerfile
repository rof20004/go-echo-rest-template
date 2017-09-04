# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Install Glide
RUN curl https://glide.sh/get | sh

# Create project folders
RUN mkdir -p /go/src/github.com/rof20004

# Enter project folder
WORKDIR /go/src/github.com/rof20004

# Clone git repository
RUN git clone https://github.com/rof20004/go-echo-rest-template.git

# Install project dependencies
WORKDIR /go/src/github.com/rof20004/go-echo-rest-template
RUN glide install

# Enter GOPATH
WORKDIR /go/src

# Copy the local package files to the container's workspace.
# ADD . /go/src/github.com/golang/example/outyet

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/rof20004/go-echo-rest-template

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-echo-rest-template

# Document that the service listens on port 8081.
EXPOSE 8081
