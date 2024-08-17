package main

import seohelperai "github.com/Fibig/seo-helper-ai/internal/seo-helper-ai"

func main() {
	server := seohelperai.NewServer()
	server.Run(":9000")
}
