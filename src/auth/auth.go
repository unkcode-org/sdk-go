package auth

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/unkcode-org/sdk-go/src/config"
)

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type VerifyRequestBody struct {
	Data string `json:"data"`
}

type VerifyRequestParams struct {
	Auth string `json:"auth"`
}

type VerifyLicenseRequestParams struct {
	Token string `json:"token"`
	Mac   string `json:"macHash,omitempty"`
}

// LoginWithLicense realiza el proceso de inicio de sesión utilizando una licencia y dirección MAC (opcional).
func LoginWithLicense(token string, name string, regularValidation bool, license string, macAddress string, cb func()) bool {
	// Crear una instancia del cliente HTTP de resty
	client := resty.New()

	// Generar el hash MD5 del token para la autenticación
	authHash := md5.Sum([]byte(token))

	// Construir la URL de la pre verificación de licencia
	verifyURL := fmt.Sprintf("%s%s", config.BaseURL, name)

	// Construir la estructura del cuerpo de la peticion
	verifyRequestParams := VerifyRequestParams{
		Auth: fmt.Sprintf("%x", authHash),
	}

	// Codificar la estructura verifyRequestParams en JSON
	authJSON, err := json.Marshal(verifyRequestParams)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Codificar el JSON en base64
	authB64 := base64.StdEncoding.EncodeToString(authJSON)

	// Crear el cuerpo de solicitud para la pre verificación
	verifyRequestBody := VerifyRequestBody{
		Data: authB64,
	}

	// Realizar la primer solicitud POST con el cuerpo de solicitud codificado
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(verifyRequestBody).
		Post(verifyURL)

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Decodificar la respuesta JSON
	var preVerifyData AuthResponse
	err = json.Unmarshal(resp.Body(), &preVerifyData)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Verificar si el mensaje contiene "SUCN"
	if !containsSUCN(preVerifyData.Message) {
		fmt.Println("[-] Error on pre verification request")
		return false
	}

	// Construir la estructura del cuerpo de la peticion
	verifyLicenseRequestParams := VerifyLicenseRequestParams{
		Token: preVerifyData.Token,
	}

	// Añadir el mac Address si la app lo requiere
	if macAddress != "" {
		verifyLicenseRequestParams.Mac = macAddress
	}

	// Codificar la estructura verifyRequestParams en JSON
	licenseJSON, err := json.Marshal(verifyLicenseRequestParams)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Codificar el JSON en base64
	licenseB64 := base64.StdEncoding.EncodeToString(licenseJSON)

	fmt.Println("STR")
	fmt.Println(string(licenseB64))

	// Crear el cuerpo de solicitud para la verificación
	verifyLicenseRequestBody := VerifyRequestBody{
		Data: licenseB64,
	}

	// Construir la URL de verificación de licencia
	verifyLicenseURL := fmt.Sprintf("%s/%s", verifyURL, license)

	// Realizar la solicitud POST con el cuerpo de solicitud codificado
	verifyResp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(verifyLicenseRequestBody).
		Post(verifyLicenseURL)

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Decodificar la respuesta JSON
	var data map[string]interface{}
	err = json.Unmarshal(verifyResp.Body(), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	fmt.Println("xd")
	fmt.Println(data)

	// Verificar si el mensaje contiene "SUCN"
	if message, ok := data["message"].(string); !ok || !containsSUCN(message) {
		fmt.Println(message)
		return false
	}

	// Verificar si la aplicacion realiza validacion de licencia recurrente
	if regularValidation {
		CreateValidationRoutine(token, name, regularValidation, license, macAddress, cb)
	}

	return true
}

func CreateValidationRoutine(token string, name string, regularValidation bool, license string, macAddress string, cb func()) {
	retries := 0
	maxRetries := 5

	for retries < maxRetries {
		time.Sleep(20 * time.Second)
		fmt.Println("Rutina ejecutada")
		logged := LoginWithLicense(token, name, regularValidation, license, macAddress, cb)

		if !logged {
			retries++
		}
	}

	cb()
}

func containsSUCN(message string) bool {
	return len(message) >= 4 && message[:4] == "SUCN"
}
