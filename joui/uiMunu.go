package joui

import (
	"syscall"
	"unsafe"
)

// 创建菜单
func (d *DLL) CreateMenu() (syscall.Handle, error) {
	addr, err := d.GetProcAddress("jo_UImenu_Create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return syscall.Handle(ret), nil // HMENU 是 HANDLE 类型
}

// 销毁菜单
func (d *DLL) DestroyMenu(hMenu syscall.Handle) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImenu_Destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
	)
	return ret != 0, nil
}

// 禁止菜单项目
func (d *DLL) EnableMenuItem(hMenu syscall.Handle, uIDNewItem uint32, wEnable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImenu_EnableItem")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(uIDNewItem),
		boolToUintptr(wEnable),
	)
	return ret != 0, nil
}

// 弹出菜单
func (d *DLL) PopupMenu(hMenu syscall.Handle, uFlags uint32, x, y int32, nReserved uintptr, handle int32, crText, crBackground ARGB, lpRC *RECT, pfnCallback MsgPROC, dwFlags uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImenu_Popup")
	if err != nil {
		return false, err
	}
	var callback uintptr
	if pfnCallback != nil {
		callback = syscall.NewCallback(pfnCallback)
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(x),
		uintptr(y),
		nReserved,
		uintptr(handle),
		uintptr(crText),
		uintptr(crBackground),
		uintptr(unsafe.Pointer(lpRC)),
		callback, // 使用转换后的回调
		uintptr(dwFlags),
	)
	return ret != 0, nil
}

// 关闭菜单
func (d *DLL) EndMenu() (bool, error) {
	addr, err := d.GetProcAddress("jo_UImenu_End")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return ret != 0, nil
}

// 添加菜单
func (d *DLL) AppendMenu(hMenu syscall.Handle, dwFlags uint32, uIDNewItem uint32, newItem string, hbmpItem UIImage) (bool, error) {
	lpNewItem, err := syscall.UTF16PtrFromString(newItem)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UImenu_Append")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(dwFlags),
		uintptr(uIDNewItem),
		uintptr(unsafe.Pointer(lpNewItem)),
		uintptr(hbmpItem),
	)
	return ret != 0, nil
}

// 获取菜单条目标题
func (d *DLL) GetMenuString(hMenu syscall.Handle, uIDNewItem uint32, lpNewString *string) (int32, error) {
	addr, err := d.GetProcAddress("jo_UImenu_GetString")
	if err != nil {
		return 0, err
	}
    var str *uint16 //LPCWSTR*
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(uIDNewItem),
        uintptr(unsafe.Pointer(&str)), // 传递指针的指针
	)
    // 转换为 Go 字符串
    *lpNewString = syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(str))[:])

    //注意这里可能需要释放内存

	return int32(ret), nil
}

// 修改菜单
func (d *DLL) SetMenuItem(hMenu syscall.Handle, dwFlags uint32, uIDNewItem uint32, newItem string, hbmpItem UIImage) (bool, error) {
	lpNewItem, err := syscall.UTF16PtrFromString(newItem)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UImenu_SetItem")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(dwFlags),
		uintptr(uIDNewItem),
		uintptr(unsafe.Pointer(lpNewItem)),
		uintptr(hbmpItem),
	)
	return ret != 0, nil
}

// 删除菜单
func (d *DLL) RemoveMenu(hMenu syscall.Handle, uIDNewItem uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImenu_Remove")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hMenu),
		uintptr(uIDNewItem),
	)
	return ret != 0, nil
}