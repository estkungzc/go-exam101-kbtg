package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	idRange := flag.Int("id_range", 0, "range id to call api")
	empId := flag.String("employee_id", "", "employee id")

	routines := flag.Int("routines", 10, "routines threads")

	flag.Parse()

	if *idRange <= 0 || *idRange >= 250 {
		panic(fmt.Errorf("id_range must be in 0 < id_range < 250"))
	}

	if *empId == "" {
		panic(fmt.Errorf("employee_id is required"))
	}

	if *routines <= 0 || *routines > 50 {
		*routines = 10
	}

	pipelineOut := make(chan Photo)

	defer close(pipelineOut)

	var wg sync.WaitGroup
	wg.Add(*idRange)

	var photoArr []Photo
	//Channel get output
	go func() {
		for v := range pipelineOut {
			fmt.Println("data in channel :", v)
			photoArr = append(photoArr, v)
			wg.Done() // for handle GetDataFromPhotoAPI success before append data (change from GetDataFromPhotoAPI)
		}
	}()

	//Loop to feed input to worker
	for i := 1; i <= *idRange; i++ {
		go GetDataFromPhotoAPI(i, &wg, pipelineOut, *empId)
	}

	wg.Wait()
	fmt.Println("call api complete ")
	//time.Sleep(1*time.Second)
	fmt.Println(len(photoArr))
	WriteXlsxPhoto(*empId,"gobasic_exam3", photoArr)
}
