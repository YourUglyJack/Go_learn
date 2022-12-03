package search

import "log"

/*
搜索数据用的默认匹配器
*/

// defaultMatcher 实现了默认匹配器
type defaultMatcher struct {
	//空结构在创建实例时，不会分配任何内存。这种结构很适合创建没有任何状态的类型。
	//对于默认匹配器来说，不需要维护任何状态，所以我们只要实现对应的接口就行。
}

// init 将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现默认匹配器的行为，使用defaultMatcher类型的值作为接收者
func (m defaultMatcher) Search(feed *Feed, searchItem string) ([]*Result, error) {
	// 由于defaultMatcher是空结构，没有内存，不需要维护状态，所以传值就可以，一般情况下是指针
	return nil, nil
}

// Register 注册匹配器，供程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatal(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
