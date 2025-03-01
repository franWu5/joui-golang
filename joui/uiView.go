package joui

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 注册控件
func (d *DLL) RegisterControlClass(className string, dwStyle, dwStyleEx, dwTextFormat int32, hCursor string, pfnObjProc ClsPROC) (bool, error) {
	lpClsname, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return false, err
	}
	var lpszCursor *uint16
	if hCursor != "" {
		lpszCursor, err = syscall.UTF16PtrFromString(hCursor)
		if err != nil {
			return false, err
		}
	}

	addr, err := d.GetProcAddress("jo_UIview_RegClass")
	if err != nil {
		return false, err
	}
	var callback uintptr
	if pfnObjProc != nil {
		callback = syscall.NewCallback(pfnObjProc)
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpClsname)),
		uintptr(dwStyle),
		uintptr(dwStyleEx),
		uintptr(dwTextFormat),
		uintptr(unsafe.Pointer(lpszCursor)),
		callback, // 使用转换后的回调
	)
	return ret != 0, nil
}

// 创建控件
func (d *DLL) CreateControl(className, title string, dwStyle int32, x, y, width, height, hParent, nID int32) (UIView, error) {
	lpClsname, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	lpTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}

	addr, err := d.GetProcAddress("jo_UIview_Create")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpClsname)),
		uintptr(unsafe.Pointer(lpTitle)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(hParent),
		uintptr(nID),
	)
	if ret == 0 { // 假设返回 0 表示失败
		return 0, fmt.Errorf("jo_UIview_Create failed")
	}
	return UIView(ret), nil
}

func (d *DLL) CreateControlEx(dwStyleEx int32, className, title string, dwStyle int32, x, y, width, height, hParent, nID, dwTextFormat int32,
	lParam uintptr, lpfnMsgProc MsgPROC) (UIView, error) {

	lpClsname, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	lpTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}
	var callback uintptr
	if lpfnMsgProc != nil {
		callback = syscall.NewCallback(lpfnMsgProc)
	}

	addr, err := d.GetProcAddress("jo_UIview_CreateEx")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dwStyleEx),
		uintptr(unsafe.Pointer(lpClsname)),
		uintptr(unsafe.Pointer(lpTitle)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(hParent),
		uintptr(nID),
		uintptr(dwTextFormat),
		lParam,
		callback,
	)
	if ret == 0 { // 假设返回 0 表示失败
		return 0, fmt.Errorf("jo_UIview_CreateEx failed")
	}
	return UIView(ret), nil
}

// jo_UIview_Create_this
func (d *DLL) CreateControl_this(dwStyleEx int32, className, title string, dwStyle int32, x, y, width, height, hParent, nID, dwTextFormat int32,
	lParam uintptr, hts unsafe.Pointer, lpfnMsgProc MsgPROC) (UIView, error) {

	lpClsname, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	lpTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}
	var callback uintptr
	if lpfnMsgProc != nil {
		callback = syscall.NewCallback(lpfnMsgProc)
	}

	addr, err := d.GetProcAddress("jo_UIview_Create_this")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(dwStyleEx),
		uintptr(unsafe.Pointer(lpClsname)),
		uintptr(unsafe.Pointer(lpTitle)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(hParent),
		uintptr(nID),
		uintptr(dwTextFormat),
		lParam,
		uintptr(hts),
		callback,
	)
	if ret == 0 {
		return 0, fmt.Errorf("jo_UIview_Create_this failed")
	}
	return UIView(ret), nil
}

// 销毁控件
func (d *DLL) DestroyControl(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Destroy")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 设置布局句柄
func (d *DLL) SetLayout(handle int32, hLayout UILayout, fUpdate bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_LayoutSet")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(hLayout),
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 获取布局句柄
func (d *DLL) GetLayout(handle int32) (UILayout, error) {
	addr, err := d.GetProcAddress("jo_UIview_LayoutGet")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
	)
	if ret == 0 {
		return 0, syscall.GetLastError() // 或者自定义的错误
	}
	return UILayout(ret), nil
}

// 更新布局
func (d *DLL) UpdateLayoutControl(handle int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_LayoutUpdate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
	)
	return ret != 0, nil
}

