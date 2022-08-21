package logger

import "go.uber.org/zap"

func NewLogger(paths ...string) (*zap.Logger, error) {
	// 配置开发模式日志输出
	cfg := zap.NewDevelopmentConfig()
	if len(paths) > 0 {
		cfg.OutputPaths = paths
	} else {
		cfg.OutputPaths = []string{
			"stderr",
		}
	}

	// 使用自定义配置构建一个新的logger
	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	// 替换全局的
	_ = zap.ReplaceGlobals(l)

	return l, nil
}
