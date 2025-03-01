package joui

import (
	"syscall"
	"unsafe"
)
//注意：由于DLL里面使用了std::vector<LPCTSTR>，在go里面无法直接使用，所以这个函数无法使用，或者需要修改DLL。
/*
// 枚举系统字体族列表  (无法直接在 Go 中实现，需要修改 C++ 代码)
func (d *DLL) EnumerateFonts() ([]string, error) {
	// 由于 C++ 的 std::vector<LPCTSTR> 在 Go 中没有直接对应的类型，
	// 因此无法直接调用此函数。
	// 你需要修改 C++ 代码，提供一个不使用 std::vector 的接口。
	// 例如，可以提供一个回调函数，或者返回一个 C 风格的字符串数组。
	return nil, errors.New("jo_UIfont_EnumerateFonts cannot be called directly from Go")
}
*/

// 创建默认字体
func (d *DLL) CreateDefaultFont() (UIFont, error) {
	addr, err := d.GetProcAddress("jo_UIfont_create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return UIFont(ret), nil
}

// 创建字体自字体族
func (d *DLL) CreateFontFromFamily(fontFace string, fontSize int32, fontStyle uint32) (UIFont, error) {
	lpwzFontFace, err := syscall.UTF16PtrFromString(fontFace)
	if err != nil {
		return 0, err
	}
	addr, err := d.GetProcAddress("jo_UIfont_createfromfamily")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpwzFontFace)),
		uintptr(fontSize),
		uintptr(fontStyle),
	)
	return UIFont(ret), nil
}

// 创建字体自逻辑字体
func (d *DLL) CreateFontFromLogFont(logFont *LOGFONTW) (UIFont, error) {
	addr, err := d.GetProcAddress("jo_UIfont_createfromlogfont")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(logFont)),
	)
	return UIFont(ret), nil
}

// 创建字体文件自内存
func (d *DLL) CreateFontFromMemory(fontData []byte, fontFace string) (bool, error) {
	lpwzFontFace, err := syscall.UTF16PtrFromString(fontFace)
	if err != nil {
		return false, err
	}
    var fontDataPtr unsafe.Pointer
    var fontDataSize uintptr
    if len(fontData) > 0 {
        fontDataPtr = unsafe.Pointer(&fontData[0])
        fontDataSize = uintptr(len(fontData))
    }
	addr, err := d.GetProcAddress("jo_UIfont_createFromMem")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(fontDataPtr),
		fontDataSize,
		uintptr(unsafe.Pointer(lpwzFontFace)),
	)
	return ret != 0, nil
}

// 字体是否存在
func (d *DLL) IsFontFamilyAvailable(fontFamilyName string, dwFontStyle uint32, fontWeight int32) (bool, error) {
	lpwzFontFamilyName, err := syscall.UTF16PtrFromString(fontFamilyName)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIfont_IsFontFamilyAvailable")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpwzFontFamilyName)),
		uintptr(dwFontStyle),
		uintptr(fontWeight),
	)
	return ret != 0, nil
}

// 字体销毁
func (d *DLL) DestroyFont(hFont UIFont) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfont_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hFont),
	)
	return ret != 0, nil
}

// 字体取描述表
func (d *DLL) GetFontContext(hFont UIFont) (unsafe.Pointer, error) {
	addr, err := d.GetProcAddress("jo_UIfont_getcontext")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hFont),
	)
	return unsafe.Pointer(ret), nil
}

// 字体取高度
func (d *DLL) GetFontHeight(hFont UIFont) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIfont_Height")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hFont),
	)
	return int32(ret), nil
}

// 字体旋转
func (d *DLL) SetFontEscapement(hFont UIFont, lfEscapement int32) error {
	addr, err := d.GetProcAddress("jo_UIfont_Capement")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hFont),
		uintptr(lfEscapement),
	)
    return nil
}

// 字体复制
func (d *DLL) CopyFont(hFont UIFont) (UIFont, error) {
	addr, err := d.GetProcAddress("jo_UIfont_copy")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hFont),
	)
	return UIFont(ret), nil
}

// 获取逻辑字体
func (d *DLL) GetLogFont(hFont UIFont, logFont *LOGFONTW) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfont_getlogfont")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hFont),
		uintptr(unsafe.Pointer(logFont)),
	)
	return ret != 0, nil
}