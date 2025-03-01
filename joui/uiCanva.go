package joui

import (
	"syscall"
	"unsafe"
)

// 画布创建
func (d *DLL) CreateCanvas(handle int32, width, height int32) (UICanvas, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(width),
		uintptr(height),
	)
	return UICanvas(ret), nil
}

// 画布销毁
func (d *DLL) DestroyCanvas(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 画布输出为Bloom效果
func (d *DLL) OutCanvasBloom(hCanvas, canvas UICanvas, scRGB ARGB, padding H_POINT, fblur float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_OutCanvasBloom")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(canvas),
		uintptr(scRGB),
		uintptr(unsafe.Pointer(&padding)),
		uintptr(fblur), //float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布保存为图像
func (d *DLL) CanvasToImage(hCanvas UICanvas, dstImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_Toimg")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(unsafe.Pointer(dstImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 画布模糊
func (d *DLL) BlurCanvas(hCanvas UICanvas, fDeviation float32, lprc *RECT) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_blur")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(fDeviation),//float32和float64都兼容
		uintptr(unsafe.Pointer(lprc)),
	)
	return ret != 0, nil
}

// 画布测量文本
func (d *DLL) CalcTextSize(hCanvas UICanvas, hFont UIFont, text string, dwLen int, dwDTFormat uint32, layoutWidth, layoutHeight float32, lpWidth, lpHeight *float32) (bool, error) {
	lpwzText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIcanvas_calctextsize")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hFont),
		uintptr(unsafe.Pointer(lpwzText)),
		uintptr(dwLen),
		uintptr(dwDTFormat),
		uintptr(layoutWidth), //float32和float64都兼容
		uintptr(layoutHeight),//float32和float64都兼容
		uintptr(unsafe.Pointer(lpWidth)),
		uintptr(unsafe.Pointer(lpHeight)),
	)
	return ret != 0, nil
}

// 画布开始绘制
func (d *DLL) BeginDraw(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_begindraw")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 画布结束绘制
func (d *DLL) EndDraw(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_enddraw")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 画布清除
func (d *DLL) ClearCanvas(hCanvas UICanvas, color ARGB) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_clear")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(color),
	)
	return ret != 0, nil
}

// 画布置剪辑区
func (d *DLL) ClipRect(hCanvas UICanvas, left, top, right, bottom int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_cliprect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
	)
	return ret != 0, nil
}

// 画布重置剪辑区
func (d *DLL) ResetClip(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_resetclip")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 画布从画布复制
func (d *DLL) DrawCanvas(hCanvas, sCanvas UICanvas, dstLeft, dstTop, dstRight, dstBottom, srcLeft, srcTop int32, dwAlpha, dwCompositeMode uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawcv")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(sCanvas),
		uintptr(dstLeft),
		uintptr(dstTop),
		uintptr(dstRight),
		uintptr(dstBottom),
		uintptr(srcLeft),
		uintptr(srcTop),
		uintptr(dwAlpha),
		uintptr(dwCompositeMode),
	)
	return ret != 0, nil
}

// 画布画图片
func (d *DLL) DrawImage(hCanvas UICanvas, hImage UIImage, left, top float32, alpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawimage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImage),
		uintptr(left), //float32和float64都兼容
		uintptr(top),//float32和float64都兼容
		uintptr(alpha),
	)
	return ret != 0, nil
}

// 画布画图像矩形
func (d *DLL) DrawImageRect(hCanvas UICanvas, hImage UIImage, left, top, right, bottom float32, alpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawimagerect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImage),
		uintptr(left),//float32和float64都兼容
		uintptr(top),//float32和float64都兼容
		uintptr(right),//float32和float64都兼容
		uintptr(bottom),//float32和float64都兼容
		uintptr(alpha),
	)
	return ret != 0, nil
}

//画布画图像矩形和旋转角度
func (d *DLL) DrawImageRotate(hCanvas UICanvas, hImage UIImage, pMatrix UIMatrix, left, top, fAngle float32, alpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawimagerotate")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImage),
		uintptr(pMatrix),
		uintptr(left),  //float32和float64都兼容
		uintptr(top),   //float32和float64都兼容
		uintptr(fAngle),//float32和float64都兼容
		uintptr(alpha),
	)
	return ret != 0, nil
}

