package matchers

import (
	"../search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

/*
1.搜索rss源的匹配器
2.这个程序里的匹配器，是指包含特定信息、用于处理某类数据源的实例
*/

// 解码RSS文档需要用到的4个结构类型

type (
	// item 根据item字段的标签，将定义的字段与rss文档的字段关联
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image 根据image字段标签，将定义的字段与rss文档的字段关联
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel 根据channel字段的标签，将定义的字段与rss文档的字段关联
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher 实现matcher接口
type rssMatcher struct {
}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// retrive 发送HTTP Get请求获取rss数据并解码
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("no rss feed URI provided")
	}

	// 从网络获得rss数据文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 一旦从函数返回，关闭该响应的连接
	defer resp.Body.Close()

	// 检查状态码是否为200
	if resp.StatusCode != 200 {
		// 返回一个自定义类型的错误
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// 将rss数据源文档解码到我们定义的结果类型里
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

// Search 在文档中查找特定的搜索项
func (m rssMatcher) Search(feed *search.Feed, searchItem string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URi[%s]\n\n", feed.Type, feed.Name, feed.URI)

	// 获取想要搜索的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	// 在标题和描述中搜索
	for _, channelItem := range document.Channel.Item {
		// 检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchItem, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 如果找到匹配的项
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title})
		}

		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchItem, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
