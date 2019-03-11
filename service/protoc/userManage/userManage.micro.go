// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: userManage.proto

/*
Package userManage is a generated protocol buffer package.

It is generated from these files:
	userManage.proto

It has these top-level messages:
	GetUserListReq
	UserMesssage
	UserListRsp
	ChangeUserReq
	ChangeUserRsp
	AddUserReq
	AddUserRsp
	DeleteUserReq
	DeleteUserRsp
*/
package userManage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserManage service

type UserManageService interface {
	GetUserListByOffstAndLimit(ctx context.Context, in *GetUserListReq, opts ...client.CallOption) (*UserListRsp, error)
	UpdateUserById(ctx context.Context, in *ChangeUserReq, opts ...client.CallOption) (*ChangeUserRsp, error)
	AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRsp, error)
	DeleteUserById(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*DeleteUserRsp, error)
}

type userManageService struct {
	c    client.Client
	name string
}

func NewUserManageService(name string, c client.Client) UserManageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "usermanage"
	}
	return &userManageService{
		c:    c,
		name: name,
	}
}

func (c *userManageService) GetUserListByOffstAndLimit(ctx context.Context, in *GetUserListReq, opts ...client.CallOption) (*UserListRsp, error) {
	req := c.c.NewRequest(c.name, "UserManage.GetUserListByOffstAndLimit", in)
	out := new(UserListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManageService) UpdateUserById(ctx context.Context, in *ChangeUserReq, opts ...client.CallOption) (*ChangeUserRsp, error) {
	req := c.c.NewRequest(c.name, "UserManage.UpdateUserById", in)
	out := new(ChangeUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManageService) AddUser(ctx context.Context, in *AddUserReq, opts ...client.CallOption) (*AddUserRsp, error) {
	req := c.c.NewRequest(c.name, "UserManage.AddUser", in)
	out := new(AddUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManageService) DeleteUserById(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*DeleteUserRsp, error) {
	req := c.c.NewRequest(c.name, "UserManage.DeleteUserById", in)
	out := new(DeleteUserRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserManage service

type UserManageHandler interface {
	GetUserListByOffstAndLimit(context.Context, *GetUserListReq, *UserListRsp) error
	UpdateUserById(context.Context, *ChangeUserReq, *ChangeUserRsp) error
	AddUser(context.Context, *AddUserReq, *AddUserRsp) error
	DeleteUserById(context.Context, *DeleteUserReq, *DeleteUserRsp) error
}

func RegisterUserManageHandler(s server.Server, hdlr UserManageHandler, opts ...server.HandlerOption) error {
	type userManage interface {
		GetUserListByOffstAndLimit(ctx context.Context, in *GetUserListReq, out *UserListRsp) error
		UpdateUserById(ctx context.Context, in *ChangeUserReq, out *ChangeUserRsp) error
		AddUser(ctx context.Context, in *AddUserReq, out *AddUserRsp) error
		DeleteUserById(ctx context.Context, in *DeleteUserReq, out *DeleteUserRsp) error
	}
	type UserManage struct {
		userManage
	}
	h := &userManageHandler{hdlr}
	return s.Handle(s.NewHandler(&UserManage{h}, opts...))
}

type userManageHandler struct {
	UserManageHandler
}

func (h *userManageHandler) GetUserListByOffstAndLimit(ctx context.Context, in *GetUserListReq, out *UserListRsp) error {
	return h.UserManageHandler.GetUserListByOffstAndLimit(ctx, in, out)
}

func (h *userManageHandler) UpdateUserById(ctx context.Context, in *ChangeUserReq, out *ChangeUserRsp) error {
	return h.UserManageHandler.UpdateUserById(ctx, in, out)
}

func (h *userManageHandler) AddUser(ctx context.Context, in *AddUserReq, out *AddUserRsp) error {
	return h.UserManageHandler.AddUser(ctx, in, out)
}

func (h *userManageHandler) DeleteUserById(ctx context.Context, in *DeleteUserReq, out *DeleteUserRsp) error {
	return h.UserManageHandler.DeleteUserById(ctx, in, out)
}