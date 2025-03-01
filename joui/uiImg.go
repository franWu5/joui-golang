package joui

import (
	"syscall"
	"unsafe"
)

// 加载位图对象自内存
func (d *DLL) LoadBitmapFromMemory(data []byte, width, height uint32, retBitmap *syscall.Handle) (bool, error) {
	var dataPtr unsafe.Pointer
	var dataLen uintptr
	if len(data) > 0 {
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = uintptr(len(data))
	}
	addr, err := d.GetProcAddress("jo_UIloadBitMapFromMemory")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dataPtr),
		dataLen,
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(retBitmap)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 加载图像对象自内存
func (d *DLL) LoadImageFromMemory(data []byte, uType, nIndex int32) (unsafe.Pointer, error) {

	var dataPtr unsafe.Pointer
	var dataLen uintptr
	if len(data) > 0 {
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = uintptr(len(data))
	}

	addr, err := d.GetProcAddress("jo_UIloadImageFromMemory")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dataPtr),
		dataLen,
		uintptr(uType),
		uintptr(nIndex),
	)
	return unsafe.Pointer(ret), nil
}

// 创建图像
func (d *DLL) CreateImage(width, height int32, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_create")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自文件
func (d *DLL) CreateImageFromFile(filename string, phImg *UIImage) (bool, error) {
	lpwzFilename, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIimg_createfromfile")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpwzFilename)),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自位图句柄
func (d *DLL) CreateImageFromHBitmap(hBitmap, hPalette unsafe.Pointer, fPreAlpha bool, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_createfromhbitmap")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hBitmap),
		uintptr(hPalette),
		boolToUintptr(fPreAlpha),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自图标句柄
func (d *DLL) CreateImageFromHIcon(hIcon syscall.Handle, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_createfromhicon")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hIcon),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自内存、缓冲区
func (d *DLL) CreateImageFromMemory(data []byte, phImg *UIImage) (bool, error) {
	var dataPtr unsafe.Pointer
	var dataLen uintptr
	if len(data) > 0 {
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = uintptr(len(data))
	}

	addr, err := d.GetProcAddress("jo_UIimg_createfrommemory")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dataPtr),
		dataLen,
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}
func (d *DLL) CreateImageFromPngBits(nWidth, nHeight int32, pbBuffer *byte, dstImg *UIImage)(bool,error){
    addr,err := d.GetProcAddress("jo_UIimg_createfrompngbits")
    if err != nil {
        return false,err
    }
    ret,_,_ := syscall.SyscallN(addr,
        uintptr(nWidth),
        uintptr(nHeight),
        uintptr(unsafe.Pointer(pbBuffer)),
        uintptr(unsafe.Pointer(dstImg)), // 传递指针的指针
    )
    return ret != 0, nil
}

// 图像创建自内存缩放
func (d *DLL) CreateThumbnailImage(data []byte, dstWidth, dstHeight int32, phImg *UIImage) (bool, error) {
	var dataPtr unsafe.Pointer
	var dataLen uintptr
	if len(data) > 0 {
		dataPtr = unsafe.Pointer(&data[0])
		dataLen = uintptr(len(data))
	}

	addr, err := d.GetProcAddress("jo_UIimg_createThumbnail")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dataPtr),
		dataLen,
		uintptr(dstWidth),
		uintptr(dstHeight),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自资源包 zFile为0取默认主题
func (d *DLL) CreateImageFromZip(hRes UIZip, fileName string, dstWidth, dstHeight int32, phImg *UIImage) (bool, error) {
	lpwzFileName, err := syscall.UTF16PtrFromString(fileName)
	if err != nil {
		return false, err
	}
	addr,err := d.GetProcAddress("jo_UIimg_createZip")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hRes),
		uintptr(unsafe.Pointer(lpwzFileName)),
		uintptr(dstWidth),
		uintptr(dstHeight),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像创建自SVG
func (d *DLL) CreateImageFromSvg(input string, color ARGB, dstWidth, dstHeight int32, phImg *UIImage) (bool, error) {

    inputPtr := unsafe.StringData(input)
	addr,err := d.GetProcAddress("jo_UIimg_createSvg")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(inputPtr)),
		uintptr(color),
		uintptr(dstWidth),
		uintptr(dstHeight),
		uintptr(unsafe.Pointer(phImg)), //传递指针
	)

	return ret != 0, nil
}

// 根据代码创建SVG
func (d *DLL) CreateImageFromSvgCode(svgCode string, color ARGB, dstWidth, dstHeight int32, phImg *UIImage) (bool, error) {
	lpwzSvgCode, err := syscall.UTF16PtrFromString(svgCode)
	if err != nil {
		return false, err
	}
	addr,err := d.GetProcAddress("jo_UIimg_createCodeSvg")
    if err != nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpwzSvgCode)),
		uintptr(color),
		uintptr(dstWidth),
		uintptr(dstHeight),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 创建自字节流
