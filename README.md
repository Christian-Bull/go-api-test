[![Go Report Card](https://goreportcard.com/badge/github.com/Christian-Bull/go-api-test)](https://goreportcard.com/report/github.com/Christian-Bull/go-api-test)
![badge](https://action-badges.now.sh/Christian-Bull/go-api-test)

# go-api-test

This is just a super simple api that I use for testing  

## Build instructions  
  
`git clone https://github.com/Christian-Bull/go-api-test.git`  
  
`docker build -t <sometagname> .`  

`docker run -p 5050:5050 <sometagname>`    
<br>
### Buildx
`docker buildx build --platform linux/arm64,linux/amd64 -t go-api-sample:v1 . `  
<br>

## Info
Send a list of strings to the end-point `/reverse` and you'll receive them back but reversed

```
[
    {
        "data": "hello"
    },
    {
        "data": "hola"
    },
    {
        "data": "hallo"
    }
]
```

```
[
    {
        "data": "olleh"
    },
    {
        "data": "aloh"
    },
    {
        "data": "ollah"
    }
]
```
