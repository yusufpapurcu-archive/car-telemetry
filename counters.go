package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Battery Namespace
var (
	battery_temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "battery",
		Subsystem: "temperature",
		Name:      "battery_temperature",
		Help:      "Gauge for Battery Temperature",
	})

	battery_usage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "battery",
		Subsystem: "usage",
		Name:      "battery_usage",
		Help:      "Gauge for Battery Usage",
	})

	cell_voltages = map[int]prometheus.Gauge{}

	battery_percent = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "battery",
		Subsystem: "percent",
		Name:      "battery_percent",
		Help:      "Gauge for Battery Percent",
	})

	battery_current = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "battery",
		Subsystem: "current",
		Name:      "battery_current",
		Help:      "Gauge for Battery Current",
	})
)

// Engine Namespace
var (
	engine_temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "engine",
		Subsystem: "temperature",
		Name:      "engine_temperature",
		Help:      "Gauge for Engine Temperature",
	})
)

// Counters without spesific namespace
var car_speed = promauto.NewGauge(prometheus.GaugeOpts{
	Namespace: "car",
	Subsystem: "speed",
	Name:      "car_speed",
	Help:      "Gauge for Car Speed",
})
