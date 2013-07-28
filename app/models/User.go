package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/base64"
	"fmt"
	"github.com/pa001024/MoeWorker/util"
	r "github.com/robfig/revel"
	"regexp"
	"time"

	"net/url"
	"strings"
)

type User struct {
	Id          int64
	Username    string `qbs:"size:32,unique,notnull"`
	DisplayName string `qbs:"size:64"`
	Password    string `qbs:"size:80,notnull"`
	Email       string `qbs:"size:100,unique,notnull"`
	AvatarEmail string `qbs:"size:100"`
	Url         string `qbs:"size:100"`
	Status      int8   // 0: 未验证 1: 用户 2: 组织
	Created     time.Time
	Updated     time.Time
}

// enum User.Status
const (
	UnauthedUser = iota
	UnauthedTeam
	AuthedUser
	AuthedTeam
	SysAdmin
)

var (
	userRegex  = regexp.MustCompile(`^[A-z0-9_\-]*$`)
	emailRegex = regexp.MustCompile(`\w+(?:[-+.]\w+)*@\w+(?:[-.]\w+)*\.\w+(?:[-.]\w+)*`)
)

func (this *User) Validate(v *r.Validation, password string) {
	const (
		EMAIL         = "电子邮件地址"
		USERNAME      = "用户名"
		PASSWORD      = "密码"
		INVALID       = "不符合要求"
		NOEMPTY       = "不能为空"
		ALREADYEXISTS = "已存在"
	)
	v.Required(this.Email).Message(EMAIL + NOEMPTY)
	v.Check(this.Email, r.MinSize{6}, r.MaxSize{100}, r.Match{emailRegex}).
		Message(EMAIL + INVALID)
	v.Required(this.Username).Message(USERNAME + NOEMPTY)
	v.Check(this.Username, r.MinSize{2}, r.MaxSize{32}, r.Match{userRegex}).
		Message(USERNAME + INVALID)
	v.Required(password).Message(PASSWORD + NOEMPTY)
	v.Check(password, r.MaxSize{32}, r.MinSize{6}).
		Message(PASSWORD + INVALID)

}

func (this *User) ValidatePassword(v *r.Validation, password string) {
	bin := []byte(_APPSECRET + password + this.Username)
	hbin, _ := base64.StdEncoding.DecodeString(this.Password)
	err := bcrypt.CompareHashAndPassword(hbin, bin)
	if err != nil {
		v.Error("密码不匹配")
	}
}

func (this *User) GeneratePassword(password string) string {
	bin := []byte(_APPSECRET + password + this.Username)
	b, _ := bcrypt.GenerateFromPassword(bin, bcrypt.DefaultCost)
	s := base64.StdEncoding.EncodeToString(b) // len = 80 = 64 hash 16 salt
	return s
}

func (this *User) GetAvatarUrl(size string) string {
	return fmt.Sprintf("https://secure.gravatar.com/avatar/%s?%s",
		util.Md5String(strings.ToLower(this.AvatarEmail)),
		(url.Values{
			"s": {size},
			"d": {"retro"}, // TODO: 自定义
		}).Encode())
}

func (this *User) String() string {
	return fmt.Sprintf("User(%s)", this.Username)
}

func (this *User) Logined() string {
	return this.Updated.Format("2006年1月2日")
}
func (this *User) Joined() string {
	return this.Created.Format("2006年1月2日")
}
