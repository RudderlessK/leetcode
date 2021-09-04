package main

import (
	"PracticeProject/tryit"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

func run(fn func()) {
	// runtime.FuncForPC(reflect.ValueOf(action).Pointer()).Name(), err)
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
}

func fuck() {
	defer func() {
		cost := timeCost()
		fmt.Println("cost: ", cost)
	}()
	defer func() {
		time.Sleep(2 * time.Second)
		fmt.Println("???")
	}()
	fmt.Println("fuck")
	time.Sleep(1 * time.Second)
}

func timeCost() time.Duration {
	start := time.Now()
	return time.Since(start)
}

type A struct {
	b *B
}

func (a *A) B(key string) *B {
	if a.b == nil || a.b.key != key {
		fmt.Println("not equal")
		a.b = &B{
			key: key,
		}
	}
	return a.b
}

type B struct {
	key string
	a   *A
}

func (b *B) print() {
	fmt.Println(b.key)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//run(fuck)
	//fuck()
	//a := A{}
	//fmt.Printf("%p\n", &a)
	//b := a.B("asd")
	//c := a.B("asda")
	//fmt.Printf("%p\n", &a)
	//d := a.B("asda")
	//fmt.Printf("%p\n", &a)
	////b.print()
	//c.print()
	//d.print()
	//a.b.print()
	//
	//fmt.Printf("%p\n", b)
	//fmt.Printf("%p\n", c)
	//fmt.Printf("%p\n", d)
	//fmt.Printf("%p\n", a.b)
	//
	//defer func() {
	//	fmt.Printf("jjj- %p\n", &a)
	//	fmt.Printf("%p\n", c)
	//}()
	tryit.Run()
	fmt.Println(FuncName(timeCost))
	fmt.Println(reflect.TypeOf(&B{}).String())

	errs := make(map[string]error, 3)
	fmt.Println(errs)

	unMarshal()

	cfg := &GatewayConfig{
		Stage:     "1",
		SecretID:  1,
		SecretKey: "ad",
	}

	fmt.Println(structToMap(cfg))

	fmt.Println(ToString(GatewayConfig{
		Stage:     "2",
		SecretID:  2,
		SecretKey: "add",
	}))

	//tryit.FullGoroutine()
	//tryit.PtrReferGoroutine()

	fmt.Println(len(strings.Split("H W", " ")))

	fmt.Println(fmt.Sprintf("%%%s%%", "d"))

	a := "asdasd"
	fmt.Println(a[0:0])
	for _, c := range []rune(a) {
		fmt.Printf("%c", c)
	}



	c := []string{"abc", "absa", "aabe"}
	sort.Strings(c)
	fmt.Println(c)

	b := make([]string, len(c))
	b = append(b, c...)
	fmt.Printf("%p %p", c, b)
}

// StageValue 大同流程单stage变更消息内容(仅包含dwd和ods需要关注的字段信息)
type StageValue struct {
	ID          string   `json:"id"`
	FlowID      string   `json:"flowId"`
	AppID       string   `json:"appId"`
	Version     string   `json:"version"`
	PreStatus   string   `json:"preStatus"`
	PostStatus  string   `json:"postStatus"`
	Operation   string   `json:"operation"`
	Operator    string   `json:"operator"`
	Handlers    []string `json:"handlers"`
	StoryID     string   `json:"storyId"`
	WorkspaceID string   `json:"workspaceId"`
}

func unMarshal() {
	str := "{\"test\":[{\"test1\":1,\"test2\":\"gfds\"},{\"test2\":1,\"test3\":\"gfds\"}]}"
	type A struct {
		Test []interface{} `json:"test"`
	}
	var a A
	_ = json.Unmarshal([]byte(str), &a)
	fmt.Println(a.Test)
}

func FuncName(x func() time.Duration) string {
	return runtime.FuncForPC(reflect.ValueOf(x).Pointer()).Name()
}

type GatewayConfig struct {
	Stage     string
	SecretID  int
	SecretKey string
}

func structToMap(s interface{}) (map[string]string, error) {
	//if reflect.TypeOf(s).Elem().Kind() != reflect.Struct {
	//	return nil, fmt.Errorf("%v is not struct", reflect.TypeOf(s).Kind())
	//}
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var out map[string]string
	err = json.Unmarshal(data, &out)

	return out, err
}

func ToString(s interface{}) (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
