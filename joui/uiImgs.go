package joui

import (
	"syscall"
	"unsafe"
)

// 图片组添加图片从数据指针
func (d *DLL) AddImageFromMemory(hImageList UIImageList, pImg []byte, nIndex uintptr) (uintptr, error) {
	var pImgPtr unsafe.Pointer
	var dwBytes uintptr
	if len(pImg) > 0 {
		pImgPtr = unsafe.Pointer(&pImg[0])
		dwBytes = uintptr(len(pImg))
	}

	addr, err := d.GetProcAddress("jo_UIfromImg_add")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		uintptr(pImgPtr),
		dwBytes,
		nIndex,
	)
	return ret, nil
}

// 图片组添加图片从图片句柄
func (d *DLL) AddImageFromHandle(hImageList UIImageList, hImg UIImage, nIndex uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_addimage")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		uintptr(hImg),
		nIndex,
	)
	return ret, nil
}

// 获取图片组图片数量
func (d *DLL) GetImageCount(hImageList UIImageList) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_count")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
	)
	return int32(ret), nil
}

// 创建图片组
func (d *DLL) CreateImageList(width, height int32) (UIImageList, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_create")
	if err != nil {
		return nil, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(width),
		uintptr(height),
	)
	return UIImageList(ret), nil
}

// 图片组删除图片
func (d *DLL) DeleteImage(hImageList UIImageList, nIndex uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_del")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		uintptr(nIndex),
	)
	return ret != 0, nil
}

// 销毁图片组
func (d *DLL) DestroyImageList(hImageList UIImageList) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
	)
	return ret != 0, nil
}

// 清空图片组
func (d *DLL) EmptyImageList(hImageList UIImageList) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_empty")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
	)
	return ret != 0, nil
}

// 绘制图片从图片组
func (d *DLL) DrawImageFromList(hImageList UIImageList, nIndex uintptr, hCanvas UICanvas, nLeft, nTop, nRight, nBottom int32, nAlpha uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_draw")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		nIndex,
		uintptr(hCanvas),
		uintptr(nLeft),
		uintptr(nTop),
		uintptr(nRight),
		uintptr(nBottom),
		uintptr(nAlpha),
	)
	return ret != 0, nil
}

// 获取图片组图片句柄
func (d *DLL) GetImageFromList(hImageList UIImageList, nIndex uintptr) (UIImage, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_get")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		nIndex,
	)
	return UIImage(ret), nil
}

// 设置图片组自内存图片
func (d *DLL) SetImageFromMemory(hImageList UIImageList, nIndex uintptr, pImg []byte) (bool, error) {
	var pImgPtr unsafe.Pointer
	var dwBytes uintptr
	if len(pImg) > 0 {
		pImgPtr = unsafe.Pointer(&pImg[0])
		dwBytes = uintptr(len(pImg))
	}
	addr, err := d.GetProcAddress("jo_UIfromImg_set")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		nIndex,
		uintptr(pImgPtr),
		dwBytes,
	)
	return ret != 0, nil
}

// 设置图片组图片
func (d *DLL) SetImageListImage(hImageList UIImageList, nIndex uintptr, hImg UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_setimage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		nIndex,
		uintptr(hImg),
	)
	return ret != 0, nil
}

// 获取图片组图片尺寸
func (d *DLL) GetImageListSize(hImageList UIImageList, pWidth, pHeight *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIfromImg_size")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hImageList),
		uintptr(unsafe.Pointer(pWidth)),
		uintptr(unsafe.Pointer(pHeight)),
	)
	return ret != 0, nil
}