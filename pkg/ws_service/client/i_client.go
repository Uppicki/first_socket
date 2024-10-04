package wsserviceclient

import wsservicemessage "first_socket/pkg/ws_service/ws_message"

type IWSClient interface {
	Run()
	GetReceivedChan() <-chan wsservicemessage.IWSMessage
	Send(wsservicemessage.IWSMessage)
}
