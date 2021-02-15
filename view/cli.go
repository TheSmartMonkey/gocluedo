package main

import (
	"log"
	"flag"

	c "github.com/TheSmartMonkey/gocluedo/controller"
)


func main() {
	log.Println("CLUDEO")
	c.Game()
	cli()
}

func cli() {
	textPtr := flag.String("text", "", "Text to parse.")
    metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
    uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
    flag.Parse()

    log.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}
