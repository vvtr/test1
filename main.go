package main

import (
	"fmt"
	"strconv"
	"strings"
	"test/rwmutex"
	"time"
)

func setarr(arr *[3]int) {
	(*arr)[0] = 9
}
func main() {
	str1 := "golang你好"
	for i, v := range str1 {
		fmt.Printf("%d - %c\n", i, v)
	}
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}
	fmt.Println()
	n, _ := strconv.Atoi("55")
	fmt.Printf("%T\n", n)
	n1, _ := strconv.ParseInt("66", 10, 64)
	fmt.Printf("%T\n", n1)
	str2 := strconv.FormatInt(33, 10)
	fmt.Printf("%T\n", str2)
	str3 := strconv.Itoa(22)
	fmt.Printf("%T\n", str3)
	fmt.Println(strings.Contains("abcdef", "ab")) //ac不是字串
	fmt.Println(strings.Count("abcdbcebc", "bc"))
	fmt.Println(strings.EqualFold("helL", "HELL")) //不区分大小写
	fmt.Println(strings.Index("abcded", "a"))
	fmt.Println(strings.Replace("goigoLgoss", "go", "HA", -1)) //n=-1表示全部替换，替换两个n就是2
	arr := strings.Split("a-b-c-de-ff", "-")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(strings.ToLower("HAHA"))
	fmt.Println(strings.ToUpper("uuius"))
	fmt.Println(strings.TrimSpace("   de "))
	fmt.Println(strings.Trim("---ddd--", "-")) //ddd 多少个-都会被去除
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println(now)
	//数组的初始化
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{4, 5, 6}
	var arr3 = [...]int{7, 8, 9}
	var arr4 = [...]int{2: 3, 5: 6}
	setarr(&arr1)
	fmt.Print(arr1, arr2, arr3, arr4)
	fmt.Println()
	a := make(map[int]string, 10)
	a[1] = "a"
	a[9] = "z"
	fmt.Println(a)
	var cnt rwmutex.Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				cnt.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for {
		cnt.Incr()
		time.Sleep(time.Second)
		if cnt.Count() == 5 {
			break
		}
	}

}
