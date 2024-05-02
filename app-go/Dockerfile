FROM golang:1.22.1 as base 

# BUILD THE APP
FROM base as builder 
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY *.go .

RUN go build -o go-app .

# RUN THE APP
FROM base as runner
WORKDIR /app

RUN addgroup --system --gid 1001 golang
RUN adduser --system --uid 1001 gopher
RUN chown -R gopher:golang /app

COPY --from=builder --chown=gopher:golang /app/go-app /app/go-app

USER gopher
EXPOSE 8080 

CMD ["./go-app"]