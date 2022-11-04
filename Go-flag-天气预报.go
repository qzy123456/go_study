package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JohannesKaufmann/html-to-markdown"
	"github.com/MichaelMure/go-term-markdown"
)

func GetWeather(city string) (string, error) {
	url := "https://wttr.in/" + city
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	weather := string(all)
	md := getMD(weather)
	result := markdown.Render(md, 280, 6)
	return string(result), nil
}

func getMD(html string) string {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return ""
	}
	return markdown
}
func main() {
	var weather string
	flag.StringVar(&weather, "w", "shanghai", "区域")
	flag.StringVar(&weather, "weather", "shanghai", "区域")
	flag.Parse()
	str, err := GetWeather(weather)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
