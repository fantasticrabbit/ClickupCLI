package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func FormatJSON(jsonresp []byte) string {
	var formattedJSON bytes.Buffer
	if err := json.Indent(&formattedJSON, jsonresp, "", "    "); err != nil {
		log.Fatalln(err)
	}
	return formattedJSON.String()
}

func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".clickup")
}

func GetConfigFile() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".clickup", "config.yaml")
}
