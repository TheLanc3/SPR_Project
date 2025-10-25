FROM golang:1.25.3-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache gcc musl-dev

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/*.go ./

RUN CGO_ENABLED=1 go build -o ./main -trimpath -ldflags "-s -w"


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/main /main

ENV FIBER_PORT=8080

EXPOSE $FIBER_PORT

CMD ["/main"]