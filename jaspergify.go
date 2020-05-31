// Package jaspergify contains an HTTP Cloud Function.
package jaspergify

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Run(w http.ResponseWriter, r *http.Request) {

	var response struct {
		Challenge string `json:"challenge"`
	}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		fmt.Fprint(w, "ukMkofynvX8SiGDxp1APZBXG7TorVFVqjaxKzsPnjJY64XIzl7TT")
		return
	}
	if response.Challenge == "" {
		fmt.Fprint(w, "ukMkofynvX8SiGDxp1APZBXG7TorVFVqjaxKzsPnjJY64XIzl7TT")
		return
	}
	fmt.Fprint(w, html.EscapeString(response.Challenge))
}
