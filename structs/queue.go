package structs

import (
	"container/list"
	"sync"

	"github.com/laurence6/telegram-bot-api"
)

type messageWithID struct {
	senderID string
	message  *tgbotapi.Message
}

// MessageQueue is a thread safe queue for message.
//
// It makes sure that the messages from the same user will be returned in order by recording the senderID of the message that are being processed. It will not return the message from a user whose former message is being processed.
//
// When finishes processing a message, you must call Done().
type MessageQueue struct {
	MaxLength int

	*list.List
	processing map[string]bool

	*sync.Mutex
	cond *sync.Cond
}

// NewMessageQueue returns a MessageQueue instance with max length = 0.
func NewMessageQueue() *MessageQueue {
	return NewMessageQueueWithMaxLength(0)
}

// NewMessageQueueWithMaxLength returns a MessageQueue instance with specified max length.
func NewMessageQueueWithMaxLength(maxLength int) *MessageQueue {
	return &MessageQueue{
		maxLength,
		list.New(),
		map[string]bool{},
		&sync.Mutex{},
		sync.NewCond(&sync.Mutex{}),
	}
}

// Put puts a message into the queue.
//
// It requires a senderID.
func (queue *MessageQueue) Put(senderID string, message *tgbotapi.Message) {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	for queue.MaxLength != 0 && queue.Len() >= queue.MaxLength {
		queue.cond.Wait()
	}

	queue.Lock()
	queue.PushBack(&messageWithID{
		senderID: senderID,
		message:  message,
	})
	queue.Unlock()

	queue.cond.Broadcast()
}

// Get returns a message and puts the senderID into processing list to avoid returning the later messages from this user.
//
// Get will block if no message can be returned.
func (queue *MessageQueue) Get() *tgbotapi.Message {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	for {
		for queue.Len() == 0 {
			queue.cond.Wait()
		}

		queue.Lock()
		for i := queue.Front(); i != nil; i = i.Next() {
			message := i.Value.(*messageWithID)

			senderID := message.senderID
			if _, ok := queue.processing[senderID]; ok {
				continue
			}

			queue.Remove(i)
			queue.processing[senderID] = true

			queue.Unlock()

			queue.cond.Broadcast()
			return message.message
		}
		queue.Unlock()

		queue.cond.Wait()
	}
}

// Done removes the senderID from the processing list.
func (queue *MessageQueue) Done(senderID string) {
	queue.Lock()
	delete(queue.processing, senderID)
	queue.Unlock()

	queue.cond.Broadcast()
}
