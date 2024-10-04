package wsservicemessage

import "first_socket/pkg/ws_service/payload/response"

type IWSMessage interface {
	Map()
	ToResponse() (wsservicereponses.IWSResponse, error)
}
