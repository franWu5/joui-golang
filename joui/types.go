package joui

import (
	"syscall"
	"unsafe"
)

// --- C/C++ 类型别名 ---
type (
	UnzFile  = uintptr
	ZPOS64_T = uint64
	ARGB     = uint32 // COLORREF 通常是 DWORD，即 uint32
	HATOM    = int32  //或者int 根据实际使用情况
	UIWnd    = int32
	UILayout = int32
	UIView   = int32
	UICanvas  = int32
	UIBrush  = unsafe.Pointer
	UIImageList = unsafe.Pointer
	UIImage     = int32 //或者int
	UIFont      = int32//或者int
	UIPath      = int32//或者int
	UIRgn       = unsafe.Pointer
	UIEasing    = unsafe.Pointer
	UIArray     = unsafe.Pointer
	UIMatrix    = unsafe.Pointer
    UIZip       = uintptr // 使用 uintptr，因为 UIZip 在其他文件中已定义为 uintptr

	CHANNEL = byte


	//回调函数
    WinMsgPROC =  func(hWnd syscall.Handle, UIWnd UIWnd, msg, nType int, wParam, lParam uintptr) uintptr //LRESULT(CALLBACK*)(HWND, UIWnd, INT, INT, WPARAM, LPARAM);
    MsgPROC    =   func(hWnd syscall.Handle, hObj UIView, msg, nType int, wParam, lParam uintptr) uintptr//LRESULT(CALLBACK*)(HWND, UIView, INT, INT, WPARAM, LPARAM);
    ClsPROC    =    func(hWnd syscall.Handle, hObj UIView, msg, nType int, wParam, lParam uintptr) uintptr//LRESULT(CALLBACK*)(HWND, UIView, INT, WPARAM, LPARAM);
    EventHandlerPROC =  func(hObj UIView, nType, msg int, wParam, lParam uintptr) uintptr //LRESULT(CALLBACK*)(UIView, INT, INT, WPARAM, LPARAM);
    EnumPropsPROC    = func(hObj UIView, a, b, c uintptr) uintptr //LRESULT(CALLBACK*)(UIView, size_t, size_t, size_t);
    EnumFileCallback = func(path string, lParam uintptr) bool//BOOL(CALLBACK*)(LPCTSTR, LPARAM);
    ArrayComparePROC =  func(hArray UIArray, a bool, b, c uintptr)//void(CALLBACK*)(UIArray, BOOL, size_t, size_t);
)


// --- 枚举 ---
type JouiWndStyleFlags uint32
const (
	WwsButtonClose       JouiWndStyleFlags = 1        // 关闭按钮
	WwsButtonMax                      = 1 << 1   // 最大化按钮
	WwsButtonMin                      = 1 << 2   // 最小化按钮
	WwsButtonMenu                     = 1 << 3   // 菜单按钮
	WwsButtonSkin                     = 1 << 4   // 皮肤按钮
	WwsButtonSetting                  = 1 << 5   // 设置按钮
	WwsButtonHelp                     = 1 << 6   // 帮助按钮
	WwsHasicon                       = 1 << 7   // 图标
	WwsTitle                         = 1 << 8   // 标题
	WwsFullscreen                    = 1 << 9   // 全屏模式
	WwsSizeable                      = 1 << 10  // 允许调整尺寸
	WwsMoveable                      = 1 << 11  // 允许随意移动
	WwsNoshadow                      = 1 << 12  // 不显示窗口阴影
	WwsNoinheritbkg                  = 1 << 13  // 不继承父窗口背景
	WwsNotabborder                   = 1 << 14  // 不显示TAB焦点边框
	WwsEscexit                       = 1 << 15  // ESC关闭窗口
	WwsMainwindow                    = 1 << 16  // 主窗口
	WwsCenterwindow                  = 1 << 17  // 窗口居中
	WwsNocaptiontopmost              = 1 << 18  // 标题栏取消置顶
	WwsPopupwindow                   = 1 << 19  // 弹出式窗口
	WwsNotitlebar                    = 1 << 22  // 取消标题栏
	WwsModal                         = 1 << 23  // 模态窗口
)

type JouiWndStyleExFlags uint32
const (
    StyleexNoshadow      JouiWndStyleExFlags = 0x80000000 // 不显示阴影 信息框/菜单有效
	StyleexNoinheritbkg                  = 0x20000000 // 不继承父窗口背景数据 信息框有效
	StyleexCentewindow                   = 0x40000000 // 居父窗口中间 信息框有效
	StyleexWindowicon                    = 0x80000000 // 显示窗口图标 信息框有效
)

type JouiObjSaniFlags int
const (
	FshowObjaniNone    JouiObjSaniFlags = iota // 无
	FshowObjaniStleft                          // 左移
	FshowObjaniStright                         // 右移
	FshowObjaniSttop                           // 上移
	FshowObjaniStdown                          // 下移
	FshowObjaniStfadein                        // 渐隐渐显
)

type JouiWndLongFlags int
const (
	WlHtheme      JouiWndLongFlags = 1  // 主题包句柄
	WlBlur                         = 2  // 背景模糊
	WlTitle                        = 3  // 标题 set为文本指针；get为标题控件句柄
	WlMsgproc                      = 4  // 窗口消息过程
	WlAlpha                        = 5  // 窗口透明度
	WlHwnd                         = 6  // 窗口句柄
	WlLparam                       = 7  // 自定义参数
	WlTitlecolor                   = 8  // 标题颜色
	WlStyle                        = 9  // 界面风格
	WlHwndstyle                    = 10 // 窗口风格
	WlHwndstyleex                  = 11 // 窗口扩展风格
	WlCrbkg                        = 12 // 背景颜色
	WlCrborder                     = 13 // 边框颜色
	WlMinheight                    = 14 // 最小高度
	WlMinwidth                     = 15 // 最小宽度
	WlBtncolor                     = 16 // 控制按钮颜色 参数1，参数2，参数3分别为(默认、点燃、按下)纯色主题或无主题下有效
	WlBtnkgcolor                   = 17 // 控制按钮背景颜色 参数1，参数2，参数3分别为(默认、点燃、按下)纯色主题或无主题下有效
	WlHuerotation                  = 18 // 色调旋转
	WlObjfocus                     = 19 // 焦点组件句柄
	WlObjtitlebkg                  = 20 // 标题栏背景颜色、参数1，参数2支持两种颜色渐变，参数3为时钟周期(前提参数1、参数2颜色不能为空)
	WlObjcaption                   = 21 // 标题栏组件句柄
	WlSysfps                       = 22 // 参数1为真时启用/获取FPS 需要先启用才能获取
	WlRadius                       = 23 // 设置圆角度
	WlBordersize                   = 24 // 边框&阴影大小
	WlCrshadow                     = 25 // 阴影颜色&与边框颜色二选一
)
type JouiObjLongFlags int

const (
	OlNodeid       JouiObjLongFlags = 1  // 节点ID
	OlBlur                          = 2  // 模糊系数
	OlObjproc                       = 3  // 事件回调
	OlAlpha                         = 4  // 透明度
	OlThis                          = 5
	OlLparam                        = 6  // 附加参数
	OlObjparent                     = 7  // 父句柄
	OlThisclass                     = 8
	OlTextformat                    = 9  // 文本格式
	OlID                            = 10 // 组件ID
	OlStyle                         = 11 // 基本风格
	OlCursor                        = 12 // 光标类型 IDC_ARROW
	OlHfont                         = 13 // 字体句柄
	OlExstyle                       = 14 // 扩展风格
	OlUserdata                      = 15 // 用户数据
	OlHcanvas                       = 16 // 画布句柄
	OlOwner                         = 17 // 控件数据
	OlState                         = 18 // 控件状态
	OlHuerotation                   = 19 // 色调旋转 等同于jo_UIcanvas_rotate_hue
	OlRotate                        = 20 // 控件旋转 参数1为旋转角度0-360(-1自动计算) 参数2为类型(0、普通旋转 1、3D旋转)
	OlLpwztitle                     = 21 // 标题内容,如果为编辑框控件可用于加入文本（参数1,参数2,参数3分别为：内容,颜色,是否末尾）
	OlHtheme                        = 22 // 主题句柄
	OlClsnameex                     = 23 // 控件遍历类名 为0不遍历
	OlClsname                       = 24 // 控件类名
	OlHwnd                          = 25 // 窗口句柄
	OlNomovecaption                 = 26 // 禁止更新标题栏按钮位置
	OlShadowobj                     = 27 // 阴影句柄
)

type JouiFontStyleFlags uint32
const (
	EfsDefault    JouiFontStyleFlags = 0      // 普通
	EfsBold                          = 1      // 加粗
	EfsItalic                        = 2      // 倾斜
	EfsUnderline                     = 1 << 2 // 下划线
	EfsStrickout                     = 1 << 3 // 删除线
)

type JouiGetRectFlags int
const (
	GrtDefault  JouiGetRectFlags = iota // 相对位置矩形
	GrtClient                           // 客户区矩形
	GrtWindow                           // 窗口矩形
	GrtDirty                            // 脏区域矩形
	GrtText                             // 文本矩形
	GrtTextoff                          // 文本偏移矩形
)
type JouiMsgwmExFlags int
const (
	WmExLclick    JouiMsgwmExFlags = -3  // 左键单击组件
	WmExRclick                     = -4  // 右键单击组件
	WmExMclick                     = -5  // 中键单击组件
	WmExInitpopup                  = -6  // 弹出式窗口初始化完毕
	WmExExitpopup                  = -7  // 弹出式窗口即将销毁 wParam=0:即将销毁 wParam=1:是否可销毁, 返回1则取消销毁
	WmExEasing                     = -8  // 缓动	 发给控件用这个,窗口是EMT_EASING转WMM_EASING
	WmExDrop                       = -9  // 控件接收到拖放 lParam为H_DROPINFO结构体,若处理后应当返回 DROPEFFECT_开头的常量
	WmExProps                      = -11 // 属性消息 lParam为H_OBJ_PROPS结构体
	WmExLdclick                    = -12 // 左键双击
	WmExTrayicon                   = -22 // 托盘消息 lParam 为传递msg
	WmExCommand                    = -23 // 菜单消息 wParam 选中菜单ID lParam菜单句柄 该消息响应在弹出时 指定父句柄为控件消息 如果父句柄为窗口 用WM_COMMAND
)

