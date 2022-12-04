package helpers

import (
	"errors"
	"io"
	"net/http"
)

func DownloadFromUrl(URL string) (error, []byte) {
	response, err := http.Get(URL)
	if err != nil {
		return err, nil
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code"), nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}
