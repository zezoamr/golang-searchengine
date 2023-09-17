package parsing

import (
	"fmt"
	"io"
	"net/http"
)

// getPage retrieves the HTML content from the specified URL.
//
// It takes a single parameter:
// - url: a string representing the URL to fetch the HTML from.
//
// The function returns a string containing the HTML content.
func getPage(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("url is empty")
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		//fmt.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	// following reads all at once
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	formattedString := string(body)
	return formattedString, nil

	// following reads in chunks
	// buf := make([]byte, 1024)
	// for {
	// 	n, err := resp.Body.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		log.Fatalln("Error:", err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	fmt.Print(string(buf[:n]))
	// }
}
