package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input-file and output-file must be set")
	}

	return cfg
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadCurrencies(path string) []Currency {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel

	var vc ValCurs
	if err := dec.Decode(&vc); err != nil {
		panic(err)
	}

	result := make([]Currency, 0, len(vc.Valutes))

	for _, v := range vc.Valutes {
		num, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			panic(err)
		}

		val, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(v.Value), ",", ".", 1), 64)
		if err != nil {
			panic(err)
		}

		result = append(result, Currency{num, strings.TrimSpace(v.CharCode), val})
	}

	return result
}

func saveJSON(path string, data []Currency) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "path to the YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("Doesn't add flag -config")
	}

	cfg := loadConfig(*configPath)
	currencies := loadCurrencies(cfg.InputFile)

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveJSON(cfg.OutputFile, currencies)
}
