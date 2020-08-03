package errorhandle

import "log"

// RecoverServiceError | Recover Service Error
func RecoverServiceError() {
	if err := recover(); err != nil {
		log.Printf("%v", err)
	}
}
