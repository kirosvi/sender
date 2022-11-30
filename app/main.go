// Service to send messages by webhooks to telegram
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

var configData = make(map[string]ConfigParams)
var (
	configPath    string
	templatesPath string
	port          string
	address       string
	httpMessage   string
	httpResponse  int
)

// ConfigParams is struct for params from config.yaml
type ConfigParams struct {
	ChatID   int64  `yaml:"chatid"`
	Token    string `yaml:"token"`
	Template string `yaml: "template"`
}

// Message is struct for params for message that sent to telegram
type Message struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func init() {
	flag.StringVar(&configPath, "c", "config.yaml", "Path to a config file")
	flag.StringVar(&templatesPath, "t", "templates", "Path to a template files")
	flag.StringVar(&address, "address", "0.0.0.0", "address to bind")
	flag.StringVar(&port, "port", "8080", "port number")
}

func main() {
	flag.Parse()

	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Problem reading configuration file: %v", err)
	}

	err = yaml.Unmarshal(configContent, &configData)
	if err != nil {
		log.Fatalf("Error parsing configuration file: %v", err)
	}

	addressToBind := fmt.Sprintf("%s:%s", address, port)
	log.Println("Bind Address:", addressToBind)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", GETHandling)
	router.POST("/alert/:chatname", POSTHandling)
	router.Run(addressToBind)
}

// GETHandling manage get requests, basicaly its healthcheck
func GETHandling(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// POSTHandling manage post requests
func POSTHandling(c *gin.Context) {

	chatname := c.Param("chatname")
	action := c.Query("action")

	rawData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
	}

	status, responseMsg := executeAlerting(chatname, action, rawData)

	r := string(responseMsg)

	if action == "dry-run" {
		r = "Dry-run output: " + r
	}

	log.Println(status, r)
	c.String(http.StatusOK, "%s", r)

}

func executeAlerting(chatname string, action string, rawData []byte) (int, string) {
	parsedData := importJSON(rawData)

	renderedData := renderTemplate(chatname, parsedData)

	msg := &Message{
		ChatID: configData[chatname].ChatID,
		Text:   renderedData}

	url := getURL(configData[chatname].Token)

	status, responseMsg := SendMessage(url, msg, chatname, action)

	return status, responseMsg
}

func loadTemplate(templateName string) *template.Template {

	pathFile := filepath.Join(templatesPath, templateName)

	tmpl, err := template.New("template").Funcs(sprig.FuncMap()).ParseFiles(pathFile)

	if err != nil {
		log.Printf("Problem reading parsing template file: %v", err)
	} else {
		log.Printf("Load template file: %s", templateName)
	}

	return tmpl
}

func renderTemplate(chatname string, messageData map[string]interface{}) string {

	tmpl := loadTemplate(configData[chatname].Template)

	var tpl bytes.Buffer
	_ = tmpl.ExecuteTemplate(&tpl, configData[chatname].Template, messageData)

	result := tpl.String()

	return result

}

func getURL(token string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
}

// SendMessage function to send message based on chat name and state of action parametr
func SendMessage(url string, message *Message, chatname string, action string) (int, string) {

	payload, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	if action != "dry-run" {
		response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Println(err)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}

		if response.StatusCode != http.StatusOK {
			log.Printf("failed to send successful request. Status was %q", response.Status)
			return response.StatusCode, response.Status
		}
		responseMsg := string(body)
		return response.StatusCode, responseMsg

	}
	return 200, string(payload)

}

func importJSON(data []byte) map[string]interface{} {

	m := map[string]interface{}{}

	if err := json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}

	return m

}
