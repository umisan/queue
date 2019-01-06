package main

import (
	"sync"
)

//キュー構造体
///////////////////////////////
/*
関数
Enqueue
Dequeue
Empty
Len
Front
Back
*/
type Queue struct {
	mutex sync.Mutex
	elems Tasks
}

//キューにデータを追加
func (q *Queue) Enqueue(task Task) {

}

//キューからデータを取り除く
//空の場合はエラーを返す
func (q *Queue) Dequeue() error {
	return nil
}

//キューが空の場合はtrue, それ以外はfalseを返す
func (q *Queue) Empty() bool {
	return true
}

//キューに含まれている要素数を返す
func (q *Queue) Len() int64 {
	return 0
}

//キューの先頭の要素を返す
func (q *Queue) Front() Task {
	return TestTask{}
}

//キューの末尾の要素を返す
func (q *Queue) Back() Task {
	return TestTask{}
}
