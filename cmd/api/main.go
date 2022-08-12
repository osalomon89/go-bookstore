package main

import (
	"fmt"
	"sync"

	inventorysrv "github.com/osalomon89/go-bookstore/inventory"
)

func main() {
	fmt.Println("Hello World")
	wg := sync.WaitGroup{}

	wg.Add(1)
	go inventorysrv.InventoryService(&wg)
	wg.Wait()
}
