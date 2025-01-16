// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe2-skeleton/internal/model"
)

type (
	ISysConfig interface {
		// InitConfig 初始化系统配置
		InitConfig(ctx context.Context)
		// LoadConfig 加载系统配置
		LoadConfig(ctx context.Context) (err error)
		// GetLoadToken 获取本地token配置
		GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
