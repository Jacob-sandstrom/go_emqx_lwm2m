package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/Jacob-sandstrom/go_emqx_lwm2m/pkg/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	ReqPayload struct {
		ReqID   string    `json:"reqId"`
		MsgType string    `json:"msgType"`
		Data    DataTypes `json:"data"`
	}
	RespPayload struct {
		ReqID   string    `json:"reqId,omitempty"`
		MsgType string    `json:"msgType"`
		Data    DataTypes `json:"data"`
		Imei    string    `json:"imei,omitempty"`
		Imsi    string    `json:"imsi,omitempty"`
		SeqNum  string    `json:"seqNum,omitempty"`
	}

	DataTypes interface {
		printData()
	}

	Path struct {
		Path string `json:"path"`
	}

	Attributes struct {
		Path
		PMin int `json:"pmin,omitempty"`
		PMax int `json:"pmax,omitempty"`
		GT   int `json:"gt,omitempty"`
		LT   int `json:"lt,omitempty"`
		ST   int `json:"st,omitempty"`
	}

	Register struct {
		Ep              string   `json:"ep"`
		LT              int      `json:"lt"`
		Sms             string   `json:"sms"`
		LWM2M           string   `json:"lwm2m"`
		B               string   `json:"b"`
		AlternativePath string   `json:"alternativePath"`
		ObjectList      []string `json:"objectList"`
	}
)

func (d Path) printData() {
	fmt.Printf(`data: %v`, d)
}

func (d Attributes) printData() {
	fmt.Printf(`data: %v`, d)
}

func (d Register) printData() {
	fmt.Printf(`data: %v`, d)
}

func main() {

	// jsonReq := []byte(`
	// 	{
	// 		"reqId": 1,
	// 		"msgType": "read",
	// 		"data": {
	// 			"path": "/3/0"
	// 		}
	// 	}`)

	// req, err := models.UnmarshalReadReq(jsonReq)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// fmt.Printf("%+v", req)

	// jsonReq = []byte(`{"reqID":3097,"msgType":"discover","data":{"reqPath":"/3/0","content":["</3/0>","</3/0/0>,</3/0/1>,</3/0/2>,</3/0/3>,</3/0/4>,</3/0/5>,</3/0/6>;dim=1,</3/0/7>;dim=1,</3/0/8>;dim=1,</3/0/9>,</3/0/10>,</3/0/11>;dim=1,</3/0/13>,</3/0/14>,</3/0/15>,</3/0/16>,</3/0/17>,</3/0/18>,</3/0/19>,</3/0/20>,</3/0/21>,</3/0/22>;dim=1"],"codeMsg":"content","code":"2.05"}}`)

	// resp, err := models.UnmarshalDiscoverResp(jsonReq)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// fmt.Printf("\n\n%+v\n", resp)
	// resp.Print()

	//

	//

	//

	//

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server := "tcp://35.228.122.91:1883"
	qos := 0

	connOpts := mqtt.NewClientOptions().AddBroker(server).SetOrderMatters(false)
	connOpts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe("lwm2m/+/up/resp", byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		if token := c.Subscribe("lwm2m/+/up/notify", byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := mqtt.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}

	<-c
}

func Publish(client mqtt.Client, endpoint string, payload ReqPayload) {
	p, err := json.Marshal(payload)
	if err != nil {
		fmt.Print(err.Error())
	}

	token := client.Publish(fmt.Sprintf("lwm2m/%v/dn", endpoint), 0, false, p)
	if token.Error() != nil {
		fmt.Print(token.Error())
	}
}

func observe(client mqtt.Client, endpoint string, path string) {
	fmt.Printf("\n%v %v %v", "observe", endpoint, path)
	payload := ReqPayload{ReqID: "1", MsgType: "observe", Data: Attributes{Path: Path{path}}}
	Publish(client, endpoint, payload)
}

func writeAttr(client mqtt.Client, endpoint string, path string) {
	fmt.Printf("\n%v %v %v", "write-attr", endpoint, path)
	payload := ReqPayload{ReqID: "1", MsgType: "write-attr", Data: Attributes{
		Path: Path{path},
		PMin: 10,
		PMax: 20,
	}}
	Publish(client, endpoint, payload)
}

func discover(client mqtt.Client, endpoint string, path string) {
	fmt.Printf("\n%v %v %v", "discover", endpoint, path)
	payload := ReqPayload{ReqID: "1", MsgType: "discover", Data: Attributes{Path: Path{path}}}
	Publish(client, endpoint, payload)
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	topic := message.Topic()
	payload := message.Payload()
	// endpoint := topic[6 : len(topic)-8]

	fmt.Printf("\n%v", topic)

	match, err := regexp.MatchString(`lwm2m\/.*\/up\/resp$`, topic)
	if err != nil {
		panic(err.Error())
	}
	if match {
		var p map[string]interface{}
		err = json.Unmarshal(payload, &p)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Print(p)
		// fmt.Printf("\n%v", p[`msgType`])

		if p[`msgType`] == "register" {
			var r models.RegisterResp
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%+v", r)
			// discover(client, endpoint, "/3")
			// writeAttr(client, endpoint, "/3")
			// observe(client, endpoint, "/3")
		}
		if p[`msgType`] == "write-attr" {
			var r models.WriteAttrResp
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%+v", r)
			// discover(client, endpoint, "/3")
			// writeAttr(client, endpoint, "/3")
			// observe(client, endpoint, "/3")
		}
	}

	match, err = regexp.MatchString(`lwm2m\/.*\/up\/notify$`, topic)
	if err != nil {
		panic(err.Error())
	}
	if match {

	}

}
