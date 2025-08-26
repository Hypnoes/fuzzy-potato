package api_proxy

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func CallProxy(method string, args interface{}) (interface{}, error) {
	queue, err := lookupQueueFromManager(method)

	if err != nil {
		return nil, err
	}

	jsonRpcRequest := &JsonRpcRequest{
		Method: method,
		Params: args,
	}

	publishToRabbitMQ(queue, jsonRpcRequest)
	return nil, nil
}

func publishToRabbitMQ(queue any, request *JsonRpcRequest) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	body, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	ch.Publish("", queue.(string), false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

func lookupQueueFromManager(method string) (any, error) {
	panic("unimplemented")
}
