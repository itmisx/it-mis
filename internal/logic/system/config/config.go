package system

import (
	"it-mis/init/db"
	"it-mis/internal/model"
)

type Config struct {
}

// 设置配置项
func (c Config) Set(key string, value string, encrypType string) (string, error) {
	var (
		conf model.SystemConfig
		db   = db.New()
	)
	// FirstOrCreate
	db.Where(&model.SystemConfig{
		Key: key,
	}).Attrs()

	err := db.Model(&model.SystemConfig{}).
		Where("key = ?", key).
		Take(&conf).Error
	if err != nil {
		return "", err
	}
	return conf.Value, nil
}

// 获取配置项
func (Config) Get(key string) (string, error) {
	var (
		conf model.SystemConfig
		db   = db.New()
	)
	err := db.Model(&model.SystemConfig{}).
		Where("key = ?", key).
		Take(&conf).Error
	if err != nil {
		return "", err
	}
	return conf.Value, nil
}

type SecureToken struct {
}

// EnableSecureLogin 启用安全登录验证
func (SecureToken) EnableSecureLogin() {

}

// DisableSecureLogin 禁用安全登录验证
func (SecureToken) DisableSecureLogin() {

}

// GetSecureLoginConfig 获取安全登录配置
// 如果开启了安全登录，则需要二次验证（邮箱）
func (SecureToken) SecureLoginEnableStatus() (enabled bool) {
	return true
}
