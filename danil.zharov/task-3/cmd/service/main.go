package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v4"
	"golang.org/x/net/html/charset"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CommaFloat float64

func (f *CommaFloat) UnmarshalText(text []byte) error {
	normalized := strings.Replace(string(text), ",", ".", 1)
	value, err := strconv.ParseFloat(normalized, 64)

	if err != nil {
		return fmt.Errorf("failed to parse number %q: %w", text, err)
	}
	*f = CommaFloat(value)

	return nil
}

type Currency struct {
	NumCode  int        `xml:"NumCode"`
	CharCode string     `xml:"CharCode"`
	Value    CommaFloat `xml:"Value"`
}

type ExchangeRates struct {
	Currencies []Currency `xml:"Valute"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &cfg, nil
}

func loadCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var rates ExchangeRates
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("failed to decode input file: %w", err)
	}

	return rates.Currencies, nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := loadCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", currencies)
}