// 清空对象布局信息
func (d *DLL) ClearLayout(handle int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_LayoutClear")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
	)
	return ret != 0, nil
}

// 设置控件焦点
func (d *DLL) SetFocusControl(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetFocus")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 获取焦点控件
func (d *DLL) GetFocusControl(handle int32) (UIView, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetFocus")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
	)
	if ret == 0 { // 根据实际情况，可能需要检查返回值
		return 0, syscall.GetLastError()
	}
	return UIView(ret), nil
}

// 销毁控件焦点
func (d *DLL) KillFocus(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_KillFocus")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 发送消息
func (d *DLL) SendMessage(parent UIView, uMsg int32, wParam, lParam uintptr) uintptr {
	addr, err := d.GetProcAddress("jo_UIview_SendMessage")
	if err != nil {
		return 0 // 或者返回一个表示错误的特定值
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(uMsg),
		wParam,
		lParam,
	)
	return ret
}

// 投递消息
func (d *DLL) PostMessage(parent UIView, uMsg int32, wParam, lParam uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_PostMessage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(uMsg),
		wParam,
		lParam,
	)
	return ret != 0, nil
}

// 设置控件是否可以重画
func (d *DLL) SetRedraw(parent UIView, fCanbeRedraw bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetRedraw")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fCanbeRedraw),
	)
	return ret != 0, nil
}

// 锁定控件矩形
func (d *DLL) LockRect(parent UIView, left, top, right, bottom, width, height int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_LockRect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		uintptr(width),
		uintptr(height),
	)
	return ret != 0, nil
}

// 获取矩形
func (d *DLL) GetRect(parent UIView, lpRect *RECT, nType int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetRect")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(nType),
	)
	return ret != 0, nil
}

// 设置重画区域
func (d *DLL) Redraw(parent UIView, hUpdate bool, lprcRedraw *RECT) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Redraw")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(hUpdate),
		uintptr(unsafe.Pointer(lprcRedraw)),
	)
	return ret != 0, nil
}

// 设置控件文本偏移内边距
func (d *DLL) SetPadding(parent UIView, left, top, right, bottom int32, fRedraw bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetPadding")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		boolToUintptr(fRedraw),
	)
	return ret != 0, nil
}

// 客户区坐标到窗口坐标
func (d *DLL) ClientToWindow(parent UIView, x, y *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ClientToWindow")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)),
	)
	return ret != 0, nil
}

// 客户区坐标到屏幕坐标
func (d *DLL) ClientToScreen(parent UIView, x, y *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ClientToScreen")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)),
	)
	return ret != 0, nil
}

// 设置控件禁止.
func (d *DLL) EnableControl(parent UIView, fEnable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Enable")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fEnable),
	)
	return ret != 0, nil
}

// 查询控件是否禁止状态
func (d *DLL) IsControlEnabled(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_IsEnable")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 查询控件是否可视
func (d *DLL) IsControlVisible(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_IsVisible")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 设置控件可视
func (d *DLL) ShowControl(parent UIView, fShow bool, fAniType int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Show")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fShow),
		uintptr(fAniType),
	)
	return ret != 0, nil
}

// 查询控件是否有效
func (d *DLL) IsControlValid(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_IsValidate")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 获取父控件
func (d *DLL) GetParentControl(parent UIView) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetParent")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return int32(ret), nil
}

// 获取窗口UI句柄
func (d *DLL) GetUIWnd(parent UIView, hUIWnd *UIWnd) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetUIWnd")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(hUIWnd)),
	)
	return ret != 0, nil
}

// 获取控件参数
func (d *DLL) GetControlLong(parent UIView, nIndex int32) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetLong")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nIndex),
	)
	return ret, nil
}

