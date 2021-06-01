package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Requset(urls string, bodyReader io.Reader, token string) (ress string, err error, errMessage RequestErr) {
	req, err := http.NewRequest("POST", urls, bodyReader)
	if err != nil {
		fmt.Printf("http.NewRequest failed,err: %v \n", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if len(token) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do failed,err: %v \n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll failed,err: %v \n", err)
		return
	}
	if res.StatusCode != 201 && res.StatusCode != 204 {
		requestErr := RequestErr{}
		json.Unmarshal(body, &requestErr)
		err = errors.New(requestErr.Message)
		errMessage = requestErr
	}
	ress = string(body)
	return
}

//RequestPost Post请求
func RequestPostStr(urls string, dataStr string, token string) (ress string, err error, errMessage RequestErr) {

	return Requset(urls, strings.NewReader(dataStr), token)
}

//RequestPost Post请求
func RequestPost(urls string, data map[string]interface{}, token string) (ress string, err error, errMessage RequestErr) {
	dataStr, _ := json.Marshal(data)

	return Requset(urls, strings.NewReader(string(dataStr)), token)

}

func RequestGet(urls string, token string) (ress string, err error, errMessage RequestErr) {

	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		fmt.Printf("http.NewRequest failed,err: %v \n", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if len(token) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do failed,err: %v \n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll failed,err: %v \n", err)
		return
	}
	if res.StatusCode != 200 {
		requestErr := RequestErr{}
		json.Unmarshal(body, &requestErr)
		err = errors.New(requestErr.Message)
		errMessage = requestErr
	}
	ress = string(body)
	return
}
