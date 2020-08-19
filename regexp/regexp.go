package main

import (
	"fmt"
	"regexp"
)

func Match() {
	fmt.Printf("regexp.Match 判断一个 byte 数组中是否包含指定 pattern；该 pattern 可以是一个正则表达式或者字符串\n")
	str := "asdf123123"
	pattern := "[0-9]+"
	ans, _ := regexp.Match(pattern, []byte(str))
	fmt.Printf("regexp.Match: %t\n", ans)
	pattern = "123"
	ans, _ = regexp.Match(pattern, []byte(str))
	fmt.Printf("regexp.Match: %t\n", ans)
	ans, _ = regexp.MatchString(pattern, str)
	fmt.Printf("regexp.MatchString: %t\n", ans)
}

func Regexp() {
	fmt.Printf("regexp.Compile 将一个字符串编译成一个 regexp 的 struct\n")
	str := "123, hello"
	pattern := "[0-9]+"
	regex, _ := regexp.Compile(pattern)
	fmt.Printf("regexp.Compile: %+v\n", regex)
	regex = regexp.MustCompile(pattern)
	fmt.Printf("regexp.MustCompile: %+v\n", regex)
	posix := regexp.MustCompilePOSIX(pattern)
	fmt.Printf("regexp.MustCompilePOSIX: %+v\n", posix)
	fmt.Printf("regexp.String: %+v\n", regex.String())
	fmt.Printf("regexp.NumSubexp: %d\n", regex.NumSubexp())
	fmt.Printf("regexp.SubexpNames: %+v\n", regex.SubexpNames())
	fmt.Printf("regexp.MatchString: %t\n", regex.MatchString(str))
	fmt.Printf("regexp.ReplaceAllString: %s\n", regex.ReplaceAllString(str, "world"))
	a := func(s string) string {
		res := []byte{}
		for i := 0; i < len(s); i++ {
			res = append(res, 'a')
		}
		return string(res)
	}
	fmt.Printf("regexp.ReplaceAllStringFunc: %s\n", regex.ReplaceAllStringFunc(str, a))
	fmt.Printf("regexp.ReplaceAll: %+v\n", regex.ReplaceAll([]byte(str), []byte("world")))
	fmt.Printf("regexp.QuoteMeta: %s\n", regexp.QuoteMeta(pattern))
	fmt.Printf("regexp.Find: %s\n", string(regex.Find([]byte(str))))
	fmt.Printf("regexp.FindIndex: %+v\n", regex.FindIndex([]byte(str)))
	fmt.Printf("regexp.FindString: %s\n", string(regex.FindString(str)))
	fmt.Printf("regexp.FindStringIndex: %+v\n", regex.FindStringIndex(str))
	pattern = "[0-9]{2,2}([0-9])"
	regex = regexp.MustCompile(pattern)
	fmt.Printf("regexp.FindAllString: %+v\n", regex.FindAllString(str, -1))
	fmt.Printf("regexp.FindAllSubMatch: %+v\n", regex.FindAllStringSubmatch(str, -1))
	pattern = ","
	regex = regexp.MustCompile(pattern)
	fmt.Printf("regexp.Split: %+v\n", regex.Split(str, -1))
}

func main() {
	Match()
	Regexp()
}