type JouiObjStateFlags uint32
const (
    StateNormal           JouiObjStateFlags = 0       // 正常
	StateDisable                        = 1       // 禁止
	StateSelect                         = 1 << 1  // 选择
	StateFocus                          = 1 << 2  // 焦点
	StateDown                           = 1 << 3  // 按下
	StateChecked                        = 1 << 4  // 选中
	StateHalfselect                     = 1 << 5  // 半选中
	StateReadonly                       = 1 << 6  // 只读
	StateHover                          = 1 << 7  // 悬浮
	StateDefault                        = 1 << 8  // 默认
	StateSubitemVisiable                = 1 << 9  // 子项目_可视
	StateSubitemHidden                  = 1 << 10 // 子项目_隐藏
	StateBusy                           = 1 << 11 // 繁忙中
	StateRolling                        = 1 << 12 // 滚动中
	StateAnimating                      = 1 << 13 // 动画中
	StateHidden                         = 1 << 14 // 隐藏
	StateAllowsize                      = 1 << 15 // 允许修改尺寸
	StateAllowdrag                      = 1 << 16 // 允许拖动
	StateAllowfocus                     = 1 << 17 // 允许焦点
	StateAllowselect                    = 1 << 18 // 允许选择
	StateHyperlinkHover                = 1 << 19 // 超链接_悬浮
	StateHyperlinkVisited              = 1 << 20 // 超链接_已访问
	StateAllowmultiple                  = 1 << 21 // 允许多选
	StatePassword                       = 1 << 22 // 密码模式
)

type JouiObjStyleFlags uint32
const (
	EosScrollAmds        JouiObjStyleFlags = 1 << 24 // 自动隐藏滚动条
	EosScrollDisableno                   = 1 << 25 // 隐藏滚动条
	EosSizebox                          = 1 << 26 // 可调整尺寸
	EosDisabled                         = 1 << 27 // 禁止
	EosHidden                           = 1 << 28 // 隐藏
	EosScrollV                          = 1 << 30 // 垂直滚动条
	EosScrollH                          = 1 << 31 // 水平滚动条
)

type JouiObjStyleExFlags uint32
const (
	EosExAutosize    JouiObjStyleExFlags = 1 << 22 // 自适应尺寸
	EosExTransparent                   = 1 << 23 // 鼠标穿透
	EosExDragdrop                      = 1 << 25 // 允许拖拽
	EosExAcceptfiles                   = 1 << 26 // 接收文件拖放
	EosExFocusable                     = 1 << 27 // 允许焦点
	EosExTabstop                       = 1 << 28 // 允许TAB焦点
	EosExTopmost                       = 1 << 29 // 总在最前
	EosExComposited                    = 1 << 30 // 背景混合 控件透明度无效
	EosExCustomdraw                    = 1 << 31 // 自定义绘制
)

type JouiObjCheckFlags uint32
const (
	EosEbsCheckbutton JouiObjCheckFlags = 1 // 复选按钮
	EosEbsRadiobutton                  = 2 // 单选按钮
)

type JouiObjbuttonFlags uint32
const (
    EosEbsTextoffset  JouiObjbuttonFlags = 4 // 文本偏移
	EosEbsEx                           = 8 // 扩展
	// 消息
	BtexMsgSetlong = 4000 // 置参数 lParam为H_OBJ_PROPS结构
	BtexMsgSeticon = 4001 // 置图标 lParam为图标句柄 wParam 图标位置 [忽略/0：左; 1：右; 2:上]48
	BtexMsgSetimage = 4002 // 置背景 lParam为H_IMAGEINFO结构 wParam可设置九宫矩形
	BtexMsgSetrect = 4003 // 置九宫矩形 lParam为RECT
)

type JouiObjsbFlags uint32
const (
	EssHorizontalscroll JouiObjsbFlags = 0       // 水平滚动条
	EssVerticalscroll                   = 1       // 垂直滚动条
	EssLefttopalign                     = 1 << 1  // 左顶对齐
	EssRightbottomalign                 = 1 << 2  // 右底对齐
	EssControlbutton                    = 1 << 3  // 控制按钮
)

type JouiObjLdingFlags uint32
const (
	EosEldsIcuar JouiObjLdingFlags = 1         // 转圈
	EosEldsIline                   = 1 << 1    // 直线
	EosEldsWin11                   = 1 << 2    // win11加载动画
	EosEldsFillcolor6              = 1 << 4    // 6色彩圈
	// 消息
	LoadingMsgColor = 4100 // 置颜色
	LoadingMsgSize  = 4101 // 置圆大小
)

type JouiObjPageFlags uint32
const (
	TabMsgAddpage          JouiObjPageFlags = 4000 // 绑定现有页面
	TabMsgInserpage                         = 4001 // 添加页面
	TabMsgGetcount                          = 4002 // 获取数量
	TabMsgGethandle                         = 4003 // 获取子夹句柄
	TabMsgDelepage                          = 4004 // 删除子夹
	TabMsgDeleteall                         = 4005 // 清空子夹
	TabMsgSetpagename                       = 4006 // 置子夹标题
	TabMsgGetpagename                       = 4007 // 取子夹标题
	TabMsgSetselectionmark                  = 4008 // 现行子夹
	TabMsgSetcolor                          = 4009 // 置颜色
	TabMsgGetheaderhandle                   = 4010 // 取表头句柄
	TabMsgGetspinhandle                     = 4011 // 取调节器句柄

	EosTabHorizontal = 0x200  // 选择夹表头风格_横向
	EosTabVertical   = 0x400  // 选择夹表头风格_纵向
	EosTabSimple     = 0x800  // 选择夹表头风格_简洁模式、需要组合横向或纵向
	EosTabSpin       = 0x1000 // 选择夹表头风格_调节器
	// 事件
	WmmTabSelchanger = 1700 // 现行选中 lParam=选中ID,wParam=type
	// 调节器
	EosSpinNoedit = 0x800 // 调节器风格_无编辑框
	// 事件
	WmmSpinClick = 1700 // 按钮被单击 lParam=1左边按钮 2右边按钮
)

type JouiObjEditFlags uint32
const (
	EosEditDisabledrag   JouiObjEditFlags = 1        // 允许拖拽
	EosEditUsepassword                   = 1 << 1   // 密码输入
	EosEditHideselection                 = 1 << 2   // 显示选择文本
	EosEditRichtext                      = 1 << 3   // 富文本
	EosEditAllowbeep                     = 1 << 4   // 允许鸣叫
	EosEditReadonly                      = 1 << 5   // 只读
	EosEditNewline                       = 1 << 6   // 回车换行
	EosEditNumericinput                  = 1 << 7   // 数值输入
	EosEditAutowordsel                   = 1 << 8   // 自动选择字符
	EosEditDisablemenu                   = 1 << 9   // 禁用右键默认菜单
	EosEditParseurl                      = 1 << 10  // 解析URL
	EosEditAllowtab                      = 1 << 11  // 允许TAB字符
	EosEditShowtipsalways                = 1 << 12  // 总是显示提示文本
	EosEditHiddencaret                   = 1 << 13  // 隐藏插入符
	EosEditDisablectrl                   = 1 << 14  // 禁用ctrl快捷键 包括复制、粘贴、全选等

	// 编辑框消息事件
	EmMsgUndor       = 199  // 撤销
	EmMsgSetcuebanner = 5377 // 设置提示文本(wParam:提示文本颜色,lParam:文本指针)
	EmMsgLoadrtf      = 6001 // 加载RTF文件(lParam:LPCTSTR路径)
	EmMsgSavertf      = 6002 // 保存为RTF文件(lParam:LPCTSTR保存路径)
	EmMsgLoadimage    = 6003 // 加载图片(wParam:图像句柄 lParam左边缩进)
	EmMsgInputtype    = 5379 // 输入方式 -1获取 (lParam:EOS_EES_)
	EmMsgUsepassword  = 5380 // 恢复*遮盖默认字符
	EmMsgExSetlong   = -11  // 扩展编辑框有效 置参数、wParam为H_OBJ_PROPS结构

	// 编辑框选中行段落格式
	PfmAlnment = 0x00000008 // 段落对齐方式
)
type JouiObjStaticFlags uint32
const (
    EosStaticDline     JouiObjStaticFlags = 16   // 分割线
	EosStaticRoll                        = 64   // 滚动
	EosStaticMirrorshadow                = 128  // 镜像阴影
	EosStaticBlurtext                    = 256  // 模糊文字
	EosStaticPro                         = 512  // 高性能文本
	EosStaticFlow                        = 1024 // 流动渐变
	EosStaticEx                          = 2048 // 高级标签

	// 消息
	StaticMsgRollIngtime      = 1771 // 置滚动时间
	StaticMsgRollType         = 1772 // 滚动类型 roll_type_
	StaticMsgRollDelay        = 1773 // 滚动等待延迟 默认300
	StaticMsgCvwidth          = 1774 // 设置画布宽度 挂接WMM_LABEL_DRAWITEM事件可实现自定义绘制
	StaticMsgSetblur          = 1777 // 设置模糊值 eos_static_blurtext风格有效
	StaticMsgBordersize       = 1778 // 设置流动边框大小 eos_static_flow风格有效
	StaticMsgFlowstcr         = 1779 // 设置流动边框颜色 eos_static_flow风格有效 lParam为H_FLOWST
	StaticMsgAddstring        = 1780 // 添加内容 lParam为H_STATICEX_INFO eos_static_ex风格有效

	LabelMoUserdata = 80078
	// 滚动类型
	RollTypeLeft  = 0 // 左边
	RollTypeAbout = 1 // 左右滚动
	// 事件
	WmmLabelDrawitem = 20010 // 自定义绘制
)

