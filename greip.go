package greip

import (
	"errors"
	"net/url"
	"strings"
)

var availableGeoIPParams = []string{"location", "security", "timezone", "currency", "device"}
var availableCountryParams = []string{"language", "flag", "currency", "timezone"}
var baseUrl = "https://greipapi.com/"

// NewGreip initializes a new Greip instance
func NewGreip(apiToken string, test ...bool) *Greip {
	//? If the user provides a value for test, use it; otherwise, default to false.
	testValue := false
	if len(test) > 0 {
		testValue = test[0]
	}

	return &Greip{
		token:   apiToken,
		BaseURL: baseUrl,
		test:    testValue,
	}
}

// Lookup performs an IP lookup request to the Greip API to retrieve details
// about the specified IP address, such as location, security status, and more.
//
// Parameters:
//   - ip (string): The IP address to look up. This can be either an IPv4 or IPv6 address.
//   - params ([]string): An optional list of parameters to include in the lookup request.
//     The available parameters are: "location", "security", "timezone", "currency", and "device".
//   - lang (string): An optional parameter to specify the language for the response data.
//     The default language is English ("EN"). The supported languages are: "EN", "ES", "FR", "DE", "IT", "PT", "RU".
//
// Returns:
//
//   - *ResponseLookup: A pointer to a ResponseLookup struct containing the API response data
//     about the IP address. The ResponseLookup struct includes various fields such as IP type,
//     continent, country, region, city, ASN (Autonomous System Number) details, timezone,
//     security information (e.g., proxy, Tor, hosting), and device details (e.g., browser, OS).
//
//   - error: An error object if any issues occur during the lookup request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing an IP lookup without additional parameters
//	response, err := greipInstance.Lookup("1.1.1.1", nil)
//	if err != nil {
//	    log.Fatalf("Error performing IP lookup: %v", err)
//	}
//	fmt.Printf("IP Lookup Result: %+v\n", response)
//
//	// Performing an IP lookup with additional parameters
//	response, err = greipInstance.Lookup("1.1.1.1", []string{"device", "security"}, "EN")
//	if err != nil {
//	    log.Fatalf("Error performing IP lookup: %v", err)
//	}
//	fmt.Printf("IP Lookup with Params Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - If the `params` parameter is nil, the API will return the default set of data.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, malformed IP address).
func (g *Greip) Lookup(ip string, params []string, lang ...string) (*ResponseLookup, error) {
	//? If no params are provided, params will be an empty slice
	if params == nil {
		params = []string{} // Optional, as it will be nil by default
	}

	//? If the user provides a value for lang, use it; otherwise, default to "en".
	langValue := "EN"
	if len(lang) > 0 {
		langValue = lang[0]
	}

	payload := map[string]interface{}{
		"ip":     ip,
		"params": strings.Join(params, ","),
		"lang":   strings.ToUpper(langValue),
	}

	//? Validate the input IP
	if ip == "" {
		return nil, errors.New("you must provide the `ip` parameter")
	}

	//? Validate the params list
	if params == nil {
		params = []string{}
	}

	//? Validate the params
	if err := validateParams(params, availableGeoIPParams); err != nil {
		return nil, err
	}

	//? Validate the language
	if err := validateLang(langValue); err != nil {
		return nil, err
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("ip", ip)
	query.Set("params", strings.Join(params, ","))
	query.Set("lang", strings.ToUpper(langValue))

	//? Make the HTTP request
	var response ResponseLookup
	err := g.getRequest("IPLookup", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Threats performs a threat lookup using the Greip API to check if the specified IP address
// is associated with any known threats, such as proxy connection, hosting/cloud, Tor exit node, etc.
//
// Parameters:
//   - ip (string): The IP address to check for threats. This can be either an IPv4 or IPv6 address.
//
// Returns:
//
//   - *ResponseThreats: A pointer to a ResponseThreats struct containing the API response data
//     about the threat status of the IP address. The ResponseThreats struct includes fields such as
//     proxy status, Tor status, hosting/cloud status, etc.
//
//   - error: An error object if any issues occur during the threat lookup request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing a threat lookup for an IP address
//	response, err := greipInstance.Threats("1.1.1.1")
//	if err != nil {
//	    log.Fatalf("Error performing threat lookup: %v", err)
//	}
//	fmt.Printf("Threat Lookup Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, malformed IP address).
func (g *Greip) Threats(ip string) (*ResponseThreats, error) {
	payload := map[string]interface{}{
		"ip": ip,
	}

	//? Validate the input IP
	if ip == "" {
		return nil, errors.New("you must provide the `ip` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("ip", ip)

	//? Make the HTTP request
	var response ResponseThreats
	err := g.getRequest("threats", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs a bulk IP lookup using the Greip API to retrieve details about multiple IP addresses
// in a single request. The function takes a list of IP addresses and optional parameters to include in the lookup.
//
// Parameters:
//   - ips ([]string): A list of IP addresses to look up. Each IP address can be either an IPv4 or IPv6 address.
//   - params ([]string): An optional list of parameters to include in the lookup request.
//     The available parameters are: "location", "security", "timezone", "currency", and "device".
//   - lang (string): An optional parameter to specify the language for the response data.
//     The default language is English ("EN"). The supported languages are: "EN", "ES", "FR", "DE", "IT", "PT", "RU".
//
// Returns:
//
//   - *map[string]ResponseLookup: A pointer to a map containing the API response data for each IP address.
//     The map key is the IP address, and the value is a ResponseLookup struct with details about the IP address.
//
//   - error: An error object if any issues occur during the bulk lookup request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing a bulk IP lookup without additional parameters
//	ips := []string{"1.1.1.1", "2.2.2.2"}
//	response, err := greipInstance.BulkLookup(ips, nil, "EN")
//	if err != nil {
//	    log.Fatalf("Error performing bulk IP lookup: %v", err)
//	}
//	fmt.Printf("Bulk IP Lookup Result: %+v\n", response)
//
//	// Performing a bulk IP lookup with additional parameters
//	response, err = greipInstance.BulkLookup(ips, []string{"device", "security"}, "EN")
//	if err != nil {
//	    log.Fatalf("Error performing bulk IP lookup: %v", err)
//	}
//	fmt.Printf("Bulk IP Lookup with Params Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - If the `params` parameter is nil, the API will return the default set of data.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, malformed IP address).
func (g *Greip) BulkLookup(ips []string, params []string, lang ...string) (*map[string]ResponseLookup, error) {
	//? If no params are provided, params will be an empty slice
	if params == nil {
		params = []string{} // Optional, as it will be nil by default
	}

	//? If the user provides a value for lang, use it; otherwise, default to "en".
	langValue := "EN"
	if len(lang) > 0 {
		langValue = lang[0]
	}

	payload := map[string]interface{}{
		"ips":    strings.Join(ips, ","),
		"params": strings.Join(params, ","),
		"lang":   strings.ToUpper(langValue),
	}

	//? Validate the input IPs
	if ips == nil {
		return nil, errors.New("you must provide the `ips` parameter")
	}

	//? Validate the params list
	if params == nil {
		params = []string{}
	}

	//? Validate the params
	if err := validateParams(params, availableGeoIPParams); err != nil {
		return nil, err
	}

	//? Validate the language
	if err := validateLang(langValue); err != nil {
		return nil, err
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("ips", strings.Join(ips, ","))
	query.Set("params", strings.Join(params, ","))
	query.Set("lang", strings.ToUpper(langValue))

	//? Make the HTTP request
	var response map[string]ResponseLookup
	err := g.getRequest("BulkLookup", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs a country lookup using the Greip API to retrieve details about the specified country code,
// such as the country's language, flag, currency, and timezone.
//
// Parameters:
//   - countryCode (string): The ISO 3166-1 alpha-2 country code to look up (e.g., "US" for United States).
//   - params ([]string): An optional list of parameters to include in the lookup request.
//     The available parameters are: "language", "flag", "currency", and "timezone".
//   - lang (string): An optional parameter to specify the language for the response data.
//     The default language is English ("EN"). The supported languages are: "EN", "ES", "FR", "DE", "IT", "PT", "RU".
//
// Returns:
//
//   - *ResponseCountry: A pointer to a ResponseCountry struct containing the API response data
//     about the country. The ResponseCountry struct includes fields such as the country's name,
//     language, flag URL, currency code, and timezone.
//
//   - error: An error object if any issues occur during the country lookup request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing a country lookup without additional parameters
//	response, err := greipInstance.Country("US", nil, "EN")
//	if err != nil {
//	    log.Fatalf("Error performing country lookup: %v", err)
//	}
//	fmt.Printf("Country Lookup Result: %+v\n", response)
//
//	// Performing a country lookup with additional parameters
//	response, err = greipInstance.Country("US", []string{"language", "flag"}, "EN")
//	if err != nil {
//	    log.Fatalf("Error performing country lookup: %v", err)
//	}
//	fmt.Printf("Country Lookup with Params Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - If the `params` parameter is nil, the API will return the default set of data.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, malformed country code).
func (g *Greip) Country(countryCode string, params []string, lang ...string) (*ResponseCountry, error) {
	//? If no params are provided, params will be an empty slice
	if params == nil {
		params = []string{} // Optional, as it will be nil by default
	}

	//? If the user provides a value for lang, use it; otherwise, default to "en".
	langValue := "EN"
	if len(lang) > 0 {
		langValue = lang[0]
	}

	payload := map[string]interface{}{
		"CountryCode": countryCode,
		"params":      strings.Join(params, ","),
		"lang":        strings.ToUpper(langValue),
	}

	//? Validate the input countryCode
	if countryCode == "" {
		return nil, errors.New("you must provide the `countryCode` parameter")
	}

	//? Validate the params list
	if params == nil {
		params = []string{}
	}

	//? Validate the params
	if err := validateParams(params, availableCountryParams); err != nil {
		return nil, err
	}

	//? Validate the language
	if err := validateLang(langValue); err != nil {
		return nil, err
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("CountryCode", countryCode)
	query.Set("params", strings.Join(params, ","))
	query.Set("lang", strings.ToUpper(langValue))

	//? Make the HTTP request
	var response ResponseCountry
	err := g.getRequest("Country", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs a profanity check using the Greip API to detect any profane or inappropriate language
// in the specified text. The function returns a profanity score indicating the level of profanity
// detected in the text.
//
// Parameters:
//   - text (string): The text to check for profanity.
//
// Returns:
//
//   - *ResponseProfanity: A pointer to a ResponseProfanity struct containing the API response data
//     about the profanity score of the text. The ResponseProfanity struct includes fields such as
//     the profanity score, the list of bad words detected, and the original text.
//
//   - error: An error object if any issues occur during the profanity check request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing a profanity check on text
//	response, err := greipInstance.Profanity("This is a test message with some bad words.")
//	if err != nil {
//	    log.Fatalf("Error performing profanity check: %v", err)
//	}
//	fmt.Printf("Profanity Check Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty text).
func (g *Greip) Profanity(text string) (*ResponseProfanity, error) {
	payload := map[string]interface{}{
		"text": text,
	}

	//? Validate the input text
	if text == "" {
		return nil, errors.New("you must provide the `text` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("text", text)

	//? Make the HTTP request
	var response ResponseProfanity
	err := g.getRequest("badWords", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs an ASN lookup using the Greip API to retrieve details about the specified Autonomous System Number (ASN).
//
// Parameters:
//   - asn (string): The ASN to look up.
//
// Returns:
//
//   - *ResponseASN: A pointer to a ResponseASN struct containing the API response data
//     about the ASN. The ResponseASN struct includes fields such as the ASN number, name, country code,
//     description, and other details.
//
//   - error: An error object if any issues occur during the ASN lookup request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing an ASN lookup for an ASN number
//	response, err := greipInstance.AsnLookup("AS13335")
//	if err != nil {
//	    log.Fatalf("Error performing ASN lookup: %v", err)
//	}
//	fmt.Printf("ASN Lookup Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty ASN).
func (g *Greip) AsnLookup(asn string) (*ResponseASN, error) {
	payload := map[string]interface{}{
		"asn": asn,
	}

	//? Validate the input ASN
	if asn == "" {
		return nil, errors.New("you must provide the `asn` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("asn", asn)

	//? Make the HTTP request
	var response ResponseASN
	err := g.getRequest("ASNLookup", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs an email validation using the Greip API to check the validity of the specified email address.
//
// Parameters:
//   - email (string): The email address to validate.
//
// Returns:
//
//   - *ResponseEmail: A pointer to a ResponseEmail struct containing the API response data
//     about the email validation status. The ResponseEmail struct includes fields such as
//     the email address, domain, validity status, and other details.
//
//   - error: An error object if any issues occur during the email validation request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing an email validation for an email address
//	response, err := greipInstance.Email("
//	if err != nil {
//	    log.Fatalf("Error performing email validation: %v", err)
//	}
//	fmt.Printf("Email Validation Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty email address).
func (g *Greip) Email(email string) (*ResponseEmail, error) {
	payload := map[string]interface{}{
		"email": email,
	}

	//? Validate the input email
	if email == "" {
		return nil, errors.New("you must provide the `email` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("email", email)

	//? Make the HTTP request
	var response ResponseEmail
	err := g.getRequest("validateEmail", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs a phone validation using the Greip API to check the validity of the specified phone number.
//
// Parameters:
//   - phone (string): The phone number to validate.
//   - countryCode (string): The ISO 3166-1 alpha-2 country code for the phone number (e.g., "US" for United States).
//
// Returns:
//
//   - *ResponsePhone: A pointer to a ResponsePhone struct containing the API response data
//     about the phone validation status. The ResponsePhone struct includes fields such as
//     the phone number, country code, validity status, and other details.
//
//   - error: An error object if any issues occur during the phone validation request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing a phone validation for a phone number
//	response, err := greipInstance.Phone("123456789", "US")
//	if err != nil {
//	    log.Fatalf("Error performing phone validation: %v", err)
//	}
//	fmt.Printf("Phone Validation Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty phone number).
func (g *Greip) Phone(phone string, countryCode string) (*ResponsePhone, error) {
	payload := map[string]interface{}{
		"phone":       phone,
		"countryCode": countryCode,
	}

	//? Validate the input phone
	if phone == "" {
		return nil, errors.New("you must provide the `phone` parameter")
	}

	//? Validate the input countryCode
	if countryCode == "" {
		return nil, errors.New("you must provide the `countryCode` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("phone", phone)
	query.Set("countryCode", countryCode)

	//? Make the HTTP request
	var response ResponsePhone
	err := g.getRequest("validatePhone", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs an IBAN validation & lookup using the Greip API to check the validity of the specified IBAN (International Bank Account Number).
//
// Parameters:
//   - iban (string): The IBAN to validate.
//
// Returns:
//
//   - *ResponseIBAN: A pointer to a ResponseIBAN struct containing the API response data
//     about the IBAN validation status and details. The ResponseIBAN struct includes fields such as
//     the IBAN number, country code, validity status, and other details.
//
//   - error: An error object if any issues occur during the IBAN validation request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	// Performing an IBAN validation for an IBAN number
//	response, err := greipInstance.IBAN("GB82WEST12345698765432")
//	if err != nil {
//	    log.Fatalf("Error performing IBAN validation: %v", err)
//	}
//	fmt.Printf("IBAN Validation Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty IBAN).
func (g *Greip) IBAN(iban string) (*ResponseIBAN, error) {
	payload := map[string]interface{}{
		"iban": iban,
	}

	//? Validate the input iban
	if iban == "" {
		return nil, errors.New("you must provide the `iban` parameter")
	}

	//? Construct the query parameters
	query := url.Values{}
	query.Set("iban", iban)

	//? Make the HTTP request
	var response ResponseIBAN
	err := g.getRequest("validateIBAN", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Performs a payment fraud detection using the Greip API to check the payment data for potential fraud indicators.
//
// Parameters:
//   - data (map[string]interface{}): A map containing the payment data to check for fraud. The payment data
//     should include fields such as card number, expiry date, CVV, billing address, etc.
//
// Returns:
//
//   - *ResponsePayment: A pointer to a ResponsePayment struct containing the API response data
//     about the payment fraud status. The ResponsePayment struct includes fields such as the payment data,
//     fraud score, risk level, and other details.
//
//   - error: An error object if any issues occur during the payment fraud detection request, such as
//     network failures or invalid responses from the API. It returns nil if the request succeeds.
//
// Example usage:
//
//	  // Performing a payment fraud detection for payment data
//	  paymentData := map[string]interface{}{
//	      "cardNumber": "4111111111111111",
//	      "customer_id": "123456",
//	      "customer_firstname": "John",
//	      "customer_lastname": "Doe",
//		     "customer_email": "name@domain.com",
//	      "customer_ip": "1.1.1.1"
//	  }
//	  response, err := greipInstance.Payment(paymentData)
//	  if err != nil {
//	      log.Fatalf("Error performing payment fraud detection: %v", err)
//	  }
//	  fmt.Printf("Payment Fraud Detection Result: %+v\n", response)
//
// Notes:
//   - This function uses the provided API token stored in the Greip instance to authorize
//     the request. Ensure that a valid token is set when initializing the Greip instance.
//   - It is recommended to handle any errors returned by this function to ensure robust code execution.
//
// Errors:
//   - Network-related errors (e.g., timeouts, unreachable server).
//   - API-related errors (e.g., invalid API token, empty payment data).
func (g *Greip) Payment(data map[string]interface{}) (*ResponsePayment, error) {
	//? Validate the input data
	if data == nil {
		return nil, errors.New("you must provide the `data` parameter")
	}

	//? Create new variable for the data
	payload := map[string]interface{}{
		"data": data,
	}

	//? Make the HTTP request
	var response ResponsePayment
	err := g.postRequest("paymentFraud", &response, payload)
	if err != nil {
		return nil, err
	}

	return &response, err
}
