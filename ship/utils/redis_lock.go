package utils

import (
    "errors"
    "fmt"
    "github.com/go-redis/redis/v7"
    "github.com/satori/go.uuid"
    "time"
)

func GetLock(redisLY *redis.Client, lockName string, acquireTimeout, lockTimeOut time.Duration) (string, error) {
    code := uuid.NewV4().String()
    // endTime := util.FwTimer.CalcMillis(time.Now().Add(acquireTimeout))
    endTime := time.Now().Add(acquireTimeout).UnixNano()
    // for util.FwTimer.CalcMillis(time.Now()) <= endTime {
    for time.Now().UnixNano() <= endTime {
        if success, err := redisLY.SetNX(lockName, code, lockTimeOut).Result(); err != nil && err != redis.Nil {
            return "", err
        } else if success {
            return code, nil
        } else if redisLY.TTL(lockName).Val() == -1 {
            redisLY.Expire(lockName, lockTimeOut)
        }
        time.Sleep(time.Millisecond)
    }
    return "", errors.New("timeout")
}

// var count = 0  // test assist
func ReleaseLock(redisLY *redis.Client, lockName, code string) bool {
    txf := func(tx *redis.Tx) error {
        if v, err := tx.Get(lockName).Result(); err != nil && err != redis.Nil {
            return err
        } else if v == code {
            _, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
                // count++
                // fmt.Println(count)
                pipe.Del(lockName)
                return nil
            })
            return err
        }
        return nil
    }

    for {
        if err := redisLY.Watch(txf, lockName); err == nil {
            return true
        } else if err == redis.TxFailedErr {
            fmt.Println("watch key is modified, retry to release lock. err:", err.Error())
        } else {
            fmt.Println("err:", err.Error())
            return false
        }
    }
}
