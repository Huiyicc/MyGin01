package common

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse_str(str string) map[string]interface{} {
	str_array := strings.Split(str,"&")
	var strMap = make(map[string]interface{})
	for _ ,value := range str_array {
		if value != "=" && value != "" {
			if strStr(value, "=") > 0 {
				//截取字符串 value[x:x]
				//获取 = 所在string的位置 strings.Index(value, "=")
				strMap[value[0:strings.Index(value, "=")]] = value[strings.Index(value, "=") + 1 :len([]rune(value))]
			} else {
				strMap[value] = ""
			}
		}

	}
	return strMap
}

func strStr(haystack string, needle string) int {
	//特殊情况，当needle是空字符串，那就返回0
	if len(needle) == 0 {
		return 0
	}
	//整段判断是否相同
	//判断的时候需要注意数组越界，所以小于haystack长度减去needle的长度
	//判断的时候也需要注意两个长度是一样的
	nlen := len(needle)
	for i := 0; i <= len(haystack)-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

//表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

//表情转换
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
