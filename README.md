# UNKCode SDK Go

This is the official UNKCode SDK for Go applications. The SDK allows you to easily integrate UNKCode's license verification into your applications.

## Installation

To use the SDK, you need to add it as a dependency in your Go project. You can do this using the `go get` command:

```bash
go get github.com/unkcode-org/sdk-go
```

## Usage

Below is how you can use the SDK in your application to perform license verification with UNKCode.

```go
package main

import (
	"fmt"
	"github.com/unkcode-org/sdk-go/interface"
)

func main() {
	// Create an instance of AppInterface with the required configuration
	app := &appinterface.AppInterface{
		Name:                  "YourAppName",
		Token:                 "YourSecretKey",
		RegularVerification:   true, // Enable recurring verification
		FailedVerificationCallback: func() {
			fmt.Println("License regular verification failed!")
		},
	}

	// Perform login with the license and MAC address (optional)
	success := app.Login("UserLicense", "MACAddress")
	if success {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed!")
	}
}
```

Make sure to replace `"YourAppName"` and `"YourSecretKey"` with the appropriate values provided by UNKCode. If you don't need recurring verification, you can set `RegularVerification` to `false`.

## Contributions

If you find any issues or have suggestions to improve the SDK, feel free to open an issue or submit a pull request on the GitHub repository.

## License

This SDK is available under the MIT License. See the `LICENSE` file for more information.
