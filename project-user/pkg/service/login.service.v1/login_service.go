package login_service_v1

import (
	"context"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1.获取参数
	mobile := msg.Mobile
	// 2.校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 3.生成验证码
	code := "123456"
	// 4.调用短信平台(三方，放入go协程中执行，接口可以快速响应)
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信")

		// 5.存储验证码 redis中 过期时间15min
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("验证码存入Redis出错，cause by : %v\n", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, code)
	}()
	return &CaptchaResponse{Code: code}, nil
}
