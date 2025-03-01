package joui

import (
	"syscall"
)

// 区域销毁
func (d *DLL) DestroyRegion(hRgn UIRgn) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hRgn),
	)
	return ret != 0, nil
}

// 区域创建自圆角矩形
func (d *DLL) CreateRegionFromRoundRect(left, top, right, bottom, radiusX, radiusY float32) (UIRgn, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_createfromroundrect")
	if err != nil {
		return nil, err // 返回 nil 和 error
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(left),   //float32和float64都兼容
		uintptr(top),    //float32和float64都兼容
		uintptr(right),  //float32和float64都兼容
		uintptr(bottom), //float32和float64都兼容
		uintptr(radiusX),//float32和float64都兼容
		uintptr(radiusY),//float32和float64都兼容
	)
	if ret == 0 {
		return nil, syscall.GetLastError() // 或者自定义错误
	}
	return UIRgn(ret), nil
}

// 区域创建自矩形
func (d *DLL) CreateRegionFromRect(left, top, right, bottom float32) (UIRgn, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_createfromrect")
	if err != nil {
		return nil, err // 返回 nil 和 error
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(left),  //float32和float64都兼容
		uintptr(top),   //float32和float64都兼容
		uintptr(right), //float32和float64都兼容
		uintptr(bottom),//float32和float64都兼容
	)
    if ret == 0 {
		return nil, syscall.GetLastError() // 或者自定义错误
	}
	return UIRgn(ret), nil
}

// 区域创建自路径
func (d *DLL) CreateRegionFromPath(hPath UIPath) (UIRgn, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_createfrompath")
	if err != nil {
		return nil, err // 返回 nil 和 error
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
	)
    if ret == 0 {
		return nil, syscall.GetLastError() // 或者自定义错误
	}
	return UIRgn(ret), nil
}

// 区域合并
func (d *DLL) CombineRegions(hRgnSrc, hRgnDst UIRgn, nCombineMode, dstOffsetX, dstOffsetY int32) (UIRgn, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_combine")
	if err != nil {
		return nil, err // 返回 nil 和 error
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hRgnSrc),
		uintptr(hRgnDst),
		uintptr(nCombineMode),
		uintptr(dstOffsetX),
		uintptr(dstOffsetY),
	)

    if ret == 0 {
		return nil, syscall.GetLastError() // 或者自定义错误
	}
	return UIRgn(ret), nil
}

// 区域命中测试
func (d *DLL) RegionHitTest(hRgn UIRgn, x, y float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIrgn_hittest")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hRgn),
		uintptr(x),//float32和float64都兼容
		uintptr(y),//float32和float64都兼容
	)
	return ret != 0, nil
}