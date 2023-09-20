FROM golang:1.21-alpine AS builder

ARG GH_LOGIN
ENV GH_LOGIN=$GH_USER

ARG GH_ACCESS_TOKEN
ENV GH_ACCESS_TOKEN=$GH_ACCESS_TOKEN

ARG GH_PATH
ENV GH_PATH=$GH_PATH

RUN apk add git

RUN echo "machine github.com login ${GH_LOGIN} password ${GH_ACCESS_TOKEN}" > ~/.netrc
ENV GOPRIVATE "${GH_PATH}/*"
RUN go env -w GOPRIVATE="${GH_PATH}/*"
ENV GO111MODULE=on

# RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./go-service ./cmd/go-service/main.go

FROM gcr.io/distroless/static-debian11 AS runner

COPY --from=builder --chown=nonroot:nonroot /app/go-service /

EXPOSE 80

ENTRYPOINT ["/go-service"]
