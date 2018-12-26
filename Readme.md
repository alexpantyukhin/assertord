# Go asserts for checking collections order.
[![Build Status](https://travis-ci.org/alexpantyukhin/assertord.svg?branch=master
)](https://travis-ci.org/alexpantyukhin/assertord)
[![GoDoc](https://godoc.org/alexpantyukhin/assertord?status.svg)](https://godoc.org/github.com/alexpantyukhin/assertord)

The lib contains a list of functions for checking ordering of slices and arrays.
The following checks are implemented:
   - IsOrdered
   - IsNotOrdered
   - IsOrderedAsc
   - IsNotOrderedAsc
   - IsOrderedDesc
   - IsNotOrderedDesc

# Usage example

```go
package main

import (
	"testing"

	"github.com/alexpantyukhin/assertord"
)

func TestIsOrdered(t *testing.T) {
	arr := []int{1, 2, 3, 4}

	assertord.IsOrdered(t, arr)
}
```
