package sharable

import "os"

var (
	ApiKey          = os.Getenv("API_KEY")
	PORT            = os.Getenv("PORT")
	AllowedOrigins1 = os.Getenv("ALLOWED_ORIGINS_1")
	AllowedOrigins2 = os.Getenv("ALLOWED_ORIGINS_2")
)
