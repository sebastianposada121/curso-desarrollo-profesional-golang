package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
	"web/models"
	"web/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/plutov/paypal"
)

// -----------------------WEBPAY-----------------

func PaymentGateway(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	template.Execute(response, nil)
}

func Webpay(response http.ResponseWriter, request *http.Request) {

	errorGodotenv := godotenv.Load()

	if errorGodotenv != nil {
		fmt.Println(errorGodotenv)
		return
	}

	webpay_url := os.Getenv("WEBPAY_URL")
	webpay_id := os.Getenv("WEBPAY_ID")
	webpay_secret := os.Getenv("WEBPAY_SECRET")

	client := &http.Client{}

	order := map[string]string{
		"buy_order":  "ordenCompra12345678",
		"session_id": "sesion1234557545",
		"amount":     "10000",
		"return_url": "http://localhost:8081/payment-gateway-webpay-info",
	}

	jsonOrder, _ := json.Marshal(order)

	webpayRequest, errorWebpay := http.NewRequest("POST", webpay_url, bytes.NewBuffer(jsonOrder))
	webpayRequest.Header.Set("Content-Type", "application/json")
	webpayRequest.Header.Set("Tbk-Api-Key-Id", webpay_id)
	webpayRequest.Header.Set("Tbk-Api-Key-Secret", webpay_secret)
	if errorWebpay != nil {
		fmt.Println("tenemos un error")
	}

	responseWebpay, errorResponseWebpay := client.Do(webpayRequest)

	if errorResponseWebpay != nil {
		fmt.Println(errorResponseWebpay)
	}
	defer responseWebpay.Body.Close()

	body, _ := io.ReadAll(responseWebpay.Body)
	webpay := models.WebpayModel{}

	errorUnmarshal := json.Unmarshal(body, &webpay)

	if errorUnmarshal != nil {
		fmt.Println(errorUnmarshal)
	}

	templateData := make(map[string]interface{})
	templateData["webpay"] = webpay

	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	template.Execute(response, templateData)
}

func WebpayInfo(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	token := params["token_ws"]
	errorGodotenv := godotenv.Load()

	if errorGodotenv != nil {
		fmt.Println(errorGodotenv)
		return
	}

	webpay_url := os.Getenv("WEBPAY_URL") + "/" + token
	webpay_id := os.Getenv("WEBPAY_ID")
	webpay_secret := os.Getenv("WEBPAY_SECRET")

	client := &http.Client{}

	webpayRequest, errorWebpay := http.NewRequest("GET", webpay_url, nil)
	webpayRequest.Header.Set("Content-Type", "application/json")
	webpayRequest.Header.Set("Tbk-Api-Key-Id", webpay_id)
	webpayRequest.Header.Set("Tbk-Api-Key-Secret", webpay_secret)
	if errorWebpay != nil {
		fmt.Println("tenemos un error")
	}

	responseWebpay, errorResponseWebpay := client.Do(webpayRequest)

	fmt.Println(responseWebpay.Status)
	if errorResponseWebpay != nil {
		fmt.Println(errorResponseWebpay)
	}
	defer responseWebpay.Body.Close()

	body, _ := io.ReadAll(responseWebpay.Body)
	webpay := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &webpay)

	if errorUnmarshal != nil {
		fmt.Println(errorUnmarshal)
	}

	templateData := make(map[string]interface{})
	templateData["webpayInfo"] = webpay

	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	template.Execute(response, templateData)
}

// -------------------------PAYPAL----------------------

func getAccessToken() string {

	PAYPAL_CLIENT_ID := os.Getenv("PAYPAL_CLIENT_ID")
	PAYPAL_SECRET := os.Getenv("PAYPAL_SECRET")
	paypalClient, errorPaypalClient := paypal.NewClient(PAYPAL_CLIENT_ID, PAYPAL_SECRET, paypal.APIBaseSandBox)
	if errorPaypalClient != nil {
		fmt.Println(errorPaypalClient)
	}

	accessToken, errorAccessToken := paypalClient.GetAccessToken()

	if errorAccessToken != nil {
		fmt.Println(errorAccessToken)
	}

	return accessToken.Token
}

