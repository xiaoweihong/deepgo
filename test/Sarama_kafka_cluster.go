package main

import (
	"deepgo/go_kafka"
	"flag"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
	"os/signal"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("v", "1")
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "kafka ip")
	port := flag.String("port", "32400", "kafka port")
	topic := flag.String("topic", "index-vehicle", "kafka topic")
	flag.Parse()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	brokers := []string{*ip + ":" + *port}
	topics := []string{*topic}

	consumer, err := cluster.NewConsumer(brokers, "deepglint", topics, config)

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	genericObj := &go_kafka.GenericObj{}
	vehicleobj := &go_kafka.VehicleObj{}
	//recVehicle := &go_kafka.RecVehicle{}
	//nonMotorObj := &go_kafka.NonMotorVehicleObj{}
	//pedestrianObj:=go_kafka.PedestrianObj{}
	face_obj := &go_kafka.FaceObj{}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s", err.Error())
		}
	}()

	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}

	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//fmt.Println(string(msg.Value))
				proto.Unmarshal(msg.Value, genericObj)
				if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_CAR {
					proto.Unmarshal(genericObj.GetBinData(), vehicleobj)
					//glog.Infoln("vehicle num--->",len(vehicleobj.GetVehicle()))
					glog.Infoln("vehicle img--->", vehicleobj.GetVehicle()[0].GetImg().GetImg().GetURI())

				} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_BICYCLE {
					//glog.Info("ObjType_OBJ_TYPE_BICYCLE")
				} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_TRICYCLE {

				} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_PEDESTRIAN {

				} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_FACE {
					proto.Unmarshal(genericObj.GetBinData(), face_obj)
					//glog.Infoln("face num---->",len(face_obj.GetFaces()))
					glog.Infoln("face img---->", face_obj.GetFaces()[0].GetImg().GetImg().GetURI())

				}
			}
		case <-signals:
			return
		}

	}

}
