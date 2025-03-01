package joui

import (
	"fmt"
	"syscall"
	"unsafe"
	
)

// 辅助函数：将bool转换为uintptr
func boolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// GetDisplayFrequency 获取屏幕刷新率
func (d *DLL) GetDisplayFrequency() (uint32, error) {
	addr, err := d.GetProcAddress("jo_UIrefresh_display_frequency")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr)
	return uint32(ret), nil
}

// Init 初始化引擎渲染接口
func (d *DLL) Init(config *H_INITINFO, device int) error {
	addr, err := d.GetProcAddress("jo_UIinit")
	if err != nil {
		return err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(config)),
		uintptr(device))

	if ret == 0 {
		return fmt.Errorf("初始化引擎失败")
	}
	return nil
}

// Run 执行信号槽
func (d *DLL) Run() error {
	addr, err := d.GetProcAddress("jo_UIrun")
	if err != nil {
		return err
	}

	_, _, _ = syscall.SyscallN(addr)
	return nil
}

// ReadZipSourceByName 使用字符串名称读取zip资源, 并处理返回的资源数据
func (d *DLL) ReadZipSourceByName(zFile UIZip, name string) ([]byte, uint32, error) {
    namePath, err := syscall.UTF16PtrFromString(name)
    if err != nil {
        return nil, 0, fmt.Errorf("转换文件名失败: %v", err)
    }

    var retData uintptr
    var uncompressedSize uint32

    addr, err := d.GetProcAddress("jo_UIzip_ReadSource")
    if err != nil {
        return nil, 0, err
    }

    ret, _, _ := syscall.SyscallN(addr,
        uintptr(zFile),
        uintptr(unsafe.Pointer(namePath)),
        uintptr(unsafe.Pointer(&retData)),
        uintptr(unsafe.Pointer(&uncompressedSize)))

    if ret == 0 {
        return nil, 0, fmt.Errorf("读取zip资源失败")
    }

    // 将retData转换为[]byte
    data := make([]byte, uncompressedSize)
    copy(data, (*[1<<30]byte)(unsafe.Pointer(retData))[:uncompressedSize:uncompressedSize])

    return data, uncompressedSize, nil
}
// EnumZipFileByPath 使用字符串路径枚举zip资源文件
func (d *DLL) EnumZipFileByPath(hRes UIZip, path string, callback func(string, uintptr) bool, param uintptr) error {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return fmt.Errorf("转换路径失败: %v", err)
	}

	// 创建一个包装回调函数，将*uint16转换为string并处理返回值类型
	wrappedCallback := func(pathPtr *uint16, param uintptr) uintptr {
		path := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(pathPtr))[:])
		if callback(path, param) {
			return 1
		}
		return 0
	}

	addr, err := d.GetProcAddress("jo_UIzip_EnumFile")
	if err != nil {
		return err
	}

    	_, _, _ = syscall.SyscallN(addr,
		uintptr(hRes),
		uintptr(unsafe.Pointer(pathPtr)),
		syscall.NewCallback(wrappedCallback), // 使用NewCallback
		uintptr(param))


	return nil
}


// LoadTheme 从文件加载主题包（无密码版本）
func (d *DLL) LoadTheme(path string, isDefault bool) (UIZip, error) {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, fmt.Errorf("转换路径失败: %v", err)
	}
    addr, err := d.GetProcAddress("jo_UItheme_LoadPsw")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(pathPtr)),
		boolToUintptr(isDefault))

	return UIZip(ret), nil
}

// LoadZipByPath 使用字符串路径从文件加载zip资源
func (d *DLL) LoadZipByPath(path string, password string) (UIZip, error) {
    pathBytes := []byte(path + "\x00") //直接go string 转成c的char
    var passwordPtr *byte = nil

    if password != "" {
        passwordBytes := []byte(password + "\x00") //直接go string 转成c的char
        passwordPtr = &passwordBytes[0]
    }
    addr, err := d.GetProcAddress("jo_UIzip_Load")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(&pathBytes[0])),
		uintptr(unsafe.Pointer(passwordPtr)))

	return UIZip(ret), nil
}

