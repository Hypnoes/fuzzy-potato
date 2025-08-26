package api_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/streadway/amqp"
)

func RegisterWithManager(managerURL string, serviceInfo ServiceConfig) {
	resp, _ := http.PostForm(managerURL+"/register", url.Values{
		"name":    {serviceInfo.Name},
		"queue":   {serviceInfo.Queue},
		"methods": {fmt.Sprintf("%v", serviceInfo.Method)},
	})

	if resp.StatusCode != 200 {
		log.Fatal(("Failed to register with APIManager"))
	}
}

func StartConsumer(queue string) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare(queue, false, false, false, false, nil)
	ch.Qos(1, 0, false)
	msgs, _ := ch.Consume(q.Name, "", false, false, false, false, nil)

	for msg := range msgs {
		var req struct {
			Method string      `json:"method"`
			Params interface{} `json:"params"`
		}
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			log.Printf("Failed to decode message: %v", err)
			msg.Nack(false, false)
			continue
		}
		HandleRequest(req.Method, req.Params)
		msg.Ack(false)
	}
}

func HandleRequest(method string, params interface{}) {
	v := reflect.ValueOf(service)
	m := v.MethodByName(method)
	if !m.IsValid() {
		log.Fatal("Method not found")
	}
	m.Call([]reflect.Value{reflect.ValueOf(params)})
}
