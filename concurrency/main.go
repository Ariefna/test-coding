package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	DataResult := getData(url)
	people1 := Graduate{}
	jsonErr := json.Unmarshal(DataResult, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	var a []string
	for i := 0; i < len(people1.Result.Records); i++ {
		files1, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files1 {
			a = append(a, f.Name())
		}

		if contains(a, people1.Result.Records[i].Year+".csv") == true {
			file, err := os.OpenFile(people1.Result.Records[i].Year+".csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			var data [][]string
			data = append(data, []string{
				strconv.Itoa(people1.Result.Records[i].Ide),
				people1.Result.Records[i].Sex,
				people1.Result.Records[i].No,
				people1.Result.Records[i].Course,
				people1.Result.Records[i].Year,
			})
			w := csv.NewWriter(file)
			if err := w.WriteAll(data); err != nil {
				log.Fatalln("error writing record to file", err)
			}
		} else {
			file, err := os.Create(people1.Result.Records[i].Year + ".csv")
			defer file.Close()
			if err != nil {
				log.Fatalln("failed to open file", err)
			}
			w := csv.NewWriter(file)
			defer w.Flush()
			row := []string{
				strconv.Itoa(people1.Result.Records[i].Ide),
				people1.Result.Records[i].Sex,
				people1.Result.Records[i].No,
				people1.Result.Records[i].Course,
				people1.Result.Records[i].Year,
			}
			if err := w.Write(row); err != nil {
				log.Fatalln("error writing record to file", err)
			}
		}
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func getData(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	return data
}

// Model section
type Graduate struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}

type Result struct {
	Resource_id string    `json:"resource_id"`
	Fields      []Fields  `json:"fields"`
	Records     []Records `json:"records"`
}

type Fields struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Records struct {
	Ide    int    `json:"_id"`
	Sex    string `json:"sex"`
	No     string `json:"no_of_graduates"`
	Course string `json:"type_of_course"`
	Year   string `json:"year"`
}

// Worker section

func worker(id int, jobs <-chan int, results chan<- int, url string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		getData(url)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func job(workerNum int, totalJob int, url string) {
	jobs := make(chan int, totalJob)
	results := make(chan int, totalJob)

	for i := 0; i < workerNum; i++ {
		go worker(i, jobs, results, url)
	}

	for i := 0; i < totalJob; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < totalJob; i++ {
		<-results
	}
}
