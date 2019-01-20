package main

import (
	"deepgo/go_kafka"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

func init() {
	//log.SetOutput(os.Stdout)
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("v", "1")

}

func main() {
	ip := flag.String("ip", "127.0.0.1", "kafka ip")
	port := flag.String("port", "32400", "kafka port")
	topic := flag.String("topic", "index-vehicle", "kafka topic")
	flag.Parse()

	brokers := []string{*ip + ":" + *port}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.MaxVersion

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		panic("err get consumer")
	}
	defer consumer.Close()
	genericObj := &go_kafka.GenericObj{}
	vehicleobj := &go_kafka.VehicleObj{}
	partitonsConsumer, err := consumer.ConsumePartition(*topic, 1, sarama.OffsetOldest)
	if err != nil {
		fmt.Println("err get partition consumer", err)
	}
	defer partitonsConsumer.Close()
	for {
		select {
		case msg := <-partitonsConsumer.Messages():
			proto.Unmarshal(msg.Value, genericObj)

			if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_CAR {
				proto.Unmarshal(genericObj.GetBinData(), vehicleobj)
				glog.Infoln(vehicleobj.Vehicle)

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_BICYCLE {
				//glog.Info("ObjType_OBJ_TYPE_BICYCLE")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_TRICYCLE {

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_PEDESTRIAN {

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_FACE {
				//glog.Info("ObjType_OBJ_TYPE_FACE")
			}

		case err := <-partitonsConsumer.Errors():
			glog.Infoln(err.Err)
		}
	}

}
