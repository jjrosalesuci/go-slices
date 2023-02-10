# go-slices
Libreria para el menajo de arreglos con generics

## Como instalar 

```
go get github.com/jjrosalesuci/go-slices
```

## Funciones disponibles
### Filter

```go
package main

import (
	"fmt"
	slicesUtil "github.com/jjrosalesuci/go-slices"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	divisibleBy2 := slicesUtil.Filter(numbers, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(divisibleBy2)
}
```


### Reduce

```go
numbers := []int{1, 2, 3}
sum := slicesUtil.Reduce(numbers, func(acc, current int) int {
	return acc + current
}, 0)

```

### Map

```go
numbers := []float64{4, 9}
newNumbers := slicesUtil.Map(numbers, math.Sqrt)
```

## Some
```go
numbers := []float64{4, 9, 20}

toFind := 2
result := Some(numbers, func(numberCursor float64) bool {
	return numberCursor == float64(toFind)
})

```