package errcode

/**
* @Author: super
* @Date: 2021-02-08 16:20
* @Description:
**/

var (
	ErrorGetCategoriesListFail = NewError(40010001, "获取分类列表失败")
	ErrorAddCategoryFail = NewError(40010002, "添加分类失败")
)
