package utils

const (
	Error           = iota // 内部错误
	Success                //成功
	Fail                   // 失败
	RegisterSuccess        //注册成功
	RegisterFail           //注册失败
	LoginSuccess           //登录成功
	LoginFail              //登录失败
	OrderSuccess           //下单成功
	OrderFail              //下单失败
	PurchaseSuccess        //购买成功
	PurchaseFail           //购买失败
	ForwardSuccess         //转发成功
	ForwardFail            //转发失败
	TokenSuccess           //Token正确
	TokenFail              //Token错误
)
