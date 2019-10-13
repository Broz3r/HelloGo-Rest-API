package models

import "time"

type HelloWorld struct {
	Value string `json:"value"`
	Date time.Time `json:"date"`
}