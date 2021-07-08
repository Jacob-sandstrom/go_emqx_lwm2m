package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models"
	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models/base_models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type RegisterMessageHandler func(models.RegisterResp)

type GeneralMessageHandler func(models.Resp)

type MessageHandler func(Client, mqtt.Message)

type Client interface {
	mqtt.Client
	ConnectMqtt()
	read(string) error
}

type client struct {
	mqtt.Client
	Endpoint      string
	RespHandler   MessageHandler
	NotifyHandler MessageHandler
}

func NewClient(endpoint string) Client {
	return &client{
		Client:        NewMqttClient("tcp://127.0.0.1:1883"),
		Endpoint:      endpoint,
		RespHandler:   onRespReceived,
		NotifyHandler: onNotifyReceived,
	}
}

func NewMqttClient(server string) mqtt.Client {
	server = "tcp://35.228.122.91:1883"
	qos := 0

	connOpts := mqtt.NewClientOptions().AddBroker(server).SetOrderMatters(false)
	connOpts.OnConnect = func(c mqtt.Client) {
		// Uses emqx_lwm2m default topics
		if token := c.Subscribe("lwm2m/+/up/resp", byte(qos), onRespReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		if token := c.Subscribe("lwm2m/+/up/notify", byte(qos), onNotifyReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	return mqtt.NewClient(connOpts)
}

func (c *client) ConnectMqtt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		op := c.OptionsReader()
		fmt.Printf("Connected to %s\n", op.Servers())
	}

	<-ch
}

func main() {
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	// server := "tcp://35.228.122.91:1883"
	// qos := 0

	// connOpts := mqtt.NewClientOptions().AddBroker(server).SetOrderMatters(false)
	// connOpts.OnConnect = func(c mqtt.Client) {
	// 	if token := c.Subscribe("lwm2m/+/up/resp", byte(qos), onRespReceived); token.Wait() && token.Error() != nil {
	// 		panic(token.Error())
	// 	}
	// 	if token := c.Subscribe("lwm2m/+/up/notify", byte(qos), onNotifyReceived); token.Wait() && token.Error() != nil {
	// 		panic(token.Error())
	// 	}
	// }
	// mqttClient := mqtt.NewClient(connOpts)

	// if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// } else {
	// 	fmt.Printf("Connected to %s\n", server)
	// }

	client := NewClient("js_lwm2m_demo")

	client.ConnectMqtt()

	// <-ch
}

func (c *client) PublishMqtt(endpoint string, payload []byte) error {
	fmt.Print("publish\n")
	token := c.Publish(fmt.Sprintf("lwm2m/%v/dn", endpoint), 0, false, payload)
	if token.Error() != nil {
		fmt.Print(token.Error())
		return token.Error()
	}
	return nil
}

func (c *client) observe(path string) error {
	fmt.Printf("\n%v %v %v", "observe", c.Endpoint, path)
	req := models.ObserveReq{ReqID: 1, MsgType: "observe", Data: models.ObserveReqData{Path: path}}
	payload, err := req.Marshal()
	if err != nil {
		return err
	}
	return c.PublishMqtt(c.Endpoint, payload)
}

func (c *client) writeAttr(values map[string]interface{}) error {
	fmt.Printf("\n%v %v %v", "write-attr", c.Endpoint, values[`path`])
	req := models.WriteAttrReq{ReqID: 1, MsgType: "write-attr", Data: models.WriteAttrData{
		Path: values[`path`].(string),
		Pmin: int64(values[`pmin`].(int)),
		Pmax: int64(values[`pmax`].(int)),
	}}
	payload, err := req.Marshal()
	if err != nil {
		return err
	}
	return c.PublishMqtt(c.Endpoint, payload)
}

func (c *client) discover(path string) error {
	fmt.Printf("\n%v %v %v", "discover", c.Endpoint, path)
	req := models.DiscoverReq{base_models.ReadDiscoverDeleteReq{ReqID: 1, MsgType: "discover", Data: base_models.ReadDiscoverDeleteData{Path: path}}}
	payload, err := req.Marshal()
	if err != nil {
		return err
	}
	return c.PublishMqtt(c.Endpoint, payload)
}

func (c *client) read(path string) error {
	fmt.Printf("\n%v %v %v", "read", c.Endpoint, path)
	req := models.ReadReq{base_models.ReadDiscoverDeleteReq{ReqID: 1, MsgType: "read", Data: base_models.ReadDiscoverDeleteData{Path: path}}}
	payload, err := req.Marshal()
	if err != nil {
		return err
	}
	return c.PublishMqtt(c.Endpoint, payload)
}

func onRespReceived(mqttClient Client, message mqtt.Message) {
	topic := message.Topic()
	payload := message.Payload()
	endpoint := topic[6 : len(topic)-8]

	fmt.Printf("\n%v", topic)

	var p map[string]interface{}
	err := json.Unmarshal(payload, &p)
	if err != nil {
		panic(err.Error())
	}

	switch p[`msgType`] {
	case "register":
		r, err := models.UnmarshalRegisterResp(payload)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v", r)

		// c := NewClient(endpoint)
		// c.ConnectMqtt()
		// c.discover("/3")
		// c.writeAttr(map[string]interface{}{"path": "/3", "pmin": 10, "pmax": 20})
		mqttClient.read("/3")
		// c.observe("/3")

	case "read":
		r, err := models.UnmarshalResp(payload)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v", r)

	case "discover":
		r, err := models.UnmarshalResp(payload)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v", r)

	case "write":
	case "write-attr":
		r, err := models.UnmarshalResp(payload)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v", r)

	case "execute":
	case "create":
	case "delete":

	default:

	}

}

func onNotifyReceived(mqttClient mqtt.Client, message mqtt.Message) {
	topic := message.Topic()
	payload := message.Payload()

	fmt.Printf("\n%v", topic)
	fmt.Printf("\n%v", payload)

}

func registerHandler(reg models.RegisterResp) {

}
