package snowflake

import (
	"fmt"
	"testing"
)

func TestGenID(t *testing.T) {
	// 初始化 Snowflake 配置，设定纪元时间为 "2022-07-01"，机器 ID 为 1
	if err := Init("2022-07-01 00:00:00", 1); err != nil {
		t.Fatalf("Initialization failed: %v", err) // 使用 t.Fatalf 报告错误并停止测试
	}

	// 测试生成多个 ID 是否唯一
	id1, _ := GenID()
	id2, _ := GenID()

	// 输出生成的 ID
	fmt.Println("Generated ID 1:", id1)
	fmt.Println("Generated ID 2:", id2)

	// 验证生成的 ID 是否唯一
	if id1 == id2 {
		t.Errorf("Expected different IDs, but got the same: %v", id1) // 生成的 ID 应该是唯一的
	}
}
