package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"net/http"
	"os"
	"time"
)

var channel = make(chan string, 1)

type Data struct {
	Temperature string `json:"temperature"`
	Time        string `json:"time"`
}

func main() {
	godotenv.Load(".env")
	SERIAL_PORT := os.Getenv("SERIAL_PORT")
	serialPort, _ := serial.Open(SERIAL_PORT, &serial.Mode{BaudRate: 9600})
	defer serialPort.Close()
	fmt.Println("Connected to serial port")
	go func() {
		for {
			fmt.Println("Waiting for data")
			time.Sleep(10 * time.Minute)
			select {
			case res := <-channel:
				fmt.Println("Got data", res)
				var time = time.Now().Format("2006-01-02 15:04:05")
				var data = Data{Temperature: res, Time: time}
				var jsonStr, _ = json.Marshal(data)
				resp, err := http.Post("http://localhost:8000/api/iot", "application/json", bytes.NewReader(jsonStr))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(resp)
				if resp.StatusCode == 200 {
					fmt.Println("Success")
				} else {
					fmt.Println("Failed")
				}
			default:
				fmt.Println("No data")
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		buff := make([]byte, 100)

		_, err := serialPort.Read(buff)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(buff[0:4]), time.Now().Format("15:04:05"))
		if len(channel) == 0 {
			channel <- string(buff[0:4])
		}
	}

}
