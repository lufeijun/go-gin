package v1

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"

	ginconfig "gin/config"
	response "gin/structs"
)

func Producer(c *gin.Context) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder("吉鹏"),
		Value: sarama.StringEncoder("吉鹏123"),
	}
	msg.Topic = "quickstart-events"

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{ginconfig.KafkaBroker}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	// 延迟执行
	defer client.Close()

	res := response.Response{}

	// 发送消息
	// pid, offset, err := client.SendMessage(msg)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		// fmt.Println("send msg failed, err:", err)
		res.Message = err.Error()
		res.ToClientData()
		c.JSON(200, res)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

	c.JSON(200, res)
}

func Consumer(c *gin.Context) {
	res := response.Response{}

	consumer, err := sarama.NewConsumer([]string{ginconfig.KafkaBroker}, nil)
	if err != nil {
		res.Message = "fail to start consumer, err:" + err.Error()
		res.ToClientData()
		c.JSON(200, res)
		return
	}

	// partitionList, err := consumer.Partitions("quickstart-events") // 根据topic取到所有的分区
	partitionList, err := consumer.Partitions("quickstart-events") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("quickstart-events", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		// 同步
		for msg := range pc.Messages() {
			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)

			writeInLog("kafka 队列消息 AAAAAAAAA")

			writeInLog(string(msg.Key))
			writeInLog("kafka 队列消息 ========")
			writeInLog(string(msg.Value))

			writeInLog("kafka 队列消息 BBBBBBBB")

		}

		// 异步从每个分区消费信息
		// go func(sarama.PartitionConsumer) {
		// 	for msg := range pc.Messages() {
		// 		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)

		// 		writeInLog("kafka 队列消息")

		// 		writeInLog(string(msg.Key))
		// 		writeInLog(string(msg.Value))

		// 	}
		// }(pc)
	}

	c.JSON(200, res)

}

func Index(c *gin.Context) {

	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}

func writeInLog(msg string) {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(msg)
}
