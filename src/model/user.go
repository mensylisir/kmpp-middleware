package model

import (
	"errors"

	"github.com/mensylisir/kmpp-middleware/src/model/common"
	"github.com/mensylisir/kmpp-middleware/src/util/encrypt"
	uuid "github.com/satori/go.uuid"
)

var (
	AdminCanNotDelete = "ADMIN_CAN_NOT_DELETE"
	LdapCanNotUpdate  = "LDAP_CAN_NOT_UPDATE"
)

type User struct {
	common.BaseModel
	ID       string `json:"id" gorm:"type:varchar(64)"`
	Name     string `json:"name" gorm:"type:varchar(256);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(256)"`
	IsAdmin  bool   `json:"-" gorm:"type:boolean;default:false"`
	Role     int    `json:"role" gorm:"type:varchar(64)"`
	IsActive bool   `json:"-" gorm:"type:boolean;default:true"`
	Type     string `json:"type" gorm:"type:varchar(64)"`
}

type Token struct {
	Token string `json:"access_token"`
}

func (u *User) BeforeCreate() (err error) {
	u.ID = uuid.NewV4().String()
	return err
}

func (u *User) BeforeDelete() (err error) {
	if u.Name == "admin" {
		return errors.New(AdminCanNotDelete)
	}
	return nil
}

func (u *User) ValidateOldPassword(password string) (bool, error) {
	oldPassword, err := encrypt.StringDecrypt(u.Password)
	if err != nil {
		return false, err
	}
	if oldPassword != password {
		return false, err
	}
	return true, err
}
