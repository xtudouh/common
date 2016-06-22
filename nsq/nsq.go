package nsq

import (
    "bytes"
    "github.com/bitly/go-nsq"
    "xtudouh/common/conf"
    "xtudouh/common/log"
)

var (
    prodAddr, consAddr string
    NsqProducer        *nsq.Producer
    config             *nsq.Config
    consumers          []MessageHandler
    l = log.NewLogger()
)

type MessageHandler struct {
    Topic   string
    Channel string
    Handler nsq.HandlerFunc
}

func (mh *MessageHandler) ListenAndServe() error {
    consumer, err := nsq.NewConsumer(mh.Topic, mh.Channel, config)
    if err != nil {
        return err
    }
    consumer.AddHandler(mh.Handler)
    err = consumer.ConnectToNSQD(consAddr)
    if err != nil {
        return err
    }

    return nil
}

func Push(topic string, msg ...*nsq.Message) {
    if len(msg) == 0 {
        return
    }
    msgBodies := make([][]byte, len(msg))
    for i, m := range msg {
        buf := new(bytes.Buffer)
        m.WriteTo(buf)
        msgBodies[i] = buf.Bytes()
    }
    NsqProducer.MultiPublish(topic, msgBodies)
}

func RegisterConsumers(mh *MessageHandler) {
    consumers = append(consumers, *mh)
}

func Init() {
    l.Debug("nsq module start initializing.")
    prodAddr = conf.String("nsq", "PROD_ADDRESS")
    consAddr = conf.String("nsq", "CONS_ADDRESS")
    config = nsq.NewConfig()
    var err error
    NsqProducer, err = nsq.NewProducer(prodAddr, config)
    if err != nil {
        panic(err)
    }
    for _, con := range consumers {
        err = con.ListenAndServe()
        if err != nil {
            panic(err)
        }
    }
    l.Debug("redis module initialize successfully.")
}
