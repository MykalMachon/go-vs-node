# App Go

This is a simple golang application

## Run in dev 

```bash
go install github.com/cortesi/modd/cmd/modd@latest
modd
```

## Build for prod

```bash
docker build -t app-go:1.0 .
```

## Run in prod

```bash
docker run -p "8080:8080" app-go:1.0
```