// 画布画图像缩放到矩形
func (d *DLL) DrawImageRectRect(hCanvas UICanvas, hImage UIImage, dstLeft, dstTop, dstRight, dstBottom, srcLeft, srcTop, srcRight, srcBottom float32, alpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawimagerectrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImage),
		uintptr(dstLeft),//float32和float64都兼容
		uintptr(dstTop),//float32和float64都兼容
		uintptr(dstRight),//float32和float64都兼容
		uintptr(dstBottom),//float32和float64都兼容
		uintptr(srcLeft),//float32和float64都兼容
		uintptr(srcTop),//float32和float64都兼容
		uintptr(srcRight),//float32和float64都兼容
		uintptr(srcBottom),//float32和float64都兼容
		uintptr(alpha),
	)
	return ret != 0, nil

}

// 画布画九宫矩形
func (d *DLL) DrawImageFromGrid(hCanvas UICanvas, hImage UIImage, dstLeft, dstTop, dstRight, dstBottom, srcLeft, srcTop, srcRight, srcBottom,
	gridPaddingLeft, gridPaddingTop, gridPaddingRight, gridPaddingBottom float32, dwFlags int32, dwAlpha uint32) (bool, error) {

	addr, err := d.GetProcAddress("jo_UIcanvas_drawimagefromgrid")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImage),
		uintptr(dstLeft),         //float32和float64都兼容
		uintptr(dstTop),          //float32和float64都兼容
		uintptr(dstRight),        //float32和float64都兼容
		uintptr(dstBottom),       //float32和float64都兼容
		uintptr(srcLeft),         //float32和float64都兼容
		uintptr(srcTop),          //float32和float64都兼容
		uintptr(srcRight),        //float32和float64都兼容
		uintptr(srcBottom),       //float32和float64都兼容
		uintptr(gridPaddingLeft), //float32和float64都兼容
		uintptr(gridPaddingTop),    //float32和float64都兼容
		uintptr(gridPaddingRight), //float32和float64都兼容
		uintptr(gridPaddingBottom),//float32和float64都兼容
		uintptr(dwFlags),
		uintptr(dwAlpha),
	)
	return ret != 0, nil
}

//
func (d *DLL) DrawShadow(hCanvas UICanvas, fLeft, fTop, fRight, fBottom, fShadowSize float32, crShadow ARGB,
	radiusTopLeft, radiusTopRight, radiusBottomLeft, radiusBottomRight, offsetX, offsetY float32) (bool, error) {

	addr, err := d.GetProcAddress("jo_UIcanvas_drawshadow")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(fLeft),             //float32和float64都兼容
		uintptr(fTop),              //float32和float64都兼容
		uintptr(fRight),            //float32和float64都兼容
		uintptr(fBottom),           //float32和float64都兼容
		uintptr(fShadowSize),       //float32和float64都兼容
		uintptr(crShadow),          //float32和float64都兼容
		uintptr(radiusTopLeft),   //float32和float64都兼容
		uintptr(radiusTopRight),  //float32和float64都兼容
		uintptr(radiusBottomLeft),  //float32和float64都兼容
		uintptr(radiusBottomRight), //float32和float64都兼容
		uintptr(offsetX),           //float32和float64都兼容
		uintptr(offsetY),         //float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画泛光特效
func (d *DLL) DrawBloom(hCanvas, sCanvas UICanvas, hImg UIImage, shadowColor ARGB, isBloom bool, pBlur float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawBloom")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(sCanvas),
		uintptr(hImg),
		uintptr(shadowColor),
		boolToUintptr(isBloom),
		uintptr(pBlur), //float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画图像圆角
func (d *DLL) DrawRoundedImage(hCanvas UICanvas, hImg UIImage, left, top, width, height, radiusX, radiusY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawroundedimage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hImg),
		uintptr(left),   //float32和float64都兼容
		uintptr(top),    //float32和float64都兼容
		uintptr(width),  //float32和float64都兼容
		uintptr(height), //float32和float64都兼容
		uintptr(radiusX),//float32和float64都兼容
		uintptr(radiusY),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画线
func (d *DLL) DrawLine(hCanvas UICanvas, hBrush UIBrush, x1, y1, x2, y2, strokeWidth float32, strokeStyle uint32, isRadius bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawline")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		uintptr(strokeStyle),
		boolToUintptr(isRadius),
	)
	return ret != 0, nil
}

// 画布画点
func (d *DLL) DrawPoint(hCanvas UICanvas, hBrush UIBrush, x, y, strokeWidth float32, isRadius bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawPoint")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x),//float32和float64都兼容
		uintptr(y),//float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		boolToUintptr(isRadius),
	)
	return ret != 0, nil
}

