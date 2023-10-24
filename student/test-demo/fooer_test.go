package test_demo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFooer(t *testing.T) {

	configEnv := os.Getenv("GO_ENV")
	fmt.Println("configEnv: ", configEnv)
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooerSkiped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func FuzzTest11(f *testing.F) {
	var ch = make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for data := range ch {
				fmt.Println("取出数据", data)
			}
		}()
	}

	go func() {
		for true {
			ch <- 1
			time.Sleep(2 * time.Second)

		}
	}()

	time.Sleep(60 * time.Second)
}
