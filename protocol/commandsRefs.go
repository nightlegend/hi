package protocol

// CONST : define const
const (
	LoginRequest = 0x0001
	LoginSuccess = 0x0010
	LoginFailed  = 0x0100

	UserMessageRequest = 0x0002
	UserMessageSuccess = 0x0020
	UserMessageFail    = 0x00200
)
