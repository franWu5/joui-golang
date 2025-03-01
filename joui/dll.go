// Package joui provides Go bindings for the jo_UI library.
package joui

import (
	"fmt"
	"syscall"
)

// DLL 结构体用于管理动态链接库
type DLL struct {
	handle syscall.Handle
	path   string
}

// LoadDLL 加载指定路径的DLL文件
func LoadDLL(dllPath string) (*DLL, error) {
	handle, err := syscall.LoadLibrary(dllPath)
	if err != nil {
		return nil, fmt.Errorf("加载DLL失败: %v", err)
	}

	return &DLL{
		handle: handle,
		path:   dllPath,
	}, nil
}

// GetProcAddress 获取DLL中的函数地址
func (d *DLL) GetProcAddress(procName string) (uintptr, error) {
	addr, err := syscall.GetProcAddress(d.handle, procName)
	if err != nil {
		return 0, fmt.Errorf("获取函数地址失败: %v", err)
	}
	return addr, nil
}

// Free 卸载DLL
func (d *DLL) Free() error {
	if d.handle != 0 {
		err := syscall.FreeLibrary(d.handle)
		if err != nil {
			return fmt.Errorf("卸载DLL失败: %v", err)
		}
		d.handle = 0
	}
	return nil
}

// IsLoaded 检查DLL是否已加载
func (d *DLL) IsLoaded() bool {
	return d.handle != 0
}

// GetPath 获取DLL文件路径
func (d *DLL) GetPath() string {
	return d.path
}