package TokenBucket

import (
	"fmt"
	"testing"
	"time"
)

func TestAllow(t *testing.T) {
	p := newLimiter(2, 3)
	for {
		for i := 0; i < 3; i++ {
			go func(i int) {
				if p.Allow() {
					fmt.Println(" 可以的", i)
				} else {
					fmt.Println("不可以的", i)
				}

			}(i)

		}
		time.Sleep(time.Second)
		fmt.Println("--------------")
	}

}
