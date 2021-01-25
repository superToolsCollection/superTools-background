package errcode

/**
* @Author: super
* @Date: 2020-11-24 18:45
* @Description:
**/

var (
	ErrorUserSignInFail   = NewError(20060001, "用户登录失败")
	ErrorUserUpdateFail   = NewError(20060002, "用户信息更新失败")
	ErrorUserRegisterFail = NewError(20060003, "用户注册失败")
	ErrorUserListFail     = NewError(20060004, "获取用户列表失败")
	ErrorAddUser          = NewError(20060005, "添加用户失败")
)
