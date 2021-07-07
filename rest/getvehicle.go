package rest

import (
	"encoding/json"
	"net/http"
	"pb2/entity"
)

func GetMyVehicle(rw http.ResponseWriter, r *http.Request) {
	vehtype := r.URL.Query().Get("type")
	var vehicle entity.Vehicle
	switch vehtype {
	case "car":
		vehicle = entity.Car{Name: "BMW", Type: "Car", Model: "2010"}
	case "bus":
		vehicle = entity.Bus{Name: "Mercedes", Type: "Bus", Model: "2010"}
	case "bike":
		vehicle = entity.Bike{Name: "Merida", Type: "Bike", Model: "2018"}
	default:
		rw.Write([]byte("Internal error"))
		return
	}

	vehicleBytes, err := json.Marshal(vehicle)
	if err != nil {
		rw.Write([]byte("Internal error"))
		return
	}

	rw.Write(vehicleBytes)
}
