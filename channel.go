package main

import (
	context2 "context"
	"fmt"
	"runtime"
	"time"
)

func channeling(){
	c := make(chan int)

	go func(){
		for i:=0; i<10;i++{
			c <-i
		}
		close(c)
	}()

	for n:= range c{
		fmt.Print(n)
	}
	/**
	alternative way using sleep to wait until channel is done
	go func(){
		for {
			fmt.Print(<-c)
		}
	}()
	time.Sleep(time.Second)
	*/
}

func directionalChan(){
	c := make(chan int)

	go sender(c)
	receiever(c)
	fmt.Println("done")
}

func sender(c chan<- int){
	c <-42
}

func receiever (c <-chan int){
	fmt.Println(<-c)
}

func context(){
	ctx, cancel := context2.WithCancel(context2.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())


}


func anotherChannel(){
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}