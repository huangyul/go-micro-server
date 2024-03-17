package handler

import (
	"context"
	"go-micro-server/user_srv/global"
	"go-micro-server/user_srv/model"
	"go-micro-server/user_srv/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

func modelToUserResponse(user model.User) *proto.UserInfoResponse {
	res := &proto.UserInfoResponse{
		Id:       user.ID,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		res.Birthday = uint64(user.Birthday.Unix())
	}
	return res
}

// GetUserList 获取用户列表
func (u *UserServer) GetUserList(ctx context.Context, request *proto.ListRequest) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(model.User{})
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = uint32(result.RowsAffected)

	global.DB.WithContext(ctx).Find(&users).Limit(int(request.PageSize)).Offset((int(request.PageIndex) - 1) * int(request.PageSize))

	for _, user := range users {
		rsp.Data = append(rsp.Data, modelToUserResponse(user))
	}

	return rsp, nil

}

// GetUserByMobile 通过手机号获取用户信息
func (u *UserServer) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.WithContext(ctx).Where(&model.User{Mobile: request.Mobile}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	userInfo := modelToUserResponse(user)
	return userInfo, nil
}

// GetUserById 通过id获取用户信息
func (u *UserServer) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	res := global.DB.WithContext(ctx).First(&user, request.Id)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return modelToUserResponse(user), nil
}

// CreateUser 创建用户接口
func (u *UserServer) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.WithContext(ctx).Where(&model.User{Mobile: request.Mobile}).First(&user)
	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "手机号已存在")
	}
	user.Mobile = request.Mobile
	user.NickName = request.NickName

	// 密码加密
	hashPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "系统错误")
	}
	user.Password = string(hashPass)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return modelToUserResponse(user), nil
}

// UpdateUser 更新用户信息
func (u *UserServer) UpdateUser(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateResponse, error) {
	var user model.User
	result := global.DB.WithContext(ctx).First(&user, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	user.NickName = request.NickName
	user.Gender = request.Gender
	birthday := time.Unix(int64(request.Birthday), 0)
	user.Birthday = &birthday

	result = global.DB.Save(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &proto.UpdateResponse{
		Success: true,
	}, nil
}

// CheckPassword 校验密码
func (u *UserServer) CheckPassword(ctx context.Context, request *proto.PasswordCheckRequest) (*proto.CheckResponse, error) {
	err := bcrypt.CompareHashAndPassword([]byte(request.EncryptedPassword), []byte(request.Password))
	var res *proto.CheckResponse
	if err != nil {
		res.Success = false
		return res, nil
	}
	res.Success = true
	return res, nil
}
