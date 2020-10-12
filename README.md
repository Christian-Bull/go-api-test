# go-api-test

This is just a super simple api that I use for testing  

### Build instructions  
  
`git clone https://github.com/Christian-Bull/go-api-test.git`  
  
`docker build -t <sometagname> .`  

`docker run -p 5050:5050 <sometagname>`    
<br>

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