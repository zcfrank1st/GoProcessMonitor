package data

import (
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
)

func GetMetricData() map[string]interface{} {
    resp, err := http.Get("http://localhost:8123/debug/vars")

    defer resp.Body.Close()

    if err != nil {
        log.Fatal("request get mem data error!")
    }

    body, err := ioutil.ReadAll(resp.Body)

    data := make(map[string]interface{})
    json.Unmarshal(body, &data)

    return data
}