package errorpkg

import "log"

// CheckErrors handle the err message
func CheckErrors(err error) {
	if err != nil {
		log.Println("Server occur error:", err)
	}
}
