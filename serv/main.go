package main

import (
	"serv/pkg"
)

func main() {
	device := &pkg.Device{Name: "Device-1"}
	updateSvc := &pkg.UpdateDataService{Name: "Update-1"}
	dataSvc := &pkg.DataService{}
	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)
	data := &pkg.Data{}
	device.Execute(data)

}
