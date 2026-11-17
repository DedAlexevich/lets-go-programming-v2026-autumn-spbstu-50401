package main

import "flag"

func main() {
	configPath := flag.String("config", "", "path to the YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("Doesn't add flag -config")
	}
}
