package userservicelogic

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"math/rand"
	"mxshop/app/usercenter/model/cache"
	"mxshop/app/usercenter/model/mysql"
	"mxshop/common/xerr"
	"time"

	"mxshop/app/usercenter/cmd/rpc/internal/svc"
	"mxshop/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/pbkdf2"
)

const (
	saltMinLen = 8
	saltMaxLen = 32
	iter       = 1000
	keyLen     = 32
)

type SignOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignOutLogic {
	return &SignOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SignOut 用户注册
func (l *SignOutLogic) SignOut(in *pb.SignOutReq) (*pb.SignOutResp, error) {
	//1.验证电话是否认证过
	u, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.GetMobile())
	if err != nil && !errors.Is(err, mysql.ErrNotFound) {
		l.Logger.WithContext(l.ctx).Errorf("find user by mobile error, mobile:%s, err:%v", in.GetMobile(), err.Error())
		return nil, err
	}

	//号码已经注册
	if u != nil {
		return nil, xerr.NewErrCode(xerr.USER_AlREADY_REGISTERED)
	}

	//1.加密
	pwd, err := EncryptPwd(in.GetPassword())
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("encrypt pwd error, pwd:%s, err:%v", in.GetPassword(), err.Error())
		return nil, err
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &mysql.User{
		Uid:      l.getRandomUid(),
		Name:     in.GetName(),
		Password: pwd,
		Mobile:   in.GetMobile(),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("sign up error, err:%v", err.Error())
		return nil, err
	}
	return &pb.SignOutResp{}, nil
}

// EncryptPwd 加密密码
func EncryptPwd(pwd string) (encrypt string, err error) {
	// 1、生成随机长度的盐值
	salt, err := randSalt()
	if err != nil {
		return
	}

	// 2、生成加密串
	en := encryptPwdWithSalt([]byte(pwd), salt)
	en = append(en, salt...)

	// 3、合并盐值
	encrypt = base64.StdEncoding.EncodeToString(en)
	return
}

func randSalt() ([]byte, error) {
	// 生成8-32之间的随机数字
	salt := make([]byte, rand.Intn(saltMaxLen-saltMinLen)+saltMinLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func encryptPwdWithSalt(pwd, salt []byte) (pwdEn []byte) {
	pwd = append(pwd, salt...)
	pwdEn = pbkdf2.Key(pwd, salt, iter, keyLen, sha256.New)
	return
}

func (l *SignOutLogic) getRandomUid() int64 {
	//1.从redis中获取
	uid, err := l.svcCtx.RedisModel.IncrBy(cache.RandomUidKey, int64(rand.Intn(5)))
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("incr by random uid error, err:%v", err.Error())
		return 0
	}
	if uid > 10 {
		return uid
	}

	//取数据表中的最大uid
	initUid, err := l.svcCtx.UserModel.GetMaxUid()
	if err != nil && !errors.Is(err, mysql.ErrNotFound) {
		l.Logger.WithContext(l.ctx).Errorf("get max uid error, err:%v", err.Error())
		return 0
	}
	if initUid == 0 {
		initUid = 10000000
	}

	uid, err = l.svcCtx.RedisModel.IncrBy(cache.RandomUidKey, initUid)
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("incr by random uid error, err:%v", err.Error())
		return 0
	}
	return uid
}
