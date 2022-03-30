package test

import (
	// pachong "github.com/PuerkitoBio/goquery"

	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Anjuke(c *gin.Context) {
	// baidu()
	// anjuke()

	str := "海淀 - 万柳 - 万柳华府北街9号	"

	str1 := strings.Split(str, "-")

	for i := 0; i < len(str1); i++ {
		fmt.Println(strings.TrimSpace(str1[i]))
	}

}

func anjuke() {

	res, err := http.Get("https://beijing.anjuke.com/community/p1/")

	if err != nil {
		fmt.Println("11:" + err.Error())
		return
	}
	// defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("错误码：" + res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("======")
	fmt.Println(doc.Find(".list-left .list-cell .li-row").Length())
	doc.Find(".list-left .list-cell .li-row").Each(func(i int, s *goquery.Selection) {
		// test := s.Text()
		// fmt.Println(test)
		href, _ := s.Attr("href")

		title := s.Find(".li-info .li-title .li-community-title").Text()
		addr := s.Find(".li-info .props span").Last().Text()

		fmt.Println(href + "==" + title + "===" + addr)

	})

	// c.HTML(http.StatusOK, "12", doc)

	return
}

// baidu
func baidu() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// fmt.Println("status code error: %d %s", res.StatusCode, res.Status)

	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doc.Find(".s_form_wrapper").Length())

	doc.Find(".s-hotsearch-content .hotsearch-item").Each(func(i int, s *goquery.Selection) {
		content := s.Find(".title-content-title").Text()
		fmt.Printf("%d: %s\n", i, content)
	})
}

// 安居客，添加 header Cookie 信息，防止被屏蔽
func anjukeTest2() (findcount int) {

	url := "https://beijing.anjuke.com/community/chaoyang/p1/"

	client := &http.Client{}

	// client.Jar.Cookies()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print(err)
	}
	// req.Header.Set("Cookie", resp.Cookies()) //数组报错
	// req.Header.Set("Pragma", "no-cache")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.88")
	// req.Header.Set("Cache-Control", "no-cache")
	// req.Header.Set("Upgrade-insecure-requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
	req.Header.Set("Cookie", "sessid=E89CA661-1012-94BA-68CA-FAD8A26011CE; aQQ_ajkguid=EE71D22E-69D6-9DE1-F6BD-7982BA2033C8; id58=CpQDkGIpkqxnnsScG9mkAg==; 58tj_uuid=b36241bf-fce1-4267-ac2d-5178ba54527c; als=0; _ga=GA1.2.1770530048.1646891693; wmda_uuid=0c3d4f3e69112d36053d44d6b822b682; wmda_new_uuid=1; wmda_visited_projects=%3B6289197098934; __xsptplus8=8.1.1646891725.1646891725.1%234%7C%7C%7C%7C%7C%23%2341e590Ylc4OKmxHrY6tlUD-j-_87GS7Z%23; cmctid=1; _gid=GA1.2.315449468.1648450315; ctid=14; twe=2; ajk-appVersion=; fzq_js_anjuke_ershoufang_pc=c105e23f5d626a0035b0ca59e150d9a4_1648533299727_24; ajk_member_verify=6JmaHasGdXvH%2B%2FFrueFijL8Bkz8ZHhxfRCDStFt%2FVF4%3D; ajk_member_verify2=MjQ1NTA1MzM1fENQOVRPRWp8MQ%3D%3D; fzq_h=6da521e20881740bcda352783d7ecd49_1648537300355_3194efd5ce4d4133bc619ffe08a750c4_2071798790; ajk_member_id=245505335; obtain_by=2; new_uv=9; fzq_js_anjuke_xiaoqu_pc=4609410d2416bf98e8f40a32e39b46fd_1648610292686_23; ajkAuthTicket=TT=0b5ace0bb90e7d24a1aa9311da44e971&TS=1648610293011&PBODY=WHZhV-PwH5fnfDWNPOQr8ASdtHG5lLzAdDHTCEh4sbIO2uZEmjzI33L6z9FX3flA8sw_VNzMAXEFDyNiGR9P3shQk0Z3L0vH-jxAan0uSjQzji_g22ln03uCh54VVwxbkq6b15ih9KZoob0c9CPFsjywwZSyEGw6eFD-CgoXkkY&VER=2&CUID=2MJktpi8BWgMUtXxQOttFgj5svAO2gEh; xxzl_cid=c8dffd75b17a4604a208b997fc1b9e0c; xzuid=90c34bc7-9bff-451e-81e7-7581b588bda4")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp2, err := client.Do(req)
	if err != nil {
		print(err)
	}

	// body, err := ioutil.ReadAll(resp2.Body)

	// fmt.Println(string(body))

	// return

	doc, err := goquery.NewDocumentFromReader(resp2.Body)
	// doc, err := goquery.NewDocumentFromResponse(resp2)

	if err != nil {
		fmt.Println(err.Error())
		return
		// haveerror = true
	}

	findcount = doc.Find(".list-left .list-cell .li-row").Length()

	fmt.Println(findcount)

	return
}
