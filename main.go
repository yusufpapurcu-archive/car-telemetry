package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetPrefix("[Car Telemetry]")
	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			StartPortListening(conf)
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}

func getConfig() (Config, error) {
	f, err := os.Open("Telemetry.toml")
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if _, err := toml.Decode(string(content), &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
