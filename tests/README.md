# Tests 

I use `wrk` to run my http tests here.

## Standard test 

```bash
wrk -t12 -c400 -d30 http://DOMAIN_HERE:PORT_HERE
```

## Create test 

```bash
wrk -t12 -c400 -d30 -s ./tests/create-post.lua http://DOMAIN_HERE:PORT_HERE/posts
```

## Read test

```bash
wrk -t12 -c400 -d30 -s ./tests/create-post.lua http://DOMAIN_HERE:PORT_HERE/posts
```