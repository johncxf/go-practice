package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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

// 爬取指定 url 的页面
func HttpGet(url string) (result string, err error) {
	client := &http.Client{}
	resp, err1 := http.NewRequest("GET", url, nil)
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.115")
	response, _ := client.Do(resp)

	if err1 != nil {
		err = err1
		return
	}
	defer response.Body.Close()

	buf := make([]byte, 4096)
	// 循环爬取整页数据
	for {
		n, err2 := response.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

// 写入文件
func Savefile(idx int, filmName, filmScore, peopleNum [][]string) {
	path := outputDir + "page_" + strconv.Itoa(idx) + ".txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	n := len(filmName)
	f.WriteString("电影名称" + "\t\t\t" + "评分" + "\t\t" + "评分人数" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[i][1] + "\t\t\t" + filmScore[i][1] + "\t\t" + peopleNum[i][1] + "\n")
	}
}

// 爬取一个豆瓣页面数据信息
func SpiderPage(idx int, page chan int) {
	// 获取 url
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="

	// 爬取 url 对应页面
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}

	// 电影名称
	ret1 := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	filmName := ret1.FindAllStringSubmatch(result, -1)

	// 分数
	ret2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	filmScore := ret2.FindAllStringSubmatch(result, -1)

	// 评分人数
	ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	peopleNum := ret3.FindAllStringSubmatch(result, -1)

	// 写入文件
	Savefile(idx, filmName, filmScore, peopleNum)

	page <- idx
}

// 爬取
func Spider(start, end int) {
	// 当前时间
	startTime := time.Now()
	fmt.Printf("正在爬取 %d 到 %d 页...\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页爬取完毕\n", <-page)
	}

	// 计算爬取耗时
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
