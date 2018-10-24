package function

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
)

type HomepageTokens struct {
	Dark bool
}

// Handle a serverless request
func Handle(req []byte) string {

	if os.Getenv("Http_Path") == "/cloud.png" {
		image, _ := ioutil.ReadFile("./cloud.png")
		output := base64.StdEncoding.EncodeToString(image)
		return string(output)
	}

	dark := true

	darkVal := os.Getenv("dark")
	valOut, parseErr := strconv.ParseBool(darkVal)
	if parseErr == nil {
		dark = valOut
	}

	var err error
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		return fmt.Sprintf("Internal server error with homepage: %s", err.Error())
	}
	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, HomepageTokens{
		Dark: dark,
	})
	if err != nil {
		return fmt.Sprintf("Internal server error with homepage template: %s", err.Error())
	}

	return string(tpl.Bytes())

}
