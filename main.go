package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func getStatus(value int, thresholds []int, statuses []string) string {
	for i, threshold := range thresholds {
		if value <= threshold {
			return statuses[i]
		}
	}
	return statuses[len(statuses)-1]
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1
		data := Data{Water: water, Wind: wind}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("POST Response:", string(body))

		waterStatus := getStatus(water, []int{5, 8}, []string{"aman", "siaga", "bahaya"})
		windStatus := getStatus(wind, []int{6, 15}, []string{"aman", "siaga", "bahaya"})

		fmt.Printf("status water : %s\n", waterStatus)
		fmt.Printf("status wind : %s\n", windStatus)

		time.Sleep(15 * time.Second)
	}
}
