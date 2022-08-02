## Go

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

or 

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

