package main

import (
	"deepgo/go_kafka"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"time"
)

func init() {
	//log.SetOutput(os.Stdout)
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("v", "5")

}

var kafkaChain chan []*go_kafka.RecVehicle

func main() {
	ip := flag.String("ip", "127.0.0.1", "kafka ip")
	port := flag.String("port", "32400", "kafka port")
	topic := flag.String("topic", "index-vehicle", "kafka topic")
	flag.Parse()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": *ip + ":" + *port,
		"group.id":          "deepglint",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{*topic}, nil)

	glog.Info("for.....")

	go getKafkaMessage(c)
	time.Sleep(time.Second * 2)

	msg := <-kafkaChain
	glog.Infoln(msg)
	glog.Flush()
	c.Close()
}

func getKafkaMessage(c *kafka.Consumer) {

	genericObj := &go_kafka.GenericObj{}
	vehicleobj := &go_kafka.VehicleObj{}
	msg, err := c.ReadMessage(-1)
	glog.Info("get kafka information...")
	if err == nil {
		//fmt.Printf("Message on %s:%s \n", msg.TopicPartition, msg.Value)
		proto.Unmarshal(msg.Value, genericObj)
		if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_CAR {
			proto.Unmarshal(genericObj.GetBinData(), vehicleobj)
			//a:=vehicleobj.Vehicle
			kafkaChain <- vehicleobj.Vehicle
		}
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

}
