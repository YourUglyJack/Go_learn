package search

import (
	"log"
	"sync"
)

/*
1.执行搜索的主控制逻辑
*/

// 大写开头的是公开的，小写开头的是私有的

// 注册用于搜索的匹配器映射，包级变量
var matchers = make(map[string]Matcher)

// Run执行搜索逻辑
func Run(searchTerm string) {

	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup，以便处理所有数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待的处理
	// 每个数据源的goroutine数量
	waitGroup.Add(len(feeds)) // 有几个数据源，就开启几个协程

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取匹配器
		matcher, exists := matchers[feed.TYPE]
		if !exists {
			matcher = matchers['default']
		}

		// 启动一个goroutine来搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine监控所有工作是否完成
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		close(results)
	}()

	// 启动函数，显示返回结果
	Display(results)

}
