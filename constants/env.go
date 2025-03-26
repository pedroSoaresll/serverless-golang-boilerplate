package constants

import "os"

var (
	ENV                 = os.Getenv("ENV")
	CDK_DEFAULT_ACCOUNT = os.Getenv("CDK_DEFAULT_ACCOUNT")
	CDK_DEFAULT_REGION  = os.Getenv("CDK_DEFAULT_REGION")
)
