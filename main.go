package main

import (
	"fmt"
	"hJoui/joui"
	"unsafe"
)

/*
	func buttonClick(hObj joui.UIView, nType, msg int, wParam, lParam uintptr) uintptr {
		fmt.Printf("Button clicked: hObj=%d, nType=%d, msg=%d\n", hObj, nType, msg)
		// 在这里处理按钮点击事件
		return 0
	}
*/
func main() {
	dll, err := joui.LoadDLL("JellyOrangeUI.x64.dll") //  DLL 文件的实际路径
	if err != nil {
		fmt.Println("Error loading DLL:", err)
		return
	}
	defer dll.Free()
	fmt.Printf("成功加载DLL: %s\n", dll.GetPath())
	// 初始化引擎
	config := &joui.H_INITINFO{
		DwDebug: true,
	}
	dll.Init(config, -1)

	/* 	uzip, err := dll.LoadZipByPath("./resource/tmp.zip", "")
	    if err!= nil {
	        fmt.Printf("加载zip失败: %v\n", err)
	        return
	    }
		dll.EnumZipFileByPath(uzip, "./", func(path string, param uintptr) bool {
			// 在回调函数中处理每个文件
			fmt.Printf("发现文件: %s\n", path)
			return true  // 返回true继续枚举，返回false停止枚举
		}, 0)
	*/

	hWnd, err := dll.CreateWindow(0, "", "Joui_demo😝", 0, 0, 1080, 600, 0, 0)
	if err != nil {
		fmt.Println("Error creating window:", err)
		return
	}

	wndUI, err := dll.BindWindow(hWnd, 0, uint32(joui.WwsTitle|joui.WwsSizeable|joui.WwsMoveable|joui.WwsButtonClose|joui.WwsCenterwindow|joui.WwsModal|joui.WwsButtonMax|joui.WwsButtonMin|joui.WwsButtonSkin|joui.WwsButtonHelp), 0, nil)
	if err != nil {
		fmt.Println("Error binding window:", err)
		return
	}

	dll.SetWindowLong(wndUI, joui.WlCrbkg, uintptr(joui.Argb(255, 255, 255, 255)), 0, 0)
	dll.SetWindowLong(wndUI, joui.WlTitlecolor, uintptr(joui.Argb(0, 122, 204, 240)), 0, 0)
	dll.SetWindowLong(wndUI, joui.WlCrshadow, uintptr(joui.Argb(244, 96, 108, 255)), 0, 0)
	dll.SetWindowLong(wndUI, joui.WlRadius, 10, 0, 0) //设置窗口圆角

	hObj_imagebox, err := dll.CreateControlEx(int32(joui.EosExFocusable), "form-imagebox", "", int32(joui.EosImageRound), 200, 150, 100, 100, int32(wndUI), 0, -1, 50, nil)
	var hImg joui.UIImage
	ok, err := dll.CreateImageFromFile("./resource/logo64.png", &hImg) //注意这里的 &hImg
	if err != nil {
		fmt.Println("加载图片失败:", err)
		return
	}
	if ok {
		dll.SendMessage(hObj_imagebox, joui.PicMsgSetimage, 0, uintptr(hImg)) //hImg
	}
	dll.DestroyImage(hImg) // 释放图片资源

	dll.CreateControlEx(-1, "form-button", "系统按钮", int32(joui.EosEbsTextoffset), 10, 35, 100, 30, int32(wndUI), 100, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-static", "😄7月30日，中国国家版本馆中央总馆、西安、杭州、广州分馆同步举行开馆仪式。", int32(joui.EosStaticRoll), 10, 70, 300, 40, int32(wndUI), 101, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "扩展按钮", int32(joui.EosEbsTextoffset|joui.EosEbsEx), 130, 35, 100, 30, int32(wndUI), 102, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "扩展圆角", int32(joui.EosEbsTextoffset|joui.EosEbsEx), 240, 35, 100, 30, int32(wndUI), 103, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "发送（&S）", int32(joui.EosEbsTextoffset|joui.EosEbsEx), 350, 35, 100, 30, int32(wndUI), 104, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "🆗树形框test", -1, 10, 120, 100, 30, int32(wndUI), 105, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "😄组合框test", -1, 10, 160, 100, 30, int32(wndUI), 106, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "💙布局test", -1, 10, 200, 100, 30, int32(wndUI), 107, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "💌报表test", -1, 10, 240, 100, 30, int32(wndUI), 108, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "😿模态窗口test", -1, 10, 280, 100, 30, int32(wndUI), 109, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "🙈从资源包创建界面test", -1, 10, 320, 150, 30, int32(wndUI), 110, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "😒元素列表test", -1, 10, 360, 150, 30, int32(wndUI), 111, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "分页test", -1, 10, 400, 150, 30, int32(wndUI), 112, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "3DControl", -1, 10, 440, 150, 30, int32(wndUI), 113, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "颜色选择器", -1, 10, 480, 150, 30, int32(wndUI), 114, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)
	dll.CreateControlEx(-1, "form-button", "图片框", -1, 10, 520, 150, 30, int32(wndUI), 115, int32(joui.DtVcenter|joui.DtCenter|joui.DtSingleline), 0, nil)

	// 圆角按钮 (103)
	hObj_btnex2, _ := dll.GetControlFromID(int32(wndUI), 103)
	if hObj_btnex2 != 0 {
		buttonExprops2 := joui.H_OBJ_PROPS{}
		buttonExprops2.CrBkgNormal = joui.Argb(71, 173, 203, 255)
		buttonExprops2.CrBkgHover = joui.Argb(51, 154, 184, 255)
		buttonExprops2.CrBkgDownOrChecked = joui.Argb(51, 154, 184, 255)
		buttonExprops2.Radius = 15
		dll.SendMessage(hObj_btnex2, int32(joui.BtexMsgSetlong), 0, uintptr(unsafe.Pointer(&buttonExprops2)))
		dll.SetControlColor(hObj_btnex2, int32(joui.ColorTextNormal), joui.Argb(255, 255, 255, 255), false)
		dll.SetControlColor(hObj_btnex2, int32(joui.ColorTextDown), joui.Argb(255, 255, 255, 255), false)
	}

	// 图片按钮 (104)
	hObj_btnex3, _ := dll.GetControlFromID(int32(wndUI), 104)
	if hObj_btnex3 != 0 {
		/* 	dll.SetControlColor(hObj_btnex3, int32(joui.ColorTextNormal), joui.Argb(255, 255, 255, 255), false)
		dll.SetControlColor(hObj_btnex3, int32(joui.ColorTextDown), joui.Argb(255, 255, 255, 255), false)
		*/
		// 文本偏移
		dll.SetPadding(hObj_btnex3, -1, -1, 20, -1, false)
	}

	dll.ShowWindow(wndUI, int32(joui.SwShownormal))
	dll.Run()
}
