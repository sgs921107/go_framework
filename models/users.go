/*************************************************************************
> File Name: book.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 20:27:15 星期四
> Content: This is a desc
*************************************************************************/

package models

type Users struct {
	BaseModel
	// 用户名
	Username string `gorm:"unique;column:username;type:varchar(100);not null;comment:用户名" json:"username"`
	// 昵称
	Nickname string `gorm:"unique;column:nickname;type:varchar(16);not null;comment:昵称" json:"nickname"`
	// 密码
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码sha256" json:"password"`
	// 邮箱
	Email string `gorm:"unique;column:email;type:varchar(100);not null;default:'';comment:邮箱" json:"email"`
	// 手机号
	Phone string `gorm:"unique;column:phone;type:varchar(100);not null;default:'';comment:手机号" json:"phone"`
	// 性别
	Gender int `gorm:"column:gender;type:tinyint;not null;default:0;check:gender in (0,1,2);comment:性别:0未知1男2女" json:"gender"`
	// 头像url
	Avatar string `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像url" json:"avatar"`
	// 角色
	Role int `gorm:"column:role;type:tinyint;not null;default:0;comment:用户角色:0普通用户/1商家/2管理员" json:"role"`
}

type UserOutput struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
}

// 插入用户
func (u *Users) Insert() error {
	return db.Create(&u).Error
}

// 更新数据
func (u *Users) Update(data interface{}) error {
	return db.Model(&u).Where("id = ?", u.ID).Updates(data).Error
}

// 通过用户名查询用户
func (u *Users) GetByUsername() error {
	return db.Model(&u).Select("id,password,role").Where("username = ?", u.Username).First(&u).Error
}

// 通过id查询用户
func (u *Users) GetByID() error {
	return db.Model(&u).Select(
		"username,password,nickname,email,phone,gender,avatar",
	).Where("id = ?", u.ID).First(&u).Error
}

func (u *Users) Output() *UserOutput {
	return &UserOutput{
		ID:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Email:    u.Email,
		Phone:    u.Phone,
		Gender:   u.Gender,
		Avatar:   u.Avatar,
	}
}
