package pylon

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"pylon/frd"
)

func Feed(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")
	id := r.URL.Query().Get("id")
	format := r.URL.Query().Get("format")
	if strings.ToUpper(format) == "FRD" {
		if period != "m5" {
			fmt.Println("please support other period")
		}
		b, _ := ioutil.ReadAll(r.Body)
		frd.Feed(id, string(b))
	}
}