type JouiObjListFlags uint32

const (
	EosElvsVerticallist  JouiObjListFlags = 0      // 纵向列表
	EosElvsHorizontallist                 = 1      // 横向列表
	EosElvsAllowmultiple                  = 0x08   // 允许多选
	EosElvsAllowskn                       = 0x800  // 关闭主题绘制
	EosElvsItemtracking                   = 0x10   // 表项跟踪
	EosElvsShowcheck                      = 0x40   // 取消默认检查框 此时多选以ctrl按住触发 用于普通列表框
	EosElvsButton                         = 0x400  // 表项以按钮形式呈现

	// 列表命中
	LvhtNowhere = 1  // 未命中
	LvhtOnitem  = 14 // 命中表项

	// 列表事件
	WmmLvnItemchanged       = -101 // 事件_列表_现行选中项被改变
	WmmLvnItemselectd       = -102 // 事件_列表_表项选中状态
	WmmLvnItemselectc       = -103 // 事件_列表_表项选中状态取消
	WmmLvnItemdragdropBegin = -106 // 事件_列表_表项正在被拖拽 wParam为选中项目，lParam为当前坐标
	WmmLvnItemdragdropEnd   = -107 // 事件_列表_表项结束拖拽
	WmmLvnHottrack          = -121 // 事件_列表_表项热点跟踪
	// 列表消息
	IlvmMsgSetitemsize = 11001 // 设置表项尺寸
)
type ElementViewFlags int
const (
    //元素列表消息
	TlvmMsgItemCreate       ElementViewFlags = 10010 // 创建 返回值将作为列表项控件 wParam:nIndex
	TlvmMsgItemCreated                       = 10011 // 创建完毕 wParam:nIndex
	TlvmMsgItemDestroy                       = 10012 // 销毁
	TlvmMsgItemFill                          = 10013 // 填充数据 wParam:nIndex,lParam:hObjItem
	TlvmMsgItemHovercolor                    = 100022 // 设置表项热点背景色 lParam:ARGB颜色
	TlvmMsgItemSelectcolor                   = 100023 // 设置表项选中背景色 lParam:ARGB颜色
	TlvmMsgSettemplate                       = 10020 // 置元素数据 wParam:cbSize,lParam:pTemplate
	TlvmMsgGetitemobj                        = 10021 // 取项目句柄 wParam:表项索引,返回表项容器句柄(不在可视区返回0)
	TlvmMsgGetcuiindex                       = 10022 // 取鼠标位置表项索引 wParam:X  lParam:Y
)
type JouiCanvasFlags uint32
const (
	Ecvf           JouiCanvasFlags = 0
	EcvfCliped                     = 0x80000000 // 重置剪辑区

	// 画布信息类型
	CsinfoHdc          = 1
	CsinfoBitmap       = 2
	CsinfoScanvas      = 3
	CsinfoD2Dcontext = 4

	// 混合模式
	CsCompositeModeSrcover = 0 // 覆盖
	CsCompositeModeSrccopy = 1 // 拷贝
)

type JouiColourFlags int
const (
	ColorBackground     JouiColourFlags = iota // 背景颜色
	ColorBorder                                // 边框颜色 自定义绘制情况下无效
	ColorBorderHover                           // 边框颜色_点燃
	ColorTextNormal                            // 文本颜色_正常
	ColorTextHover                             // 文本颜色_点燃
	ColorTextDown                              // 文本颜色_按下
	ColorTextFocus                             // 文本颜色_焦点
	ColorTextChecked                           // 文本颜色_选中
	ColorTextSelect                            // 文本颜色_选择
	ColorTextBan                               // 文本颜色_禁止
	ColorTextShadow                            // 文本颜色_阴影
	ColorEditCaret      = 30                   // 编辑框_光标原色
	ColorEditBanner     = 31                   // 编辑框_提示文本颜色
)

type JouiBackgFlags uint32
const (
	BifDefault              JouiBackgFlags = 0        // 默认
	BifPlayimage                           = 1        // 播放动画
	BifDisablescale                        = 1 << 1   // 禁用缩放
	BifGridExclusionCenter                 = 1 << 2   // 九宫矩形_排除中间区域
	BifPositionYPercent                    = 1 << 3   // Y使用百分比单位
	BifPositionXPercent                    = 1 << 4   // X使用百分比单位
	BifCvscale                             = 1 << 5   // 按画布尺寸缩放
	// 背景重复模式
	BirDefault    = 0 // 默认(适应缩放)
	BirEpault     = 1 // 等比缩放
	BirNoRepeat  = 2 // 平铺不重复
	BirRepeat     = 3 // 水平垂直重复平铺
	BirRepeatX    = 4 // 水平重复平铺
	BirRepeatY    = 5 // 垂直重复平铺
	BirRepeatCv   = 6 // 按画布尺寸缩放
)

type JouiEventFlags int

const (
	WmmCreate           JouiEventFlags = -99  // 创建
	WmmDestroy                         = -98  // 销毁
	WmmCalcsize                        = -97  // 计算尺寸
	WmmMove                            = -96  // 控件移动
	WmmSize                            = -95  // 尺寸被改变
	WmmEnable                          = -94  // 禁止状态被改变
	WmmShow                            = -93  // 可视状态被改变
	WmmLup                             = -92  // 左键被放开
	WmmLeave                           = -91  // 离开组件
	WmmTimer                           = -90  // 时钟
	WmmCheck                           = -89  // 选中
	WmmTrayicon                        = -88  // 托盘图标
	WmmIntdlg                          = -87  // 对话框初始化完毕
	WmmEasing                          = -86  // 缓动
	WmmRup                             = -85  // 右键被放开
	WmmClick                           = -2   // 左键被单击
	WmmDblclk                          = -3   // 左键被双击
	WmmRclick                          = -5   // 右键被单击
	WmmRdblclk                         = -6   // 右键被双击
	WmmSetfocus                        = -7   // 设置焦点
	WmmKillfocus                       = -8   // 失去焦点
	WmmWrapperraw                      = -11  // 空项目绘制提示 lParam为ps_context结构
	WmmCustomdraw                      = -12  // 自定义绘制 wParam为当前项目、lParam为ps_context结构
	WmmHover                           = -13  // 进入组件
	WmmNchittest                       = -14  // 点击测试
	WmmKeydown                         = -15  // 按下某键
	WmmReleasedcapture                 = -16  // 取消鼠标捕获
	WmmChar                            = -18  // 字符输入
	WmmTooltipscreated                 = -19  // 提示框即将弹出
	WmmLdown                           = -20  // 左键被按下
	WmmRdown                           = -21  // 右键被按下
	WmmFontchanged                     = -23  // 字体被改变
	WmmExCustomdraw                    = -24  // 扩展自定义绘制 lParam为ps_customdraw结构
)

type JouiLayoutFlags int
const (
    EltNull    JouiLayoutFlags = iota // 无
	EltLinear                       // 线性
	EltFlow                         // 流式
	EltPage                         // 页面
	EltTable                        // 表格
	EltRelative                     // 相对
	EltAbsolute                     // 绝对

	// 通用布局属性
	ElpPaddingLeft   = -1 // 通用_内间距_左
	ElpPaddingTop    = -2 // 通用_内间距_顶
	ElpPaddingRight  = -3 // 通用_内间距_右
	ElpPaddingBottom = -4 // 通用_内间距_底
	ElcpMarginLeft   = -1 // 通用_外间距_左
	ElcpMarginTop    = -2 // 通用_外间距_顶
	ElcpMarginRight  = -3 // 通用_外间距_右
	ElcpMarginBottom = -4 // 通用_外间距_底

	// 线性布局属性
	ElpLinearDirection          = 1 // 排布方向
	ElcpLinearSize              = 1 // 尺寸 [-1或未填写为组件当前尺寸]
	ElcpLinearAlign             = 2 // 另外一个方向对齐方式
	ElpLinearDalign             = 2 // 布局方向对齐方式
	ElpLinearDalignLeftTop      = 0 // 左上
	ElpLinearDalignCenter       = 1 // 居中
	ElpLinearDalignRightBottom = 2 // 右下

	// 线性布局另一个方向对齐方式
	ElcpLinearAlginFill     = 0 // 填满
	ElcpLinearAlignLeftTop  = 1 // 左上
	ElcpLinearAlignCenter    = 2 // 居中
	ElcpLinearAlignRightTop = 3 // 右上
	ElcpLinearAlignRightBottom = 4 //右下

    //布局排布方向
	ElpDirectionH = 0 // 水平
	ElpDirectionV = 1 // 垂直

	// 流式布局属性
	ElpFlowDirection    = 1 // 排布方向
	ElcpFlowSize        = 1 // 尺寸 [-1或未填写为组件当前尺寸]
	ElcpFlowNewLine     = 2 // 组件强制换行

	// 页面布局属性
	ElpPageCurrent  = 1 // 当前显示页面索引
	ElcpPageFill    = 1 // 是否填充整个布局

	// 表格布局属性
	ElpTableArrayRow   = 1 // 行高数组
	ElpTableArrayCell  = 2 // 列宽数组
	ElcpTableRow       = 1 // 所在行
	ElcpTableCell      = 2 // 所在列
	ElcpTableRowSpan   = 3 // 跨行数
	ElcpTableCellSpan  = 4 // 跨列数
	ElcpTableFill      = 5 // 是否填满

	// 相对布局属性件
	ElcpRelativeLeftOf         = 1  // 左侧于组件
	ElcpRelativeTopOf          = 2  // 之上于组件
	ElcpRelativeRightOf        = 3  // 右侧于组件
	ElcpRelativeBottomOf       = 4  // 之下于组件
	ElcpRelativeLeftAlignOf    = 5  // 左对齐于组件
	ElcpRelativeTopAlignOf     = 6  // 顶对齐于组件
	ElcpRelativeRightAlignOf   = 7  // 右对齐于组件
	ElcpRelativeBottomAlignOf  = 8  // 底对齐于组件
	ElcpRelativeCenterParentH  = 9  // 水平居中于父
	ElcpRelativeCenterParentV  = 10 // 垂直居中于父

	// 绝对布局属性
	ElcpAbsoluteLeft        = 1  // 左侧
	ElcpAbsoluteLeftType    = 2  // 位置类型_左侧
	ElcpAbsoluteTop         = 3  // 顶部
	ElcpAbsoluteTopType     = 4  // 位置类型_顶部
	ElcpAbsoluteRight       = 5  // 右侧
	ElcpAbsoluteRightType   = 6  // 位置类型_右侧
	ElcpAbsoluteBottom      = 7  // 底部
	ElcpAbsoluteBottomType  = 8  // 位置类型_底部
	ElcpAbsoluteWidth       = 9  // 宽度（优先级低于右侧）
	ElcpAbsoluteWidthType   = 10 // 位置类型_宽度
	ElcpAbsoluteHeight      = 11 // 高度（优先级低于底部）
	ElcpAbsoluteHeightType  = 12 // 位置类型_高度
	ElcpAbsoluteOffsetH     = 13 // 水平偏移量
	ElcpAbsoluteOffsetHType = 14 // 位置类型_水平偏移量
	ElcpAbsoluteOffsetV     = 15 // 垂直偏移量
	ElcpAbsoluteOffsetVType = 16 // 位置类型_垂直偏移量

	// 绝对布局位置类型
	ElcpAbsoluteTypeUnknown = 0 // 未知 (未设置或保持不变)
	ElcpAbsoluteTypePx      = 1 // 像素
	ElcpAbsoluteTypePs      = 2 // 百分比
	ElcpAbsoluteTypeObjps   = 3 // 组件尺寸百分比，仅OFFSET可用
	// 布局事件
	ElnGetpropscount        = 1  // 获取布局父属性个数
	ElnGetchildpropcount    = 2  // 获取布局子属性个数
	ElnInitprops            = 3  // 初始化父属性列表
	ElnUninitprops          = 4  // 释放父属性列表
	ElnInitchildprops       = 5  // 初始化子属性列表
	ElnUninitchildprops     = 6  // 释放子属性列表
	ElnCheckpropvalue       = 7  // 检查属性值是否正确,wParam为propID，lParam为值
	ElnCheckchildpropvalue  = 8  // 检查子属性值是否正确,wParam低位为nIndex，高位为propID，lParam为值
	ElnUpdate               = 15 // 更新布局
)

