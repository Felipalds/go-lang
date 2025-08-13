package main

import "fmt"

var greetings = map[string]string{
	"pt": "Olá, ",
	"en": "Hello, ",
	"es": "Hola, ",
}

var world = map[string]string{
	"pt": "mundo",
	"en": "world",
	"es": "mundo",
}

func Hello(name, lang string) string {
	if lang == "" {
		lang = "en"
	}
	if name == "" {
		name = world[lang]
	}
	return greetings[lang] + name
}

func main() {
	fmt.Println(Hello("Luiz", "pt"))
}
