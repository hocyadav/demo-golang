package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { // main is itself is one go routine
	//working - never reach 2nd foo
	//foo("sheep")// this will only run
	//foo("cat")// never run


	//working - 2 go routine: one is main and other one is below foo-sheep
	//go foo("sheep")// go makes run in background, so now we have 2 go-routine one is main and one is foo()
	//foo("cat")//this is part of main go-routine


	//total 3 go routine: one is main go-routine
	//problem : problem will exit and not excute below lines because main go routine will complete very fast and we cant se result
	//go foo("sheep")// go makes run in background, so now we have 2 go-routine one is main and one is foo()
	//go foo("cat")//this is part of main go-routine


	//solve above problem, so add blocking call by adding sleep or fmt.Scanf()
	//still some problem , we have fixed time to see output
	//go foo("sheep")
	//go foo("cat")
	//time.Sleep(time.Millisecond * 4000) //wait for five second,
	//or fmt.Scanf()


	//solution1 of above problem - use sync package counter
	//we increment the counter based on how many we have go-routine and
	//2nd then we will decrement the counter after each go-routine completed its execution ,
	//3rd in end we will call waitgroup wait() method to check final count is 0 or not if 0 then it will make main fun to complete its execution
	//waitGroupCounter := &sync.WaitGroup{}// wait group is simply a counter
	//waitGroupCounter.Add(1)// increment wait group counter,
	//go foo2("sheep", &waitGroupCounter)
	//waitGroupCounter.Wait()


	//solution 2: in above we are passing counter as param, so write anonymous function and add decrement count inside
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {// anonymous function
		foo("sheep")
		waitGroup.Done()// decrements counter
	}()
	waitGroup.Wait()//wait to see final count value is zero or not




}

func foo(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Second * 1) // sleep for 1 sec
	}
}

func foo2(str string, counter **sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Second * 1) // sleep for 1 sec
	}
	(*counter).Done() // after completed this method execution decrement wait group counter, DOne() means decrement
}
