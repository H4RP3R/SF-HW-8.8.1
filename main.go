package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"SF-HW-8.8.1/electronic"
)

var w = tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

func printCharacteristics(p electronic.Phone) {
	var osName string = "-"
	var btnCtn int = -1

	if phone, ok := p.(electronic.Smartphone); ok {
		osName = phone.OS()
	}

	if phone, ok := p.(electronic.StationPhone); ok {
		btnCtn = phone.ButtonsCount()
	}

	ternary := map[bool]string{true: fmt.Sprintf("%d", btnCtn), false: "-"}

	str := fmt.Sprintf("%s\t%s\t%s\t%s\t%s", p.Brand(), p.Model(), p.Type(), osName, ternary[btnCtn != -1])
	fmt.Fprintln(w, str)
}

func main() {
	iphone := electronic.NewApplePhone("Apple iPhone 15", "iOS 17")
	xperia := electronic.NewAndroidPhone("sony", "Sony Xperia 1 III", "Android 11")
	panasonic := electronic.NewRadioPhone("panasonic", "Panasonic KX-TG2512RUN", 20)

	phones := []electronic.Phone{iphone, xperia, panasonic}
	fmt.Fprintln(w, strings.ToUpper("brand\tmodel\ttype\tos\tbtn count"))
	for _, phone := range phones {
		printCharacteristics(phone)
	}
	w.Flush()
}