// 设置控件参数
func (d *DLL) SetControlLong(parent UIView, nIndex int32, dwNewLong1, dwNewLong2, dwNewLong3 uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetLong")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nIndex),
		dwNewLong1,
		dwNewLong2,
		dwNewLong3,
	)
	return ret, nil
}

// 保存为图像
func (d *DLL) SaveControlToImage(parent UIView, phImg *UIImage) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Savetoimg")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(phImg)), // 传递指针的指针
	)
	return ret != 0, nil
}

// 设置控件位置
func (d *DLL) SetControlPos(parent, hObjInsertAfter UIView, x, y, width, height, flags int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetPos")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(hObjInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)
	return ret != 0, nil
}

// 移动控件
func (d *DLL) MoveControl(parent UIView, x, y, width, height int32, bRepaint bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_Move")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		boolToUintptr(bRepaint),
	)
	return ret != 0, nil
}

// 分发消息
func (d *DLL) DispatchMessage(parent UIView, uMsg int32, wParam, lParam uintptr) uintptr {
	addr, err := d.GetProcAddress("jo_UIview_DispatchMessage")
	if err != nil {
		return 0 // 或者其他错误值
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(uMsg),
		wParam,
		lParam,
	)
	return ret
}

// 分发事件
func (d *DLL) DispatchNotify(parent UIView, nCode int32, wParam, lParam uintptr) uintptr {
	addr, err := d.GetProcAddress("jo_UIview_DispatchNotify")
	if err != nil {
		return 0 // 或者其他错误值
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nCode),
		wParam,
		lParam,
	)
	return ret
}

// 获取控件相关颜色
func (d *DLL) GetControlColor(parent UIView, nIndex int32) (ARGB, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetColor")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nIndex),
	)
	return ARGB(ret), nil
}

// 设置控件相关颜色
func (d *DLL) SetControlColor(parent UIView, nIndex int32, dwColor ARGB, fRedraw bool) (ARGB, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetColor")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nIndex),
		uintptr(dwColor),
		boolToUintptr(fRedraw),
	)
	return ARGB(ret), nil // 返回原来的颜色
}

// 获取字体
func (d *DLL) GetControlFont(parent UIView, lfFaceName *unsafe.Pointer, lfSize *int32) (UIFont, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetFont")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lfFaceName)), // 传递指针的指针
		uintptr(unsafe.Pointer(lfSize)),     // 传递指针的指针
	)
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return UIFont(ret), nil
}

// 设置字体
func (d *DLL) SetControlFont(parent UIView, hFont UIFont, fRedraw bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetFont")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(hFont),
		boolToUintptr(fRedraw),
	)
	return ret != 0, nil
}

// 从字体名称设置控件文本字体
func (d *DLL) SetFontFromFamily(parent UIView, fontfamily string, fontSize, fontStyle int32, fRedraw bool) (bool, error) {
	lpszFontfamily, err := syscall.UTF16PtrFromString(fontfamily)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIview_SetFontFromFamily")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpszFontfamily)),
		uintptr(fontSize),
		uintptr(fontStyle),
		boolToUintptr(fRedraw),
	)
	return ret != 0, nil
}

// 设置控件文本.
func (d *DLL) SetControlText(parent UIView, text string, fRedraw bool) (bool, error) {
	lpString, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIview_SetText")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpString)),
		boolToUintptr(fRedraw),
	)
	return ret != 0, nil
}

// 获取控件文本
func (d *DLL) GetControlText(parent UIView, buffer []uint16) (uintptr, error) {

	addr, err := d.GetProcAddress("jo_UIview_GetText")
	if err != nil {
		return 0, err
	}
	nMaxCount := len(buffer)
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(nMaxCount),
	)
	return ret, nil
}

// 获取控件文本长度
func (d *DLL) GetControlTextLength(parent UIView) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetTextLength")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret, nil
}

