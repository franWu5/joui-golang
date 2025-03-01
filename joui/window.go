package joui

import (
	"syscall"
	"unsafe"
)

// 创建单窗口
func (d *DLL) CreateWindow(hWndParent syscall.Handle, className, windowName string, x, y, width, height int, dwStyle, dwStyleEx int32) (syscall.Handle, error) {
	var lpwzClassName, lpwzWindowName *uint16
	var err error

	if className != "" {
		lpwzClassName, err = syscall.UTF16PtrFromString(className)
		if err != nil {
			return 0, err
		}
	} // else lpwzClassName remains nil, which is equivalent to C's NULL

	if windowName != "" {
		lpwzWindowName, err = syscall.UTF16PtrFromString(windowName)
		if err != nil {
			return 0, err
		}
	} // else lpwzWindowName remains nil

	addr, err := d.GetProcAddress("jo_UIwnd_Create")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWndParent),
		uintptr(unsafe.Pointer(lpwzClassName)), //  NULL if className is empty
		uintptr(unsafe.Pointer(lpwzWindowName)), //  NULL if windowName is empty
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(dwStyle),
		uintptr(dwStyleEx),
	)
	return syscall.Handle(ret), nil
}

// 从XML布局创建
func (d *DLL) LoadFromXml(hWndParent syscall.Handle, hTheme, hRes UIZip, lpLayoutXml []byte, nCmdShow int) (bool, syscall.Handle, UIWnd, error) {

    var rethWnd syscall.Handle
    var retParentUI UIWnd

    addr, err := d.GetProcAddress("jo_UIwnd_loadXml")
	if err != nil {
		return false, 0, 0, err
	}

	var lpLayoutXmlPtr unsafe.Pointer
    var dwLen uintptr
    if len(lpLayoutXml) > 0 {
        lpLayoutXmlPtr = unsafe.Pointer(&lpLayoutXml[0])
        dwLen = uintptr(len(lpLayoutXml))
    }

    ret, _, _ := syscall.SyscallN(addr,
        uintptr(hWndParent),
        uintptr(hTheme),
        uintptr(hRes),
        uintptr(lpLayoutXmlPtr),
        dwLen,
        uintptr(nCmdShow),
        uintptr(unsafe.Pointer(&rethWnd)),    // 注意这里传递指针的指针
		uintptr(unsafe.Pointer(&retParentUI)), // 注意这里传递指针的指针
    )
    return ret != 0, rethWnd, retParentUI, nil
}

// 保存窗口控件到XML文件
func (d *DLL) SaveXmlFile(handle int32, savePath string) (bool, error) {
	pszSavePath, err := syscall.UTF16PtrFromString(savePath)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIwnd_SaveXmlFile")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(unsafe.Pointer(pszSavePath)),
	)
	return ret != 0, nil
}

// 注册XML事件
func (d *DLL) RegXmlCallback(name string, lpfnMsgProc unsafe.Pointer) (bool, error) {
	// Go 中通常不直接使用 C 风格的函数指针，而是使用 Go 的函数值。
	// 这里假设你有一个 Go 函数与 C 的 MsgPROC 类型兼容。
	// 你需要使用 syscall.NewCallback 将 Go 函数转换为 C 回调。

    namePtr := unsafe.StringData(name)
	addr, err := d.GetProcAddress("jo_UIwnd_RegXmlCallback")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(lpfnMsgProc), //直接传入，不需要NewCallback
	)
	return ret != 0, nil
}

// 绑定窗口
func (d *DLL) BindWindow(hWnd syscall.Handle, hTheme UIZip, wwsStyle uint32, lParam uintptr, lpfnMsgProc WinMsgPROC) (UIWnd, error) {

	addr, err := d.GetProcAddress("jo_UIwnd_Bind")
	if err != nil {
		return 0, err
	}

	var callback uintptr
    if lpfnMsgProc != nil {
        callback = syscall.NewCallback(lpfnMsgProc)
    }

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
		uintptr(hTheme),
		uintptr(wwsStyle),
		lParam,
		callback, // 使用转换后的回调函数指针
	)

	return UIWnd(ret), nil
}

// 销毁窗口
func (d *DLL) DestroyWindow(hWnd syscall.Handle) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_Destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
	)
	return ret != 0, nil
}

// 注册窗口类名
func (d *DLL) RegClass(className string, hIcon, hIconSm syscall.Handle) (uint16, error) {
	lpwzClassName, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	addr, err := d.GetProcAddress("jo_UIwnd_RegClass")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpwzClassName)),
		uintptr(hIcon),
		uintptr(hIconSm),
	)

	return uint16(ret), nil
}

// 注册窗口Hook消息 成功返回一个HK_THUNK_DATA结构
func (d *DLL) RegHook(hWnd syscall.Handle, pfnProc ThunkPROC, dwData unsafe.Pointer) (unsafe.Pointer, error) {

	addr, err := d.GetProcAddress("jo_UIwnd_RegHook")
	if err != nil {
		return nil, err
	}
    var callback uintptr
    if pfnProc != nil {
        callback = syscall.NewCallback(pfnProc)
    }

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
		callback,
		uintptr(dwData),
	)

	return unsafe.Pointer(ret), nil // 返回的指针需要根据实际情况转换回 HK_THUNK_DATA
}

// 取窗口功能键
func (d *DLL) GetWindowKeys() (int32, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_GetKeys")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return int32(ret), nil
}

// 设置窗口参数
func (d *DLL) SetWindowLong(hParentUI UIWnd, nIndex int32, dwNewlong1, dwNewlong2, dwNewlong3 uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_SetLong")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(nIndex),
		dwNewlong1,
		dwNewlong2,
		dwNewlong3,
	)
	return ret, nil
}

