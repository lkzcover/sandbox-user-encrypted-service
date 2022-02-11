package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lkzcover/sandbox-user-encrypted-data/lib/httpreq"

	"github.com/lkzcover/sandbox-user-encrypted-data/data"
)

func main() {

	DriverListDB := make(map[string]data.Driver)

	// Имитация БД
	DriverListDB = map[string]data.Driver{
		"L1": {
			Name:      "Name1",
			Rating:    10,
			Cars:      "Car1",
			Encrypted: "sKXNF90OWY/ZfgZfyMfA+sCXpx3Zfd5X5zq+t5a8++c",
		},
		"L2": {
			Name:      "Name2",
			Rating:    20,
			Cars:      "Car2",
			Encrypted: "msbhDiQm/BdSJR0Uw0UAjwwaHwz+2ka+k1Xmq38TMpQ",
		},
		"L3": {
			Name:      "Name3",
			Rating:    30,
			Cars:      "Car3",
			Encrypted: "KuXDSrghTr3isFl0b+VQE0kes5u2WetM4E8vOj20Pao",
		},
	}

	http.HandleFunc("/driverCompletedOrder", func(resp http.ResponseWriter, req *http.Request) {

		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)

			return
		}

		defer req.Body.Close()

		var orderReq httpreq.DriveCompletedOrderReq
		if err := json.Unmarshal(reqBody, &orderReq); err != nil {
			resp.WriteHeader(http.StatusBadRequest)

			return
		}

		dreiver, ok := DriverListDB[orderReq.Login]
		if !ok {
			resp.WriteHeader(http.StatusBadRequest)

			return
		}

		dreiver.Addr = orderReq.Addr
		dreiver.OrderID = &orderReq.OrderID

		DriverListDB[orderReq.Login] = dreiver

		log.Printf("Login Driver %s", orderReq.Login)

	})

	http.HandleFunc("/listDriver", func(resp http.ResponseWriter, req *http.Request) {
		var respData []data.Driver

		for _, driver := range DriverListDB {
			if driver.OrderID != nil {
				respData = append(respData, driver)
			}
		}

		respBody, err := json.Marshal(respData)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)

			return
		}

		_, _ = resp.Write(respBody)
	})

	http.HandleFunc("/startClientOrder", func(resp http.ResponseWriter, req *http.Request) {
	})

	http.HandleFunc("/startDriverOrder", func(resp http.ResponseWriter, req *http.Request) {

	})

	log.Println("Start")

	http.ListenAndServe(":80", nil)

}