type JouiObjFlags uint32
const (
	EopDefault JouiObjFlags = 0x80000000 // 组件位置默认值
)

type JouiCBEventFlags int

const (
	WmmCbnSelchanger    JouiCBEventFlags = 2    // 现行选中 lParam为当前选中ID 返回1可拦截
	WmmCbnEditchanger                    = 5    // 编辑内容被改变
	WmmCbnDropdownr                      = 7    // 即将弹出列表
	WmmCbnCloseupr                       = 8    // 即将关闭列表
	WmmCbnPopuplistwindow                = 2001 // 弹出下拉列表
	WmmCbnCustomdraw                     = 2002 // 自定义绘制列表 wParam为绘制项目 lParam为ps_context指针
)

type JouiLSEventFlags int

const (
	LbnClick JouiLSEventFlags = 1 // 单击 wParam 索引
	LbnCheck                  = 2 // 选中 wParam 索引,lParam 状态
)

type JOEffect int

const (
	JeDefault JOEffect = iota //默认

	// 二次渐变
	QuadraticIn
	QuadraticOut
	QuadraticInOut

	// 正弦渐变
	SinusoidalIn
	SinusoidalOut
	SinusoidalInOut

	// 指数渐变
	ExponentialIn
	ExponentialOut
	ExponentialInOut

	// 圆曲线
	CircularIn
	CircularOut
	CircularInOut

	// 三次方
	CubicIn
	CubicOut
	CubicInOut

	// 四次方
	QuarticIn
	QuarticOut
	QuarticInOut

	// 五次方
	QuinticIn
	QuinticOut
	QuinticInOut

	// 指数衰减正弦曲线
	ElasticIn
	ElasticOut
	ElasticInOut

	// 回退
	BackIn
	BackOut
	BackInOut

	// 弹性缓动
	BounceIn
	BounceOut
	BounceInOut
)
type JouiTreeViewFlags int

const (
	TviFirst JouiTreeViewFlags = -65535 // 首节点
	TviLastt                    = -65534 // 尾节点
	TviRoott                    = -65536 // 根节点
	TviSortt                    = -65533 // 排序

	// 树形框风格
	EosTreeShowaddandsub = 4096 // 显示加减号
	EosTreeShowcable     = 8192 // 显示连接线

	// 消息树形框
	RtvmMsgSetwaddandsubColor = 4000  // 设置加减号颜色
	RtvmMsgSetcableColor      = 4001  // 设置连线颜色
	RtvmMsgDeleteitem        = 4353  // 删除节点及所有子孙 (lParam为项目ID,TVI_ROOTT为删除所有)
	RtvmMsgExpand            = 4354  // 展开收缩 (wParam为是否展开,lParam为设置的节点句柄)
	RtvmMsgCleareitem        = 4355  // 清空节点 wParam为项目ID
	RtvmMsgGetitemrect       = 4356  // 取节点矩形 (wParam为项目ID,lParam为 RECT指针)
	RtvmMsgGetcount          = 4357  // 取节点数
	RtvmMsgGetindent         = 4358  // 取留白宽度
	RtvmMsgSetindent         = 4359  // 设置留白宽度 取相关节点(wParam为 TVGN_ 开头的常量)
	RtvmMsgGetimagelist      = 4360  // 获取图片组
	RtvmMsgSetimagelist      = 4361  // 设置图片组(wParam为是否更新表项宽高,lParam为图片组句柄)
	RtvmMsgGetnextitem       = 4362  // 取相关节点(wParam为 TVGN_ 开头的常量,lParam为项目ID) 返回H_TREEVIEW_NODEITEM指针
	RtvmMsgSelectitem        = 4363  // 现行选中 lParam为0 获取
	RtvmMsgGetvisiblecount   = 4368  // 取现行选中项 返回选中项目ID
	RtvmMsgHittest           = 4369  // 命中测试 (wParam低位为x高位为y[相对控件],lParam为 返回#TVHT_开头常量 的指针,消息返回值为命中的节点句柄)
	RtvmMsgGethotitem        = 4371  // 取鼠标所在节点项目ID
	RtvmMsgEnsurevisible     = 4372  // 保证显示 (lParam为显示的项目ID)
	RtvmMsgInsertitem        = 4352  // 插入节点 (lParam为 H_TREEVIEW_ROWINFO 指针)
	RtvmMsgGetitem           = 4364  // 取节点信息 (wParam为项目ID,lParam为 H_TREEVIEW_NODEITEM 指针 不要自行释放，pwzText为Unicode)
	RtvmMsgSetitem           = 4365  // 设置节点信息 (wParam为项目ID,lParam为 H_TREEVIEW_NODEITEM 指针)
	RtvmMsgUpdate            = 5001  // 更新树形框
	RtvmMsgSetitemheight     = 5091  // 设置行高 (lParam为新行高)
	RtvmMsgGetitemheight     = 5092  // 获取行高
	RtvmMsgGetnodefromindex  = 5093  // 获取节点句柄 wParam为项目ID 返回 H_TREEVIEW_NODEITEM 指针 不要自行释放
	RtvmMsgSetid             = 5094  // 设置ID选中
	RtvmMsgSetitemtextw      = 14414 // 设置节点标题 (wParam为项目ID,lParam为 文本指针,Unicode)
	RtvmMsgGetitemtextw      = 14415 // 获取节点标题 (wParam为项目ID,返回值为标题Unicode字符串,不要自行释放)

	TvgnRoott        = 0 // 获取根节点
	TvgnNextt        = 1 // 获取下一个节点
	TvgnPrevioust    = 2 // 获取上一个节点
	TvgnParentt      = 3 // 获取父节点
	TvgnChildt       = 4 // 获取子节点
	TvgnNextvisiblet = 6 // 获取下一个可见节点

	// 树形框命中测试
	TvhtNowhere      = 1  // 没有命中
	TvhtOnitemicon   = 2  // 命中图标
	TvhtOnitemlabel  = 4  // 命中标题
	TvhtOnitemindent = 8  // 命中留白
	TvhtOnitemstateicon = 64 //命中加减框

    //事件树形框
	WmmTvnDeleteitemt   = 391  // 删除节点 wParam项目ID lParam 删除类型：0所有、1单删除时
	WmmTvnItemexpanded  = 394  // 节点展开 wParam状态 lParam项目ID
	WmmTvnItemexpanding = 395  // 节点展开中 wParam状态 lParam项目ID
	WmmTvnDrawitem      = 3099 // 绘制节点 wParam绘制索引 lParam为ps_customdraw
	WmmTvnSelectall     = 1011 // 多选被选中 wParam状态 lParam项目ID
	WmmTvnSelect        = 1012 // 被选中 lParam项目ID
)
type JouigooeymenuFlags int
const (
    WmmGmenuSelect      JouigooeymenuFlags = 1012 // 被选中 lParam项目ID
	WmmGmenuOrexpan                      = 1013 // 即将展开 返回1可阻止
	WmmGmenuIsdcontract                  = 1014 // 收缩完成
	// 消息
	GmenuMsgGelect      = 10030 // 取选中ID
	GmenuMsgExpandcontract = 10033 // 展开或收缩
	GmenuMsgItemsize    = 10044 // 项目大小 lParam设置大小 默认40
)

