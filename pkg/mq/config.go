package mq

import rocketmq "github.com/apache/rocketmq-client-go/core"

type RocketMQ struct {
  AccessKey string
  SecretKey string
  Endpoint  string
  GroupId   string
  Topic     string
  Tag       string
}

type MsgProcessor func(msg *rocketmq.MessageExt) rocketmq.ConsumeStatus

func InitMq(config *RocketMQ, process MsgProcessor) {
  InitTcpConsumer(config, process)
}

func CloseMq() {
  <-mqDone
  close(mqDone)
}