// ReadRcSourceByName 使用字符串名称读取rc资源
func (d *DLL) ReadRcSourceByName(name string, resType string) ([]byte, error) {
	resTypePtr, err := syscall.UTF16PtrFromString(resType) //直接转
	if err != nil {
		return nil, fmt.Errorf("转换资源类型失败: %v", err)
	}

	nameU16, err := syscall.UTF16FromString(name)  //直接转
	if err != nil {
		return nil, fmt.Errorf("转换资源名称失败: %v", err)
	}
	if len(nameU16) == 0 {
		return nil, fmt.Errorf("资源名称不能为空")
	}
    addr, err := d.GetProcAddress("jo_UIreadRcSource")
	if err != nil {
		return nil, err
	}
	var data []byte //切片
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(nameU16[0]),
		uintptr(unsafe.Pointer(resTypePtr)),
		uintptr(unsafe.Pointer(&data)))

	if ret == 0 {
		return nil, fmt.Errorf("读取RC资源失败")
	}

	return data, nil
}


// FreeZip 释放zip资源
func (d *DLL) FreeZip(hRes UIZip) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIzip_Free")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr, uintptr(hRes))
	return ret != 0, nil
}

// GetZipCount 获取资源包文件总数
func (d *DLL) GetZipCount(hRes UIZip) (int64, error) {
	addr, err := d.GetProcAddress("jo_UIzip_GetCount")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr, uintptr(hRes))
	return int64(ret), nil
}

// LoadZipFromMemory 从内存加载zip资源
func (d *DLL) LoadZipFromMemory(data unsafe.Pointer, size uintptr, password *byte) (UIZip, error) {
	addr, err := d.GetProcAddress("jo_UIzip_Loadfrommemory")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(data),
		size,
		uintptr(unsafe.Pointer(password)))

	return UIZip(ret), nil
}


// LoadThemeFromMemory 从内存加载主题包
func (d *DLL) LoadThemeFromMemory(data unsafe.Pointer, size uintptr, password *byte, isDefault bool) (UIZip, error) {
	addr, err := d.GetProcAddress("jo_UItheme_Loadfrommemory")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(data),
		size,
		uintptr(unsafe.Pointer(password)),
		boolToUintptr(isDefault))

	return UIZip(ret), nil
}



// FreeTheme 释放主题
func (d *DLL) FreeTheme(hTheme UIZip) (bool, error) {
	addr, err := d.GetProcAddress("jo_UItheme_Free")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr, uintptr(hTheme))
	return ret != 0, nil
}

// GetThemeValuePtr 获取控件属性值指针,atomClass, atomProp传入的是字符串
func (d *DLL) GetThemeValuePtr(hTheme UIZip, atomClass, atomProp string) (unsafe.Pointer, error) {
   atomClassPtr, err := syscall.UTF16PtrFromString(atomClass)
	if err != nil {
		return nil, fmt.Errorf("转换atomClass失败: %v", err)
	}

	atomPropPtr, err := syscall.UTF16PtrFromString(atomProp)
	if err != nil {
		return nil, fmt.Errorf("转换atomProp失败: %v", err)
	}

    addr, err := d.GetProcAddress("jo_UItheme_GetValuePtr")
	if err != nil {
		return nil, err
	}

	// 调用系统函数获取主题值指针
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hTheme),
		uintptr(unsafe.Pointer(atomClassPtr)),
		uintptr(unsafe.Pointer(atomPropPtr)))

	// 检查返回值是否为空,并做安全转换
	if ret == 0 {
		return nil, fmt.Errorf("获取主题值指针失败")
	}

	return unsafe.Pointer(ret), nil
}


