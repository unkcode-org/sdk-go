# UNKCode SDK Go

This is the official UNKCode SDK for Go applications. The SDK allows you to easily integrate UNKCode's license verification into your applications.

## Usage

To use the UNKCode SDK in your Go application, you can import the necessary packages and use their functionalities. Make sure your project is using Go Modules.

### 1. Add a reference to the SDK in your `go.mod` file:

```go
require (
    github.com/unkcode-org/sdk-go v1.0.0
)
```

Replace `v1.0.0` with the appropriate version of the SDK.

### 2. Import and use the SDK packages in your Go code:

```go
package main

import (
    "fmt"
    "github.com/unkcode-org/sdk-go/appinterface"
)

func main() {
    app := &appinterface.UNKCodeAppInterface{
        // Configuration
    }

    // Use the SDK packages as needed
}
```

## Installation

The UNKCode SDK does not require a separate installation step. You can import the SDK packages directly into your Go project using the steps mentioned above. The Go Modules system will automatically manage the dependencies and make the SDK packages available for use in your project.

## Example

Below is how you can use the SDK in your application to perform license verification with UNKCode.

```go
package main

import (
	"fmt"
	"github.com/unkcode-org/sdk-go/interface"
)

func main() {
	// Create an instance of AppInterface with the required configuration
	app := &appinterface.UNKCodeAppInterface{
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

Contributions to the UNKCode Go SDK are welcome! If you find any issues or would like to add new features, feel free to open an issue or submit a pull request on the GitHub repository: https://github.com/unkcode-org/sdk-go.

## License

This SDK is available under the MIT License. See the `LICENSE` file for more information.
