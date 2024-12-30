package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	fileName := "downloaded_" + url[strings.LastIndex(url, "/")+1:]
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Downloaded:", fileName)
}

func main() {
	urls := []string{
		"https://filesamples.com/sample_640%C3%97426.jpg",
		"https://filesamples.com/sample1.heif",
		"https://filesamples.com/sample1.webp",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go downloadFile(url, &wg)
	}

	wg.Wait()

	fmt.Println("All downloads complete.")
}
