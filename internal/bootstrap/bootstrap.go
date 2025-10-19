package bootstrap

import (
	"fmt"
	"os"
)

// setEnvIfExists 尝试设置环境变量，路径不存在则忽略。
func setEnvIfExists(key, path string) {
	if st, err := os.Stat(path); err == nil && st.IsDir() {
		if err := os.Setenv(key, path); err == nil {
			fmt.Printf("[bootstrap] %s=%s\n", key, path)
		}
	}
}

func init() {
	// 容错设置宿主机路径
	setEnvIfExists("HOST_PROC", "/host_proc")
	setEnvIfExists("HOST_SYS", "/host_sys")
	setEnvIfExists("HOST_ETC", "/host_etc")
	setEnvIfExists("HOST_VAR", "/host_var")
	setEnvIfExists("HOST_RUN", "/host_run")
	setEnvIfExists("HOST_ROOT", "/host_root")

	// 确保至少有默认值
	if os.Getenv("HOST_PROC") == "" {
		os.Setenv("HOST_PROC", "/proc")
	}
	if os.Getenv("HOST_SYS") == "" {
		os.Setenv("HOST_SYS", "/sys")
	}
	if os.Getenv("HOST_ROOT") == "" {
		os.Setenv("HOST_ROOT", "/")
	}
}
