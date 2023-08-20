package appinterface

import "github.com/unkcode-org/sdk-go/src/auth"

type AppInterface struct {
	Name                       string
	Token                      string
	RegularVerification        bool
	FailedVerificationCallback func()
}

func (this *AppInterface) Login(license string, macAddress string) bool {
	return auth.LoginWithLicense(this.Token, this.Name, this.RegularVerification, license, macAddress, this.FailedVerificationCallback)
}
