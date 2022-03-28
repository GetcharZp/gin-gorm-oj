package define

import "os"

var (
	DefaultPage = "1"
	DefaultSize = "20"
)

var MailPassword = os.Getenv("MailPassword")
