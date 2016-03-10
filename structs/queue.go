package structs

import (
	"container/list"
	"sync"

	"github.com/laurence6/telegram-bot-api"

	"github.com/laurence6/telegrambot-go/utils"
)

/*MessageQueue A thread safe queue for tgbotapi.Message.
*
* It make sure that the messages from the same user will be returned in order.
*
* It records the messages that are being processed. It will not return the message from a user whose former message is being processed. When finishes processing a message, you must call Done().
 */
type MessageQueue struct {
	MaxLength int

	*list.List
	processing map[string]bool

	*sync.Mutex
	cond *sync.Cond
}

func NewMessageQueue() *MessageQueue {
	return NewMessageQueueWithMaxLength(0)
}

func NewMessageQueueWithMaxLength(maxLength int) *MessageQueue {
	return &MessageQueue{
		maxLength,
		list.New(),
		map[string]bool{},
		&sync.Mutex{},
		sync.NewCond(&sync.Mutex{}),
	}
}

/*Put Put a message in the queue
 */
func (queue *MessageQueue) Put(message *tgbotapi.Message) {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	for queue.MaxLength != 0 && queue.Len() >= queue.MaxLength {
		queue.cond.Wait()
	}

	queue.Lock()
	queue.PushBack(message)
	queue.Unlock()

	queue.cond.Broadcast()
}

/*Get Return a message
*
* Get will block if no message can be returned.
*
* It will put the user into processing list to avoid returning the later messages from this user.
 */
func (queue *MessageQueue) Get() *tgbotapi.Message {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	for {
		for queue.Len() == 0 {
			queue.cond.Wait()
		}

		queue.Lock()
		for i := queue.Front(); i != nil; i = i.Next() {
			message := i.Value.(*tgbotapi.Message)

			id := utils.GetMessageChatUserID(message)
			if _, ok := queue.processing[id]; ok {
				continue
			}

			queue.Remove(i)
			queue.processing[id] = true

			queue.Unlock()

			queue.cond.Broadcast()
			return message
		}
		queue.Unlock()

		queue.cond.Wait()
	}
}

/*Done Remove the user from the processing list
 */
func (queue *MessageQueue) Done(message *tgbotapi.Message) {
	id := utils.GetMessageChatUserID(message)

	queue.Lock()
	delete(queue.processing, id)
	queue.Unlock()

	queue.cond.Broadcast()
}
