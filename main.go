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

// Data struct untuk menyimpan nilai water dan wind
type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// getStatus mengembalikan status berdasarkan nilai dan ambang batas yang diberikan
func getStatus(value int, thresholds []int, statuses []string) string {
	for i, threshold := range thresholds {
		if value <= threshold {
			return statuses[i]
		}
	}
	return statuses[len(statuses)-1]
}

func main() {
	// URL untuk mengirim POST request
	url := "https://jsonplaceholder.typicode.com/posts"

	// Mulai loop yang akan berjalan setiap 15 detik
	for {
		// Generate nilai acak untuk water dan wind antara 1 dan 100
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Buat objek data dengan nilai water dan wind yang dihasilkan
		data := Data{Water: water, Wind: wind}

		// Marshaling data ke dalam format JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		// Mengirim POST request dengan data JSON
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer resp.Body.Close()

		// Membaca response dari POST request
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("POST Response:", string(body))

		// Mendapatkan status water dan wind berdasarkan nilai yang dihasilkan
		waterStatus := getStatus(water, []int{5, 8}, []string{"aman", "siaga", "bahaya"})
		windStatus := getStatus(wind, []int{6, 15}, []string{"aman", "siaga", "bahaya"})

		// Menampilkan status water dan wind
		fmt.Printf("status water : %s\n", waterStatus)
		fmt.Printf("status wind : %s\n", windStatus)

		// Tunggu 15 detik sebelum melanjutkan loop
		time.Sleep(15 * time.Second)
	}
}
