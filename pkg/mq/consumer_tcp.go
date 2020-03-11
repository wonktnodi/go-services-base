package mq

import (
  rocketmq "github.com/apache/rocketmq-client-go/core"
  "github.com/wonktnodi/go-services-base/pkg/logging"
  "github.com/wonktnodi/go-services-base/pkg/utils"
)

var mqDone = make(chan bool)

func InitTcpConsumer(config *RocketMQ, process MsgProcessor) {
  pConfig := &rocketmq.PushConsumerConfig{
    ClientConfig: rocketmq.ClientConfig{
      //您在阿里云 RocketMQ 控制台上申请的 GID
      GroupID: config.GroupId,
      //设置 TCP 协议接入点，从阿里云 RocketMQ 控制台的实例详情页面获取
      NameServer: config.Endpoint,
      Credentials: &rocketmq.SessionCredentials{
        //您在阿里云账号管理控制台中创建的 AccessKeyId，用于身份认证
        AccessKey: config.AccessKey,
        //您在阿里云账号管理控制台中创建的 AccessKeySecret，用于身份认证
        SecretKey: config.SecretKey,
        //用户渠道，默认值为：ALIYUN
        Channel: "ALIYUN",
      },
    },
    //设置使用集群模式
    Model: rocketmq.Clustering,
    //设置该消费者为普通消息消费
    ConsumerModel: rocketmq.CoCurrently,
  }
  pConfig.LogC = &rocketmq.LogConfig{
    Path:    "./log/mq",
    FileNum: 10,
    Level:   rocketmq.LogLevelTrace,
  }
  go ConsumeWithPush(config.Topic, pConfig, process)
}

func ConsumeWithPush(topic string, config *rocketmq.PushConsumerConfig, process MsgProcessor) {
  consumer, err := rocketmq.NewPushConsumer(config)
  if err != nil {
    logging.Errorf("create Consumer failed, error:", err)
    return
  }
  // ********************************************
  // 1. 确保订阅关系的设置在启动之前完成
  // 2. 确保相同 GID 下面的消费者的订阅关系一致
  // *********************************************
  consumer.Subscribe(topic, "*", func(msg *rocketmq.MessageExt) rocketmq.ConsumeStatus {
    return process(msg)
  })
  err = consumer.Start()
  if err != nil {
    println("consumer start failed,", err)
    return
  }
  logging.Debugf("consumer: %s started...", consumer)

  utils.CheckServiceStatus()
  //请保持消费者一直处于运行状态
  err = consumer.Shutdown()
  if err != nil {
    println("consumer shutdown failed")
    return
  }
  logging.Debugf("consumer has shutdown.")
  mqDone <- true
}

//func ConsumerTcp() {
//  rocketMQ := configMQ.Settings.RocketMQ
//
//  pConfig := &rocketmq.PushConsumerConfig{
//    ClientConfig: rocketmq.ClientConfig{
//      //您在阿里云 RocketMQ 控制台上申请的 GID
//      GroupID: rocketMQ.GroupId,
//      //设置 TCP 协议接入点，从阿里云 RocketMQ 控制台的实例详情页面获取
//      NameServer: rocketMQ.NameSrvAddr,
//      Credentials: &rocketmq.SessionCredentials{
//        //您在阿里云账号管理控制台中创建的 AccessKeyId，用于身份认证
//        AccessKey: rocketMQ.AccessKey,
//        //您在阿里云账号管理控制台中创建的 AccessKeySecret，用于身份认证
//        SecretKey: rocketMQ.SecretKey,
//        //用户渠道，默认值为：ALIYUN
//        Channel: "ALIYUN",
//      },
//    },
//    //设置使用广播模式
//    //Clustering：表示集群消费
//    //Broadcasting：表示广播消费
//    Model: rocketmq.Clustering,
//    //设置该消费者为普通消息消费
//    //CoCurrently：表示普通消息消费者
//    //Orderly：表示顺序消息消费者
//    ConsumerModel: rocketmq.CoCurrently,
//    //将消费者线程数固定为 50 个
//    ThreadCount: 50,
//  }
//  ConsumeWithPush(pConfig)
//}
//
//func ConsumeWithPush(config *rocketmq.PushConsumerConfig) {
//  fmt.Println("come on rocketMQ")
//  fmt.Println("hello from users_center rocketMQ")
//  rocketMQ := configMQ.Settings.RocketMQ
//
//  consumer, err := rocketmq.NewPushConsumer(config)
//  if err != nil {
//    logging.Info("create Consumer failed, error:", err)
//    return
//  }
//  //ch := make(chan interface{})
//  //var count = (int64)(1000000)
//  // ********************************************
//  // 1. 确保订阅关系的设置在启动之前完成
//  // 2. 确保相同 GID 下面的消费者的订阅关系一致
//  // *********************************************
//  consumer.Subscribe(rocketMQ.Topic, rocketMQ.Tag, func(msg *rocketmq.MessageExt) rocketmq.ConsumeStatus {
//    logging.Info("A message received, MessageID:%s, Body:%s \n", msg.MessageID, msg.Body)
//    //if atomic.AddInt64(&count, -1) <= 0 {
//    //  //ch <- "quit"
//    //}
//    var orderInfo vo.OrderInfo
//    bytes := []byte(msg.Body)
//    err = json.Unmarshal(bytes, orderInfo)
//    if err != nil {
//      logging.Errorf("failed to generate ConsumeData data:%s ;orderInfo:%s", msg.Body, orderInfo)
//      return rocketmq.ReConsumeLater
//    }
//    code := models.CreateMember(&orderInfo)
//    if code != errors.SUCCESS {
//      logging.Errorf("Consumption failure,ResponseCode :%s  ;orderInfo:%s", code, orderInfo)
//      logging.Errorf("ReConsumeLater retry")
//      return rocketmq.ReConsumeLater
//    }
//    //消费成功回复 ConsumeSuccess，消费失败回复 ReConsumeLater。此时会触发消费重试
//    return rocketmq.ConsumeSuccess
//  })
//  err = consumer.Start()
//  if err != nil {
//    logging.Info("consumer start failed,", err)
//    return
//  }
//  logging.Info("consumer: %s started...\n", consumer)
//  //<-ch
//  //请保持消费者一直处于运行状态
//  //err = consumer.Shutdown()
//  //if err != nil {
//  //  println("consumer shutdown failed")
//  //  return
//  //}
//  //println("consumer has shutdown.")
//}
