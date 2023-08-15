package hltv

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

const HLTV_URL = "https://www.hltv.org"

func Fetch(url string) (resp *cycletls.Response, err error) {
	client := cycletls.Init()
	response, err := client.Do(HLTV_URL+url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
	}, "GET")
	if err != nil {
		fmt.Println("Request Failed: " + err.Error())
		return nil, err
	}
	if response.Status != 200 {
		return nil, errors.New("Request Failed: " + strconv.Itoa(response.Status))
	}
	return &response, nil
}
