package errcode

/**
* @Author: super
* @Date: 2021-02-07 20:42
* @Description:
**/

var (
	ErrorGetRoleListFail = NewError(30010001, "获取角色列表失败")
	ErrorAddRoleFail     = NewError(30010002, "添加角色失败")
	ErrorGetRoleByIdFail = NewError(30010003, "获取角色失败")
	ErrorUpdateRoleFail  = NewError(30010004, "更新角色失败")
	ErrorDeleteRoleFail  = NewError(30010005, "删除角色失败")
)
