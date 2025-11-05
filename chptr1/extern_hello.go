package main
import "fmt"

// Now to try and import something from pkg.go.dev
import "rsc.io/quote"


func main() {
	fmt.Println(quote.Go())
}