func ordenPaypal(token string) map[string]interface{} {
	PAYPAL_URL := os.Getenv("PAYPAL_URL") + "/v2/checkout/orders"
	time := strings.Split(time.Now().String(), " ")
	order_id := string(time[4][6:14])
	order := `{
		"purchase_units": [
			{
				"amount": {
					"currency_code": "USD",
					"value": "10.00"
				},
				"reference_id": "orden_1"
			}
		],
		"intent": "CAPTURE",
		"payment_source": {
			"paypal": {
				"experience_context": {
					"payment_method_preference": "IMMEDIATE_PAYMENT_REQUIRED",
					"payment_method_selected": "PAYPAL",
					"brand_name": "Tamila",
					"locale": "es-ES",
					"landing_page": "LOGIN",
					"shipping_preference": "NO_SHIPPING",
					"user_action": "PAY_NOW",
					"return_url": "http://localhost:8081/payment-gateway-paypal-info",
					"cancel_url": "http://localhost:8081/payment-gateway-paypal-cancel"
				}
			}
		}
	}`

	byte_arr := []byte(order)
	client := &http.Client{}
	resquestOrder, errOrder := http.NewRequest("POST", PAYPAL_URL, bytes.NewBuffer(byte_arr))
	resquestOrder.Header.Set("Content-Type", "application/json")
	resquestOrder.Header.Set("PayPal-Request-Id", "orden_"+order_id)
	resquestOrder.Header.Set("Authorization", "Bearer "+token)

	if errOrder != nil {
		fmt.Println("error al crear orden", errOrder.Error())
	}

	responseOrder, errorResponse := client.Do(resquestOrder)
	defer resquestOrder.Body.Close()

	if errorResponse != nil {
		fmt.Println("Error respuesta paypal")
		return nil
	}

	body, _ := io.ReadAll(responseOrder.Body)

	data := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &data)

	if errorUnmarshal != nil {
		fmt.Println(errorUnmarshal)
	}

	return data

}

func Paypal(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	token := getAccessToken()

	order := ordenPaypal(token)

	fmt.Println(order)

	template.Execute(response, map[string]interface{}{
		"PaypalToken": token,
		"PaypalId":    order["id"],
	})
}

func PaypalInfo(response http.ResponseWriter, request *http.Request) {
	orderToken := request.URL.Query().Get("token")
	fmt.Println(orderToken)
	token := getAccessToken()

	data := paypalCapture(token, orderToken)

	fmt.Println(data)

	templateData := make(map[string]interface{})
	templateData["PaypalStatus"] = data["status"]
	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	template.Execute(response, templateData)
}

func PaypalCancel(response http.ResponseWriter, request *http.Request) {
	orderToken := request.URL.Query().Get("token")

	templateData := make(map[string]interface{})
	if orderToken != "" {
		templateData["PaypalStatus"] = "Cancel"
	}

	template := template.Must(template.ParseFiles("templates/payment-gateway.html", utils.PATH_TEMPLATE))
	template.Execute(response, templateData)
}

func paypalCapture(token string, orderToken string) map[string]interface{} {
	fmt.Println(token, "-----", orderToken)
	errorGodotenv := godotenv.Load()

	if errorGodotenv != nil {
		fmt.Println(errorGodotenv)
		return nil
	}

	webpay_url := os.Getenv("PAYPAL_URL") + "/v2/checkout/orders/" + orderToken + "/capture"

	client := &http.Client{}

	resquestCapture, errorCapture := http.NewRequest("POST", webpay_url, bytes.NewBuffer([]byte("{}")))
	resquestCapture.Header.Set("Content-Type", "application/json")
	resquestCapture.Header.Set("Authorization", "Bearer "+token)

	if errorCapture != nil {
		fmt.Println("tenemos un error")
		return nil
	}

	responseCapture, errorResponseCapture := client.Do(resquestCapture)

	fmt.Println(responseCapture.Status)
	if errorResponseCapture != nil {
		fmt.Println(errorResponseCapture)
	}
	defer responseCapture.Body.Close()

	body, _ := io.ReadAll(responseCapture.Body)
	capture := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &capture)

	if errorUnmarshal != nil {
		fmt.Println(errorUnmarshal)
	}

	return capture
}
