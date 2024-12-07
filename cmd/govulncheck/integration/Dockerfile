FROM golang:1.22.3-alpine

# This Dockerfile sets up an image for repeated integration testing.
# This assumes the build context, i.e., CWD is vuln/

# ---- Step 0: Setup shared build tools. ----
RUN apk update && apk add bash git gcc musl-dev linux-headers gcompat

# ---- Step 1: Build govulncheck ----
COPY . /go/src/golang.org/x/vuln
WORKDIR /go/src/golang.org/x/vuln/cmd/govulncheck/integration
RUN go install github.com/StevenACoffman/invuln/cmd/govulncheck

# ---- Step 2: Build other test binaries ----
RUN go install github.com/StevenACoffman/invuln/cmd/govulncheck/integration/k8s
RUN go install github.com/StevenACoffman/invuln/cmd/govulncheck/integration/stackrox-scanner
