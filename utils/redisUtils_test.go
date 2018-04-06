package utils

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	result := Set("name", "david")
	if !result {
		t.Fatal("set new k-v failed.")
	}
}

func TestGet(t *testing.T) {
	result := Get("name")
	t.Log(result)
}

func TestLPush(t *testing.T) {
	result := LPush("david", "c")
	if !result {
		t.Fatal(result)
	}
}

func TestLRange(t *testing.T) {
	result := LRange("david")
	fmt.Println(result)
	for _, val := range result {
		fmt.Println(val)
	}
}
