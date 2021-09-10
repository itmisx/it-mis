package monitor

import (
	"it-mis/init/db"
	"it-mis/internal/model"
	"it-mis/internal/pkg/errorx"
	"it-mis/internal/scopes"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct{}

// 任务列表
func (t Task) List(c *gin.Context) (total int64, list []model.MonitorTask) {
	var (
		count int64
		db    = db.New()
		cond  model.MonitorTask
	)
	c.ShouldBindBodyWith(&cond, binding.JSON)
	db = db.Model(&model.MonitorTask{})

	name := strings.Trim(cond.Name, " ")
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	monitorType := cond.MonitorType
	if monitorType > 0 {
		db = db.Where("monitor_type = ?", monitorType)
	}

	status := cond.Status
	if status > 0 {
		db = db.Where("status = ?", cond.Status)
	}
	tx := db.Session(&gorm.Session{})
	// 查询总条数
	tx.Count(&count)
	// 获取分页数据
	tx.Scopes(scopes.Paginate(c)).
		Find(&list)
	return count, list
}

// 新建任务
func (t Task) Add(c *gin.Context) error {
	var (
		count    int64
		task     model.MonitorTask
		validate = validator.New()
	)
	c.ShouldBindBodyWith(&task, binding.JSON)
	// 任务名称不能为空
	if strings.Trim(task.Name, " ") == "" {
		return errorx.New("名称不能为空", 305001)
	}
	// 监控类型是否合法
	if task.MonitorType == 0 {
		return errorx.New("监控类型不能为空", 305002)
	}
	if validate.Var(task.Port, "gt=0,lt=65535") != nil {
		return errorx.New("", 305003)
	}
	// 监控地址是否合法
	if validate.Var(task.Host, "required,ip|url") != nil {
		return errorx.New("监控地址格式有误", 305004)
	}
	// 创建时间
	task.CreateTime = time.Now().Unix()
	// 更新时间
	task.UpdateTime = time.Now().Unix()
	count = db.New().Model(&model.MonitorTask{}).Create(&task).RowsAffected
	if count > 0 {
		return nil
	} else {
		return errorx.New("新增监控失败", 305005)
	}
}

// 编辑任务
func (t Task) Edit(c *gin.Context) {
}

// 使能任务
func (t Task) Enable(c *gin.Context) {
}

// 删除任务
func (t Task) Delete(c *gin.Context) {
}
