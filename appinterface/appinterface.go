package appinterface

import "github.com/unkcode-org/sdk-go/auth"

type UNKCodeAppInterface struct {
	Name                       string
	Token                      string
	RegularVerification        bool
	FailedVerificationCallback func()
}

func (app *UNKCodeAppInterface) Login(license string, macAddress string) bool {
	return auth.LoginWithLicense(app.Token, app.Name, app.RegularVerification, license, macAddress, app.FailedVerificationCallback)
}
