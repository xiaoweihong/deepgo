package main

import (
	"deepgo/go_kafka"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	brokerList = flag.String("brokers", "127.0.0.1:32400", "The comma separated list of brokers in the Kafka cluster")
	topic      = flag.String("topic", "index-vehicle", "REQUIRED: the topic to consume")
	partitions = flag.String("partitions", "all", "The partitions to consume, can be 'all' or comma-separated numbers")
	offset     = flag.String("offset", "newest", "The offset to start with. Can be `oldest`, `newest`")
	verbose    = flag.Bool("verbose", false, "Whether to turn on sarama logging")
	bufferSize = flag.Int("buffer-size", 256, "The buffer size of the message channel.")

	logger             = log.New(os.Stderr, "", log.LstdFlags)
	genericObj         = &go_kafka.GenericObj{}
	vehicleobj         = &go_kafka.VehicleObj{}
	NonMotorVehicleObj = &go_kafka.NonMotorVehicleObj{}
	pedestrianObj      = &go_kafka.PedestrianObj{}
	faceObj            = &go_kafka.FaceObj{}
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("logtostderr", "true")
	flag.Set("v", "1")
}

func main() {
	flag.Parse()

	if *brokerList == "" {
		printUsageErrorAndExit("no -brokers specified. Alternatively, set the KAFKA_PEERS environment variable")
	}

	if *topic == "" {
		printUsageErrorAndExit("no -topic specified")
	}

	if *verbose {
		sarama.Logger = logger
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	var initialOffset int64
	switch *offset {
	case "oldest":
		initialOffset = sarama.OffsetOldest
	case "newest":
		initialOffset = sarama.OffsetNewest
	default:
		printUsageErrorAndExit("-offset should be `oldest` or `newest`")
	}

	c, err := sarama.NewConsumer(strings.Split(*brokerList, ","), nil)

	if err != nil {
		printErrorAndExit(69, "Failed to start consumer: %s", err)
	}
	partitionList, err := getPartitions(c)
	if err != nil {
		printErrorAndExit(69, "Failed to get the list of partitions: %s", err)
	}
	var (
		messages = make(chan *sarama.ConsumerMessage, *bufferSize)
		closing  = make(chan struct{})
		wg       sync.WaitGroup
		count    int64
	)
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt)
		<-signals
		logger.Println("Initiating shutdown of consumer...")
		close(closing)
	}()

	for _, partition := range partitionList {
		pc, err := c.ConsumePartition(*topic, partition, initialOffset)
		if err != nil {
			printErrorAndExit(69, "Failed to start consumer for partition %d: %s", partition, err)
		}

		go func(pc sarama.PartitionConsumer) {
			<-closing
			pc.AsyncClose()
		}(pc)

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				messages <- message
			}
		}(pc)
	}

	go func() {
		for msg := range messages {

			proto.Unmarshal(msg.Value, genericObj)
			if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_CAR {
				glog.Info("ObjType_OBJ_TYPE_CAR")
				proto.Unmarshal(genericObj.GetBinData(), vehicleobj)
				glog.Infoln(timeStampToDate(vehicleobj.GetMetadata().Timestamp))
				glog.V(1).Infoln("vehicle--->" + vehicleobj.GetVehicle()[0].GetImg().GetImg().GetURI())
				glog.V(2).Infoln("fullImage--->" + vehicleobj.GetImg().GetURI())
				glog.V(6).Infoln(vehicleobj.GetVehicle())
				glog.Infoln("************************************************")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_BICYCLE {
				glog.Info("ObjType_OBJ_TYPE_BICYCLE")
				proto.Unmarshal(genericObj.GetBinData(), NonMotorVehicleObj)
				glog.Infoln(timeStampToDate(NonMotorVehicleObj.GetMetadata().Timestamp))
				glog.V(1).Infoln("bicycle--->" + NonMotorVehicleObj.GetNonMotorVehicles()[0].GetImg().GetImg().URI)
				glog.V(2).Infoln("fullImage--->" + NonMotorVehicleObj.GetImg().GetURI())
				glog.V(6).Infoln(NonMotorVehicleObj)
				glog.Infoln("************************************************")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_TRICYCLE {
				glog.Info("ObjType_OBJ_TYPE_TRICYCLE")
				proto.Unmarshal(genericObj.GetBinData(), NonMotorVehicleObj)
				glog.Infoln(NonMotorVehicleObj.GetMetadata().Timestamp)
				glog.V(1).Infoln("tricycle--->" + NonMotorVehicleObj.GetNonMotorVehicles()[0].GetImg().GetImg().URI)
				glog.V(2).Infoln("fullImage--->" + NonMotorVehicleObj.GetImg().GetURI())
				glog.V(6).Infoln(NonMotorVehicleObj)
				glog.Infoln("************************************************")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_PEDESTRIAN {
				glog.Info("ObjType_OBJ_TYPE_PEDESTRIAN")
				proto.Unmarshal(genericObj.GetBinData(), pedestrianObj)
				glog.Infoln(timeStampToDate(pedestrianObj.GetMetadata().Timestamp))
				glog.V(1).Infoln("pedestrian--->" + pedestrianObj.GetPedestrian()[0].GetImg().GetImg().URI)
				glog.V(2).Infoln("fullImage--->" + pedestrianObj.GetImg().GetURI())
				glog.V(6).Infoln(pedestrianObj)
				glog.Infoln("************************************************")
			} else if genericObj.ObjType == go_kafka.ObjType_OBJ_TYPE_FACE {
				glog.Info("ObjType_OBJ_TYPE_FACE")
				proto.Unmarshal(genericObj.GetBinData(), faceObj)
				glog.Infoln(timeStampToDate(faceObj.GetMetadata().Timestamp))
				glog.V(1).Infoln("face--->" + faceObj.GetFaces()[0].GetImg().GetImg().URI)
				glog.V(2).Infoln("fullImage--->" + faceObj.GetImg().GetURI())
				glog.V(6).Infoln(faceObj)
				glog.Infoln("************************************************")
			}
			count++
			//logger.Printf("consumed %dth\n",count)
		}
	}()

	/*
		go func() {
			for msg := range messages {
				fmt.Printf("Partition:\t%d\n", msg.Partition)
				fmt.Printf("Offset:\t%d\n", msg.Offset)
				fmt.Printf("Key:\t%s\n", string(msg.Key))
				fmt.Printf("Value:\t%s\n", string(msg.Value))
				count++
				logger.Printf("consumed %dth ",count)
				fmt.Println()
			}
		}()
	*/
	wg.Wait()
	logger.Println("Done consuming topic", *topic)
	logger.Printf("consumed total: %d", count)

	close(messages)
	glog.Flush()

	if err := c.Close(); err != nil {
		logger.Println("Failed to close consumer: ", err)
	}
}

func getPartitions(c sarama.Consumer) ([]int32, error) {
	if *partitions == "all" {
		return c.Partitions(*topic)
	}
	tmp := strings.Split(*partitions, ",")
	var pList []int32
	for i := range tmp {
		val, err := strconv.ParseInt(tmp[i], 10, 32)
		if err != nil {
			return nil, err
		}
		pList = append(pList, int32(val))
	}

	return pList, nil
}

func printUsageErrorAndExit(message string) {
	fmt.Fprintln(os.Stderr, "ERROR:", message)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Available command line options:")
	flag.PrintDefaults()
	os.Exit(64)
}

func printErrorAndExit(code int, format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}

func timeStampToDate(timestamp int64) string {

	return time.Unix(timestamp/1e3, 0).Format("2006-01-02 15:04:05")

}
