package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var outputDir = "./tmp/"

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 初始化环境
func InitEnv() {
	if FileExists(outputDir) {
		rErr := os.RemoveAll(outputDir)
		if rErr != nil {
			fmt.Println("清理目录失败:", rErr)
			return
		}
		fmt.Println("清理目录")
	}
	mErr := os.Mkdir(outputDir, os.ModePerm)
	if mErr != nil {
		fmt.Println("创建目录失败:", mErr)
		return
	}
	fmt.Println("创建目录:", outputDir)
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	time.Sleep(time.Second)

	// 循环读取 网页数据， 传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		// 累加每一次循环读到的 buf 数据，存入result 一次性返回。
		result += string(buf[:n])
	}
	return
}

// 爬取单个页面的函数
func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}

	// 将读到的整网页数据，保存成一个文件
	f, err := os.Create(outputDir + "page_" + strconv.Itoa(i) + ".html")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	f.WriteString(result)
	f.Close() // 保存好一个文件，关闭一个文件。

	page <- i // 与主go程完成同步。
}

// 爬取页面操作。
func Spider(start, end int) {
	startTime := time.Now()
	fmt.Printf("正在爬取第%d页到%d页....\n", start, end)

	page := make(chan int)

	// 循环爬取每一页数据
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 个页面爬取完成\n", <-page)
	}
	endTime := time.Now()
	consume := endTime.Sub(startTime).Seconds()
	fmt.Println("爬取耗时(s)：", consume)
}

func main() {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页（>=start）:")
	fmt.Scan(&end)

	// 初始化目录
	InitEnv()

	// 爬取
	Spider(start, end)
}
