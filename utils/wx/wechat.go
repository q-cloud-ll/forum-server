package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"forum-server/dao/redis"
	"forum-server/global"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"
)

func GetAccess() (string, error) {
	token, err := global.GVA_WX.GetAccessToken()
	if err != nil {
		return token, err
	}
	err = redis.FrmSetAccessToken(token)
	if err != nil {
		return token, err
	}
	return token, nil
}

// SaveWxHead 处理保存微信头像
func SaveWxHead(imgUrl string) (string, error) {
	//imgUrl := "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIomz1vPJX1xeY2pMaOvXQ6ticGJfQWaJw6wjoiaicYoIjwAOg2vFvhMdOianQ7A4OxicJ8Ml76N2an8Nw/132"
	//获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)
	guid := uuid.New().String()
	fileName := guid + ".png"
	singleFile := "static/uploadfile/" + fileName
	err = ioutil.WriteFile(singleFile, data, 0666) //buffer输出文件中（不做处理，直接写到文件）
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
