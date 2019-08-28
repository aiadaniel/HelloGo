package engine

import "log"

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Loop()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	WorkerCnt int
	Scheduler Scheduler
	ItemChan  chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Loop()

	for i := 0; i < e.WorkerCnt; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCnt := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//var memitem = item.Payload.(model.Memitem)
			//log.Printf("Get item:#%d %v\n",itemCnt,memitem)
			itemCnt++
			//快速轻量，直接丢给saver去做存储。
			//更复杂的做法可以抽象出存储的task去做
			var tempItem = item //NOTE!: 这里需要一个临时变量，解决go闭包传参问题（变量有效期、逃逸等问题）
			go func() { e.ItemChan <- tempItem }()
		}

		//增加方法告知爬虫结束

		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				log.Printf("==error duplicate url: %s", r.Url)
				continue
			}
			//var tempR = r
			//log.Printf("==r.Url: %s",r.Url)
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		//in := make(chan Request)
		for {
			//1 tell scheduler I'am ready
			ready.WorkerReady(in)

			r := <-in
			result, err := Worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
