package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dhaliwal-h/go-consume-pexelsapi/models"
	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error 1: ", err)
	}
	TOKEN := os.Getenv("PEXELS_TOKEN")
	c := NewClient(TOKEN)
	result, err := c.SearchPhotos("waves", 15, 1)
	if err != nil {
		log.Fatal("Error 2: ", err)
	}
	if result.Page == 0 {
		fmt.Println("search result wrong")
	}
	fmt.Printf("Result is %v\n", result)
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}

}

func (c *Client) SearchPhotos(query string, perPage int, page int) (models.SearchResult, error) {
	var sr models.SearchResult
	reqUrl := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", reqUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sr, err
	}

	err = json.Unmarshal(data, &sr)
	if err != nil {
		log.Fatal(err)
	}
	return sr, nil
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	res, err := c.hc.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	times, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return res, err
	} else {
		c.RemainingTimes = int32(times)
	}
	return res, nil
}
