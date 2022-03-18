# Bubble teas

## functional paradigm
Bubble tea uses the elm programming paradign to lay out the code.
What happens within the Elm program though? It always breaks into three parts:

    * Model — the state of your application
    * View — a way to turn your state into HTML
    * Update — a way to update your state based on messages

## Golang learning

### Golang pointers
background [info](https://gobyexample.com/pointers) about pointers

### Value receiver vs. pointer receiver
background [ info ](https://gobyexample.com/methods) about methods about methods
Go automatically handles conversion between values and pointers for method calls.

```go
type T struct {
    a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver
```

### closures (lambda functions)
background [info](https://gobyexample.com/closures) about closures


