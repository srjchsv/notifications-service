package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/srjchsv/notifications-service/internal/notifications"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeConnection(w http.ResponseWriter, r *http.Request, notifService *notifications.NotificationService) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Get the current length of the notifications slice
	currentNotifCount := len(notifService.GetAll())

	// Continuously check for new notifications
	for {
		// Get the current list of notifications
		allNotifs := notifService.GetAll()

		// If a new notification has arrived
		if len(allNotifs) > currentNotifCount {
			// Send the new notification
			notif := allNotifs[len(allNotifs)-1]
			err := conn.WriteJSON(notif)
			if err != nil {
				log.Printf("Error sending notification: %v", err)
				break
			}

			// Update the current notification count
			currentNotifCount = len(allNotifs)
		}

		// Sleep for a short duration before checking again
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
