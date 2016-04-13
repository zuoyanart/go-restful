package tools

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type Tools struct {
}


var (
	t    *Tools
	once sync.Once
)
/**
 * 返回单例实例
 * @method New
 */
func New() *Tools {
	once.Do(func() {//只执行一次
		t = &Tools{}
	})
	return t
}


/**
 * string转换int
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func (t *Tools) ParseInt(b string, defInt int) int {
	id, err := strconv.Atoi(b)
	if err != nil {
		return defInt
	} else {
		return id
	}
}

/**
 * int转换string
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func (t *Tools) ParseString(b int) string {
	id := strconv.Itoa(b)
	return id
}

/**
 * md5 加密
 * @method MD5
 * @param  {[type]} data string [description]
 */
func (t *Tools) MD5(data string) string {
	_m := md5.New()
	io.WriteString(_m, data)
	return fmt.Sprintf("%x", _m.Sum(nil))
}

/**
 * 结构体转成json 字符串
 * @method StruckToString
 * @param  {[type]}       data interface{} [description]
 */
func (t *Tools) StructToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}

/**
 * 生成随机字符串,n是字符串长度
 * @method func
 * @param  {[type]} t *Tools        [description]
 * @return {[type]}   [description]
 */
func (t *Tools) GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()+[]{}/<>;:=.,?"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

/**
 * 控制台打印测试
 * @method log
 * @param  {[type]} s string        [description]
 * @return {[type]}   [description]
 */
func (t *Tools) Logs(s string) {
	log.Printf(s)
}
