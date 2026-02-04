
package usecase

import "testing"

func TestConversion(t *testing.T) {
    c := 10.0
    f := c*1.8 + 32
    k := c + 273

    if f != 50 {
        t.Fail()
    }
    if k != 283 {
        t.Fail()
    }
}
