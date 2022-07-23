package util

import "testing"

// 编写测试案案例，测试addUpper是否正确
func TestAddupper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper(10) 执行错误，期望值=%v,实际值=%v\n", 55, res)
	}
	// 如果正确，输出日志
	t.Logf("addUpper(10) 执行正确...")
}

func TestGetSub(t *testing.T) {
	res := GetSub(10, 5)
	if res != 5 {
		t.Fatalf("GetSub(10,5) 执行错误，期望值=%v,实际值=%v\n", 5, res)
	}
	t.Logf("GetSub(10,5) 执行正确")

}

func TestFileUtil(t *testing.T) {
	// 调用
	mino()

	// 如果正确，输出日志
	t.Logf("TestFileUtil 执行正确...")
}
