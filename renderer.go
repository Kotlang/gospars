package gospars

import (
	"net/http"
	"io/ioutil"
)

func getTemplate(path string, callback func(error, string)) {
	go func() {
		resp, err := http.Get(path)
		if err != nil {
			callback(err, "")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		callback(nil, string(body[:]))
	}()
}
