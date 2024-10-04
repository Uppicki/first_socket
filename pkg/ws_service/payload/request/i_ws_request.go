package wsservicerequests

import "first_socket/pkg/ws_service/ws_message"

type IWSRequest interface {
	ToMessage() (wsservicemessage.IWSMessage, error)
}
