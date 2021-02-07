package errcode

/**
* @Author: super
* @Date: 2021-02-07 20:42
* @Description:
**/

var (
	ErrorGetRoleListFail = NewError(30010001, "获取角色列表失败")
	ErrorAddRoleFail = NewError(30010002, "添加角色失败")
)
