package joui

import (
	"syscall"
	"unsafe"
)

// 创建数组
func (d *DLL) CreateArray() (UIArray, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Create")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return UIArray(ret), nil
}

// 合法性检查
func (d *DLL) IsArrayLegal(pArray UIArray, nIndex uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_IsLegal")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		nIndex,
	)
	return ret != 0, nil
}

// 删除数组
func (d *DLL) DestroyArray(pArray UIArray) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
	)
	return ret != 0, nil
}

// 添加成员
func (d *DLL) AddArrayMember(pArray UIArray, value, index uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIarray_AddMember")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		value,
		index,
	)
	return ret, nil
}

// 删除成员
func (d *DLL) DeleteArrayMember(pArray UIArray, index uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_DelMember")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		index,
	)
	return ret != 0, nil
}

// 重定义数组大小
func (d *DLL) RedefineArray(pArray UIArray, size int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Redefine")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		uintptr(size),
	)
	return ret != 0, nil
}

// 清空数组
func (d *DLL) ClearArray(pArray UIArray) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Clear")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
	)
	return ret != 0, nil
}

// 获取数组成员数量
func (d *DLL) GetArrayCount(pArray UIArray) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIarray_GetCount")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
	)
	return int32(ret), nil
}

// 设置数组成员
func (d *DLL) SetArrayMember(pArray UIArray, index, value uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_SetMember")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		index,
		value,
	)
	return ret != 0, nil
}

// 获取数组成员
func (d *DLL) GetArrayMember(pArray UIArray, index uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIarray_GetMember")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		index,
	)
	return ret, nil
}

// 设置附加参数
func (d *DLL) SetArrayExtra(pArray UIArray, extra uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_SetExtra")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		extra,
	)
	return ret != 0, nil
}

// 获取附加参数
func (d *DLL) GetArrayExtra(pArray UIArray) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIarray_GetExtra")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
	)
	return ret, nil
}

// 枚举数组成员
func (d *DLL) EnumArray(pArray UIArray, callback unsafe.Pointer, pvParam uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Emum")
	if err != nil {
		return 0, err
	}
    /* 不需要NewCallback，直接传入
    var cb uintptr
    if callback != nil {
        cb = syscall.NewCallback(callback)
    }
    */

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		uintptr(callback), //直接传入
		pvParam,
	)
	return ret, nil
}

// 数组排序
func (d *DLL) SortArray(pArray UIArray, fDesc bool, fun ArrayComparePROC, extra1, extra2 uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIarray_Sort")
	if err != nil {
		return false, err
	}
    var callback uintptr
    if fun != nil {
        callback = syscall.NewCallback(fun)
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pArray),
		boolToUintptr(fDesc),
		callback, // 使用转换后的回调
		extra1,
		extra2,
	)
	return ret != 0, nil
}