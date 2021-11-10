FROM golang:alpine AS builder

LABEL maintainer="Mayukh Sarkar <mayukh2012@hotmail.com>"
# Redundant, current golang images already include ca-certificates
RUN apk --no-cache add ca-certificates

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

# 2 staged build
FROM scratch
# copy the ca-certificate.crt from the build stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

EXPOSE 9999/tcp
EXPOSE 9000/tcp
# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]