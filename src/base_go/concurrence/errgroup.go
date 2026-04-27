package concurrence

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func Errgroup() {
	// var group sync.WaitGroup

	// group.Go(func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println("hello")
	// })

	// group.Go(func() {
	// 	time.Sleep(50 * time.Millisecond)
	// 	fmt.Println("hello")
	// })

	// group.Go(func() {
	// 	time.Sleep(30 * time.Millisecond)
	// 	fmt.Println("hello")
	// })

	// group.Wait()
	var group errgroup.Group

	group.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("hello")
		return nil
	})

	group.Go(func() error {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("hello")
		return errors.New("Sorry")
	})

	group.Go(func() error {
		time.Sleep(30 * time.Millisecond)
		fmt.Println("hello")
		return errors.New("ops")
	})
	err := group.Wait()
	fmt.Println("wait result", err) // wait result ops 返回error是最先发生的error ops
}
