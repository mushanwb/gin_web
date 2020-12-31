package config

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/Utils"
)

func init() {
	Utils.Add("database", Utils.StrMap{
		"mysql": map[string]interface{}{
			// 数据库连接信息
			"host":     Utils.Env("DB_HOST", "127.0.0.1"),
			"port":     Utils.Env("DB_PORT", "3306"),
			"database": Utils.Env("DB_DATABASE", ""),
			"username": Utils.Env("DB_USERNAME", ""),
			"password": Utils.Env("DB_PASSWORD", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": Utils.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": Utils.Env("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     Utils.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}
