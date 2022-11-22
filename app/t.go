package main

import (
    "fmt"
    "flag"
//    "net/http"
    "io/ioutil"
    "log"
//    "path"
    "path/filepath"
    "bytes"
//    "os"
    "text/template"

//    "github.com/gin-gonic/gin"
    "gopkg.in/yaml.v3"
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
    ChatId       string
    Token        string
    Template     string
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
    fmt.Println(configData["pp_core"].Token)

    tmpl := loadTemplate(configData["rundeck"].Template)
//    tmpl.Execute(os.Stdout, struct {
//            Name string
//        }{"Jane Doe"})
//
//
    var tpl bytes.Buffer
    _ = tmpl.Execute(&tpl, nil)

    result := tpl.String()
    fmt.Println(result)
////////////

//    addressToBind := fmt.Sprintf("%s:%s", address, port)
//    fmt.Println("Bind Address:", addressToBind)
//
//    router := gin.Default()
//
//    router.GET("/ping", GET_Handling)
//    router.POST("/alert/:chatname", POST_Handling)
//    router.Run(addressToBind)
}
func loadTemplate(templateName string) *template.Template {

    pathFile := filepath.Join(templatesPath, templateName)
    fmt.Println(pathFile)

    tmpl, err := template.ParseFiles(pathFile)

    if err != nil {
    log.Printf("Problem reading parsing template file: %v", err)
    } else {
    log.Printf("Load template file:%s", templateName)
    }

    return tmpl
}

//func GET_Handling(c *gin.Context) {
//    c.JSON(http.StatusOK, gin.H{
//      "message": "pong",
//    })
//}
//
//func POST_Handling(c *gin.Context) {
//    b := c.FullPath() == "/alert/:chatname" // true
//
//    chatname := c.Param("chatname")
//
//    messageData, err := ioutil.ReadAll(c.Request.Body)
//    if err != nil {
//        log.Print(err)
//    }
//
//    send_data(chatname, messageData)
//
//    c.String(http.StatusOK, "%t", b)
//
//}
//
//func send_data(chatname string, messageData []byte) {
//
//    log.Printf("Bot alert post: %s", chatname)
//    log.Printf("data: %s", messageData)
//
//    tmpl := loadTemplate(configData["rundeck"].Template)
//
//    tmpl.Execute(os.Stdout, nil)
//
//
//    var doc bytes.Buffer
//    tmpl.Execute(&doc, nil)
//    s := doc.String()
//
//    fmt.Printf("template:",s)
//
//    template, err := template.ParseFiles("templates/rundeck.tpl")
//    // Capture any error
//    if err != nil {
//    log.Fatalln(err)
//    }
//    // Print out the template to std
//    template.Execute(os.Stdout, nil)
//
//
//}