// 画布画多边形
func (d *DLL) DrawPoly(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom float32, numberOfEdges uint32, angle, strokeWidth float32, strokeStyle uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawPoly")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),      //float32和float64都兼容
		uintptr(top),       //float32和float64都兼容
		uintptr(right),     //float32和float64都兼容
		uintptr(bottom),    //float32和float64都兼容
		uintptr(numberOfEdges),
		uintptr(angle),//float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		uintptr(strokeStyle),
	)

	return ret != 0, nil
}

// 画布填充多边形
func (d *DLL) FillPoly(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom float32, numberOfEdges uint32, angle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillPoly")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),   //float32和float64都兼容
		uintptr(top),    //float32和float64都兼容
		uintptr(right),  //float32和float64都兼容
		uintptr(bottom), //float32和float64都兼容
		uintptr(numberOfEdges),
		uintptr(angle),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画路径
func (d *DLL) DrawPath(hCanvas UICanvas, hPath UIPath, hBrush UIBrush, strokeWidth float32, strokeStyle uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawpath")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hPath),
		uintptr(hBrush),
		uintptr(strokeWidth), //float32和float64都兼容
		uintptr(strokeStyle),
	)
	return ret != 0, nil
}

// 画布填充路径
func (d *DLL) FillPath(hCanvas UICanvas, hPath UIPath, hBrush UIBrush) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillpath")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hPath),
		uintptr(hBrush),
	)
	return ret != 0, nil
}

// 画布画矩形
func (d *DLL) DrawRect(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom, strokeWidth float32, strokeStyle uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),      //float32和float64都兼容
		uintptr(top),       //float32和float64都兼容
		uintptr(right),     //float32和float64都兼容
		uintptr(bottom),    //float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		uintptr(strokeStyle),
	)
	return ret != 0, nil
}

// 画布填充矩形
func (d *DLL) FillRect(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),  //float32和float64都兼容
		uintptr(top),   //float32和float64都兼容
		uintptr(right), //float32和float64都兼容
		uintptr(bottom),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画圆角矩形
func (d *DLL) DrawRoundedRect(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom, radiusX, radiusY, strokeWidth float32, strokeStyle uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawroundedrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),      //float32和float64都兼容
		uintptr(top),       //float32和float64都兼容
		uintptr(right),     //float32和float64都兼容
		uintptr(bottom),    //float32和float64都兼容
		uintptr(radiusX),   //float32和float64都兼容
		uintptr(radiusY),   //float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		uintptr(strokeStyle),
	)
	return ret != 0, nil
}

// 画布填充圆角矩形
func (d *DLL) FillRoundedRect(hCanvas UICanvas, hBrush UIBrush, left, top, right, bottom, radiusX, radiusY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillroundedrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(left),    //float32和float64都兼容
		uintptr(top),     //float32和float64都兼容
		uintptr(right),   //float32和float64都兼容
		uintptr(bottom),  //float32和float64都兼容
		uintptr(radiusX), //float32和float64都兼容
		uintptr(radiusY), //float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布画曲线组
func (d *DLL) DrawCurves(hCanvas UICanvas, hBrush UIBrush, points *H_POINT, count int32, tension, strokeWidth float32, fillMode, openMode bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawcurves")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(tension),    //float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		boolToUintptr(fillMode),
		boolToUintptr(openMode),
	)
	return ret != 0, nil
}

