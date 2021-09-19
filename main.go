package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bregydoc/gtranslate"
)

func main() {
	fmt.Print("Enter text you want to translate: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Print("Translate from (Country Code): ")
	scanner.Scan()
	from := scanner.Text()
	fmt.Print("Translate to (Country Code): ")
	scanner.Scan()
	to := scanner.Text()
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: from,
			To:   to,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(from+": %s | "+to+": %s ", text, translated)

}
