package main

import (
	"fmt"
	"log"
	"github.com/otiai10/gosseract"
	"github.com/sajari/docconv/client"
)

func main() {
	// Create a new client, using the default endpoint (localhost:8000)
	c := client.New()

	res, err := client.ConvertPath(c, "heychris.docx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	out := gosseract.Must(gosseract.Params{
		Src:       "Text_Entropy.png",
		Languages: "eng",

	})
	fmt.Println(out)
}