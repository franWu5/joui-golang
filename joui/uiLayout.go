package joui

import (
	"syscall"
	"unsafe"
)

// 创建布局
func (d *DLL) CreateLayout(nType int32, hObjBind int32) (UILayout, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(nType),
		uintptr(hObjBind),
	)
	return UILayout(ret), nil
}

// 销毁布局
func (d *DLL) DestroyLayout(hLayout UILayout) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
	)
	return ret != 0, nil
}

// 更新布局
func (d *DLL) UpdateLayout(hLayout UILayout) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_update")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
	)
	return ret != 0, nil
}

// 取布局类型
func (d *DLL) GetLayoutType(hLayout UILayout) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_gettype")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
	)
	return int32(ret), nil
}

// 是否允许更新
func (d *DLL) EnableLayoutUpdate(hLayout UILayout, fUpdateable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_enableupdate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		boolToUintptr(fUpdateable),
	)
	return ret != 0, nil
}

// 分发通知
func (d *DLL) LayoutNotify(hLayout UILayout, nEvent int32, wParam, lParam uintptr) uintptr {
	addr, err := d.GetProcAddress("jo_UIlayout_notify")
	if err != nil {
		return 0 // 或者返回一个表示错误的特定值
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(nEvent),
		wParam,
		lParam,
	)
	return ret
}

// 置表格信息
func (d *DLL) SetTableInfo(hLayout UILayout, aRowHeight *int32, cRows int32, aCellWidth *int32, cCells int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_table_setinfo")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(unsafe.Pointer(aRowHeight)),
		uintptr(cRows),
		uintptr(unsafe.Pointer(aCellWidth)),
		uintptr(cCells),
	)
	return ret != 0, nil
}

// 置子属性
func (d *DLL) SetChildProperty(hLayout UILayout, parent UIView, dwPropID int32, pvValue uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_setchildprop")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(parent),
		uintptr(dwPropID),
		pvValue,
	)
	return ret != 0, nil
}

// 取子属性
func (d *DLL) GetChildProperty(hLayout UILayout, parent UIView, dwPropID int32, pvValue *uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_getchildprop")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(parent),
		uintptr(dwPropID),
		uintptr(unsafe.Pointer(pvValue)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 取子属性列表  (需要修改 C++ 代码才能在 Go 中使用)
func (d *DLL) GetChildPropertyList(hLayout UILayout, parent UIView) (unsafe.Pointer, error) {
	// 由于 C++ 的 LPVOID* 在 Go 中没有直接对应的类型，
	// 因此无法直接调用此函数。
	// 你需要修改 C++ 代码，提供一个不使用 LPVOID* 的接口。
	//return nil, errors.New("jo_UIlayout_getchildproplist cannot be called directly from Go")

	addr, err := d.GetProcAddress("jo_UIlayout_getchildproplist")
	if err != nil {
		return nil, err
	}
	var props unsafe.Pointer // 先假设可以返回
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(parent),
		uintptr(unsafe.Pointer(&props)),
	)

	if ret != 0 { //获取成功
		return props, nil
	}

	return nil, err
}

// 置属性
func (d *DLL) SetProperty(hLayout UILayout, dwPropID int32, pvValue uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_setprop")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(dwPropID),
		pvValue,
	)
	return ret != 0, nil
}

// 取属性
func (d *DLL) GetProperty(hLayout UILayout, dwPropID int32) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_getprop")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(dwPropID),
	)
	return ret, nil
}

// 取属性列表 (需要修改 C++ 代码才能在 Go 中使用)
func (d *DLL) GetPropertyList(hLayout UILayout) (unsafe.Pointer, error) {
	// 和 GetChildPropertyList 类似，需要修改 C++ 代码
	//return nil, errors.New("jo_UIlayout_getproplist cannot be called directly from Go")
	addr, err := d.GetProcAddress("jo_UIlayout_getproplist")
	if err != nil {
		return nil, err
	}
	var props unsafe.Pointer // 先假设可以返回
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(unsafe.Pointer(&props)),
	)
	if ret != 0 {
		return props, nil
	}
	return nil, err
}

// 绝对布局按当前位置锁定
func (d *DLL) LockAbsoluteLayout(hLayout UILayout, hObjChild UIView, tLeft, tTop, tRight, tBottom, tWidth, tHeight int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_absolute_lock")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(hObjChild),
		uintptr(tLeft),
		uintptr(tTop),
		uintptr(tRight),
		uintptr(tBottom),
		uintptr(tWidth),
		uintptr(tHeight),
	)
	return ret != 0, nil
}

// 绝对布局置边界信息
func (d *DLL) SetAbsoluteLayoutEdge(hLayout UILayout, hObjChild UIView, dwEdge, dwType int32, nValue uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_absolute_setedge")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(hObjChild),
		uintptr(dwEdge),
		uintptr(dwType),
		nValue,
	)
	return ret != 0, nil
}

// 添加控件
func (d *DLL) AddChild(hLayout UILayout, parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_addchild")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(parent),
	)
	return ret != 0, nil
}

// 布局加入所有控件
func (d *DLL) AddChildren(hLayout UILayout, fDesc bool, dwObjClass string, nCount *int32) (bool, error) {
    var lpwzObjClass *uint16 = nil
	if dwObjClass != "" {
		lpwzClassName, err := syscall.UTF16PtrFromString(dwObjClass)
		if err != nil {
			return false, err
		}
		lpwzObjClass = lpwzClassName
	}

	addr, err := d.GetProcAddress("jo_UIlayout_addchildren")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		boolToUintptr(fDesc),
		uintptr(unsafe.Pointer(lpwzObjClass)),
		uintptr(unsafe.Pointer(nCount)), // 传递指针
	)
	return ret != 0, nil
}

// 删除控件
func (d *DLL) DeleteChild(hLayout UILayout, parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_deletechild")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(parent),
	)
	return ret != 0, nil
}

// 删除所有控件
func (d *DLL) DeleteChildren(hLayout UILayout, dwObjClass string) (bool, error) {

    var lpwzObjClass *uint16 = nil
	if dwObjClass != ""{
		lpwzClassName, err := syscall.UTF16PtrFromString(dwObjClass)
		if err != nil {
			return false, err
		}
        lpwzObjClass = lpwzClassName
	}

	addr, err := d.GetProcAddress("jo_UIlayout_deletechildren")
    if err != nil {
        return false, err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
		uintptr(unsafe.Pointer(lpwzObjClass)),
	)
	return ret != 0, nil
}

// 取控件数量
func (d *DLL) GetChildCount(hLayout UILayout) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIlayout_getchildcount")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hLayout),
	)
	return int32(ret), nil
}