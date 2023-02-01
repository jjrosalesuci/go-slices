# go-slices
Libreria para el menajo de arreglos

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
