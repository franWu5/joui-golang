package joui

import (
	"syscall"
	"unsafe"
)

// --- 画刷操作 ---

// 画刷创建自颜色
func (d *DLL) CreateBrushFromColor(argb ARGB) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_create")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(argb),
	)
	return UIBrush(ret), nil
}

// 画刷创建自画布
func (d *DLL) CreateBrushFromCanvas(hCanvas UICanvas, alpha uint32) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_createfromcv")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(alpha),
	)
	return UIBrush(ret), nil
}

// 画刷创建自图片句柄
func (d *DLL) CreateBrushFromImage(hImg UIImage, mode int32) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_createfromimg")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(mode),
	)
	return UIBrush(ret), nil
}

// 画刷创建自线性渐变
func (d *DLL) CreateLinearGradientBrush(xStart, yStart, xEnd, yEnd float32, crBegin, crEnd ARGB) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_createlinear")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(xStart), //float32和float64都兼容
		uintptr(yStart), //float32和float64都兼容
		uintptr(xEnd),   //float32和float64都兼容
		uintptr(yEnd),   //float32和float64都兼容
		uintptr(crBegin),
		uintptr(crEnd),
	)
	return UIBrush(ret), nil
}
func (d *DLL) CreateLinearGradientBrushEx(xStart, yStart, xEnd, yEnd float32, arrStopPts *int32, arrINT *float32, cStopPts int32) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_createlinearEx")
    if err != nil {
        return nil,err
    }
    ret,_,_ := syscall.SyscallN(addr,
        uintptr(xStart), //float32和float64都兼容
        uintptr(yStart), //float32和float64都兼容
        uintptr(xEnd),   //float32和float64都兼容
        uintptr(yEnd),   //float32和float64都兼容
        uintptr(unsafe.Pointer(arrStopPts)),
        uintptr(unsafe.Pointer(arrINT)),
        uintptr(cStopPts),
        )
    return UIBrush(ret),nil
}

// 画刷创建自径向渐变
func (d *DLL) CreateRadialGradientBrush(x, y, radiusX, radiusY float32, crBegin, crEnd ARGB) (UIBrush, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_createradialgradient")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(x),       //float32和float64都兼容
		uintptr(y),       //float32和float64都兼容
		uintptr(radiusX), //float32和float64都兼容
		uintptr(radiusY), //float32和float64都兼容
		uintptr(crBegin),
		uintptr(crEnd),
	)
	return UIBrush(ret), nil
}
func (d *DLL) CreateRadialGradientBrushEx(x, y, radiusX, radiusY float32, arrStopPts *int32, arrINT *float32, cStopPts int32, gradientOriginOffset H_POINT)(UIBrush,error){
    addr, err := d.GetProcAddress("jo_UIbrush_createradialgradientEx")
    if err != nil {
        return nil,err
    }
    ret,_,_ := syscall.SyscallN(addr,
        uintptr(x),//float32和float64都兼容
        uintptr(y),//float32和float64都兼容
        uintptr(radiusX),//float32和float64都兼容
        uintptr(radiusY),//float32和float64都兼容
        uintptr(unsafe.Pointer(arrStopPts)),
        uintptr(unsafe.Pointer(arrINT)),
        uintptr(cStopPts),
        uintptr(unsafe.Pointer(&gradientOriginOffset)),
        )
    return UIBrush(ret), nil
}

// 画刷销毁
func (d *DLL) DestroyBrush(hBrush UIBrush) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hBrush),
	)
	return ret != 0, nil
}

// 画刷置颜色
func (d *DLL) SetBrushColor(hBrush UIBrush, argb ARGB) (ARGB, error) {
	addr, err := d.GetProcAddress("jo_UIbrush_setcolor")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hBrush),
		uintptr(argb),
	)
	return ARGB(ret), nil // 返回原来的颜色
}

// 设置画刷颜色不透明度
func (d *DLL) SetBrushOpacity(hBrush UIBrush, alpha float32) error {
	addr, err := d.GetProcAddress("jo_UIbrush_SetOpacity")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hBrush),
		uintptr(alpha),//float32和float64都兼容
	)
    return nil
}

// 画刷置矩阵
func (d *DLL) SetBrushTransform(hBrush UIBrush, matrix UIMatrix) error{
	addr, err := d.GetProcAddress("jo_UIbrush_settransform")
    if err != nil {
        return err
    }
    _,_,_ = syscall.SyscallN(addr,
        uintptr(hBrush),
        uintptr(matrix),
        )
    return nil
}

// 设置渐变偏移
func (d *DLL) SetShaderSkewing(hBrush UIBrush, offsetX, offsetY float32) error {
	addr, err := d.GetProcAddress("jo_UIshader_setskewing")
	if err != nil {
		return err
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(hBrush),
		uintptr(offsetX),//float32和float64都兼容
		uintptr(offsetY),//float32和float64都兼容
	)
    return nil
}

// --- 矩阵操作 ---

// 创建矩阵
func (d *DLL) CreateMatrix() (UIMatrix, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_create")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return UIMatrix(ret), nil
}

// 销毁矩阵
func (d *DLL) DestroyMatrix(pMatrix UIMatrix) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pMatrix),
	)
	return ret != 0, nil
}

// 重置矩阵
func (d *DLL) ResetMatrix(pMatrix UIMatrix) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_reset")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pMatrix),
	)
	return ret != 0, nil
}

// 矩阵旋转
func (d *DLL) RotateMatrix(pMatrix UIMatrix, fAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_rotate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pMatrix),
		uintptr(fAngle),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 矩阵缩放
func (d *DLL) ScaleMatrix(pMatrix UIMatrix, scaleX, scaleY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_scale")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pMatrix),
		uintptr(scaleX),//float32和float64都兼容
		uintptr(scaleY),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 矩阵转置
func (d *DLL) TranslateMatrix(pMatrix UIMatrix, offsetX, offsetY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UImatrix_translate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pMatrix),
		uintptr(offsetX),//float32和float64都兼容
		uintptr(offsetY),//float32和float64都兼容
	)
	return ret != 0, nil
}