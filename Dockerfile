FROM golang:1.23.4 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/cxp-api-starter/main.go

## STEP 2
FROM scratch

COPY --from=0 /app/main /app
EXPOSE 8080

CMD [ "/app" ]
