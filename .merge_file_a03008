package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"encoding/json"
)

/**
 * string转换int
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func ParseInt(b string, defInt int) int {
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
func ParseString(b int) string {
	id := strconv.Itoa(b)
	return id
}

/**
 * md5 加密
 * @method MD5
 * @param  {[type]} data string [description]
 */
func MD5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//哈希加密
func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

/**
 * 结构体转成json 字符串
 * @method StruckToString
 * @param  {[type]}       data interface{} [description]
 */
func StructToString( data interface{}) (string) {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}
