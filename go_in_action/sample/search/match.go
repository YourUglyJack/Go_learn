package search

import (
	"fmt"
	"log"
)

/*
用于支持不同匹配器的接口
*/

// Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// 接口声明了结构类型或者具名类型需要实现的行为
// 如果接口类型只包含一个方法，那么这个类型的名字以 er 结尾

// Matcher 定义了要实现的搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchItem string) ([]*Result, error)
}

// Match 为每个数据源单独启动一个goroutine来这行这个函数，并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchItem string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchItem)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display 从每个单独的goroutine接收到结果后，在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直阻塞，直到有结果写入
	// 一旦通道被关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
