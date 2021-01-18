package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64
//var mutex sync.Mutex

func concurrency(){

	wg.Add(2)
	go incrementer("foo:")
	go incrementer("bar:")
	wg.Wait()
}

func incrementer(s string){
	for i:=0; i<20; i++{
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "counter", counter)

	}
	wg.Done()
}

/** using mutex
func incrementer(s string){
	for i:=0; i<20; i++{
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "counter", counter)
		mutex.Unlock()
	}
	wg.Done()
}
*/


func foo(){
	for i:=0; i<10;i++{
		fmt.Println("foo", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar(){
	for i:=0; i<10;i++{
		fmt.Println("bar", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}