package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/CaoJiayuan/go-workflow/workflow-engine/model"
	"github.com/mumushuiding/util"
)

// UserInfo 用户信息
type UserInfo struct {
	Company string `json:"company"`
	// 用户所属部门
	Department string `json:"department"`
	Username   string `json:"username"`
	ID         string `json:"ID"`
	// 用户的角色
	Roles []string `json:"roles"`
	// 用户负责的部门
	Departments []string `json:"departments"`
	Token       string   `json:"token,omitempty"`
}

func (u UserInfo) GetToken() string {
	if u.Token != "" {
		return u.Token
	}

	h := md5.New()
	h.Write([]byte("user_token:" + u.ID))
	token := hex.EncodeToString(h.Sum(nil))
	return token
}

// GetUserinfoFromRedis GetUserinfoFromRedis
func GetUserinfoFromRedis(token string) (*UserInfo, error) {
	result, err := GetValFromRedis(token)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)
	var userinfo = &UserInfo{}
	err = util.Str2Struct(result, userinfo)
	if err != nil {
		return nil, err
	}
	return userinfo, nil
}

func SetUserInfoToRedis(u UserInfo) (string, error) {
	token := u.GetToken()
	j, e := json.Marshal(u)
	if e != nil {
		return "", e
	}
	e = model.RedisSetVal(token, string(j), -1)
	return token, e
}

// GetValFromRedis 从redis获取值
func GetValFromRedis(key string) (string, error) {
	return model.RedisGetVal(key)
}