func (d *DLL) CreateImageFromStream(lpStream unsafe.Pointer, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_createfromstream")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(lpStream),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 销毁图像
func (d *DLL) DestroyImage(hImg UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
	)
	return ret != 0, nil
}

// 复制图像
func (d *DLL) CopyImage(hImg UIImage, width, height int32, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_copy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 复制部分图像
func (d *DLL) CopyImageRect(hImg UIImage, x, y, width, height int32, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_copyrect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 取图像帧数
func (d *DLL) GetFrameCount(hImage UIImage, nFrameCount *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_getframecount")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImage),
		uintptr(unsafe.Pointer(nFrameCount)),
	)
	return ret != 0, nil
}

// 取图像帧延时
func (d *DLL) GetFrameDelay(hImg UIImage, lpDelayAry *int32, nFrames int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_getframedelay")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(lpDelayAry)),
		uintptr(nFrames),
	)
	return ret != 0, nil
}

// 获取点像素
func (d *DLL) GetPixel(hImg UIImage, x, y int32, retPixel *ARGB) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_getpixel")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(retPixel)),
	)
	return ret != 0, nil
}

// 图像设置点像素
func (d *DLL) SetPixel(hImg UIImage, x, y int32, color ARGB) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_setpixel")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(x),
		uintptr(y),
		uintptr(color),
	)
	return ret != 0, nil
}

// 获取图像尺寸
func (d *DLL) GetImageSize(hImg UIImage, lpWidth, lpHeight *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_getsize")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(lpWidth)),
		uintptr(unsafe.Pointer(lpHeight)),
	)
	return ret != 0, nil
}

// 取图像高度
func (d *DLL) GetImageHeight(hImg UIImage) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIimg_height")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
	)
	return int32(ret), nil
}

// 锁定图像
func (d *DLL) LockImage(hImg UIImage, lpRectL *RECT, flags uint32, pixelFormat int32, lpLockedBitmapData *unsafe.Pointer) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_lock")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(lpRectL)),
		uintptr(flags),
		uintptr(pixelFormat),
		uintptr(unsafe.Pointer(lpLockedBitmapData)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 解锁图像
func (d *DLL) UnlockImage(hImg UIImage, lpLockedBitmapData unsafe.Pointer) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_unlock")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(lpLockedBitmapData),
	)
	return ret != 0, nil
}

// 翻转图像
func (d *DLL) RotateFlipImage(hImg UIImage, rfType int32, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_rotateflip")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(rfType),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像保存到内存
func (d *DLL) SaveImageToMemory(hImg UIImage, lpBuffer *unsafe.Pointer, len *uintptr, format int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_savetomemory")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(lpBuffer)), // 传递指针的指针
		uintptr(unsafe.Pointer(len)),      // 传递指针的指针
		uintptr(format),
	)
	return ret != 0, nil
}

// 图像保存到文件
func (d *DLL) SaveImageToFile(hImg UIImage, wzFileName string, format int32) (bool, error) {
	lpwzFileName, err := syscall.UTF16PtrFromString(wzFileName)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIimg_savetofile")
    if err!=nil {
        return false,err
    }
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(lpwzFileName)),
		uintptr(format),
	)
	return ret != 0, nil
}

// 图像取目标
func (d *DLL) GetImageContext(hImg UIImage, pBitmapSource *unsafe.Pointer) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_getcontext")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(unsafe.Pointer(pBitmapSource)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 缩放图像
func (d *DLL) ScaleImage(hImg UIImage, width, height int32, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_scale")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(width),
		uintptr(height),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 图像设置当前活动帧
func (d *DLL) SelectActiveFrame(hImg UIImage, nIndex int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIimg_selectactiveframe")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
		uintptr(nIndex),
	)
	return ret != 0, nil
}

// 取图像宽度
func (d *DLL) GetImageWidth(hImg UIImage) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIimg_width")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImg),
	)
	return int32(ret), nil
}