# Build the manager binary
FROM golang:1.18 as builder

ARG TARGETARCH
ARG GIT_HEAD_COMMIT
ARG GIT_TAG_COMMIT
ARG GIT_LAST_TAG
ARG GIT_MODIFIED
ARG GIT_REPO
ARG BUILD_DATE

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

ENV http_proxy="http://httpproxy.bfm.com:8080"
ENV https_proxy="http://sftpproxy.bfm.com:8080"
ENV no_proxy=".local,.blkint.com,.bfm.com,.blackrock.com,.mlam.ml.com,*.local,*.blkint.com,*.bfm.com,*.blackrock.com,*.mlam.ml.com,localhost,127.0.0.1,192.168.99.100"
ENV ftp_proxy="ftp://sftpproxy.bfm.com/"
ENV HTTP_PROXY="http://httpproxy.bfm.com:8080"
ENV HTTPS_PROXY="http://sftpproxy.bfm.com:8080"
ENV NO_PROXY=".local,.blkint.com,.bfm.com,.blackrock.com,.mlam.ml.com,*.local,*.blkint.com,*.bfm.com,*.blackrock.com,*.mlam.ml.com,localhost,127.0.0.1,192.168.99.100"
ENV FTP_PROXY="ftp://sftpproxy.bfm.com/"
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY version.go version.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH GO111MODULE=on go build \
        -gcflags "-N -l" \
        -ldflags "-X main.GitRepo=$GIT_REPO -X main.GitTag=$GIT_LAST_TAG -X main.GitCommit=$GIT_HEAD_COMMIT -X main.GitDirty=$GIT_MODIFIED -X main.BuildTime=$BUILD_DATE" \
        -o manager

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]