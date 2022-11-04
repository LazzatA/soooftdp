package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (upd *DataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("Data not update")
		return
	}
	fmt.Println("Data save")

}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc
}
