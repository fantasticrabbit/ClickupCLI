package internal

import (
	"bytes"
	"encoding/json"
	"log"
)

func FormatJSON(jsonresp []byte) string {
	var formattedJSON bytes.Buffer
	if err := json.Indent(&formattedJSON, jsonresp, "", "    "); err != nil {
		log.Fatalln(err)
	}
	return formattedJSON.String()
}
