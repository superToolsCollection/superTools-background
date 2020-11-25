package errcode

/**
* @Author: super
* @Date: 2020-11-25 22:44
* @Description:
**/

var (
	ErrorAddToolFail        = NewError(20070001, "添加工具失败")
	ErrorUpdateToolInfoFail = NewError(20070002, "更新工具信息失败")
	ErrorDeleteToolFail     = NewError(20070003, "删除工具失败")
	ErrorToolOnlineFail     = NewError(20070004, "工具上线失败")
	ErrorToolOffLineFail    = NewError(20070005, "工具下线失败失败")
	ErrorGetToolByKeyFail   = NewError(20070006, "根据ID获取工具信息失败")
	ErrorGetToolByNameFail  = NewError(20070007, "根据名称获取工具失败")
	ErrorGetToolListFail    = NewError(20070008, "获取工具列表失败")
)
