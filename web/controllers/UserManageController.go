package controllers

import (
	"github.com/astaxie/beego"
	proto "service/protoc/userManage"
	"web/common"
)

type UserManageController struct {
	beego.Controller
}

func (this *UserManageController) UserManageInit() {
	this.TplName = "manage/user_manage.html"
}

func (this *UserManageController) UserManage() {
	offset, _ := this.GetInt32("offset")
	limit, _ := this.GetInt32("limit")
	//获取用户信息
	userSession := this.GetSession("user_id")
	if userSession == nil { //未登陆
		this.Ctx.Redirect(304, "/index")
		return
	}
	userId := userSession.(int64)

	//调用服务
	userManage,ctx := common.InitUserManage(this.CruSession)
	//call the userManage method
	req := proto.GetUserListReq{Offset: offset, Limit: limit, ManageId: int64(userId)}
	rsp, err := userManage.GetUserListByOffstAndLimit(ctx, &req)

	//user_data,page_num
	beego.Info("======user_list=====", rsp.UserList, "-----------------", err)
	var result map[string]interface{}
	result = make(map[string]interface{})
	result["user_data"] = rsp.UserList
	result["page_num"] = offset
	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *UserManageController) ChangeUser() {
	changeId, _ := this.GetInt64("change_id")
	userName := this.GetString("user_name")
	loginName := this.GetString("login_name")
	userPhoneNumber := this.GetString("user_phone_number")
	userJobNumber := this.GetString("user_job_number")
	userGender, _ := this.GetInt32("user_gender")

	//call the userManage method
	userManage,ctx := common.InitUserManage(this.CruSession)
	req := proto.ChangeUserReq{ChangeId: changeId, Name: userName, LoginName: loginName, PhoneNumber: userPhoneNumber, JobNumber: userJobNumber, Gender: userGender}
	rsp, err := userManage.UpdateUserById(ctx, &req)
	if err != nil {
		beego.Info("======ChangeUser=====", rsp.UserId, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return
}

func (this *UserManageController) AddUser() {
	userName := this.GetString("user_name")
	loginName := this.GetString("login_name")
	userPhoneNumber := this.GetString("user_phone_number")
	userJobNumber := this.GetString("user_job_number")
	userGender, _ := this.GetInt32("user_gender")

	//call the userManage method
	userManage,ctx := common.InitUserManage(this.CruSession)
	req := proto.AddUserReq{Name: userName, LoginName: loginName, PhoneNumber: userPhoneNumber, JobNumber: userJobNumber, Gender: userGender}
	rsp, err := userManage.AddUser(ctx, &req)
	if err != nil {
		beego.Info("======AddUser=====", rsp.UserId, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return

}

func (this *UserManageController) DeleteUserById() {
	deleteId, _ := this.GetInt64("delete_id")

	//call the userManage method
	userManage,ctx := common.InitUserManage(this.CruSession)
	req := proto.DeleteUserReq{DeleteId: deleteId}
	rsp, err := userManage.DeleteUserById(ctx, &req)
	if err != nil {
		beego.Info("======DeleteUserById=====", rsp.UserId, "-------err--------", err)
	}

	var result map[string]interface{}
	result = make(map[string]interface{})
	result["result"] = rsp.Message
	this.Data["json"] = result
	this.ServeJSON()
	return
}
