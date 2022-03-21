package test

import (
	"fmt"
	"gin/structs/response"
	"gin/tool/logger"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ginconfig "gin/config"

	"math/rand"
)

var (
	topic = "go_gin_220321"
)

func KafkaProducer(c *gin.Context) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	rand1 := strconv.Itoa(rand.Intn(1000))

	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder("key=" + rand1),
		Value: sarama.StringEncoder("value=" + rand1),
	}
	msg.Topic = topic

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

func KafkaConsumer(c *gin.Context) {
	res := response.Response{}

	consumer, err := sarama.NewConsumer([]string{ginconfig.KafkaBroker}, nil)
	if err != nil {
		res.SetMessage("fail to start consumer, err:" + err.Error())
		c.JSON(200, res)
		return
	}

	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		res.SetMessage("fail to get list of partition:err:" + err.Error())
		c.JSON(200, res)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		// 同步
		// for msg := range pc.Messages() {
		// 	fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)

		// 	mylog.WriteInLog("kafka 队列消息 AAAAAAAAA")

		// 	mylog.WriteInLog(string(msg.Key))
		// 	mylog.WriteInLog("kafka 队列消息 ========")
		// 	mylog.WriteInLog(string(msg.Value))

		// 	mylog.WriteInLog("kafka 队列消息 BBBBBBBB")

		// }

		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				// logger.ZapLog.Info
				logger.ZapLog.Info(
					"kafka 队列消息 start",
					zap.String("key=", string(msg.Key)),
					zap.String("value=", string(msg.Value)),
				)
			}
		}(pc)
	}

	// 主线程休眠一会，防止子协程挂掉
	time.Sleep(60 * time.Second)

	c.JSON(200, res)
}
