package httpreq

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lkzcover/sandbox-user-encrypted-data/data"
)

// Получение списка доступных водителей
func GetDrierList() (respData []data.Driver, err error) {
	resp, err := http.Get("http://127.0.0.1/listDriver")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &respData)

	return
}
