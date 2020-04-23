# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.14
ARG GO_VERSION=1.14

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git 

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# download swagger command line
RUN go get -u github.com/swaggo/swag/cmd/swag
# generate swagger documentation
RUN swag init -g ./cmd/api/main.go

# run unit tests
RUN CGO_ENABLED=0 go test \
-v \
-cover ./...


# Build the executable to `/api`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -installsuffix 'static' -ldflags "-X main.Version=`date -u +1.%Y%m%d.%H%M%S`" \
    -o /api ./cmd/api/

# Final stage: the running container.
FROM alpine

WORKDIR /usr

# Import the user and group files from the first stage.
COPY --from=0 /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=0 /api .

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 5000

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
CMD ["/usr/api"]