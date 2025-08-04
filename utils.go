package main

import "fmt"

func itoa(def int) string {
	return fmt.Sprintf("%d", def)
}
func ftoa(def float64) string {
	return fmt.Sprintf("%f", def)
}
