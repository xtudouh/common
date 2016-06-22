package idg

import (
    "xtudouh/common/conf"
    "xtudouh/common/log"
)

var (
    inst *IdWorker
    l = log.NewLogger()
)

func Id() (int64, error) {
    return inst.NextId()
}

func Ids(num int) ([]int64, error) {
    return inst.NextIds(num)
}

func Init() {
    l.Debug("idg module start initializing.")
    datacenterId := conf.Int64("idg", "DATACENTER")
    worker, err := NewIdWorker(int64(1), datacenterId, twepoch)
    if err != nil {
        panic(err)
    }

    inst = worker
    l.Debug("idg module initialize successfully.")
}
