package timer

import (
	"container/list"
	"container/ring"
	"time"
)

// TimeWheelCallback 时间轮回调函数类型
type TimeWheelCallback func([]interface{})

// TimeWheel 时间轮主体
type TimeWheel struct {
	//
	interval time.Duration
	// 时间轮
	ticker *time.Ticker
	// 回调函数
	callback TimeWheelCallback
	//
	buckets *ring.Ring
	//
	currentPos int
	// 请求管道
	requestChannel chan interface{}
	// 时间轮结束管道
	quitChannel chan interface{}
}

// BucketItem 时间轮触发点
type BucketItem struct {
	interval time.Duration
	value    interface{}
}

// NewTimeWhell 新建一个时间轮
func NewTimeWhell(interval time.Duration, bucketCount int, callback TimeWheelCallback) *TimeWheel {
	timeWheel := &TimeWheel{
		buckets:        ring.New(bucketCount),
		interval:       interval,
		currentPos:     0,
		callback:       callback,
		requestChannel: make(chan interface{}, 100),
		quitChannel:    make(chan interface{}, 1),
	}

	for i := 1; i <= timeWheel.buckets.Len(); i++ {
		timeWheel.buckets.Value = list.New()
		timeWheel.buckets = timeWheel.buckets.Next()
	}

	return timeWheel
}

// Start 运行一个时间轮
func (timeWheel *TimeWheel) Start() {
	timeWheel.ticker = time.NewTicker(timeWheel.interval * time.Second)
	go timeWheel.run()
}

// Add 向时间轮中添加一个触发点
func (timeWheel *TimeWheel) Add(interval time.Duration, item ...interface{}) {
	timeWheel.requestChannel <- &BucketItem{
		interval: interval,
		value:    item,
	}
}

// Stop 停止一个时间轮
func (timeWheel *TimeWheel) Stop() {
	close(timeWheel.quitChannel)
}

func (timeWheel *TimeWheel) run() {
	for {
		select {
		case <-timeWheel.quitChannel:
			timeWheel.ticker.Stop()
			break
		case <-timeWheel.ticker.C:
			if nil != timeWheel.callback {
				buckets := timeWheel.buckets.Value.(*list.List)
				userDatas := make([]interface{}, 0, buckets.Len())
				var n *list.Element
				for v := buckets.Front(); v != nil; v = n {
					if bucketItem := v.Value.(*BucketItem); nil != bucketItem {

						//insertBucket := timeWheel.buckets.Move(int(bucketItem.interval.Nanoseconds()) / int(timeWheel.interval))
						//itemList, _ := insertBucket.Value.(*list.List)
						//itemList.PushBack(bucketItem)

						userDatas = append(userDatas, bucketItem.value)
						n = v.Next()
						buckets.Remove(v)
					}
				}
				if len(userDatas) > 0 {
					timeWheel.callback(userDatas)
				}
			}
			timeWheel.buckets = timeWheel.buckets.Next()
		case item := <-timeWheel.requestChannel:
			if bucketItem, _ := item.(*BucketItem); nil != bucketItem {
				insertBucket := timeWheel.buckets.Move(int(bucketItem.interval.Nanoseconds()) / int(timeWheel.interval))
				if itemList, _ := insertBucket.Value.(*list.List); nil != itemList {
					itemList.PushBack(bucketItem)
				}
			}
		}
	}
}
