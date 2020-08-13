package main

import (
    "fmt"
    htmlquery "github.com/antchfx/xquery/html"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"
)

func main() {
    urlTemplate := "https://www.kuaidaili.com/free/inha/%d/"
    var proxies []string
    for i := 1; i < 4; i++ {
        html := getHtml(fmt.Sprintf(urlTemplate, i))
        root, _ := htmlquery.Parse(strings.NewReader(html))
        tr := htmlquery.Find(root, "//*[@id='list']/table/tbody/tr")
        for _, row := range tr {
            item := htmlquery.Find(row, "./td")
            ip := htmlquery.InnerText(item[0])
            port := htmlquery.InnerText(item[1])
            //type_ := htmlquery.InnerText(item[3])
            p := ip + ":" + port
            proxies = append(proxies, p)
        }
        time.Sleep(3 * time.Second)
    }
    fmt.Println(len(proxies), proxies[0:5])
}
func getHtml(url_ string) string {
    req, _ := http.NewRequest("GET", url_, nil)
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
    client := &http.Client{Timeout: time.Second * 5}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil && data == nil {
        log.Fatalln(err)
    }
    return fmt.Sprintf("%s", data)
}
