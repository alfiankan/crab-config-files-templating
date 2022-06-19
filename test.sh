#!/bin/sh

go run ./cmd/... \
    -f nginx-deployment.yaml \
    -r namespace=production \
    -r exposePort=8081 \
    -q traceUrl=http://trace.domain.com \
    -q authUrl=http://auth.domain.com \
    -v \
    -o result.yaml
