package main

import "net/url"

func EncodeURI(uri string, params map[string]string) string {
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	u, _ := url.ParseRequestURI(uri)
	dataEncode := data.Encode()
	u.RawQuery = dataEncode
	uriStr := u.String()
	return uriStr
}

func main() {

	url := "https://graph.microsoft.com/v1.0/sites?search="
	params := map[string]string{
		"search": "",
	}
	urlNew := EncodeURI(url, params)
	println(urlNew)

}