type JouiTableFlags uint32
const (
	EosTableShowtable            JouiTableFlags = 0x20   // 表格方案
	EosTableDrawhorizontalline                  = 0x100  // 绘制横线
	EosTableDrawverticalline                    = 0x200  // 绘制竖线
	EosTableNohead                              = 0x400  // 无表头
	EosTableAllowmultiple                       = 0x800  // 允许多选
	EosTableAllowediting                        = 0x1000 // 允许双击编辑

	// 报表表头风格
	CsTableClickable     = 1  // 可点击
	CsTableLockwidth     = 2  // 锁定宽度
	CsTableCheckbox      = 4  // 选择框
	CsTableSortableAb    = 8  // 可排序-字母
	CsTableSortableNumber = 16 // 可排序-数值

	// 事件报表
	WmmRlvnColumnclick  = 97000 // 表头被单击
	WmmRlvnDrawTr       = 97001 // 绘制表行
	WmmRlvnDrawTd       = 97002 // 绘制表项
	WmmRlvnCheck        = 97003 // 检查框点击
	WmmRlvnDeleteItem   = 97004 // 当删除表项
	WmmRlvnCentchange   = 97005 // 编辑内容改变 wParam所在列 lParam所在行

	// 消息报表
	RlvmMsgCheck        = 99001 // 检查框点击
	RlvmMsgSetcheck     = 99002 // 置检查框状态
	RlvmMsgGetcheck     = 99003 // 获取检查框状态
	RlvmMsgGethitcol    = 99004 // 获取命中列索引
	// 报表表行风格
	RsTableCheckbox   = 1 // 检查框默认状态
	RsTableCheckboxok = 2 // 检查框为选中状态
	// 项目风格
	EsTableDisableed = 1 // 禁止编辑
)

type JouiObjpagingFlags int
const (
	// 事件类型
	PagingCrbkgnormal   JouiObjpagingFlags = 1 // 背景_正常
	PagingCrbkghover                       = 2 // 背景_点燃
	PagingCrbkgchecked                     = 3 // 背景_选中
	PagingCrbkgban                         = 4 // 背景_禁止
	PagingCrbordernormal                   = 5 // 边框_正常
	PagingCrborderhover                    = 6 // 边框_点燃
	PagingCrborderchecked                  = 7 // 边框_选中
	PagingCrborderban                      = 8 // 边框_禁止
	PagingCrtextnormal                     = 9 // 文本_正常
	PagingCrtexthover                      = 10 // 文本_点燃
	PagingCrtextchecked                    = 11 // 文本_选中
	PagingCrtextban                        = 12 // 文本_禁止
	PagingCrcromitnormal                   = 13 // 省略_正常
	PagingCrcromithover                    = 14 // 省略_点燃
	// 风格类型
	EosPagingPnarrow  = 0  // 上页下页风格_箭头
	EosPagingPnatext  = 4  // 上页下页风格_文本
	EosPagingShowtips = 8  // 显示提示文本
	EosPagingShowjump = 16 // 显示跳转

	PageMsgSetcount      = 10010 // 设置总页 wParam为总页数 lParam为每页显示
	PageMsgSetcurrent    = 10012 // 设置或获取当前页 wParam为0返回当前页
	PageMsgSetcolorchange = 10020 // 设置颜色 wParam=PAGING_类型 lParam颜色值
	PageMsgUpdate        = 10022 // 刷新
	PageMsgSetitemtext   = 10030 // 设置按钮标题 wParam:类型 lParam标题指针
	PageMsgSetinterval   = 10033 // 设置按钮间隔 lParam＞0 为0获取

	WmmPageSelchanger = 10200 // 事件现行选中 lParam为当前页码，返回1可拦截
)

type JouiObjlvmlFlags int

const (
	LvmlMsgGetitemcount     JouiObjlvmlFlags = 4100  // 取表项总数
	LvmlMsgGetitem                           = 4101  // 获取表项 (LISTBUTTON wParam为项目索引, lParam为H_REPORTLIST_ITEMINFO指针或H_LISTBUTTON_ITEMINFO指针)
	LvmlMsgSetitem                           = 4102  // 设置表项 (wParam为是否重画,lParam为H_REPORTLIST_ITEMINFO或H_LISTBUTTON_ITEMINFO指针)
	LvmlMsgInsertitem                        = 4103  // 插入表项 lParam 为H_REPORTLIST_ROWINFO指针,wParam为是否立即重画,返回索引
	LvmlMsgDeleteitem                        = 4104  // 删除表项,wParam为是否立即重画，lParam为删除的索引
	LvmlMsgDeleteallitems                    = 4105  // 清空表项
	LvmlMsgGetitemrect                       = 4110  // 取表项矩形
	LvmlMsgHittest                           = 4114  // 命中测试 lParam为 返回列表命中测试
	LvmlMsgEnsurevisible                     = 4115  // 保证显示表项
	LvmlMsgRedrawitems                       = 4117  // 重画表项 wParam为起始项目,lParam 为结束项目
	LvmlMsgGetcolumn                         = 4121  // 获取列信息 (wParam为列索引,lParam为 H_REPORTLIST_COLUMNINFO 指针)
	LvmlMsgSetcolumn                         = 4122  // 设置列信息
	LvmlMsgInsertcolumn                      = 4123  // 插入列 (wParm为是否立即更新,lParam为H_REPORTLIST_COLUMNINFO指针)
	LvmlMsgDeletecolumn                      = 4124  // 删除列 (wParm为是否立即更新,lParam为列索引)
	LvmlMsgGetcolumnwidth                    = 4125  // 获取列宽
	LvmlMsgSetcolumnwidth                    = 4126  // 设置列宽 (wParam为列索引,lParam为 列宽)
	LvmlMsgGetTopindex                = 4135 //取可视区起始索引
	LvmlMsgGetCountperpage           = 4136 //取可视区表项数
	LvmlMsgUpdate                    = 4138 //更新列表框
	LvmlMsgSetitemstate              = 4139 //置表项状态
	LvmlMsgGetitemstate              = 4140 //取表项状态
	LvmlMsgGetitemimage              = 4141 //获取表项图片索引 wParam若不为0则为表项索引 返回UIImage
	LvmlMsgSetitemimage              = 4142 //设置表项图片索引 (wParam若不为0则为表项索引,lParam为新图片UIImage)
	LvmlMsgSetitemcount              = 4143 //设置表项总数 wParam为表项条数,lParmam为是否立即刷新
	LvmlMsgSortitems                 = 4144 //排序 (lParam为H_REPORTLIST_SORTINFO指针)
	LvmlMsgGetselectedcount          = 4146 //取被选择表项数
	LvmlMsgGetitemselect             = 4147 //取表项选中状态
	LvmlMsgGethotitem                = 4157 //取鼠标所在表项
	LvmlMsgGetselectionmark          = 4162 //取现行选中项
	LvmlMsgSetselectionmark          = 4163 //置现行选中项
	LvmlMsgDeleteallcolumn           = 4900 //删除所有列
	LvmlMsgGetcolumncount            = 4901 //获取列数
	LvmlMsgSetcolumntext             = 4904 //设置列标题
	LvmlMsgGetcolumntext             = 4905 //获取列标题
	LvmlMsgSetitemheight             = 4908 //设置表项高度 (lParam为新行高)
	LvmlMsgGetitemheight             = 4909 //获取表项高度
	LvmlMsgCalcitemsize              = 5150 //重新计算尺寸
	LvmlMsgSetselect                 = 5151 //全选 lParam=0取消 1选中  报表不需要发送此消息
	LvmlMsgSetitemcolor              = 5152 //设置项目背景颜色 lParam为H_LISTCOLOR
	LvmlMsgAllowmultiple             = 5153 //允许多选
	LvmlMsgSetcolourwidth            = 5551 //设置线宽和颜色 lParam设置线宽 wParam设置颜色
	LvmlMsgMerge                     = 6001 //合并单元格、合并后以第一单元内容形式表现 其它单元数据会释放 lParam为H_REPORTLIST_SELECTRECT结构
	LvmlMsgDecompose                 = 6002 //分解单元格、分解表格中指定的已经组合的单元格，行列参数必须指向被组合单元格的第一个单元格 wParam起始行索引 lParam起始列索引
	LvmlMsgGetselectrect             = 6003 //获取多选信息 返回H_REPORTLIST_SELECTRECT结构
	LvmlMsgSetselectrect             = 6004 //设置多选信息 lParam为H_REPORTLIST_SELECTRECT结构
)

type JouiObjCselFlags int

const (
	CarouselMsgInsertitem JouiObjCselFlags = 4103 // 插入
	CarouselMsgSetimgsize                  = 4104 // 设置图像尺寸 wParam宽度 lParam高度
	// 走马灯风格
	EosCarouselRound     = 1  // 圆角风格 附加参数设置圆角度
	EosCarouselIndicator = 4  // 显示指示器
	EosCarouselEasing    = 8  // 使用缓动
	// 走马灯事件
	EosCarouselSelchanger = 99001 // 被选中 lParam选中ID
)

type JouiTipsFlags uint32
const (
    TipsInfo    JouiTipsFlags = 0      // 信息
	TipsError                   = 1      // 错误
	TipsEnquire                 = 1 << 1 // 询问
	TipsSuccess                 = 1 << 2 // 成功
	TipsWarning                 = 1 << 3 // 警告

	TipsTypeCenter      = 0      // 方向 居中上
	TipsTypeLeftTop     = 1 << 4 // 方向 靠左上
	TipsTypeRightTop    = 1 << 5 // 方向 靠右上
	TipsTypeRightBottom = 1 << 6 // 方向 靠右下
	TipsTypeCenterBottom = 1 << 7 //方向 居中下
)

type CodeType int

const (
	CtAnsi CodeType = iota
	CtUtf8
	CtUtf8NoBom
	CtUtf16Le
	CtUtf16Be
	CtAuto
)

type JouiNodeFlags int

const (
	GwfHwndnext JouiNodeFlags = 2 // 下一个句柄
	GwfHwndprev                = 3 // 上一个句柄
	GwfChild                   = 5 // 子句柄
)

type JouiMsgFlags int

const (
	MsgIdok     JouiMsgFlags = 1 // 确认钮
	MsgIdcancel              = 2 // 取消钮
	MsgIdyes                 = 6 // 是钮
	MsgIdno                  = 7 // 否钮
	MsgIdclose               = 8 // 关闭钮
)

type JouiObjProressFlags int

