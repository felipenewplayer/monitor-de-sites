package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    sitesEnv := os.Getenv("SITES")
    if sitesEnv == "" {
        fmt.Println("Nenhum site definido na variável de ambiente SITES")
        return
    }

    sites := strings.Split(sitesEnv, ",")

    for _, site := range sites {
        checkSite(site)
    }
}

func checkSite(site string) {
    resp, err := http.Get(site)
    if err != nil {
        fmt.Printf("%s está fora do ar: %v\n", site, err)
        logStatus(site, "offline")
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        fmt.Printf("%s está no ar!\n", site)
        logStatus(site, "online")
    } else {
        fmt.Printf("%s retornou status %d\n", site, resp.StatusCode)
        logStatus(site, "offline")
    }
}

func logStatus(site, status string) {
    f, err := os.OpenFile("/root/logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erro ao abrir log:", err)
        return
    }
    defer f.Close()

    logLine := fmt.Sprintf("%s - %s: %s\n", time.Now().Format("02-01-2006 15:04:05"), site, status)
    if _, err := f.WriteString(logLine); err != nil {
        fmt.Println("Erro ao escrever no log:", err)
    }
}
