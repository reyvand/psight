# Psight
Simple port scanner based on Go programming language.


Installation:

```sh
> cd psight
> GOPATH=$(pwd) go build
> ls
main.go   psight    readme.md src
```

Usage:
```sh
> ./psight 

   __   __        _       _     _   
  / /   \ \      (_)     | |   | |  
 | | _ __| |  ___ _  __ _| |__ | |_ 
/ / | '_ \\ \/ __| |/ _' | '_ \| __|
\ \ | |_) / /\__ \ | (_| | | | | |_ 
 | || .__/ | |___/_|\__, |_| |_|\__|
  \_\ | /_/          __/ |          
    |_|             |___/      v.0.1

Usage:
	-t IP	Target IP
Optional:
	-p PORT	Target Port

Examples:

- Discover all open port from target [default]
  $ psight -t 127.0.0.1
- Scan single port from target
  $ psight -t 127.0.0.1 -p 22
- Scan multiple port from target
  $ psight -t 127.0.0.1 -p 21,22
- Scan range of port from target
  $ psight -t 127.0.0.1 -p 22-25
```