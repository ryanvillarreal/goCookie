# goCookie
goCookie is based on the AnomalousCookie from Coalfire Research Github page.  https://github.com/Coalfire-Research/AnomalousCookie

```
NAME:
   goCookie - Usage: goCookie.exe https://www.example.com

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   l33tllama

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --delay value, -d value    --delay (milliseconds) or -d (milliseconds) (default: "0")
   --output value, -o value   --output <output> or -o <output>  (default: "./output.txt")
   --pic value, -P value      --pic <s> or -p <s> Disable picture taking.
   --proxy value, -p value    --proxy <127.0.0.1:8080> or -f <127.0.0.1:8080>
   --request value, -r value  --request GET/POST or -p GET/POST (default: "GET")
   --target value, -t value   --target http://localhost/dir or -t http://localhost/dir
   --help, -h                 show help
   --version, -v              print the version
```

## Bugs/Progress
1.  Bug: Golang won't allow to send illeagal characters in a cookie.  
2.  ~~fuzz cookies~~
3.  ~~support GET/POST requests~~
4.  ~~proxy support~~
5.  ~~delay support~~
6.  output results to a file
7.  picture taking of web app when fuzzing cookies.  Borrowed Inspiration from goEyewitness
8.  error handling
9.  statistical analysis of different results (i.e. - response delays)
10.  multi-threading
11. add newly discovered cookies to the queue


## Installation
Right now I haven't built out the setup.sh script.  I will be working on that for future releases.  

Make sure to have your $GOPATH set `export GOPATH=$HOME/go` and then simply run: 

```
go get github.com/ryanvillarreal/goCookie
```

Now the project should live in your $HOME/go/src/github.com/ryanvillarreal/goCookie/ .  You can simply run go like a script by using 
`go run main.go`

or you can build for your current platform using `go build main.go`
