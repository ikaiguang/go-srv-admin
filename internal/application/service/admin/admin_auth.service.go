package adminsrv

import (
	"context"
	baseerror "github.com/ikaiguang/go-srv-kit/api/base/error"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"

	adminerrorv1 "github.com/ikaiguang/go-srv-admin/api/admin/v1/errors"
	resources "github.com/ikaiguang/go-srv-admin/api/admin/v1/resources"
	adminservicev1 "github.com/ikaiguang/go-srv-admin/api/admin/v1/services"
	assemblers "github.com/ikaiguang/go-srv-admin/internal/application/assembler/admin"
	repos "github.com/ikaiguang/go-srv-admin/internal/domain/repo/admin"
)

// adminAuth ...
type adminAuth struct {
	adminservicev1.UnimplementedSrvAdminAuthServer

	authConfig *confv1.App_Auth

	adminRepo         repos.AdminRepo
	adminRegEmailRepo repos.AdminRegEmailRepo
}

// NewAdminAuthService ...
func NewAdminAuthService(
	authConfig *confv1.App_Auth,
	adminRepo repos.AdminRepo,
	adminRegEmailRepo repos.AdminRegEmailRepo,
) adminservicev1.SrvAdminAuthServer {
	return &adminAuth{
		authConfig: authConfig,

		adminRepo:         adminRepo,
		adminRegEmailRepo: adminRegEmailRepo,
	}
}

// LoginByEmail Email登录
func (s *adminAuth) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq) (out *resources.LoginResp, err error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.adminRegEmailRepo.QueryOneByAdminEmail(ctx, in.Email)
	if err != nil {
		reason := baseerror.ERROR_STATUS_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// admin
	adminModel, isNotFound, err := s.adminRepo.QueryOneByAdminUuid(ctx, regEmailModel.AdminUuid)
	if err != nil {
		reason := baseerror.ERROR_STATUS_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// 验证密码
	err = passwordutil.Compare(adminModel.PasswordHash, in.Password)
	if err != nil {
		reason := adminerrorv1.ERROR_ADMIN_PASSWORD_INCORRECT.String()
		message := "密码不正确"
		err = errorutil.BadRequest(reason, message)
		//err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// generate token
	//claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtpkg.Claims{
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		ExpiresAt: jwtpkg.GenExpireTime(),
	//	},
	//	AuthId:   adminModel.Id,
	//	AuthUuid: adminModel.AdminUuid,
	//})
	//signedString, err := claims.SignedString([]byte(s.authConfig.AdminKey))
	out = &resources.LoginResp{
		AdminInfo: &resources.Info{
			Id:            adminModel.Id,
			AdminNickname: adminModel.AdminNickname,
			AdminAvatar:   adminModel.AdminAvatar,
			AdminGender:   assemblers.ToAdminGenderEnum(adminModel.AdminGender),
			AdminStatus:   assemblers.ToAdminStatusEnum(adminModel.AdminStatus),
		},
		//Token: signedString,
	}
	return out, err
}

// Ping ping pong
func (s *adminAuth) Ping(ctx context.Context, in *resources.PingReq) (out *resources.PingResp, err error) {
	//authInfo, err := authpkg.GetAdminFromContext(ctx)
	//if err != nil {
	//	return out, err
	//}
	//logutil.InfowWithContext(ctx, "authInfo", authInfo)
	out = &resources.PingResp{
		Message: in.Message,
	}
	return out, err
}