// 画布填充曲线组
func (d *DLL) FillCurves(hCanvas UICanvas, hBrush UIBrush, points *H_POINT, count int32, tension float32, openMode bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillcurves")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(tension), //float32和float64都兼容
		boolToUintptr(openMode),
	)
	return ret != 0, nil
}

// 画布画贝塞尔曲线组
func (d *DLL) DrawBeziers(hCanvas UICanvas, hBrush UIBrush, points *H_POINT, strokeWidth float32, fillMode, openMode bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawbezies")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(unsafe.Pointer(points)),
		uintptr(strokeWidth), //float32和float64都兼容
		boolToUintptr(fillMode),
		boolToUintptr(openMode),
	)
	return ret != 0, nil
}

// 画布画贝塞尔曲线
func (d *DLL) DrawBezier(hCanvas UICanvas, hBrush UIBrush, x1, y1, x2, y2, x3, y3, x4, y4, strokeWidth float32, fillMode, openMode bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawbezier")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
		uintptr(x3),//float32和float64都兼容
		uintptr(y3),//float32和float64都兼容
		uintptr(x4),//float32和float64都兼容
		uintptr(y4),//float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		boolToUintptr(fillMode),
		boolToUintptr(openMode),
	)
	return ret != 0, nil
}

// 画布画二次方贝塞尔曲线
func (d *DLL) DrawQuadraticBezier(hCanvas UICanvas, hBrush UIBrush, x1, y1, x2, y2, x3, y3, strokeWidth float32, fillMode, openMode bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawquadraticbezier")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x1),//float32和float64都兼容
		uintptr(y1),//float32和float64都兼容
		uintptr(x2),//float32和float64都兼容
		uintptr(y2),//float32和float64都兼容
		uintptr(x3),//float32和float64都兼容
		uintptr(y3),//float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		boolToUintptr(fillMode),
		boolToUintptr(openMode),
	)
	return ret != 0, nil
}

// 画布画文本
func (d *DLL) DrawText(hCanvas UICanvas, text string, hFont UIFont, left, top, right, bottom float32, crText, crShadow ARGB, dwDTFormat int32, lpWidth, lpHeight *float32) (bool, error) {
	lpwzText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIcanvas_drawtext")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(unsafe.Pointer(lpwzText)),
		uintptr(hFont),
		uintptr(left),   //float32和float64都兼容
		uintptr(top),    //float32和float64都兼容
		uintptr(right),  //float32和float64都兼容
		uintptr(bottom), //float32和float64都兼容
		uintptr(crText),
		uintptr(crShadow),
		uintptr(dwDTFormat),
		uintptr(unsafe.Pointer(lpWidth)),
		uintptr(unsafe.Pointer(lpHeight)),
	)
	return ret != 0, nil
}
//画布画文本Ex
func (d *DLL) DrawTextex(hCanvas UICanvas, lpwzText string, hFont UIFont, left, top, right, bottom float32, crText, crShadow UIBrush, dwDTFormat int32, lpWidth, lpHeight *float32)(bool,error){
    addr,err := d.GetProcAddress("jo_UIcanvas_drawtextex")
    if err != nil {
        return false,err
    }
    ret,_,_:= syscall.SyscallN(addr,
        uintptr(hCanvas),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpwzText))),
        uintptr(hFont),
        uintptr(left),//float32和float64都兼容
        uintptr(top),//float32和float64都兼容
        uintptr(right),//float32和float64都兼容
        uintptr(bottom),//float32和float64都兼容
        uintptr(crText),
        uintptr(crShadow),
        uintptr(dwDTFormat),
        uintptr(unsafe.Pointer(lpWidth)),
        uintptr(unsafe.Pointer(lpHeight)),
    )
    return ret != 0,nil
}

// 绘制字体文本、免字体创建直接指定
func (d *DLL) DrawTextFname(hCanvas UICanvas, text, fontFace string, crText UIBrush, left, top, right, bottom float32, dwFontSize int32, dwFontStyle uint32) (bool, error) {
	lpwzText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}
	lpwzFontFace, err := syscall.UTF16PtrFromString(fontFace)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIcanvas_drawtextfname")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(unsafe.Pointer(lpwzText)),
		uintptr(crText),
		uintptr(left),//float32和float64都兼容
		uintptr(top),//float32和float64都兼容
		uintptr(right),//float32和float64都兼容
		uintptr(bottom),//float32和float64都兼容
		uintptr(unsafe.Pointer(lpwzFontFace)),
		uintptr(dwFontSize),
		uintptr(dwFontStyle),
	)
	return ret != 0, nil
}

