FROM public.ecr.aws/lambda/go:latest

WORKDIR /app

COPY api/ ./api
COPY cmd/ ./cmd
COPY events/ ./events
COPY internal/ ./internal
COPY invocation/ ./invocation
COPY lambda/ ./lambda
COPY go.mod ./go.mod
COPY go.sum ./go.sum