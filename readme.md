# Roadmap Reference

Here's some roadmap reference to learn https://roadmap.sh/pdfs/go.pdf

## Go

Why Use Go ? There's Node JS

It is a compiled driven language, progress in the directory according to the https://www.openmymind.net/assets/go/go.pdf 
we can develop the application we want using this dir

```bash
mkdir {{parent_dir}}/go/src
```

Using the go programming language we can try to create the main function. 
Using the main package example code  

```go
package main 
 
func main() {
  println("Helo World")
}

```

After we write the code we can compile with  

```
go run main.go
```

or if you want to create the executable files you can create the go build main.go.

```
go build main.go
```

then we can have the executable of the source code of hello world program.  

That's basically the process so we can try to write the Makefile.

## Running every Section

To run every section of this writing excercices you can try to type this.

```bash
make build run
```

Generally I write the make file just like below snippets

```Makefile
build:
	cd ./bin/; go build ../src/main.go;

run: 
	cd ./bin; ./main; cd ..
```

so the directory of every exercices should be 

```bash
.
└── {{Exercices Name}}
    └── go
        ├── bin
        │   └── main
        ├── Makefile
        └── src
            └── main.go
```

That way we can try to run the compiled program using Makefile.

## Running locally documentation

documentation at 

```bash
godoc -http=:6060
```

## Go Modules

It's an additional features into which to simplify the package for Go Workspace Program.    

## Go Workspace

Go Workspace is opinionated for efficiency and organized project with the team. There's certain convention in the go workspace. Example for using the Folder

```bash
.
├── bin -------------------> for binary files
│   └── main --------------> the binary files
├── Makefile --------------> where to create the makefile  
├── pkg -------------------> using the package (archived files)
└── src -------------------> source code
    └── github.com --------> the control integration provider maybe github or gitlab
      └── <username> ------> the username in that provider
        └── <folder with the code for project or repo> ----> the username in that provider
```

The `pkg` folder is to put the file of compiled archived library so we don't need to recompile.  

## Go Help Command

```
go help
```

## Environment Variables

Two Important Environment Variables

```
GOPATH = workspace
GOROOT = binary installation of Go
```

Just Type

```
go env 
```

## Create the module. 

First of all we need to create the module in the project directory with this command 

```
go mod init <name of the module>
```

## Auto Formatting the Files

So Go have this features with the command `fmt` so all code will look the same. Just run 

```
go fmt <dir/files you want to format> 
```

So for examples if we create the files `src/main.go` like this

```go
package main

import ( "fmt" )

func main() {
fmt.Println("Hello World") --> this is bad
}
```

It can get some troubles if we work we a team so we can go 

```
go fmt src/main.go
```

It will automatically format the code for us just like below snippets.    

```go
package main

import ( "fmt" )

func main() {
  fmt.Println("Hello World") --> this is better
}
```

## Go Modules

Go Modules is to managing the packages we install, so if our software is dependant about a software we should just configure it with the go modules.

ref: https://go.dev/blog/using-go-modules

OCT 15th WISUDA KA OKA.