const (
	// 消息进度条
	PbmMsgSetrange    JouiObjProressFlags = 1025 // lParam设置进度条范围
	PbmMsgSetpos                          = 1026 // lParam设置进度条位置
	PbmMsgGetrange                        = 1031 // 获取进度条范围
	PbmMsgGetpos                          = 1032 // 获取进度条位置
	PbmMsgSetbarcolor                     = 1033 // lParam设置进度条颜色
	PbmMsgSetbkcolor                      = 8193 // lParam设置进度条背景颜色
	PbmMsgSetradius                       = 1027 // lParam设置进度条圆角度 -1获取
	PbmMsgSetdonut                        = 1028 // lParam设置进度条圆环圆弧大小
	PbmMsgWaveSetwidth                    = 10029 // lParam设置波纹宽度 10-150
	PbmMsgWaveSetheight                   = 10030 // lParam设置波纹高度
	PbmMsgWaveSetbubblespeed              = 10034 // lParam设置波纹气泡速度（0表示不启用）

	// 进度条风格
	EosPbmHorizontal = 0  // 横向
	EosPbmVertical   = 1  // 纵向
	EosPbmArc        = 4  // 圆弧
	EosPbmDonut      = 8  // 圆环
	EosPbmWave       = 12 // 波浪
)

type JouiObjSoliderFlags int
const (
	EosBsHorizontal    JouiObjSoliderFlags = 0x00 // 横向
	EosBsVertical                          = 0x01 // 纵向
	EosBsExtendconcise                     = 128  // 简洁版
	EosBsHotspot                           = 256  // 热点跟踪

	// 消息滑块条
	SbmMsgGetblockrect = 10010 // 取当前滑块坐标
	SbmMsgPt2Value     = 10011 // 坐标转值
	SbmMsgSetvalue     = 10014//SBM_SETPOS         // 当前值  //SBM_SETPOS已被定义
	SbmMsgSetmaxbs     = 10015 // 设置滑块大小 wParam滑块大小 lParam滑块背景大小
	SbmMsgSetblockPoint = 10016 // lParam滑块圆滑动方向 设定值：1，横向风格（从右往左）|纵向风格（从下往上）
	SbmMsgSetmax       = 10017 // lParam设置最大值 -1获取
	SbmMsgSetmin       = 10018 // lParam设置最小值 -1获取
	// 事件滑块条
	WmmSbnValue = 10010 // 值改变 事件编号(lParam=值)
)

type JouiObjPicFlags int

const (
	EosImageNormal   JouiObjPicFlags = 0      // 默认
	EosImageRotate                   = 1      // 旋转
	EosImageSvg                      = 1 << 1 // SVG
	EosImageVignette                 = 1 << 2 // vignette效果
	EosImageRound                    = 1 << 3 // 圆角
	EosImageProjection               = 1 << 4 // 投影效果
	EosImageTurbulence               = 1 << 5 // 湍流效果

	// 图片框消息
	PicMsgSetbloomcolour = 7000 // 设置投影颜色 Bloom效果有效 lParam为ARGB wParam为模糊度
	PicMsgSetimage       = 7103 // 设置图像 lParam为图像句柄
	PicMsgSetedgevalue   = 7104 // 设置淡化值
	PicMsgSetedgecolour  = 7105 // 设置淡化颜色
)

type JouiObjMbFlags int
const (
    MbTypeUrl    JouiObjMbFlags = 0 // 加载类型_URL
	MbTypeFile                  = 1 // 加载类型_文件
	MbTypeHtml                  = 2 // 加载类型_HTML
	// 消息
	MbMsgDll          = 100010 // 消息_浏览框_设置动态库路径
	MbMsgGetwebview   = 100011 // 消息_浏览框_获取浏览框句柄
	MbMsgLoad         = 100012 // 消息_浏览框_加载 wParam为加载类型：MB_TYPE_， lParam为地址
	MbMsgGoback       = 100013 // 消息_浏览框_后退
	MbMsgGoforward    = 100014 // 消息_浏览框_前进
	MbMsgReload       = 100015 // 消息_浏览框_刷新
	MbMsgExecscript   = 100016 // 消息_浏览框_执行脚本 lParam为脚本内容
	// 事件
	WmmMbTitlechanged  = 10001 // 设置标题变化的通知回调 lParam为当前标题
	WmmMbUrlchanged    = 10002 // url改变回调  lParam为当前url
	WmmMbAlert         = 10003 // 网页调用alert会走到这个接口填入的回调 lParam为内容
	WmmMbNavigation    = 10004 // wkeNavigationCallback回调的返回值，如果是false，表示可以继续进行浏览，true表示阻止本次浏览。wParam为wkeNavigationType lParam为当前标题
	WmmMbCreateview    = 10005 // 网页点击a标签创建新窗口时将触发回调 wParam为wkeWindowFeatures lParam为当前url
	WmmMbDocumentready = 10006 // 对应js里的body onload事件
	WmmMbDownload      = 10007 // 页面下载事件回调。点击某些链接，触发下载会调用 lParam为 string url
)

type JouiffvideoFlags int

const (
	// 消息
	FfMsgDll        JouiffvideoFlags = 4000 // lParam设置动态库路径
	FfMsgPlay                        = 4100 // 播放 lParam = -1判断播放状态、为0暂停、为1播放
	FfMsgStop                        = 4201 // 停止 lParam = 0判断停止状态
	FfMsgResume                      = 4203 // 继续
	FfMsgIsdecode                    = 4204 // 是否解码
	FfMsgSeek                        = 4205 // lParam设置播放位置
	FfMsgSetvolume                   = 4206 // lParam设置音量
	FfMsgGettime                     = 4207 // 获取时间
	FfMsgGetvideoprogress            = 4208 // 获取播放进度
)

type JouiObjDateFlags int

const (
	DbmMsgGetdatetime JouiObjDateFlags = 100060 // 取日期 返回H_DATETIME结构 用完需要自行释放内存
	DbmMsgSetdatetime                  = 100061 // 日期框置时间 lParam为H_DATETIME

	// 事件日期框
	WmmDbnDatetime = 100062 // 日期框已更改
)

type JouihotkeyMod int
const (
     HotkeyModNone JouihotkeyMod = iota
	 HotkeyModAlt
	 HotkeyModCtrl
	 HotkeyModShift = 4
)

type JouiObjOtherFlags int

const (
	// 菜单消息
	MnMsgSelectitem = 0x1E5 // 选择菜单项目
	// 消息颜色选择器
	ColorpickerSelectcolor = 10001 // 设置ARGB颜色 lParam为新颜色
	ColorpickerGetectcolor = 10002 // 获取ARGB颜色

	// 通用消息
	OverallMsgSelectcolor = 20001 // 设置ARGB颜色 lParam为新颜色
	OverallMsgGetectcolor = 20002 // 获取ARGB颜色
	// 事件颜色选择器
	WmmColorpickerChange = 70001 // 颜色选择器颜色已更改
	// 事件视频播放器
	WmmFfvideoEnded = 10776 // 播放结束
	// 风格dock
	EosDockDisablemenu = 0x02 // 禁用右键默认菜单
	EosDockTips        = 0x04 // 提示文本
	EosDockStartani    = 0x08 // 启用动画
	// 事件dock
	WmmDockLclick = 10777 // 左键按下被选中 wParam选中地址 lParam选中ID
	WmmDockRclick = 10778 // 右键按下被选中 wParam选中地址 lParam选中ID
	// 分组框消息
	GrolMsgInsertgroup = 4100 // 插入信息 lParam 为H_GROUPBOX指针
	// hotkey
	HotkeySetkeyw = 10000 // 设置组合热键文本 lParam为组合热键文本 组合键尽量统一小写 如“ctrl+f”
	HotkeySetkey  = 10001 // 设置组合热键 wParam为vKey值 lParam为wModifiers组合键 参考JouihotkeyMod
	HotkeyGetkey  = 10002 // 返回vKey值
	HotkeyGetmodifiers = 10003 //返回wModifiers组合键值 参考JouihotkeyMod
    //blurb
	EosBlurbMov = 16 //自移动模式
	//adjtool
	EosAdjtoolMov = 16 //方案二
)

type JouiWaveRingView int
const (
	//消息
	WrviewMsgSetbassdll     JouiWaveRingView = 10000 // lParam设置bass动态库路径
	WrviewMsgSethandle                       = 10001 // lParam设置播放文件句柄
	WrviewMsgSettype                         = 10002 // lParam设置显示类型 1动感の圆、2波纹siri、3六芒星
	WrviewMsgSetrotate                       = 10003 // lParam设置是否旋转 ,-1为获取
	WrviewMsgSetrandom                       = 10004 // lParam设置是否随机角度 ,-1为获取
	WrviewMsgSetdrawbase                     = 10005 // lParam设置是否绘制基线 ,-1为获取
	WrviewMsgSetdrawwave                     = 10006 // lParam设置是否绘制波纹 ,-1为获取
	WrviewMsgSetdrawpoint                    = 10007 // lParam设置是否绘制能量点 ,-1为获取
	WrviewMsgSetpoweroffset                  = 10008 // lParam设置是否均衡能量强度 ,-1为获取
	WrviewMsgSetspread                       = 10009 // lParam设置能量点是否扩散 ,-1为获取
	WrviewMsgSetdrawtext                     = 10010 // lParam设置是否绘制文本 ,-1为获取
	WrviewMsgSettextmove                     = 10011 // lParam设置文本、头像是否抖动 ,-1为获取
	WrviewMsgSettextrandommove               = 10012 // lParam设置文本是否随机抖动 ,-1为获取
	WrviewMsgSetscope                        = 10013 // lParam设置能量阈值 1-20,-1为获取
	WrviewMsgSetdistance                     = 10014 // lParam设置能量点与圆环距离 ,-1为获取
	WrviewMsgSetspeed                        = 10015 // lParam设置能量点扩散速度 1-20,-1为获取
	WrviewMsgSetdrawfx                       = 10016 // lParam设置是否绘制特效 ,-1为获取
	WrviewMsgSetfxcolor                      = 10017 // lParam设置特效颜色
	WrviewMsgSetfxbloom                      = 10018 // lParam设置特效是否开启Bloom
	WrviewMsgSetvalue                        = 10019 // lParam设置能量取值 ,-1为获取
	WrviewMsgSetisfill                       = 10020 // lParam设置是否填充
	WrviewMsgSetplayfile                     = 10021 // lParam设置播放文件路径 由内部打开
	WrviewMsgSetalbumsize                    = 10022 // lParam设置专辑图像大小
	WrviewMsgSetalbumimages                  = 10023 // lParam设置专辑图像 必须先设置专辑图像大小
)

