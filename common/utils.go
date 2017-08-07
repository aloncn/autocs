package common
//常用方法
import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"regexp"
	"strings"
	mr "math/rand"
	"path/filepath"
	"os"
)

//md5方法
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Guid方法
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//判断是否包含
func SliceContains(src []string,value string)bool{
	isContain := false
	for _,srcValue := range src  {
		if(srcValue == value){
			isContain = true
			break
		}
	}
	return isContain
}

//判断key是否存在
func MapContains(src map[string]int ,key string) bool{
	if _, ok := src[key]; ok {
		return true
	}
	return false
}

func RemoveDuplicate(list *[]int) []int {
	var x []int = []int{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

func CheckRepeat(list []int,id int) bool  {
	if len(list) == 0 {
		return false
	}
	for _, v := range list{
		if v == id{
			return true
		}
	}
	return false
}


//去除html标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	//去除 &nbsp
	re, _ = regexp.Compile("\\&nbsp;")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}


//从数组中随机取一个
func RandGetArray(a []string) string {
	len := len(a)
	i := mr.Intn(len)
	return a[i]
}


//当前项目根目录
var APP_ROOT string
// 获取项目路径
func GetPath() string {

	if APP_ROOT != "" {
		return APP_ROOT
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		print(err.Error())
	}

	APP_ROOT = strings.Replace(dir, "\\", "/", -1)
	return APP_ROOT
}