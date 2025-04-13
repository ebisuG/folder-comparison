### Problem-Solution
#### The way to handle CLI input in a testable way
My first idea is here:
- Recieve folder path through arguments of function.
- Handle input as byte data to accept CLI and input from test code.
This is difficult to test because passing values to function process its input immediately and return the result. What I want to do was keep input in somewhere and invoke successively appropriate methods depending on a previous process for following one. I need some steps.

Solution:
```go
type CLI struct {
	args []string
}

	cli := &CLI{args: os.Args[1:]} 
```

Go's interface can keep state of something. In this case, it keeps string from CLI input. Golang provides `os.Args` method. This catches all strigns in a CLI line.\
What important is CLI is defined as just string slice. I guess for testing, a good way is to define an input as simple.

### Command
#### Run specific test
 go test -run Test_receiveArguments


Define test case as data


### Calculate hash
https://gobyexample.com/sha256-hashes
I need to calculate hash of a file and add it to the other hash. Hash is slice of byte type. Handle data in byte slice, finally convert result to human readable hex format.
```go
//a variable satisfying hash.Hash keeps hash data.
h := sha256.New()

h.Sum(nil)

```

### struct that has only function and interface
Go doesn't have class inheritance. There are only struct and interface. Struct can have only function and this looks interface, but imagine :
- a factory with many robots(robot A, robot B, ...)
- Each robot is a struct (concrete type, knows how to work).
- define a Worker interface to `Assemble()`
```go
type robotA struct{
    name string
}

func (robot *robotA)Assemble(){
    //do something
}

type Worker interface{
    Assemble()
}

func StartShift(w Worker) {
    w.Assemble()
}
```
This code doesn't care which robot is. robot A satisfies Worker interface by implementing Assebmel(). Golang takes different approach about OOP, so there are struct and interface.

https://go.dev/doc/faq#inheritance

#### Convert byte slice to hex string
Reference:\
https://schadokar.dev/to-the-point/convert-byte-to-hex-and-hex-to-byte-in-golang/\




### format string
https://gobyexample.com/string-formatting

### For loops improvement
https://go.dev/blog/loopvar-preview

### sync vs channel to wait finish of process
https://go.dev/wiki/MutexOrChannel


### Reference

Go言語でテストしやすいコマンドラインツールをつくる
https://deeeet.com/writing/2014/12/18/golang-cli-test/

nil slices vs non-nil slices vs empty slices in Go language
https://stackoverflow.com/questions/44305170/nil-slices-vs-non-nil-slices-vs-empty-slices-in-go-language