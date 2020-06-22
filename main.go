package main

import (
	"fmt"
	//"log"
	"net/url"
	"strings"
	//"gopkg.in/yaml.v2"
)

/*
var data = `
a: Easy!
b:
  c: 2
  d: [3, 4, 5]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:"...`
	}
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
	url, e := url.Parse("http://bing.com/search?q=dotnet&a=tanaka")

	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(url.Path)
	fmt.Println(url.RawQuery)
	fmt.Println(url.String())
}
*/
func main() {
	url := "http://bing.com/search?q=dotnet&a=tanaka"
	fmt.Println(url)
	fmt.Println(FilterUrlParameter(url))
}

// URLの中からGETパラメータをmapにして返す実装
func FilterUrlParameter(value string) map[string]string {
	params, _ := url.Parse(value)
	elements := strings.Split(params.RawQuery, "&")
	urlMap := make(map[string]string)
	for i := 0; i < len(elements); i += 1 {
		values := strings.Split(elements[i], "=")
		urlMap[values[0]] = values[1]
	}
	return urlMap
}

// URLのGETパラメータを元にして、YAML文字列の生成
func ConvertUrlGetParamerToYaml(map[string]string) {

	var resultYamlString string
	// TODO https://qiita.com/spiegel-im-spiegel/items/16ab7dabbd0749281227
	content := make([]byte, 0)
	emptyLine := "\r\n"
	indentValue := " "

	return resultYamlString
}
