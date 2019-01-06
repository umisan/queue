package main

import "fmt"

//Taskインターフェイス
//処理を実行するExecute関数が必要
type Task interface {
	Execute() error
}

//Taskのスライス
type Tasks []Task

//テスト用の関数
type TestTask struct {
}

func (test TestTask) Execute() error {
	fmt.Println("%p", test)
	return nil
}
