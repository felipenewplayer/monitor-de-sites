package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var sites = []string{
	"https://google.com",
	"https://github.com",
	"https://front-mini-erp.vercel.app/",
}

func main() {
	for {
		checkSites()
		time.Sleep(30 * time.Second)
	}
}

func checkSites() {
	for _, site := range sites {
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println(site, "está fora do ar:", err)
			logStatus(site, false)
			continue
		}
		if resp.StatusCode == 200 {
			fmt.Println(site, "está no ar!")
			logStatus(site, true)
		} else {
			fmt.Println(site, "respondeu com status:", resp.StatusCode)
			logStatus(site, false)
		}
		resp.Body.Close()
	}
}

func logStatus(site string, status bool) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	defer f.Close()

	statusStr := "offline"
	if status {
		statusStr = "online"
	}

	log := fmt.Sprintf("%s - %s: %s\n", time.Now().Format("02-01-2006 15:04:05"), site, statusStr)
	f.WriteString(log)
}
