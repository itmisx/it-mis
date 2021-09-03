package member

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/errorx"

	"github.com/gin-gonic/gin"
)

// 成员管理
type Member struct{}

type UserInfo struct {
	ID     int64  `json:"id"`
	User   string `json:"user"`
	Status int    `json:"status"`
}

// Add 添加成员
func (m Member) Add(c *gin.Context) {
}

// Delete 删除成员
func (m Member) Delete(c *gin.Context) {
}

// Update 更新成员
func (m Member) Update(c *gin.Context) {
}

// List 成员列表
func (m Member) List(c *gin.Context) {
}

// Info 成员信息
func (m Member) Info(userID int64) (userInfo *UserInfo, err error) {
	var (
		db   = db.New()
		user = &model.User{}
	)
	db.Model(&model.User{}).
		Select("id", "user", "status").
		Where("id = ?", userID).
		Find(user)

	if user.ID <= 0 {
		return nil, errorx.New("", 301001)
	}
	return &UserInfo{
		ID:     user.ID,
		User:   user.User,
		Status: user.Status,
	}, nil
}

// 部门管理
type Department struct{}

// Add 添加成员
func (d Department) Add(c *gin.Context) {
}

// Delete 删除成员
func (d Department) Delete(c *gin.Context) {
}

// Update 更新
func (d Department) Update(c *gin.Context) {
}

// List 成员列表
func (d Department) List(c *gin.Context) {
}

// 组织管理
type Group struct{}

// Add 添加成员
func (g Group) Add(c *gin.Context) {
}

// Delete 删除成员
func (g Group) Delete(c *gin.Context) {
}

// Update 更新
func (g Group) Update(c *gin.Context) {
}

// List 成员列表
func (g Group) List(c *gin.Context) {
}
