**[wrap-mockgen.sh](scripts/wrap-mockgen.sh)** is a helper-script for more convenient use of //go:generate with gomock.

### What is the problem with direct gomock call

For example

```go
//go:generate mockgen -source=my_interface.go -destination=mocks/my_interface_generated.go -mock_names myInterface=MockMyInterface -typed -package=mocks myInterface
type myInterface interface {
    Method()
}
```
Imagine for each new interface you have to copy the line and then change 5(!) parameters. It is too easy to make a mistake by pasting something incorrectly.

### Compare with the alternative I suggest

```go
//go:generate wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type myInterface interface {
    Method()
}
```
The same code for all the interfaces.
See the [example](example/) folder for simple and table test cases. Notice the naming and general test layout.

### Installation

1. Place the file [scripts/wrap-mockgen.sh](scripts/wrap-mockgen.sh) in a directory that is included in your $PATH.
2. Set executable permissions ```chmod +x path/to/wrap-mockgen.sh ```
3. Install mockgen if not already installed ```go install go.uber.org/mock/mockgen@latest```

Happy mocking! :) 

### Usage

1. Locally installed wrap-mockgen.sh: 
```go
//go:generate wrap-mockgen.sh -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type myInterface interface {
    Method()
}
```

2. Prebuilt docker image:
```go
//go:generate docker run -v ${PWD}:/w rogozhka/go-generate-mockgen -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type myInterface interface {
    Method()
}
```
