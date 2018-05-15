// You can edit this code!
// Click here and start typing.
package main

import "testing"
import "fmt"

func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { fmt.Println("This") })
    t.Run("A=2", func(t *testing.T) { fmt.Println("is") })
    t.Run("B=1", func(t *testing.T) { fmt.Println("Go!!!")})
    // <tear-down code>
}
