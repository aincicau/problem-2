package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pb2/entity"
	"strconv"
)

func CanDrive(rw http.ResponseWriter, r *http.Request) {
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

	typeOfVeh := objmap["type"].(string)

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

	value := vehicle.CanDrive()
	valueS := strconv.FormatBool(value)

	rw.Write([]byte(valueS))
}
