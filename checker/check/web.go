package check

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func GetWebCode(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	return string(contents), nil
}

func GetWebQuery(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	return goquery.NewDocumentFromReader(res.Body)
}
