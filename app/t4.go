package main

import (
    "bytes"
    "fmt"
    "reflect"
    "net/http"
    "flag"
    "os"
    "io/ioutil"
    "log"
    "encoding/json"
    "text/template"
    "gopkg.in/yaml.v3"
    "github.com/Masterminds/sprig/v3"
)

var configData = make(map[string]ConfigParams)
var (
    configPath  string
    templatesPath string
    port string
    address string
    http_message string
    http_response int
)

type ConfigParams struct {
    ChatId       int64
    Token        string
    Template     string
}

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

     for k, v := range configData {

          fmt.Println(k, v)
     }

    fmt.Println(configData["rundeck"].ChatId)
    fmt.Println(configData["ppcore"].ChatId)
    t, err := template.New("1").Funcs(sprig.FuncMap()).ParseFiles("templates/ppcore.tpl")
    if err != nil {
        fmt.Println(err)
    }
    data := `test message`
    jsondata := `{"alerts":[{"status":"resolved","labels":{"custom_label":"custom_label1","alertname":"testalert","severity":"info","kubernetes_name":"prometheusexporter","kubernetes_namespace":"namespace","kubernetes_node":"node"},"annotations":{"description":"testdescription","summary":"Testalert"},"generatorURL":"http://prometheus.io"},{"status":"firing","labels":{"custom_label2":"custom_label2","alertname":"testalert2","severity":"error","kubernetes_name":"prometheusexporter","kubernetes_namespace":"namespace","kubernetes_node":"node"},"annotations":{"description":"testdescription2"},"generatorURL":"http://prometheus.io"}]}`

    fmt.Println(reflect.TypeOf(data))
    fmt.Println(reflect.TypeOf(jsondata))

    m := map[string]interface{}{}

    if err := json.Unmarshal([]byte(jsondata), &m); err != nil {
        fmt.Println(err)
    }

    fmt.Println(reflect.TypeOf(m))

    if err := t.ExecuteTemplate(os.Stdout,"ppcore.tpl", m); err != nil {
        fmt.Println(err)
    }
    var tpl bytes.Buffer
    _ = t.ExecuteTemplate(&tpl, configData["ppcore"].Template, m)


    msg := &Message{
        ChatID : configData["ppcore"].ChatId,
        Text: tpl.String()}

    url := getUrl(configData["ppcore"].Token)
    SendMessage(url, msg)
}

func getUrl(token string) string {
  return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
}

func SendMessage(url string, message *Message) error {

    payload, err := json.Marshal(message)
    if err != nil {
      return err
    }

    response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err != nil {
      return err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
      return err
    }

    if response.StatusCode != http.StatusOK {
      return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
    } else {
        fmt.Printf("Response JSON: %s", string(body))
    }

    return nil
}


