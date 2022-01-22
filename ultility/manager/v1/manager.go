package v1

import (
	"encoding/json"
	"errors"
	"gin/config"
	"gin/database"
	managerModel "gin/models/manager/v1"

	"gin/tool/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 登录
func Login(email, pwd string, c *gin.Context) (err error) {
	manager := managerModel.Manager{}
	result := database.GormDB.Where("email = ? ", email).First(&manager)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("邮箱有误")
	} else if result.Error != nil {
		err = errors.New("登录失败")
	} else if manager.Pwd != common.StrToMd5(pwd) {
		err = errors.New("密码错误")
	} else {
		err = storeManagerSession(manager, c)

	}

	return
}

func storeManagerSession(manager managerModel.Manager, c *gin.Context) (err error) {
	session := sessions.Default(c)

	str, err := json.Marshal(manager)

	if err != nil {
		return
	}

	// 存值
	session.Set(config.SESSION_KEY, str)
	session.Options(sessions.Options{
		MaxAge: config.SESSION_MAX_AGE,
	})

	err = session.Save()
	return
}

func getManagerSession(c *gin.Context) (manager managerModel.Manager, err error) {
	session := sessions.Default(c)
	sessionMsg := session.Get(config.SESSION_KEY)
	err = json.Unmarshal(sessionMsg.([]byte), &manager)
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(config.SESSION_KEY, nil)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
}

func Add(name, phone, email, pwd string) (id uint, err error) {

	manager := managerModel.Manager{
		Name:  name,
		Phone: phone,
		Email: email,
		Pwd:   common.StrToMd5(pwd),
	}

	result := database.GormDB.Create(&manager)

	if result.Error != nil {
		err = errors.New("插入失败")
		return
	}

	id = manager.ID
	return
}