// 文件对话框
func (d *DLL) FileOpenDialog(hWnd syscall.Handle, bOpenFileDialog bool, filter, defExt, fileName string) (bool, string, error) {
	lpszFilter, err := syscall.UTF16PtrFromString(filter)
	if err != nil {
		return false, "", err
	}
	lpszDefExt, err := syscall.UTF16PtrFromString(defExt)
	if err != nil {
		return false, "", err
	}
	lpszFileName, err := syscall.UTF16PtrFromString(fileName)
	if err != nil {
		return false, "", err
	}

	var retstrFile *uint16 // 用于接收返回文件名的指针

	addr, err := d.GetProcAddress("jo_UIfileOpenDlg")
	if err != nil {
		return false, "", err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
		boolToUintptr(bOpenFileDialog),
		uintptr(unsafe.Pointer(lpszFilter)),
		uintptr(unsafe.Pointer(lpszDefExt)),
		uintptr(unsafe.Pointer(lpszFileName)),
		uintptr(unsafe.Pointer(&retstrFile)), // 传递指针的指针
	)

	if ret == 0 {
		return false, "", nil
	}

	// 将返回的 *uint16 转换为 Go 字符串
	result := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(retstrFile))[:])

    // 释放内存 (如果 DLL 分配了内存)
    // 注意：这里需要根据 DLL 的具体实现来决定是否需要释放内存，以及如何释放。
    //       如果 DLL 使用了 LocalAlloc，这里应该使用 syscall.LocalFree。
    //       如果 DLL 使用了其他分配器，你需要找到匹配的释放函数。
    //       如果没有文档说明，你需要反编译 DLL 来确定。
    //       下面是一个假设的释放函数：
    //  _, _, _ = syscall.SyscallN(freeFuncAddr, uintptr(unsafe.Pointer(retstrFile)))

	return true, result, nil
}

// 对话框_浏览文件夹
func (d *DLL) BrowseForFolder(hWnd syscall.Handle, title string) (bool, string, error) {
	lpszTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return false, "", err
	}

	var retstrFile *uint16

	addr, err := d.GetProcAddress("jo_UIfileOpenFolder")
    if err != nil {
        return false,"", err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpszTitle)),
		uintptr(unsafe.Pointer(&retstrFile)), // 传递指针的指针
	)

	if ret == 0 {
		return false, "", nil
	}
	result := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(retstrFile))[:])

     // 释放内存 (如果 DLL 分配了内存，参见 FileOpenDialog 中的注释)
	return true, result, nil
}

// 信息框
func (d *DLL) MessageBox(handle int32, text, caption string, uType int32, checkBox string, checkBoxChecked *bool, dwFlags int32,
	crText, crBackground ARGB, dwMilliseconds int32, lpfnMsgProc MsgPROC) (int32, error) {

	lpText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return 0, err
	}
    var lpCaption *uint16 = nil
	if caption != "" {
		lpCaption, err = syscall.UTF16PtrFromString(caption)
		if err != nil {
			return 0, err
		}
	}

    var lpCheckBox *uint16 = nil

	if checkBox != "" {
		lpCheckBox, err = syscall.UTF16PtrFromString(checkBox)
		if err != nil {
			return 0, err
		}
	}

	var checked uintptr
	if checkBoxChecked != nil {
		checked = uintptr(unsafe.Pointer(checkBoxChecked)) //直接传递指针
	}

    var callback uintptr
    if lpfnMsgProc != nil {
        callback = syscall.NewCallback(lpfnMsgProc)
    }

	addr, err := d.GetProcAddress("jo_UImsg")
	if err != nil {
		return 0, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType),
		uintptr(unsafe.Pointer(lpCheckBox)),
		checked,
		uintptr(dwFlags),
		uintptr(crText),
		uintptr(crBackground),
		uintptr(dwMilliseconds),
		callback,
	)
	return int32(ret), nil
}

// 提示框
func (d *DLL) MessageTips(handle int32, text string, uType int32, dwMilliseconds, ptOffset int32) (bool, error) {
	lpText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr,err := d.GetProcAddress("jo_UImsgTips")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(uType),
		uintptr(dwMilliseconds),
		uintptr(ptOffset),
	)
	return ret != 0, nil
}

// 加载光标资源
func (d *DLL) LoadCursor(data []byte, cursorName string) error {
	lpCursorName, err := syscall.UTF16PtrFromString(cursorName)
	if err != nil {
		return err
	}
   var dataPtr unsafe.Pointer
   var dataLen uintptr
    if len(data) > 0 {
        dataPtr = unsafe.Pointer(&data[0])
        dataLen = uintptr(len(data))
    }
	addr, err := d.GetProcAddress("jo_UIloadCursor")
    if err != nil {
        return err
    }
	_, _, _ = syscall.SyscallN(addr,
		uintptr(dataPtr),
		dataLen,
		uintptr(unsafe.Pointer(lpCursorName)),
	)
    return nil //jo_UIloadCursor 没有返回值
}