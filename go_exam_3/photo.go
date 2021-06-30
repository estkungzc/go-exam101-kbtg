package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"sync"
	"time"
)

type Photo struct {
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`

	GoRoutinesId int
	EmployeeId   string
	Error        error
}

func GetDataFromPhotoAPI(id int, wg *sync.WaitGroup, pipelineOut chan Photo, empId string) {
	url := fmt.Sprintf("%s/%d", "https://jsonplaceholder.typicode.com/photos", id)

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error json unmarshall : ", err)
		pipelineOut <- Photo{
			Id: id,
			GoRoutinesId: id,
			Error: err,
		}

		return
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("error call ", err)
		pipelineOut <- Photo{
			Id: id,
			GoRoutinesId: id,
			Error: err,
		}
		return
	}

	defer resp.Body.Close()

	var respData Photo
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		fmt.Println("error decode json user :", err)
		pipelineOut <- Photo{
			Id: id,
			GoRoutinesId: id,
			Error: err,
		}
		return
	}

	respData.GoRoutinesId = id
	respData.EmployeeId = empId
	pipelineOut <- respData
}

func WriteXlsxPhoto(fileName string, sheetName string, photoArr []Photo) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	//var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet(sheetName)
	if err != nil {
		fmt.Println("error cannot add sheet name")
		return
	}

	row = sheet.AddRow()
	row.AddCell().SetValue("EmployeeId")
	row.AddCell().SetValue("RoutineId")
	row.AddCell().SetValue("AlbumId")
	row.AddCell().SetValue("Id")
	row.AddCell().SetValue("Title")
	row.AddCell().SetValue("Url")
	row.AddCell().SetValue("ThumbnailUrl")
	row.AddCell().SetValue("Error")

	for _, po := range photoArr {
		row = sheet.AddRow()
		row.AddCell().SetValue(po.EmployeeId)
		row.AddCell().SetValue(po.GoRoutinesId)
		row.AddCell().SetValue(po.AlbumId)
		row.AddCell().SetValue(po.Id)
		row.AddCell().SetValue(po.Title)
		row.AddCell().SetValue(po.Url)
		row.AddCell().SetValue(po.ThumbnailUrl)
		row.AddCell().SetValue(po.Error)

	}
	filePath := fmt.Sprintf("%s.xlsx", fileName)
	err = file.Save(filePath)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
