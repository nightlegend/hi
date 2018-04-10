package protocol

// CONST : define const
const (
	LoginRequest = 0x0001
	LoginSuccess = 0x0010
	LoginFailed  = 0x0100

	LogoutRequest = 0x0002
	LogoutSuccess = 0x0020

	UserMessageRequest = 0x0003
	UserMessageSuccess = 0x0030
	UserMessageFail    = 0x00300
)
