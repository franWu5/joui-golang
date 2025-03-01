package joui

import (
	"syscall"
	"unsafe"
)

// 创建路径
func (d *DLL) CreatePath() (UIPath, error) {
	addr, err := d.GetProcAddress("jo_UIpath_create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr)
	return UIPath(ret), nil
}

// 创建多边形路径
func (d *DLL) CreatePolygonPath(left, top, right, bottom float32, numberOfEdges uint32, angle float32) (UIPath, error) {
	addr, err := d.GetProcAddress("jo_UIpath_creaTepolygon")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(left),          //float32和float64都兼容
		uintptr(top),           //float32和float64都兼容
		uintptr(right),         //float32和float64都兼容
		uintptr(bottom),        //float32和float64都兼容
		uintptr(numberOfEdges), //float32和float64都兼容
		uintptr(angle),         //float32和float64都兼容
	)
	if ret == 0 { // 假设返回 0 表示失败
		return 0, syscall.GetLastError() // 或者自定义的错误
	}
	return UIPath(ret), nil
}

// 销毁路径
func (d *DLL) DestroyPath(hPath UIPath) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
	)
	return ret != 0, nil
}

// 重置路径
func (d *DLL) ResetPath(hPath UIPath) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_reset")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
	)
	return ret != 0, nil
}

// 取路径边界矩形
func (d *DLL) GetPathBounds(hPath UIPath, lpBounds *RECTF) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_getbounds")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(unsafe.Pointer(lpBounds)),
	)
	return ret != 0, nil
}

// 打开路径
func (d *DLL) OpenPath(hPath UIPath) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_open")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
	)
	return ret != 0, nil
}

// 关闭路径
func (d *DLL) ClosePath(hPath UIPath) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_close")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
	)
	return ret != 0, nil
}

// 开始新图形
func (d *DLL) BeginFigure(hPath UIPath, x, y float32, figureBegin int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_beginfigure")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x),         //float32和float64都兼容
		uintptr(y),         //float32和float64都兼容
		uintptr(figureBegin),
	)
	return ret != 0, nil
}

// 结束当前图形
func (d *DLL) EndFigure(hPath UIPath, fCloseFigure bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_endfigure")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		boolToUintptr(fCloseFigure),
	)
	return ret != 0, nil
}

// 坐标是否在可见路径内
func (d *DLL) PathHitTest(hPath UIPath, x, y float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_hittest")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x),//float32和float64都兼容
		uintptr(y),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加由直线组成的多边形
func (d *DLL) AddPoly(hPath UIPath, pts *H_POINT, count int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addPoly")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(unsafe.Pointer(pts)),
		uintptr(count),
	)
	return ret != 0, nil
}

// 添加直线
func (d *DLL) AddLine(hPath UIPath, x1, y1, x2, y2 float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addline")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加直线到点
func (d *DLL) AddLineTo(hPath UIPath, x, y float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addlineTo")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x),//float32和float64都兼容
		uintptr(y),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加圆弧
func (d *DLL) AddArc(hPath UIPath, x1, y1, x2, y2, radiusX, radiusY float32, fClockwise bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addarc")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x1),      //float32和float64都兼容
		uintptr(y1),      //float32和float64都兼容
		uintptr(x2),      //float32和float64都兼容
		uintptr(y2),      //float32和float64都兼容
		uintptr(radiusX), //float32和float64都兼容
		uintptr(radiusY), //float32和float64都兼容
		boolToUintptr(fClockwise),
	)
	return ret != 0, nil
}

// 添加圆弧.扩展
func (d *DLL) AddArcOval(hPath UIPath, endX, endY, radiusX, radiusY float32, fClockwise, fArcSize bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addarcOval")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(endX),//float32和float64都兼容
		uintptr(endY),//float32和float64都兼容
		uintptr(radiusX),//float32和float64都兼容
		uintptr(radiusY),//float32和float64都兼容
		boolToUintptr(fClockwise),
		boolToUintptr(fArcSize),
	)
	return ret != 0, nil

}

//添加圆形
func (d *DLL) AddCircle(hPath UIPath, x, y, width, height, startAngle, sweepAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addCircle")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x),//float32和float64都兼容
		uintptr(y),//float32和float64都兼容
		uintptr(width), //float32和float64都兼容
		uintptr(height),//float32和float64都兼容
		uintptr(startAngle),//float32和float64都兼容
		uintptr(sweepAngle),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加矩形
func (d *DLL) AddRect(hPath UIPath, left, top, right, bottom float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(left),//float32和float64都兼容
		uintptr(top),//float32和float64都兼容
		uintptr(right),//float32和float64都兼容
		uintptr(bottom),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加圆角矩形
func (d *DLL) AddRoundedRect(hPath UIPath, left, top, right, bottom, radiusTopLeft, radiusTopRight, radiusBottomLeft, radiusBottomRight float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addroundedrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(left),         //float32和float64都兼容
		uintptr(top),          //float32和float64都兼容
		uintptr(right),        //float32和float64都兼容
		uintptr(bottom),       //float32和float64都兼容
		uintptr(radiusTopLeft),   //float32和float64都兼容
		uintptr(radiusTopRight),  //float32和float64都兼容
		uintptr(radiusBottomLeft), //float32和float64都兼容
		uintptr(radiusBottomRight),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加贝塞尔曲线
func (d *DLL) AddBezier(hPath UIPath, x1, y1, x2, y2, x3, y3 float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addbezier")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
		uintptr(x3),//float32和float64都兼容
		uintptr(y3),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加二次方贝塞尔曲线
func (d *DLL) AddQuadraticBezier(hPath UIPath, x1, y1, x2, y2 float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addquadraticbezier")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 添加贝塞尔曲线组
func (d *DLL) AddBeziers(hPath UIPath, points *H_POINT) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addbezies")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(unsafe.Pointer(points)),
	)
	return ret != 0, nil
}

// 添加闭合曲线组
func (d *DLL) AddCurve(hPath UIPath, points *H_POINT, count int32, tension float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIpath_addCurve")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hPath),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(tension),//float32和float64都兼容
	)
	return ret != 0, nil
}