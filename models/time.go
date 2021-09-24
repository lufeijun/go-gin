package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type MyTime time.Time // 导致时间不会自动更新了

func MyTimeInit() (mytime MyTime) {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	t1, _ := time.Parse("2006-01-02 15:04:05", timeNow)

	mytime = MyTime(t1)
	return
}
func (t *MyTime) UnmarshalJSON(data []byte) error {
	fmt.Println("time: un-json")
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	fmt.Println("time: MarshalJSON")

	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	fmt.Println("time: Value")
	// tTime := time.Time(t)
	// fmt.Println(t)
	tTime := time.Time(time.Now())
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	fmt.Println("time: Scan")
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	fmt.Println("string")
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
