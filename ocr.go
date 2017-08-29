package main

import (
	"fmt"
	"log"
	"github.com/otiai10/gosseract"
	"github.com/sajari/docconv"
	//"github.com/sajari/docconv/client"

)

func main() {
	// Create a new client, using the default endpoint (localhost:8000)
	//c := client.New()

	res, err := docconv.ConvertPath( "annual_report_2009.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)


	res1, err1 := docconv.ConvertPath("demo.docx")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(res1)

	//// PDF
	//pdfo, erro := docconv.ConvertPath("annual_report_2009.pdf")
	//if erro != nil {
	//	log.Fatal(erro)
	//}
	//fmt.Println(pdfo)


	// IMGE
	out := gosseract.Must(gosseract.Params{
		Src:       "Text_Entropy.png",
		Languages: "eng",
	})
	fmt.Println(out)
}