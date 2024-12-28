# ✋ BUILDPLATFORM is a build argument that specifies the platform where the image is built.
FROM --platform=$BUILDPLATFORM golang:1.23.4-alpine AS builder
WORKDIR /app
COPY go.mod .

# ✋ TARGETOS and TARGETARCH are passed as build arguments.
ARG TARGETOS
ARG TARGETARCH

RUN <<EOF
go mod tidy 
EOF

COPY main.go .

#RUN <<EOF
#GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build
#EOF

#CMD ["go", "run", "main.go"]

# docker buildx build \
#  --platform=linux/amd64,linux/arm64 \
#  --push -t philippecharriere494/paris-restaurants:0.0.1 .