// 开始绘制
func (d *DLL) BeginPaint(parent UIView, lpPS *PsContext) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_BeginPaint")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpPS)),
	)
	return ret != 0, nil
}

// 结束绘制
func (d *DLL) EndPaint(parent UIView, lpPS *PsContext) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_EndPaint")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpPS)),
	)
	return ret != 0, nil
}

// 获取特定关系的控件
func (d *DLL) GetNode(parent UIView, nCmd int32) (UIView, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetNode")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nCmd),
	)
	if ret == 0 {
		return 0, syscall.GetLastError() // 或者自定义的错误
	}
	return UIView(ret), nil
}

// 获取控件句柄自名称
func (d *DLL) GetControlFromName(handle int32, name string) (UIView, error) {
	lpName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	addr, err := d.GetProcAddress("jo_UIview_GetFromName")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(unsafe.Pointer(lpName)),
	)
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return UIView(ret), nil
}

// 获取控件句柄自ID
func (d *DLL) GetControlFromID(handle, nID int32) (UIView, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetFromID")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(nID),
	)
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return UIView(ret), nil
}

// 获取控件句柄自节点ID
func (d *DLL) GetControlFromNodeID(handle, nNodeID int32) (UIView, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetFromNodeID")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(handle),
		uintptr(nNodeID),
	)
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return UIView(ret), nil
}

// 查找控件
func (d *DLL) FindControl(hObjParent int32, hObjChildAfter UIView, className, title string) (UIView, error) {
	lpClassName, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	var lpTitle *uint16 = nil
	if title != "" {
		lpwzTitle, err := syscall.UTF16PtrFromString(title)
		if err != nil {
			return 0, err
		}
		lpTitle = lpwzTitle
	}

	addr, err := d.GetProcAddress("jo_UIview_Find")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hObjParent),
		uintptr(hObjChildAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpTitle)),
	)
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return UIView(ret), nil
}

type EnumChildCallback func(parent UIView, lParam uintptr) uintptr

// 枚举子控件
func (d *DLL) EnumChildControls(hObjParent int32, lpEnumFunc EnumChildCallback, lParam uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_EnumChild")
	if err != nil {
		return false, err
	}
	var callback uintptr
	if lpEnumFunc != nil {
		callback = syscall.NewCallback(lpEnumFunc)
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hObjParent),
		callback, // 使用转换后的回调
		lParam,
	)
	return ret != 0, nil
}

// 设置背景信息
func (d *DLL) SetControlBackgroundImage(parent UIView, image []byte, x, y int32, dwRepeat uint32, lpGrid *RECT, dwFlags int32, dwAlpha uint32, fUpdate bool) (bool, error) {
	var imagePtr unsafe.Pointer
	var imageLen uintptr

	if len(image) > 0 {
		imagePtr = unsafe.Pointer(&image[0])
		imageLen = uintptr(len(image))
	}

	addr, err := d.GetProcAddress("jo_UIview_SetBackgImage")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(imagePtr),
		imageLen,
		uintptr(x),
		uintptr(y),
		uintptr(dwRepeat),
		uintptr(unsafe.Pointer(lpGrid)),
		uintptr(dwFlags),
		uintptr(dwAlpha),
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 获取背景信息 (需要修改 C++ 代码)
func (d *DLL) GetControlBackgroundImage(parent UIView, lpBackgroundImage unsafe.Pointer) (bool, error) {

	addr, err := d.GetProcAddress("jo_UIview_GetBackgImage")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(lpBackgroundImage),
	)

	return ret != 0, nil
}

// 设置背景图片播放状态.
func (d *DLL) SetControlBackgroundPlayState(parent UIView, fPlayFrames, fResetFrame, fUpdate bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetBackgPlayState")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fPlayFrames),
		boolToUintptr(fResetFrame),
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 设置计时器
func (d *DLL) SetTimer(parent UIView, uTimeout uint32, fThread bool, lParam uintptr) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetTimer")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(uTimeout),
		boolToUintptr(fThread),
		lParam,
	)
	return ret != 0, nil
}