// 获取窗口参数
func (d *DLL) GetWindowLong(hParentUI UIWnd, nIndex int32) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_GetLong")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(nIndex),
	)

	return ret, nil
}

// 获取客户区矩形
func (d *DLL) GetClientRect(hParentUI UIWnd, lpClientRect *RECT) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_GetClientRect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(unsafe.Pointer(lpClientRect)),
	)
	return ret != 0, nil
}

// 设置可视
func (d *DLL) ShowWindow(hParentUI UIWnd, nCmdShow int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_Show")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(nCmdShow),
	)
	return ret != 0, nil
}

// 查询可视
func (d *DLL) IsWindowVisible(hWnd syscall.Handle) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_IsVisible")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
	)
	return ret != 0, nil
}

// 设置窗口启用状态
func (d *DLL) EnableWindow(hWnd syscall.Handle, bEnable, bUpdate bool) error {
	addr, err := d.GetProcAddress("jo_UIwnd_Enable")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hWnd),
		boolToUintptr(bEnable),
		boolToUintptr(bUpdate),
	)
	return nil
}

// 更新窗口
func (d *DLL) UpdateWindow(hParentUI UIWnd) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_Update")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
	)
	return ret != 0, nil
}

// 弹出托盘图标
func (d *DLL) TrayIconPopup(hParentUI UIWnd, info, infoTitle string, dwInfoFlags int32) (bool, error) {
	lpwzInfo, err := syscall.UTF16PtrFromString(info)
	if err != nil {
		return false, err
	}
	lpwzInfoTitle, err := syscall.UTF16PtrFromString(infoTitle)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIwnd_TrayIconPopup")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(unsafe.Pointer(lpwzInfo)),
		uintptr(unsafe.Pointer(lpwzInfoTitle)),
		uintptr(dwInfoFlags),
	)
	return ret != 0, nil
}

// 设置托盘图标
func (d *DLL) SetTrayIcon(hParentUI UIWnd, hIcon syscall.Handle, tips string) (bool, error) {

	lpwzTips, err := syscall.UTF16PtrFromString(tips)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIwnd_TrayIconSet")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(lpwzTips)),
	)
	return ret != 0, nil
}

// 获取鼠标所在窗口控件句柄
func (d *DLL) GetObjFromPoint(handle, x, y int32) (UIView, error) {

	addr, err := d.GetProcAddress("jo_UIwnd_GetObjFromPoint")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(x),
		uintptr(y),
	)
	return UIView(ret), nil
}

// 设置背景信息
func (d *DLL) SetBackgImage(hParentUI UIWnd, image []byte, x, y int32, dwRepeat uint32, lpGrid *RECT, dwFlags uint32, dwAlpha uint32, fUpdate bool) (bool, error) {
    var imagePtr unsafe.Pointer
    var imageLen uintptr

    if len(image) > 0 {
        imagePtr = unsafe.Pointer(&image[0])
        imageLen = uintptr(len(image))
    }
	addr, err := d.GetProcAddress("jo_UIwnd_SetBackgImage")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(imagePtr),
		imageLen,
		uintptr(x),
		uintptr(y),
		uintptr(dwRepeat),
		uintptr(unsafe.Pointer(lpGrid)),
		uintptr(dwFlags),
		uintptr(dwAlpha),
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 获取背景信息
func (d *DLL) GetBackgImage(hParentUI UIWnd, backgroundImage unsafe.Pointer) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_GetBackgImage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(backgroundImage),
	)
	return ret != 0, nil

}

// 设置背景播放状态.
func (d *DLL) SetBackgPlayState(hParentUI UIWnd, fPlayFrames, fResetFrame, fUpdate bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_SetBackgPlayState")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		boolToUintptr(fPlayFrames),
		boolToUintptr(fResetFrame),
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 设置模糊
func (d *DLL) SetBlur(hParentUI UIWnd, fDeviation float32, lprc *RECT, bRedraw bool) error {
	addr, err := d.GetProcAddress("jo_UIwnd_SetBlur")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(fDeviation), //float32和float64都兼容
		uintptr(unsafe.Pointer(lprc)),
		boolToUintptr(bRedraw),
	)
    return nil
}

// 从窗口句柄获取引擎句柄
func (d *DLL) GetUIWndFromHWnd(hWnd syscall.Handle) (UIWnd, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_FromParentUI")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
	)
	return UIWnd(ret), nil
}

// 窗口居中
func (d *DLL) CenterWindowFrom(hWnd, hWndFrom syscall.Handle, bFullScreen bool) error {
	addr, err := d.GetProcAddress("jo_UIwnd_CenterFrom")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hWnd),
		uintptr(hWndFrom),
		boolToUintptr(bFullScreen),
	)

    return nil
}

// 设置标题
func (d *DLL) SetWindowText(hParentUI UIWnd, text string, fTleRedraw bool) (bool, error) {
	lpString, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIwnd_SetText")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(unsafe.Pointer(lpString)),
		boolToUintptr(fTleRedraw),
	)
	return ret != 0, nil
}

// 获取标题
func (d *DLL) GetWindowText(hParentUI UIWnd, buffer []uint16, fTleGet bool) (uintptr, error) {

	addr, err := d.GetProcAddress("jo_UIwnd_GetText")
	if err != nil {
		return 0, err
	}
    //var buffer [4096]uint16 //固定长度
    nMaxCount := len(buffer)
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(nMaxCount),
		boolToUintptr(fTleGet),
	)
	return ret, nil
}

// 获取标题文本长度
func (d *DLL) GetWindowTextLength(hParentUI UIWnd, fTleGet bool) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIwnd_GetTextLength")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hParentUI),
		boolToUintptr(fTleGet),
	)
	return ret, nil
}