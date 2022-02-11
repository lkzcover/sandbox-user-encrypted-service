package main

import (
	"log"
	"sync"

	"github.com/lkzcover/sandbox-user-encrypted-data/data"
	"github.com/nats-io/nats.go"
)

func main() {

	type DriverDevice struct {
		Key   []byte
		Login string
		Data  data.Driver
	}

	// Список активных устройств
	drivers := []DriverDevice{
		{
			Login: "L1",
			Key:   []byte("driver1SecretKey"),
		},
		{
			Login: "L3",
			Key:   []byte("driver3SecretKey"),
		},
	}

	// NATS для имитации канала связи между устройствами
	_, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}

	startChan := make(chan struct{}, len(drivers))

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Имитируем авторизацию
	for key, driver := range drivers {

		go func(d DriverDevice, k int) {

			startChan <- struct{}{}
		}(driver, key)

	}

	for range drivers {
		<-startChan
	}

	log.Println("Success start Driver devise simulator")

	wg.Wait()
}
