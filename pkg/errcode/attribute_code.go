package errcode

/**
* @Author: super
* @Date: 2021-02-14 09:51
* @Description:
**/

var (
	ErrorGetAttributeFail        = NewError(50010001, "获取分类参数列表失败")
	ErrorAddAttributeFail        = NewError(50010002, "添加分类参数失败")
	ErrorDeleteAttributeFail     = NewError(50010003, "删除分类参数失败")
	ErrorGetAttributeByIdFail    = NewError(50010004, "根据ID获取分类参数失败")
	ErrorUpdateAttributeByIdFail = NewError(50010005, "更新分类参数失败")
)
