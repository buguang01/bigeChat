package ChatModels

import (
	"errors"
	"time"
)

const (
	defaultQueueSize = 30
)

type ChatQueue struct {
	front        int
	rear         int
	currentCount int
	queueSize    int
	curruid      int //当前标识ID
	elements     []*ChatMessage
}

func (queue *ChatQueue) GetQueueLen() int {
	return queue.queueSize
}

/**
  指定大小的初始化
*/
func NewQueueBySize(size int) *ChatQueue {
	return &ChatQueue{0, size - 1, 0, size, 0, make([]*ChatMessage, size)}
}

/**
  按默认大小进行初始化
*/
func NewQueue() *ChatQueue {
	return NewQueueBySize(defaultQueueSize)
}

/**
  向下一个位置做探测
*/
func (queue *ChatQueue) ProbeNext(i int) int {
	return (i + 1) % queue.queueSize
}

func (queue *ChatQueue) Probe(i int) int {
	return i % queue.queueSize
}

/**
  清空队列
*/
func (queue *ChatQueue) ClearQueue() {
	queue.front = 0
	queue.rear = queue.queueSize - 1
	queue.currentCount = 0
}

/**
  是否为空队列
*/
func (queue *ChatQueue) IsEmpty() bool {
	if queue.currentCount == 0 {
		return true
	}
	return false
}

/**
  队列是否满了
*/
func (queue *ChatQueue) IsFull() bool {
	if queue.currentCount == queue.queueSize {
		return true
	}
	return false
}

/**
  出队一个元素
*/
func (queue *ChatQueue) Poll() (*ChatMessage, error) {
	if queue.IsEmpty() == true {
		return nil, errors.New("the queue is empty.")
	}
	tmp := queue.front
	queue.front = queue.ProbeNext(queue.front)
	queue.currentCount = queue.currentCount - 1
	return queue.elements[tmp], nil
}

//返回列表,需要成员继承IQueueUID接口
func (queue *ChatQueue) GetArray(uidmax int) ([]*ChatMessage, error) {
	result := make([]*ChatMessage, 0, queue.queueSize)
	for index := 0; index < queue.currentCount; index++ {
		tmp := queue.elements[queue.Probe(queue.front+index)]

		if tmp.UID > uidmax {
			result = append(result, tmp)
		}
	}

	return result, nil
}

/**
  入队
*/
func (queue *ChatQueue) Offer(e *ChatMessage) error {
	if queue.IsFull() == true {
		return errors.New("the queue is full.")
	}
	queue.curruid++
	e.UID = queue.curruid
	queue.rear = queue.ProbeNext(queue.rear)
	queue.elements[queue.rear] = e
	queue.currentCount = queue.currentCount + 1
	return nil
}

//拿最新的一个成员，但是不出列
func (queue *ChatQueue) GetEndNode() (*ChatMessage, error) {
	if queue.IsEmpty() == true {
		return nil, errors.New("the queue is empty.")
	}
	tmp := queue.front
	return queue.elements[tmp], nil
}

type ChatMessage struct {
	UID        int       //主键
	MemberID   int       //用户ID
	UserInfo   string    //用户信息
	ChatNode   string    //聊天内容
	CreateTime time.Time //时间

}
