package main

import (
	"fmt"
)

// GOroutines ...

func main() {
	channelEx := make(chan int)

	go func() {
		defer close(channelEx)
		for i := 1; i < 2000; i++ {
			channelEx <- i
			fmt.Println("Jogando no channel o valor: ", i)
		}
	}()

	go worker(channelEx, 2)
	go worker(channelEx, 3)
	worker(channelEx, 1)

}

func worker(channelEx chan int, workerID int) {

	for chann := range channelEx {

		valInChannel, ok := <-channelEx
		if !ok {
			println("Saiu do for do worker")
			break
		}
		fmt.Printf("Esvaziando channel com valor: %v, executando worker: %v, quantidade de msg: %v \n", valInChannel, workerID, chann)
		//time.Sleep(time.Millisecond * 200)
	}

}