// 绘制高性能文本，适合大文本绘制
func (d *DLL) DrawTextPro(hCanvas UICanvas, text string, hFont UIFont, left, top, right, bottom float32, crText, crShadow ARGB, dwDTFormat int32, lpWidth, lpHeight *float32) (bool, error) {
	lpwzText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIcanvas_drawtextpro")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(unsafe.Pointer(lpwzText)),
		uintptr(hFont),
		uintptr(left),   //float32和float64都兼容
		uintptr(top),    //float32和float64都兼容
		uintptr(right),  //float32和float64都兼容
		uintptr(bottom), //float32和float64都兼容
		uintptr(crText),
		uintptr(crShadow),
		uintptr(dwDTFormat),
		uintptr(unsafe.Pointer(lpWidth)),
		uintptr(unsafe.Pointer(lpHeight)),
	)
	return ret != 0, nil
}

// 画布画椭圆
func (d *DLL) DrawEllipse(hCanvas UICanvas, hBrush UIBrush, x, y, radiusX, radiusY, strokeWidth float32, strokeStyle uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawellipse")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x),       //float32和float64都兼容
		uintptr(y),       //float32和float64都兼容
		uintptr(radiusX), //float32和float64都兼容
		uintptr(radiusY), //float32和float64都兼容
		uintptr(strokeWidth),//float32和float64都兼容
		uintptr(strokeStyle),
	)
	return ret != 0, nil
}

// 画布填充椭圆
func (d *DLL) FillEllipse(hCanvas UICanvas, hBrush UIBrush, x, y, radiusX, radiusY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillellipse")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(x),       //float32和float64都兼容
		uintptr(y),       //float32和float64都兼容
		uintptr(radiusX), //float32和float64都兼容
		uintptr(radiusY), //float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布填充饼
func (d *DLL) FillPieI(hCanvas UICanvas, color ARGB, x, y, width, height, startAngle, sweepAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillpiei")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(color),
		uintptr(x),        //float32和float64都兼容
		uintptr(y),        //float32和float64都兼容
		uintptr(width),    //float32和float64都兼容
		uintptr(height),   //float32和float64都兼容
		uintptr(startAngle),//float32和float64都兼容
		uintptr(sweepAngle),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布填充区域
func (d *DLL) FillRegion(hCanvas UICanvas, hRgn UIRgn, hBrush UIBrush) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_fillregion")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hRgn),
		uintptr(hBrush),
	)
	return ret != 0, nil
}

// 画布画圆弧
func (d *DLL) DrawArc(hCanvas UICanvas, hBrush UIBrush, size, x, y, width, height, startAngle, sweepAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawArc")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(hBrush),
		uintptr(size),      //float32和float64都兼容
		uintptr(x),         //float32和float64都兼容
		uintptr(y),         //float32和float64都兼容
		uintptr(width),     //float32和float64都兼容
		uintptr(height),    //float32和float64都兼容
		uintptr(startAngle),//float32和float64都兼容
		uintptr(sweepAngle),//float32和float64都兼容
	)
	return ret != 0, nil
}

//画布画圆弧I
func (d *DLL) DrawArcI(hCanvas UICanvas, hBrush UIBrush, size, x, y, width, height, startAngle, sweepAngle float32) (bool,error){
    addr, err := d.GetProcAddress("jo_UIcanvas_drawArcI")
    if err != nil {
        return false,err
    }
    ret,_,_ := syscall.SyscallN(addr,
        uintptr(hCanvas),
        uintptr(hBrush),
        uintptr(size), //float32和float64都兼容
        uintptr(x),//float32和float64都兼容
        uintptr(y),//float32和float64都兼容
        uintptr(width),//float32和float64都兼容
        uintptr(height),//float32和float64都兼容
        uintptr(startAngle),//float32和float64都兼容
        uintptr(sweepAngle),//float32和float64都兼容
        )
    return ret != 0,nil
}

