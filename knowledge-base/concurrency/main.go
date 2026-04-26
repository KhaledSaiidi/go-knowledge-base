package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mutex  sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}
	return orders
}

func processOrders(
	orderChan <-chan *Order, // receive only channel,
	processedOrderChan chan<- *Order, // send only channel,
	wg *sync.WaitGroup,
) {
	defer func() {
		wg.Done()
		close(processedOrderChan)
	}()
	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		order.Status = "Processed"
		processedOrderChan <- order
	}
}

func updateOrderStatuses(order *Order) {
	order.mutex.Lock()
	defer order.mutex.Unlock()
	time.Sleep(
		time.Duration(rand.Intn(500)) *
			time.Millisecond,
	)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]

	order.Status = status
	fmt.Printf("Updated order %d, Status: %s\n", order.ID, order.Status)

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func reportOrdersStaus(orders []*Order) {
	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %s\n",
			order.ID, order.Status,
		)
	}
	fmt.Println("--------------------------------------")
}

func main() {
	wg := sync.WaitGroup{}

	// orderChan := make(chan *Order) 		-> Unbuffered Channel No capaciy and Synchronous
	orderChan := make(chan *Order, 20) //   -> Buffered Channel has 20 capaciy and Asynchronous
	processedOrderChan := make(chan *Order, 20)
	var orders []*Order

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(orderChan)
		for _, order := range generateOrders(20) {
			orderChan <- order
			orders = append(orders, order)
		}

		fmt.Println("Done with generating orders")
	}()
	wg.Add(1)
	go processOrders(orderChan, processedOrderChan, &wg)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case processedOrder, ok := <-processedOrderChan:
				if !ok {
					fmt.Println("Processing channel closed")
					return
				}
				fmt.Printf("Processed order %d with status: %s\n", processedOrder.ID, processedOrder.Status)
			case <-time.After(10 * time.Second):
				fmt.Println("No processed orders received in the last second")
				return
			}
		}
	}()
	wg.Wait()
	for _, order := range orders {
		updateOrderStatuses(order)
	}
	reportOrdersStaus(orders)

	fmt.Println("All operations completed. Exiting.")
	fmt.Println(totalUpdates)

}
