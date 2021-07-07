package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb2/entity"
)

func PostMyVehicle(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if err != nil {
		rw.Write([]byte("Internal error"))
		return
	}

	var objmap map[string]interface{}
	err = json.Unmarshal(bodyBytes, &objmap)

	if err != nil {
		rw.Write([]byte("Internal error"))
		return
	}

	var vehicle entity.Vehicle

	typeOfVeh, ok := objmap["type"].(string)
	if ok {
		switch typeOfVeh {
		case "Car":
			vehicle = entity.Car{Name: "BMW", Type: "Car", Model: "2010"}
		case "Bus":
			vehicle = entity.Bus{Name: "Mercedes", Type: "Bus", Model: "2010"}
		case "Bike":
			vehicle = entity.Bike{Name: "Merida", Type: "Bike", Model: "2018"}
		default:
			rw.Write([]byte("Internal error"))
			return
		}
	}

	fmt.Println(vehicle)
	rw.Write(bodyBytes)
}