//--- 常量 ---
const (
	RgnCombineUnion     = 0 // 并集,采用两个区域的并集来合并这两个区域
		RgnCombineIntersect = 1 // 交集,采用两个区域的交集来合并这两个区域
	RgnCombineXor       = 2 // 异或,采用两个区域的并集，且去除重叠区域
	RgnCombineExclude   = 3 // 排除,从第一个区域中排除第二个区域
)

// --- 辅助函数 (这些在 Go 中通常不需要，但为了完整性，可以保留) ---

// GlRgb 在 Go 中，可以直接使用 (0xFF << 24) | (r << 16) | (g << 8) | b
func GlRgb(r, g, b int) uint32 {
	return (0xFF << 24) | (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
}

func RgbGetB(argb uint32) byte {
	return byte(argb)
}
func RgbGetG(argb uint32) byte {
	return byte(argb >> 8)
}

func RgbGetR(argb uint32) byte {
	return byte(argb >> 16)
}

func RgbGetA(argb uint32) byte {
	return byte(argb >> 24)
}

func Rgba(r, g, b, a byte) uint32 {
    return uint32(b) | (uint32(g) << 8) | (uint32(r) << 16) | (uint32(a) << 24)
}

func Argb(r, g, b, a byte) uint32 {
	return uint32(b) | (uint32(g) << 8) | (uint32(r) << 16) | (uint32(a) << 24) //和上面的一样
}
func Argb2Rgb(argb uint32) uint32 {
	return GlRgb(int(argb>>16), int(argb>>8), int(argb)) //复用上面的
}

func Rgb2Argb(rgb, a uint32) uint32 {
	return uint32(byte(RgbGetB(rgb))<<16) | uint32(byte(RgbGetG(rgb))<<8) | uint32(byte(RgbGetR(rgb))) | (a << 24)
}

func Argb2alphA(argb uint32, a byte) uint32 {
    return Argb(byte(argb>>16), byte(argb>>8), byte(argb), a)
}
// --- 结构体 ---
// 因为 Go 结构体字段需要大写开头，所以这里做了相应的修改。

type IDXdevice struct {
	DeviceName    *uint16   // LPCWSTR 设备名称
	DeviceIndex   int32     // 设备索引
	Description   *uint16   // LPCWSTR 设备描述
	PAdapter      uintptr   // LPVOID 适配器指针
	SharedSystem  uintptr  // size_t 共享内存
	DedicatedVideo uintptr // size_t 专用显存
	DedicatedSystem uintptr //size_t 专用系统内存
	FeatureLevel  uint32    //D3D_FEATURE_LEVEL 特性级别
	DriverType    int32      //INT 驱动类型
}

type H_NOTIFY struct {
	HObjFrom UIView // 	组件句柄
	IdFrom   int32  // 	组件ID
	NCode    int32  // 	通知消息
	WParam   uintptr  // 	无符号整数 通常是一个与消息有关的常量值，也可能是窗口或控件的句柄
	LParam   uintptr  // 	长整型 通常是一个指向内存中数据的指针
}

type H_GROUPBOX struct {
	Radius     float32 //FLOAT
	TextOffset float32 //FLOAT
	Strokewidth float32 //FLOAT
}

type RECTF struct {
	Left   float32 //FLOAT
	Top    float32
	Right  float32
	Bottom float32
}

type H_POINT struct {
	X float32 //FLOAT
	Y float32 //FLOAT
}

type H_INITINFO struct {
	HFileTheme   *uint16    // LPCWSTR 主题文件路径
	HZipTheme    uintptr    // LPVOID ZIP主题数据
	HZipThemeLen uint32      // DWORD ZIP主题数据长度
	HZipThemeKey *byte    // LPCSTR ZIP主题密钥
	DwScaledpi   float32  // FLOAT DPI缩放
	DwDebug      bool     // BOOL 调试标志
}
// RECT 结构体，与 C 语言的 RECT 结构体对应
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// PRECT, NPRECT, LPRECT 在 Go 中都是 *RECT
// 可以选择定义类型别名，也可以不定义，直接使用 *RECT

type PRECT = *RECT


type PsContext struct {
	HCanvas     UICanvas       // 	画布句柄
	HTheme      UIZip       // 	主题句柄
	DwStyle     int32     // 	风格
	DwStyleEx   int32     // 	扩展风格
	DwTextFormat int32    // 	文本格式
	HFont       UIFont        // 	字体
	DwState     int32     // 	状态
	DwOwnerData unsafe.Pointer    // 	所有数据
	UWidth      uint32        // 	宽度
	UHeight     uint32        // 	高度
	RcPaint     RECT    // 	绘制矩形
	RcText      RECT    // 	文本矩形
	DwReserved  unsafe.Pointer    // 	保留
}

type PsCustomdraw struct {
	HCanvas     UICanvas       // 	画布句柄
	HTheme      UIZip       // 	主题句柄
	DwState     uint32     // 	状态
	DwStyle     uint32     // 	风格
	DwTextFormat int32    // 	文本格式
	HFont       UIFont        // 	字体
	IItem      int32
	IItemParam uintptr
	RcPaint     RECT // 	绘制矩形
	UWidth      uint32       // 	宽度
	UHeight     uint32       // 	高度
}

type H_BITMAPDATA struct {
	Width       uint32
	Height      uint32
	Stride      int32
	PixelFormat int32
	Scan0       *byte
	Reserved    unsafe.Pointer
}

type H_FLOWST struct {
	Crstart ARGB // 开始颜色
	Crend   ARGB // 结束颜色
	Crscan  ARGB // 扫描颜色
}

type H_LISTCOLOR struct {
	ColorHover  ARGB //
	ColorDown   ARGB
	ColorSelect ARGB // 多选背景颜色
}
// 缓动信息结构
type H_EASINGINFO struct {
	NProgress float64 //  进度[0-1]  //DOUBLE
	NCurrentX float64   //  当前值X //DOUBLE
	NCurrentY float64   //  当前值Y //DOUBLE
	Param1    uintptr    //  参数1
	Param2    uintptr    //  参数2
	Param3    uintptr    //  参数3
	Param4    uintptr    //  参数4
}
type H_REPORTLIST_COLUMNINFO struct {
	PwzText     *uint16       // 表头标题 LPCWSTR
	NWidth      uint32        // 列宽度
	DwStyle     uint32        // 表头风格
	DwTextFormat int32 // 列文本格式 -1默认格式
	CrText      ARGB          // 列文本颜色
	Crbk        ARGB          // 列背景颜色
	CrbkHover   ARGB          // 列背景点燃颜色
	NImage      UIImage       // 图标
	NInsertIndex uint32 // 插入位置,0为在最后
}

type H_REPORTLIST_SELECTRECT struct {
	IRowStart uint32 // 起始行
	IRowEnd   uint32 // 终止行
	IColStart uint32 // 起始列
	IColEnd   uint32 // 终止列
}

type H_REPORTLIST_ITEMINFO struct {
	IRow      uint32    // 所在行[IN / OUT]
	ICol      uint32    // 所在列[IN / OUT]
	PwzText   *uint16   // 项目文本 LPCWSTR
	CrText    ARGB      // 项目文本颜色
	Crbk      ARGB      // 项目背景颜色
	LParam    uintptr   // 项目参数
	NImageIndex UIImage // 项目图片索引 为0清除图像;为-1保持原图像
	DwStyle   uint32    // 项目风格
	HObjItem  UIView    // 元素句柄
}

type H_REPORTLIST_ROWINFO struct {
	NInsertIndex uint32 // 插入位置,0为最后
	DwStyle      uint32 // 项目行风格(同行共用)
}
type H_REPORTLIST_SORTINFO struct {
	ICol    uint32  //
	NType   uint32  // 0:文本,1:整数
	FDesc   bool    // 是否倒序
	LParam  uintptr // 排序附加参数
}
type H_TREEVIEW_NODEITEM struct {
	ICol         uint32             // 所在列
	NID          int32              // 项目ID
	PwzText      *uint16            // 项目标题 LPCWSTR
	LParam       uintptr            // 项目附加参数
	NImageIndex  UIImage            // 图片句柄 -1清除
	FExpand      bool               // 是否展开
	NDepth       int32              // 层次
	Color        ARGB
	Select       int32             // 选中 -1保持默认 0取消选中 1选中
	PParent      *H_TREEVIEW_NODEITEM // 父节点36
	PPrev        *H_TREEVIEW_NODEITEM // 上一个节点40
	PNext        *H_TREEVIEW_NODEITEM // 下一个节点44
	PChildFirst  *H_TREEVIEW_NODEITEM // 第一个子节点48
	NCountChild  int32              // 子节点数量52
	ImgWidth    int32             // 不需要指定
	ImgHeight   int32             // 不需要指定
}

type H_TREEVIEW_ITEMINFO struct {
	NID          int32     // 项目ID0
	PwzText      *uint16   // 项目标题4 LPCWSTR
	LParam       uintptr   // 项目附加参数8
	NImageIndex  UIImage   // 图片句柄12
	FExpand      bool      // 是否展开16
	NDepth       int32     // 层次20
	Color        ARGB      //24
	Select       int32     // 选中28 -1保持默认 0取消选中 1选中
}
type H_TREEVIEW_INSERTINFO struct {
	ItemParent      *H_TREEVIEW_NODEITEM // 父项句柄（0为根项）
	ItemInsertAfter *H_TREEVIEW_NODEITEM // 插入在此项之后（必须是同层）
	ICol            uint32               // 所在列
	NID             int32                // 项目ID12
	PwzText         *uint16              // 项目标题16 LPCWSTR
	LParam          uintptr              // 项目附加参数20
	NImageIndex     UIImage              // 图片句柄24
	FExpand         bool                 // 是否展开28
	FUpdate         bool                 // 是否更新(如果循环插入建议统一用RTVM_MSG_UPDATE更新)32
	Color           ARGB
	Select          int32                //选中 -1保持默认 0取消选中 1选中
}
type H_TREEVIEW_ROWINFO struct {
	NInsertIndex uint32 // 插入位置,0为最后
	NImageIndex  UIImage // 项目图片索引
	Item *H_TREEVIEW_INSERTINFO
}
type H_TREEVIEW_COLUMNINFO struct {
	PwzText      *uint16   // 表头标题 LPCWSTR
	NWidth       uint32    // 列宽度
	DwStyle      uint32    // 表头风格 IGETS_CS_
	DwTextFormat uint32 // 列文本格式
	CrText       ARGB      // 列文本颜色
	NInsertIndex uint32 // 插入位置,0为在最后
}

type H_BACKGROUNDIMAGEINFO struct {
	DwFlags   uint32        // 标识
	HImage    UIImage       // 图片句柄
	X         int32         // 左上角横坐标
	Y         int32         // 左上角纵坐标
	DwRepeat  uint32        // 重复方式
	LpGrid    *RECT // 九宫矩形
	LpDelay   unsafe.Pointer     // 延时信息
	CurFrame  uint32        // 当前帧
	MaxFrame  uint32        // 最大帧
	DwAlpha   uint32        // 透明度
	Radius    int32
}
type RegMethod struct {
	DwStyle      int32
	DwStyleEx    int32
	DwTextFormat int32
    HCursor     *uint16 //LPCWSTR
	PfnClsProc   ClsPROC
	PClassName   *uint16 //LPCWSTR
}
type H_OBJ_PROPS struct {
	CrBkgNormal        ARGB      // 背景颜色.正常0
	CrBkgHover         ARGB      // 背景颜色.按钮悬浮4
	CrBkgDownOrChecked ARGB      // 背景颜色.按下或者选中8
	CrBkgBegin         ARGB      // 渐变背景.起点颜色ARGB12 //不适用于开关
	CrBkgEnd           ARGB      // 渐变背景.终点颜色ARGB16 //不适用于开关
	CrBorderNormal     ARGB      // 边框颜色.正常20
	CrBorderHover      ARGB      // 边框颜色.悬浮24
	CrBorderDownOrChecked ARGB   //边框颜色.按下或者选中28
	CrBorderBegin      ARGB      // 渐变边框.起点颜色ARGB32 //不适用于开关
	CrBorderEnd        ARGB      // 渐变边框.终点颜色ARGB36 //不适用于开关
	Radius             int32        // 圆角度40
	StrokeWidth        int32     // 线宽44
	ImgNormal          UIImage
	ImgDown            UIImage
}

type H_DOCK_ITEMINFO struct {
	NIndex    uint32    // 插入位置
	NImageIndex UIImage // 图片句柄
	PwzText   *uint16   // 文本 LPCWSTR
	Depth     *uint16   // 路径 LPCWSTR
	Color     ARGB
	LParam    uintptr
	Type      int32 // 类型 0应用、1其它
}
type H_GOOEYMENU_ITEMINFO struct {
	NIndex      uint32
	NImageIndex UIImage //图片句柄
	PwzText     *uint16 //LPCWSTR
	TpText      *uint16 //提示文本 LPCWSTR
	CrBkg       ARGB
	CrBorder    ARGB
	LParam      uintptr
}
type H_LISTVIEW_ITEMINFO struct {
	NIndex    uint32    //插入位置
	NImageIndex UIImage //图片句柄
	PwzText   *uint16   //文本 LPCWSTR
	LpNewStr  *uint16   //LPCWSTR
	Color     ARGB
	CrBkg     ARGB
	LParam    uint32
}
type H_COMBOX_ITEM struct {
	LpwzTitle *uint16   //LPCWSTR
	Data       uintptr
	NImageIndex UIImage //图片句柄
	NIndex      uint32  //插入位置
	Color       ARGB
}

type H_IMAGEINFO struct {
	ImgNormal        UIImage // 图像.正常
	ImgHover         UIImage // 图像.悬浮
	ImgDownOrChecked UIImage // 图像.按下或者选中
}
type H_DROPINFO struct {
	PDataObject unsafe.Pointer // 数据对象指针IDataObject*
	GrfKeyState uint32         // 功能键状态
	X           int32          // 鼠标水平位置
	Y           int32          // 鼠标垂直位置
}
type H_CHARRANGE struct {
	CpMin int32
	CpMax int32
}
type H_TEXTRANGE struct {
	Chrg    H_CHARRANGE
	PwzText *uint16 //LPCWSTR
}

type NMHDR struct {
	HwndFrom uintptr
	IdFrom   uintptr
	Code     uint32
}

type H_SELCHANGE struct {
	Nmhdr   NMHDR
	Chrg    H_CHARRANGE
	Seltyp  uint16 //WORD
}
type H_ENLINK struct {
	Nmhdr    NMHDR
	Msg      uint32
	WParam   uintptr
	LParam   uintptr
	Chrg     H_CHARRANGE
}
type H_SETTEXTEX struct {
	Flags    uint32
	CodePage uint32
}

type H_LISTBUTTON_ITEMINFO struct {
	DwMask     uint32        // 1,图片 2,标题 4,提示文本 8,状态 16,菜单 32,文本格式 64,宽度
	NType      uint32        // 项目类型   0,分隔条 1,普通按钮 2,选择按钮
	NIndex     uint32        // 插入索引
	NImage     uint32        // 图片索引
	WzText     *uint16       // 项目标题 LPCWSTR
	WzTips     *uint16       // 项目提示文本 LPCWSTR
	NLeft      uint32        // 项目左边
	NWidth     uint32        // 项目宽度
	DwState    uint32        // 项目状态   可取state_normal,state_down,state_focus,state_disable
	NMenu      syscall.Handle // HMENU 项目菜单
	TextFormat int32         // 项目文本格式
}
type H_PACK_INFO struct {
	PwzText  *uint16 // 项目标题 LPCWSTR
	ImgNormal UIImage// 项目默认图标
	CrText   ARGB
	CrBkg    ARGB
	Type   int32 //0默认页面 1自定义链接
	HPageID int32//选择夹子夹ID
}
type H_DATETIME struct {
	Year int32 // 年
	Mon  int32 // 月   1-12
	Mday int32 // 日   1-31
	Wday int32 // 星期 1-7 7=星期日
}

type H_STATICEX_INFO struct{
    PwzText *uint16 // LPCWSTR
    ColorNormal ARGB
    ColorBackg ARGB
    LpwzFontFace *uint16//LPCWSTR
    DwFontSize int32
    DwFontStyle uint32
}
type HK_THUNK_DATA struct {
	HWnd syscall.Handle
	Proc syscall.Handle//WNDPROC
	DwData unsafe.Pointer
}

type ThunkPROC = func(data *HK_THUNK_DATA, msg int, wParam, lParam uintptr) uintptr

type JouiShowFlags int32 //或者int

const (
	SwHide            JouiShowFlags = 0
	SwShownormal                    = 1
	SwNormal                        = 1
	SwShowminimized                 = 2
	SwShowmaximized                 = 3
	SwMaximize                      = 3
	SwShownoactivate                = 4
	SwShow                          = 5
	SwMinimize                      = 6
	SwShowminnoactive               = 7
	SwShowna                        = 8
	SwRestore                       = 9
	SwShowdefault                   = 10
	SwForcemaximize                 = 11
    SwMax                           = 11 //和上面重复，看情况使用
)

type PCTSTR = *uint16
type LPCTSTR = *uint16


// LOGFONTW 结构体
type LOGFONTW struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         int32
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [32]uint16 // LF_FACESIZE 通常是 32
}

