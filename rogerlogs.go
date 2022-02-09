package main

import (
	"fmt"
	"time"
	"github.com/rogercodeprojects/rogerlogs/colors"
)

func Succes(message string) {
	fmt.Println(colorString(message, colors.Green))
}

func Succeslog(message string) {
	fmt.Println(colorString(timeMessage(message), colors.Green))
}

func Fatal(message string) {
	fmt.Println(colorString(message, colors.Red))
}

func Fatallog(message string) {
	fmt.Println(colorString(timeMessage(message), colors.Red))
}

func timeMessage(message string) string {
	t := time.Now()
	return t.Format("2006-01-02T15:04:05.000-07:00") + " " + message
}

func colorString(message string, color colors.Color) string{
	return string(color) + message + string(colors.Clean)
}
