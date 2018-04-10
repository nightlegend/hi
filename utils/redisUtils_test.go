package utils

import (
	"fmt"
	"testing"
)

var cli = NewCli()

func TestSet(t *testing.T) {
	defer Close(cli)
	result := Set(cli, "name", "david")
	if !result {
		t.Fatal("set new k-v failed.")
	}
}

func TestGet(t *testing.T) {
	defer Close(cli)
	result := Get(cli, "name")
	t.Log(result)
}

func TestLPush(t *testing.T) {
	defer Close(cli)
	result := LPush(cli, "david", "c")
	if !result {
		t.Fatal(result)
	}
}

func TestLRange(t *testing.T) {
	defer Close(cli)
	result := LRange(cli, "david")
	fmt.Println(result)
	for _, val := range result {
		fmt.Println(val)
	}
}