// 销毁计时器
func (d *DLL) KillTimer(parent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_KillTimer")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return ret != 0, nil
}

// 设置圆角度
func (d *DLL) SetRadius(parent UIView, topLeft, topRight, bottomRight, bottomLeft float32, fUpdate bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetRadius")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(topLeft),     //float32和float64都兼容
		uintptr(topRight),    //float32和float64都兼容
		uintptr(bottomRight), //float32和float64都兼容
		uintptr(bottomLeft),  //float32和float64都兼容
		boolToUintptr(fUpdate),
	)
	return ret != 0, nil
}

// 设置阴影
func (d *DLL) SetShadow(parent UIView, crShadow ARGB, fShadowSize, radiusTopLeft, radiusTopRight, radiusBottomLeft, radiusBottomRight float32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetShadow")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(crShadow),
		uintptr(fShadowSize),       //float32和float64都兼容
		uintptr(radiusTopLeft),     //float32和float64都兼容
		uintptr(radiusTopRight),    //float32和float64都兼容
		uintptr(radiusBottomLeft),  //float32和float64都兼容
		uintptr(radiusBottomRight), //float32和float64都兼容
	)
	return ret != 0, nil
}

// 设置模糊
func (d *DLL) SetBlurControl(parent UIView, fDeviation float32, bRedraw bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetBlur")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(fDeviation), //float32和float64都兼容
		boolToUintptr(bRedraw),
	)
	return ret != 0, nil
}

// 设置文本格式.返回原文本格式
func (d *DLL) SetTextFormat(parent UIView, dwTextFormat uint32, bRedraw bool) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_TextFormat")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwTextFormat),
		boolToUintptr(bRedraw),
	)
	return int32(ret), nil
}

// 设置提示文本
func (d *DLL) SetTooltipsText(parent UIView, text string, crBlack, crText int32) (bool, error) {
	lpString, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIview_TooltipsSetText")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(crBlack),
		uintptr(crText),
	)
	return ret != 0, nil
}

// 弹出或关闭提示文本
func (d *DLL) PopTooltips(parent UIView, text string, crBlack, crText int32) (bool, error) {
	lpText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIview_TooltipsPop")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(crBlack),
		uintptr(crText),
	)
	return ret != 0, nil
}

// jo_UIview_TooltipsPopEx
func (d *DLL) PopTooltipsEx(parent UIView, title, text string, crBlack, crText, x, y, dwTime, nIcon int32, fShow bool) (bool, error) {
	lpTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return false, err
	}
	lpText, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return false, err
	}
	addr, err := d.GetProcAddress("jo_UIview_TooltipsPopEx")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpTitle)),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(crBlack),
		uintptr(crText),
		uintptr(x),
		uintptr(y),
		uintptr(dwTime),
		uintptr(nIcon),
		boolToUintptr(fShow),
	)
	return ret != 0, nil
}

// 创建缓动
func (d *DLL) CreateEasing(parent UIView, beginX, endX, beginY, endY int32, effect uint32, frame int32, wait bool,
	numIterations int32, alternateDirection bool, param1, param2, param3, param4 uintptr) (bool, error) {

	addr, err := d.GetProcAddress("jo_UIview_Easing")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(beginX),
		uintptr(endX),
		uintptr(beginY),
		uintptr(endY),
		uintptr(effect),
		uintptr(frame),
		boolToUintptr(wait),
		uintptr(numIterations),
		boolToUintptr(alternateDirection),
		param1,
		param2,
		param3,
		param4,
	)
	return ret != 0, nil
}

// 初始化属性列表
func (d *DLL) InitPropList(parent UIView, nPropCount int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_InitPropList")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nPropCount),
	)
	return ret != 0, nil
}

