For those who want to later code implementation.

#### How to use

```main.go
import "github.com/tkitsunai/gotodo"

func main() {
	gotodo.TODO("I am busy with the game so I will implement later.")
}
```

```
$ go run main.go 
panic: TODO: An operation is not implemented :I am busy with the game so I will implement later.

goroutine 1 [running]:
github.com/tkitsunai/gotodo.TODO(0xc000084f78, 0x1, 0x1)
```
