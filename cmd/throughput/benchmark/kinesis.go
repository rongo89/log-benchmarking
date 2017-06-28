package benchmark

import (
  "fmt"
  "math/rand"

  "github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisBenchmark struct {
  payloadSize         uint
  urls                []string
  topic               string
  recv                chan []byte
  producer            
  consumer
  msg
  errors              uint
  acked               uint
  numMsgs             uint
  sendDone            chan bool
}

func NewKinesisBenchmark(urls []string, topic string, payloadSize uint) *KinesisBenchmark {
  return &KinesisBenchmark{
    payloadSize: payloadSize,
    urls:        urls,
    topic:       topic,
    recv:        make(chan []byte, 65536),
    sendDone:    make(chan bool)
  }
}

func (k KinesisBenchmark) Setup(consumer bool, numMsgs uint) {
  k.numMsgs = numMsgs
  if consumer {
    return k.setupConsumer
  }
  return k.setupProducer()
}

func (k *KinesisBenchmark) setupProducer() error {
  
}

func (k *KinesisBenchmark) setupConsumer() error {
  
}


func (k *KinesisBenchmark) Send() error {
  return nil
}

func (k *KinesisBenchmark) Recv() <-chan []byte {
  return n.recv
}

func (k *KinesisBenchmark) Errors() uint {
  return n.errors
}

func (k *KinesisBenchmark) SendDone() <-chan bool {
  return n.sendDone
}