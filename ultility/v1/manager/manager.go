package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin/config"
	"gin/database/orm"
	"gin/models/v1/manager"
	"gin/tool/helper"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(email, pwd string, c *gin.Context) (err error) {
	manager := manager.Manager{}

	result := orm.MysqlOrm.Where("email = ?", email).First(&manager)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("登录失败")
	} else if helper.StrToMd5(pwd) != manager.Pwd {
		err = errors.New("登录失败")
	} else {
		err = storeManagerSession(manager, c)
	}

	return
}

func storeManagerSession(manager manager.Manager, c *gin.Context) (err error) {
	session := sessions.Default(c)

	// 保存到 session 中，如果将来想保存其他值，在这里修改
	str, err := json.Marshal(manager)

	if err != nil {
		return
	}

	// 存值
	session.Set(config.SESSION_KEY, str)
	session.Options(sessions.Options{
		MaxAge: config.SESSION_MAX_AGE,
		Path:   "/",
	})

	err = session.Save()
	return
}

// 获取当前登录人的 session
func GetManagerSession(c *gin.Context) (manager manager.Manager, err error) {
	session := sessions.Default(c)
	sessionMsg := session.Get(config.SESSION_KEY)
	err = json.Unmarshal(sessionMsg.([]byte), &manager)

	fmt.Println("====")
	fmt.Println(session.ID())
	fmt.Println("====")

	return
}

// 退出登录
func Logut(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(config.SESSION_KEY, nil)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
}

// 添加管理员
func Add(name, phone, email, pwd string) (id uint, err error) {

	manager := manager.Manager{
		Name:  name,
		Phone: phone,
		Email: email,
		Pwd:   helper.StrToMd5(pwd),
	}

	result := orm.MysqlOrm.Create(&manager)

	if result.Error != nil {
		err = errors.New("插入失败")
		return
	}

	id = manager.ID
	return
}