// 获取属性
func (d *DLL) GetProp(parent UIView, dwKey uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetProp")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwKey),
	)
	return ret, nil
}

// 设置属性
func (d *DLL) SetProp(parent UIView, dwKey, dwValue uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetProp")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwKey),
		dwValue,
	)
	return ret, nil
}

// 移除属性
func (d *DLL) RemoveProp(parent UIView, dwKey uintptr) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_RemoveProp")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwKey),
	)
	return ret, nil
}

// 枚举属性表
func (d *DLL) EnumProps(parent UIView, lpfnCbk EnumPropsPROC, lParam uintptr) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_EnumProps")
	if err != nil {
		return 0, err
	}
	var callback uintptr
	if lpfnCbk != nil {
		callback = syscall.NewCallback(lpfnCbk)
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		callback, // 使用转换后的回调
		lParam,
	)
	return int32(ret), nil
}

// 设置控件状态
func (d *DLL) SetUIState(parent UIView, dwState uint32, fRemove bool, lprcRedraw *RECT, fRedraw bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetUIState")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwState),
		boolToUintptr(fRemove),
		uintptr(unsafe.Pointer(lprcRedraw)),
		boolToUintptr(fRedraw),
	)
	return ret != 0, nil
}

// 获取控件状态
func (d *DLL) GetUIState(parent UIView) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetUIState")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
	)
	return int32(ret), nil
}

// 控件默认过程
func (d *DLL) DefProc(hWnd syscall.Handle, parent UIView, uMsg int32, wParam, lParam uintptr) uintptr {
	addr, err := d.GetProcAddress("jo_UIview_DefProc")
	if err != nil {
		return 0 // 或者其他错误值
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hWnd),
		uintptr(parent),
		uintptr(uMsg),
		wParam,
		lParam,
	)
	return ret
}

// 获取滚动条信息
func (d *DLL) GetScrollInfo(parent UIView, nBar int32, lpnMin, lpnMax, lpnPos, lpnTrackPos *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollGetInfo")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
		uintptr(unsafe.Pointer(lpnMin)),
		uintptr(unsafe.Pointer(lpnMax)),
		uintptr(unsafe.Pointer(lpnPos)),
		uintptr(unsafe.Pointer(lpnTrackPos)),
	)
	return ret != 0, nil
}

// 获取滚动条位置
func (d *DLL) GetScrollPos(parent UIView, nBar int32) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollGetPos")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
	)
	return int32(ret), nil
}

// 设置滚动条位置
func (d *DLL) SetScrollPos(parent UIView, nBar, nPos int32, bRedraw bool) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollSetPos")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
		uintptr(nPos),
		boolToUintptr(bRedraw),
	)
	return int32(ret), nil
}

// 设置滚动条信息
func (d *DLL) SetScrollInfo(parent UIView, nBar, mask, nMin, nMax, nPage, nPos int32, bRedraw bool) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollSetInfo")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
		uintptr(mask),
		uintptr(nMin),
		uintptr(nMax),
		uintptr(nPage),
		uintptr(nPos),
		boolToUintptr(bRedraw),
	)
	return int32(ret), nil
}

// 设置滚动条范围
func (d *DLL) SetScrollRange(parent UIView, nBar, nMin, nMax int32, bRedraw bool) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollSetRange")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
		uintptr(nMin),
		uintptr(nMax),
		boolToUintptr(bRedraw),
	)
	return int32(ret), nil
}

// 获取滚动条句柄
func (d *DLL) GetScrollBarControl(parent UIView, nBar int32) (UIView, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollGetControl")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
	)

	if ret == 0 {
		return 0, syscall.GetLastError() // 或者自定义错误
	}
	return UIView(ret), nil
}

// 获取滚动条拖动位置
func (d *DLL) GetScrollTrackPos(parent UIView, nBar int32) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollGetTrackPos")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
	)
	return int32(ret), nil
}

