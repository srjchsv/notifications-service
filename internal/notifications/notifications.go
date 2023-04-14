package notifications

import "sync"

type Notification struct {
	CreatedAt      string `json:"created_at"`
	SenderUsername string `json:"sender_username"`
	Content        string `json:"content"`
}

type NotificationService struct {
	mu            sync.Mutex
	notifications []*Notification
}

func New() *NotificationService {
	return &NotificationService{
		notifications: make([]*Notification, 0),
	}
}

func (s *NotificationService) Add(notification *Notification) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.notifications = append(s.notifications, notification)
}

func (s *NotificationService) GetAll() []*Notification {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.notifications
}

func (s *NotificationService) GetLast() *Notification {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.notifications) == 0 {
		return nil
	}

	return s.notifications[len(s.notifications)-1]
}
