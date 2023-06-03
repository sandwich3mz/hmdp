package impl

import (
	"context"
	"encoding/base64"
	"github.com/go-eden/routine"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"hmdp/dto"
	"hmdp/ent"
	"hmdp/ent/user"
	"hmdp/global"
	"hmdp/services/user_service"
	"hmdp/tools"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//var UserDTO dto.UserDTO

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) user_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) SendCode(ctx context.Context, phone string) {

	seed := time.Now().Unix()

	r := rand.New(rand.NewSource(seed))

	code := r.Intn(1000000)

	client := global.App.Redis

	client.Set(
		ctx,
		tools.LoginCodeKey+phone,
		code,
		tools.LoginCodeTtl*time.Minute,
	)

	log.Println("验证码:" + strconv.Itoa(code))

}

var nameVar = routine.NewLocalStorage()

func (p *pgImpl) Login(ctx context.Context, loginForm dto.LoginFormDTO) string {
	client := global.App.Redis
	phone := loginForm.Phone
	cacheCode, _ := client.Get(ctx, tools.LoginCodeKey+phone).Result()
	code := loginForm.Code
	if cacheCode == "" || code != cacheCode {
		//不一致：报错
		return "验证码错误"
	}

	// 一致,根据手机号查用户
	user, err := p.dbClient.User.
		Query().
		Where(user.PhoneEQ(phone)).
		Only(ctx)
	if err != nil {
		user = p.createUserWithPhone(ctx, phone)
	}

	// 保存用户信息到Redis
	// 随机生成token 作为登录令牌
	token := uuid.New().String()
	_ = copier.Copy(&global.UserDTO, &user)
	//将对象转为hash
	hash := make(map[string]interface{})
	_ = mapstructure.Decode(global.UserDTO, &hash)
	// 存储
	ctx = context.Background()
	tokenKey := tools.LoginUserKey + token
	client.Del(ctx, tokenKey)
	client.HMSet(
		ctx,
		tokenKey,
		hash,
	)
	client.Expire(ctx, tokenKey, tools.LoginUserTtl*time.Minute)
	return token
}

func (p *pgImpl) createUserWithPhone(ctx context.Context, phone string) *ent.User {

	randString, _ := generateRandomString(10)

	user, err := p.dbClient.User.
		Create().
		SetPhone(phone).
		SetNickName("user_" + randString).
		Save(ctx)
	if err != nil {
		return nil
	}
	return user
}

func generateRandomString(length int) (string, error) {
	// 计算需要生成的随机字节数组长度
	byteLength := (length*3 + 3) / 4 // 向上取整，保证最终输出长度为 length

	// 生成随机字节数组
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// 将字节数组编码成字符串
	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}
