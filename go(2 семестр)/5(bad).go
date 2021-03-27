package main
import (
    "fmt"
    "math/big"
    "github.com/skorobogatov/input"
)
func main() {
    var n int
    input.Scanf("%d", &n)
    a := big.NewInt(1)
    b := big.NewInt(1)
    for ; n > 1;{
        b.Add(a, b)
        a,b = b,a
        n--
    }
    fmt.Println(b)
}