package httpreq

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DriveCompletedOrderReq struct {
	Login   string
	Addr    string
	OrderID string
}

func DriveCompletedOrder(login string, addr, orderID string) error {
	reqData := DriveCompletedOrderReq{
		Login:   login,
		Addr:    addr,
		OrderID: orderID,
	}

	reqDataBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	_, err = http.Post("http://127.0.0.1:80/driverCompletedOrder", "application/json",
		bytes.NewBuffer(reqDataBody))

	if err != nil {
		return err
	}

	return nil
}
