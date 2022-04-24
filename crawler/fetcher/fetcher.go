package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get("http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, errors.New(fmt.Sprintf("wrong status: %d", resp.StatusCode))
	}
	return ioutil.ReadAll(resp.Body)

}
