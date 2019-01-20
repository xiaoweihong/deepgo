package main

import (
	"deepgo/go_kafka"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": *ip + ":" + *port,
		//"bootstrap.servers": "192.168.2.66:9092",
		"group.id":          "deepglint",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{*topic}, nil)
	genericObj := &go_kafka.GenericObj{}
	vehicleobj := &go_kafka.VehicleObj{}
	//recVehicle := &go_kafka.RecVehicle{}
	//nonMotorObj := &go_kafka.NonMotorVehicleObj{}
	//pedestrianObj:=go_kafka.PedestrianObj{}
	//face_obj:=&go_kafka.FaceObj{}

	for {
		msg, err := c.ReadMessage(-1)
		//glog.Info("get kafka information...")
		if err == nil {
			c.Poll(100)
			//fmt.Printf("Message on %s:%s \n", msg.TopicPartition, msg.Value)
			proto.Unmarshal(msg.Value, genericObj)
			if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_CAR {
				proto.Unmarshal(genericObj.GetBinData(), vehicleobj)
				//glog.V(1).Infoln("Big_img-->"+vehicleobj.Img.GetURI())
				//a:=vehicleobj.Vehicle
				glog.V(3).Infoln(vehicleobj.Vehicle)
				//glog.V(1).Info(vehicleobj.Vehicle[0].GetImg().GetImg().GetURI())
				//v_tmp := vehicleobj.Vehicle
				//proto.Unmarshal(recVehicle.Img,v_tmp)
				//glog.Info(len(v_tmp))
				//glog.Info(v_tmp[1].Img.Img.URI)
				//json.Marshal(go_kafka.VehicleObj{},)

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_BICYCLE {
				//glog.Info("ObjType_OBJ_TYPE_BICYCLE")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_TRICYCLE {

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_PEDESTRIAN {

			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_FACE {
				//glog.Info("ObjType_OBJ_TYPE_FACE")
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	glog.Flush()
	c.Close()
}