// PLOGFONTW, NPLOGFONTW, LPLOGFONTW 在 Go 中都是 *LOGFONTW
type PLOGFONTW = *LOGFONTW


type JouiTextFormatFlags uint32 // 或者 int, DWORD

const (
	DtTop                    JouiTextFormatFlags = 0x00000000
	DtLeft                   JouiTextFormatFlags = 0x00000000
	DtCenter                 JouiTextFormatFlags = 0x00000001
	DtRight                  JouiTextFormatFlags = 0x00000002
	DtVcenter                JouiTextFormatFlags = 0x00000004
	DtBottom                 JouiTextFormatFlags = 0x00000008
	DtWordbreak              JouiTextFormatFlags = 0x00000010
	DtSingleline             JouiTextFormatFlags = 0x00000020
	DtExpandtabs             JouiTextFormatFlags = 0x00000040
	DtTabstop                JouiTextFormatFlags = 0x00000080
	DtNoclip                 JouiTextFormatFlags = 0x00000100
	DtExternalleading        JouiTextFormatFlags = 0x00000200
	DtCalcrect               JouiTextFormatFlags = 0x00000400
	DtNoprefix               JouiTextFormatFlags = 0x00000800
	DtInternal               JouiTextFormatFlags = 0x00001000
	DtEditcontrol            JouiTextFormatFlags = 0x00002000
	DtPathEllipsis           JouiTextFormatFlags = 0x00004000
	DtEndEllipsis            JouiTextFormatFlags = 0x00008000
	DtModifystring           JouiTextFormatFlags = 0x00010000
	DtRtlreading             JouiTextFormatFlags = 0x00020000
	DtWordEllipsis           JouiTextFormatFlags = 0x00040000
	DtNofullwidthcharbreak JouiTextFormatFlags = 0x00080000
	DtHideprefix             JouiTextFormatFlags = 0x00100000
	DtPrefixonly             JouiTextFormatFlags = 0x00200000
)
