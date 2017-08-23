package data

import (
    "net/http"
    "log"
    "encoding/json"
)

func GetMetricData() map[string]interface{} {
    resp, err := http.Get("http://localhost:8123/debug/vars")

    defer resp.Body.Close()

    if err != nil {
        log.Fatal("request get mem data error!")
    }

    data := make(map[string]interface{})
    if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
        return data
    }

    return nil
}