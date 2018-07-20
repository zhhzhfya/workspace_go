package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 100)
	chSend := make(chan int)
	chConsume := make(chan int)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func(ch, quit chan int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("send to ch panic.===", err)
			}
		}()

		i := 0
		for {
			select {
			case ch <- i:
				fmt.Println("send", i)
				time.Sleep(time.Second)
				i++
			case <-quit:
				fmt.Println("send quit.")
				return
			}

		}
	}(ch, chSend)

	go func(ch, quit chan int) {
		wg.Add(1)
		for {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println("read1", i)
					time.Sleep(time.Second * 2)
				} else {
					fmt.Println("close ch1.")
				}

			case <-quit:
				for {
					select {
					case i, ok := <-ch:
						if ok {
							fmt.Println("read2", i)
							time.Sleep(time.Second * 2)
						} else {
							fmt.Println("close ch2.")
							goto L
						}
					}
				}
			L:
				fmt.Println("consume quit.")
				wg.Done()
				return

			}
		}
	}(ch, chConsume)

	<-sc

	close(ch)
	fmt.Println("close ch ")
	close(chSend)
	close(chConsume)
	wg.Wait()
}
