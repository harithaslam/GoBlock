package app

import (
	"bytes"
	"fmt"
	"net/http"

	"go.bug.st/serial"
)

func main() {
	serialPort, _ := serial.Open("COM7", &serial.Mode{BaudRate: 9600})
	defer serialPort.Close()
	fmt.Println("Connected to serial port")

	for {
		buff := make([]byte, 100)

		serialPort.Read(buff)
		resp, err := http.Post("https://finance.zapzyntax.online/api/iot", "text/plain", bytes.NewReader(buff))

	}

}

func StatsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_, err := writer.Write(
		[]byte(
			`<p style="">21 Things Have Been Collected</p>
			 <p>Hoping For More</p>
				`))
	if err != nil {

	}
	return
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))

}
