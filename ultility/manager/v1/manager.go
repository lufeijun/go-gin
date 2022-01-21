package v1

import (
	"errors"
	"gin/database"
	managerModel "gin/models/manager/v1"

	"gin/tool/common"

	"gorm.io/gorm"
)

// 登录
func Login(email, pwd string) (err error) {
	manager := managerModel.Manager{}
	result := database.GormDB.Where("email = ? ", email).First(&manager)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("邮箱有误")
	} else if result.Error != nil {
		err = errors.New("登录失败")
	} else if manager.Pwd != common.StrToMd5(pwd) {
		err = errors.New("密码错误")
	}

	return
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
