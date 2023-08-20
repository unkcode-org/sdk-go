package appinterface

import "github.com/unkcode-org/sdk-go/auth"

type AppInterface struct {
	Name                       string
	Token                      string
	RegularVerification        bool
	FailedVerificationCallback func()
}

func (app *AppInterface) Login(license string, macAddress string) bool {
	return auth.LoginWithLicense(app.Token, app.Name, app.RegularVerification, license, macAddress, app.FailedVerificationCallback)
}