// 获取滚动条范围
func (d *DLL) GetScrollRange(parent UIView, nBar int32, lpnMinPos, lpnMaxPos *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollGetRange")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nBar),
		uintptr(unsafe.Pointer(lpnMinPos)),
		uintptr(unsafe.Pointer(lpnMaxPos)),
	)
	return ret != 0, nil
}

// 显示/隐藏滚动条
func (d *DLL) ShowScrollBar(parent UIView, wBar int32, fShow bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollShow")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(wBar),
		boolToUintptr(fShow),
	)
	return ret != 0, nil
}

// 禁用/启用滚动条
func (d *DLL) EnableScrollBar(parent UIView, wSB, wArrows int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollEnable")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(wSB),
		uintptr(wArrows),
	)
	return ret != 0, nil
}

// 设置颜色
func (d *DLL) SetScrollBarColor(parent UIView, colorNormal, colorHover, colorDown ARGB, bRedraw bool) {
	addr, err := d.GetProcAddress("jo_UIview_ScrollSetColor")
	if err != nil {
		return //或者panic
	}
	_, _, _ = syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(colorNormal),
		uintptr(colorHover),
		uintptr(colorDown),
		boolToUintptr(bRedraw),
	)
}

// 滚动条消息回调
func (d *DLL) ScrollBar(parent UIView, uMsg int32, wParam, lParam uintptr, nLine int32) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_Scrollbar")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(uMsg),
		wParam,
		lParam,
		uintptr(nLine),
	)
	return int32(ret), nil
}

// 坐标转换
func (d *DLL) PointTransform(hObjSrc, hObjDst UIView, ptX, ptY *int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_PointTransform")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(hObjSrc),
		uintptr(hObjDst),
		uintptr(unsafe.Pointer(ptX)),
		uintptr(unsafe.Pointer(ptY)),
	)
	return ret != 0, nil
}

// 设置是否启用事件冒泡
func (d *DLL) EnableEventBubble(parent UIView, fEnable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_EnableEventBubble")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fEnable),
	)
	return ret != 0, nil
}

// 获取控件类信息
func (d *DLL) GetClassInfo(parent UIView, lpClassInfo *RegMethod) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetClassInfo")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(lpClassInfo)),
	)
	return ret != 0, nil
}

// 通过类名获取类信息
func (d *DLL) GetClassInfoEx(className string, lpClassInfo *RegMethod) (bool, error) {
	lpClsname, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return false, err
	}

	addr, err := d.GetProcAddress("jo_UIview_GetClassInfoEx")
	if err != nil {
		return false, err
	}

	ret, _, _ := syscall.SyscallN(addr,
		uintptr(unsafe.Pointer(lpClsname)),
		uintptr(unsafe.Pointer(lpClassInfo)),
	)
	return ret != 0, nil
}

// 挂接控件事件回调
func (d *DLL) HandleEvent(parent UIView, nEvent int32, pfnCallback EventHandlerPROC) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_HandleEvent")
	if err != nil {
		return false, err
	}
	var callback uintptr
	if pfnCallback != nil {
		callback = syscall.NewCallback(pfnCallback)
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(nEvent),
		callback, // 使用转换后的回调
	)
	return ret != 0, nil
}

// 是否允许启用输入法
func (d *DLL) EnableIME(parent UIView, fEnable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_EnableIME")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fEnable),
	)
	return ret != 0, nil
}

// 设置窗口输入法状态
func (d *DLL) SetIMEState(parent UIView, fOpen bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetIMEState")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fOpen),
	)
	return ret != 0, nil
}

// 层次
func (d *DLL) SetControlZOrder(parent UIView, hObjInsertAfter int32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_LastEx")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(hObjInsertAfter),
	)
	return ret != 0, nil
}

// 置父
func (d *DLL) SetControlParent(parent, hParent UIView) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_SetParent")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(hParent),
	)
	return ret != 0, nil
}

