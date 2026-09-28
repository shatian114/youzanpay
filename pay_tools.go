package youzan

type PayTool string

const (
	PayToolWxJs     PayTool = "WX_JS"
	PayToolAlipayJs PayTool = "ALIPAY_JS"
	PayToolWxApplet PayTool = "WX_APPLET"
)
