# App Node 

This is a simple node.js application

## Run in dev 

```bash
npm install
npm run dev
```

## Build for prod

```bash
docker build -t app-node:1.0 .
```

## Run in prod

```bash
docker run -p "3000:3000" app-node:1.0
```