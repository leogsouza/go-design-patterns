package main

import (
	"flag"
	"log"
	"os"

	"github.com/leogsouza/go-design-patterns/behavioral/strategy/shapes"
)

var output = flag.String("output", "text",
	"The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.NewPrinter(*output)

	switch *output {
	case shapes.TextStrategy:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.ImageStrategy:
		w, err := os.Create(os.TempDir() + "/image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