// 画布刷新
func (d *DLL) FlushCanvas(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_flush")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 获取画布信息
func (d *DLL) GetCanvasContext(hCanvas UICanvas, nType int32) (unsafe.Pointer, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_getcontext")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(nType),
	)
	return unsafe.Pointer(ret), nil
}

// 画布取DC
func (d *DLL) GetCanvasDC(hCanvas UICanvas) (syscall.Handle, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_getdc")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return syscall.Handle(ret), nil
}

// 画布释放DC
func (d *DLL) ReleaseCanvasDC(hCanvas UICanvas) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_releasedc")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
	)
	return ret != 0, nil
}

// 画布取尺寸
func (d *DLL) GetCanvasSize(hCanvas UICanvas, width, height *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_getsize")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)),
	)
	return ret != 0, nil
}

// 重新设置尺寸
func (d *DLL) ResizeCanvas(hCanvas UICanvas, width, height int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_resize")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(width),
		uintptr(height),
	)
	return ret != 0, nil
}

// 画布旋转
func (d *DLL) RotateCanvas(hCanvas UICanvas, fAngle, scaleX, scaleY float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_rotate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(fAngle),//float32和float64都兼容
		uintptr(scaleX),//float32和float64都兼容
		uintptr(scaleY),//float32和float64都兼容
	)
	return ret != 0, nil
}
//画布3d旋转
func (d *DLL) RotateCanvas3D(hCanvas UICanvas, fAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_rotate_3D")
    if err != nil {
        return false,err
    }
    ret,_,_ := syscall.SyscallN(addr,
        uintptr(hCanvas),
        uintptr(fAngle),//float32和float64都兼容
        )
    return ret != 0, nil
}

// 色调旋转
func (d *DLL) RotateHue(hCanvas UICanvas, fAngle float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_rotate_hue")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(fAngle),//float32和float64都兼容
	)
	return ret != 0, nil
}

// 画布置矩阵
func (d *DLL) SetTransform(hCanvas UICanvas, pMatrix UIMatrix) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_settransform")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hCanvas),
		uintptr(pMatrix),
	)
	return ret != 0, nil
}

// 绘制主题数据
func (d *DLL) DrawControlTheme(hTheme UIZip, hCanvas UICanvas, dstLeft, dstTop, dstRight, dstBottom float32, atomClass, atomSrcRect HATOM, dwAlpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIcanvas_drawcontroltheme")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hTheme),
		uintptr(hCanvas),
		uintptr(dstLeft),   //float32和float64都兼容
		uintptr(dstTop),    //float32和float64都兼容
		uintptr(dstRight),  //float32和float64都兼容
		uintptr(dstBottom), //float32和float64都兼容
		uintptr(atomClass),
		uintptr(atomSrcRect),
		uintptr(dwAlpha),
	)
	return ret != 0, nil
}

//绘制主题数据Ex
func (d *DLL) DrawControlThemeEx(hTheme UIZip, hCanvas UICanvas, dstLeft, dstTop, dstRight, dstBottom float32,
	atomClass, atomSrcRect, atomBackgroundRepeat, atomBackgroundPositon, atomBackgroundGrid, atomBackgroundFlags HATOM, dwAlpha uint32) (bool, error) {

	addr, err := d.GetProcAddress("jo_UIcanvas_drawcontrolthemeex")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hTheme),
		uintptr(hCanvas),
		uintptr(dstLeft),            //float32和float64都兼容
		uintptr(dstTop),             //float32和float64都兼容
		uintptr(dstRight),           //float32和float64都兼容
		uintptr(dstBottom),          //float32和float64都兼容
		uintptr(atomClass),
		uintptr(atomSrcRect),
		uintptr(atomBackgroundRepeat),
		uintptr(atomBackgroundPositon),
		uintptr(atomBackgroundGrid),
		uintptr(atomBackgroundFlags),
		uintptr(dwAlpha),
	)
	return ret != 0, nil
}

