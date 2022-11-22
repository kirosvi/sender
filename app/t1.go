package main

import (
    "fmt"
    "reflect"
    "os"
    "encoding/json"
    "html/template"
    "github.com/Masterminds/sprig/v3"
)


func main() {
//    t, _ := template.ParseFiles("templates/rundeck.tpl")
    t, err := template.New("1").Funcs(sprig.FuncMap()).ParseFiles("templates/rundeck.tpl")
    if err != nil {
        panic(err)
    }

    data := `test message`
    jsondata := `{"trigger": "failure","status": "failed","executionId": 5102,"execution": {"id": 5102,"href": "http://rundeck.infra.ppdev.ru/project/DevOps/execution/show/5102","permalink": null,"status": "failed","project": "DevOps","executionType": "user","user": "admin","date-started": {"unixtime": 1669198910364,"date": "2022-11-23T10:21:50Z"},"job": {"id": "6cd8db31-82a5-49d3-a6ce-ea147755045a","name": "test-tg","group": "test","project": "DevOps","description": "","href": "http://rundeck.infra.ppdev.ru/api/41/job/6cd8db31-82a5-49d3-a6ce-ea147755045a","permalink": "http://rundeck.infra.ppdev.ru/project/DevOps/job/show/6cd8db31-82a5-49d3-a6ce-ea147755045a"},"description": "exit 1","argstring": null,"serverUUID": "a14bc3e6-75e8-4fe4-a90d-a16dcc976bf6"}}`

    fmt.Println(reflect.TypeOf(data))
    fmt.Println(reflect.TypeOf(jsondata))


    m := map[string]interface{}{}

    if err := json.Unmarshal([]byte(jsondata), &m); err != nil {
        fmt.Println(err)
    }

    fmt.Println(reflect.TypeOf(m))

    if err := t.ExecuteTemplate(os.Stdout,"rundeck.tpl", m); err != nil {
        fmt.Println(err)
    }
}
