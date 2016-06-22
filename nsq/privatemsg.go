package nsq

import (
    "fmt"
    "github.com/bitly/go-nsq"
    "xtudouh/common/utils"
)

const topic = "private"

type TextMsg string

func (pm TextMsg) Push() {
    NsqProducer.Publish(topic, utils.Serialize(pm))
}

func init() {
    mh := new(MessageHandler)
    mh.Topic = topic
    mh.Channel = "android"
    mh.Handler = func(m *nsq.Message) error {
        var pm TextMsg
        utils.UnSerialize(m.Body, &pm)
        fmt.Printf("%v\n", pm)
        return nil
    }
    RegisterConsumers(mh)
}
