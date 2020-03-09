package entities

import (
	"fmt"

	"github.com/spf13/viper"
)

const recognizerURLFormat = "http://%s:%s/api/v1/analyzer/recognizers/%s"

func constructRecognizerURL(recognizerName string) string {
	var ip = viper.GetString("presidio_ip")
	var port = viper.GetString("presidio_port")

	return fmt.Sprintf(recognizerURLFormat,
		ip,
		port,
		recognizerName)
}

// CreateRecognizer creates a new recognizer
func CreateRecognizer(httpClient httpClient, name string, fileContentStr string) {
	url := constructRecognizerURL(name)
	restCommand(httpClient, create, url, fileContentStr, "")
}

// UpdateRecognizer updates an existing recognizer
func UpdateRecognizer(httpClient httpClient, name string, fileContentStr string) {
	url := constructRecognizerURL(name)
	restCommand(httpClient, update, url, fileContentStr, "")
}

// DeleteRecognizer deletes an existing recognizer
func DeleteRecognizer(httpClient httpClient, projectName string, actionName string, name string) {
	url := constructRecognizerURL(name)
	restCommand(httpClient, delete, url, "", "")
}

// GetRecognizer retrieved an existing recognizer, can be logged or saved to a file
func GetRecognizer(httpClient httpClient, projectName string, actionName string, name string, outputFilePath string) {
	url := constructRecognizerURL(name)
	restCommand(httpClient, get, url, "", outputFilePath)
}