// 设置控件是否禁止转换空格和回车为单击事件
func (d *DLL) DisableSpaceAndEnterToClick(parent UIView, fDisable bool) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_DisableToClick")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		boolToUintptr(fDisable),
	)
	return ret != 0, nil
}

// 查询拖曳信息格式
func (d *DLL) CheckDropFormat(pDataObject unsafe.Pointer, dwFormat uint32) (bool, error) {
	addr, err := d.GetProcAddress("jo_UIview_CheckDropFormat")
	if err != nil {
		return false, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pDataObject),
		uintptr(dwFormat),
	)
	return ret != 0, nil
}

// 查询拖曳文本内容
func (d *DLL) GetDropString(pDataObject unsafe.Pointer, buffer []uint16) (int32, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetDropString")
	if err != nil {
		return 0, err
	}
	cchMaxLength := len(buffer)
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(pDataObject),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(cchMaxLength),
	)
	return int32(ret), nil
}

// 查询拖曳数量
func (d *DLL) GetDropFileNumber(wParam uintptr) (uint32, error) {
	addr, err := d.GetProcAddress("jo_UIview_GetDropFileNumber")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		wParam,
	)
	return uint32(ret), nil
}

// 查询拖曳文本内容
func (d *DLL) GetDragQueryFile(wParam uintptr, iFile uint32, lpszFile []uint16) (uint32, error) {
	cch := len(lpszFile)
	addr, err := d.GetProcAddress("jo_UIview_GetDragQueryFile")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		wParam,
		uintptr(iFile),
		uintptr(unsafe.Pointer(&lpszFile[0])),
		uintptr(cch),
	)
	return uint32(ret), nil
}

// 设置编辑框选中行字符格式
func (d *DLL) EditSetSelCharFormat(parent UIView, dwMask uint32, crText ARGB, fontFace string, fontSize uint32, yOffset int32,
	bBold, bItalic, bUnderLine, bStrikeOut, bLink bool) (uintptr, error) {

	lpwzFontFace, err := syscall.UTF16PtrFromString(fontFace)
	if err != nil {
		return 0, err
	}

	addr, err := d.GetProcAddress("jo_UIview_editSetSelCharFormat")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwMask),
		uintptr(crText),
		uintptr(unsafe.Pointer(lpwzFontFace)),
		uintptr(fontSize),
		uintptr(yOffset),
		boolToUintptr(bBold),
		boolToUintptr(bItalic),
		boolToUintptr(bUnderLine),
		boolToUintptr(bStrikeOut),
		boolToUintptr(bLink),
	)
	return ret, nil
}

// 设置编辑框选中行段落格式
func (d *DLL) EditSetSelParFormat(parent UIView, dwMask uint32, wNumbering, wAlignment uint16, dxStartIndent, dxRightIndent, dxOffset int32) (uintptr, error) {
	addr, err := d.GetProcAddress("jo_UIview_editSetSelParFormat")
	if err != nil {
		return 0, err
	}
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(dwMask),
		uintptr(wNumbering),
		uintptr(dxStartIndent),
		uintptr(dxRightIndent),
		uintptr(dxOffset),
		uintptr(wAlignment),
	)
	return ret, nil
}

type stdWstring struct { //非导出
	ptr      *uint16 // 指向字符串数据的指针
	length   uintptr // 字符串长度（不包括 null 终止符）
	capacity uintptr // 分配的容量
}

// 获取编辑框内容
func (d *DLL) EditGetText(parent UIView) (string, error) {
	addr, err := d.GetProcAddress("jo_UIview_editgetText")
	if err != nil {
		return "", err
	}
	var str stdWstring
	ret, _, _ := syscall.SyscallN(addr,
		uintptr(parent),
		uintptr(unsafe.Pointer(&str)),
	)
	if ret == 0 {
		return "", fmt.Errorf("get text fail")
	}
	//转换为go string
	return syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(str.ptr))[:str.length:str.length]), nil
}
