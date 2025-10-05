package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetProxyFromEnvironment() string {
	// 首先尝试环境变量
	envVars := []string{"HTTPS_PROXY", "https_proxy", "HTTP_PROXY", "http_proxy"}
	for _, envVar := range envVars {
		if proxy := os.Getenv(envVar); proxy != "" {
			fmt.Printf("Using proxy from %s: %s\n", envVar, proxy)
			return proxy
		}
	}

	// 如果环境变量没有设置，尝试获取系统代理设置（仅限 Windows）
	if runtime.GOOS == "windows" {
		if proxy := getWindowsSystemProxy(); proxy != "" {
			fmt.Printf("Using system proxy: %s\n", proxy)
			return proxy
		}
	}

	fmt.Println("No proxy setting found")
	return ""
}

func getWindowsSystemProxy() string {
	cmd := exec.Command("powershell", "-Command", `
        $proxy = Get-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" -Name ProxyEnable -ErrorAction SilentlyContinue
        $proxyServer = Get-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" -Name ProxyServer -ErrorAction SilentlyContinue
        if ($proxy.ProxyEnable -eq 1) {
            Write-Output "http://$($proxyServer.ProxyServer)"
        }
    `)

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	proxy := strings.TrimSpace(string(output))
	if proxy != "" {
		return proxy
	}

	return ""
}

func UseSystemProxy() bool {
	_proxy := GetProxyFromEnvironment()
	if _proxy != "" {
		os.Setenv("HTTP_PROXY", _proxy)
		os.Setenv("HTTPS_PROXY", _proxy)
		return true
	} else {
		return false
	}
}

func ClearSystemProxy() {
	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("HTTPS_PROXY")
}
