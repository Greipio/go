# Greip Go Library

The Greip Go library allows you to easily interact with the Greip API to access a variety of services, including IP geolocation, threat intelligence, email validation, and more.

[Report Issue](https://github.com/Greipio/go/issues/new) ·
[Request Feature](https://github.com/Greipio/go/discussions/new?category=ideas)
· [Greip Website](https://greip.io/) · [Documentation](https://docs.greip.io/)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Greipio/go?color=green&label=Minified%20size&logo=github)
&nbsp;&nbsp;
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/apache-2-0)
&nbsp;&nbsp;
![API Status](https://img.shields.io/website?down_color=orange&down_message=down&label=API%20status&up_color=green&up_message=up&url=https%3A%2F%2Fgreipapi.com)

---

## Installation

You can install the Greip library by running:

```bash
go get github.com/Greipio/go
```

## Usage

To use the Greip library, first import the package and initialize the Greip instance with your API token. Here’s a basic example:

```go
package main

import (
    "fmt"
    "github.com/Greipio/go"
)

func main() {
    // Initialize the Greip instance with your API token
    greipInstance := greip.NewGreip("YOUR_API_TOKEN")

    // Example: Lookup IP information
    response, err := greipInstance.Lookup("1.1.1.1")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(response.IP, response.ContinentName, response.City)
}
```

## Methods

The Greip library provides various methods to interact with the API:

- **Lookup(ip string, params []string, lang ...string)**: Get geolocation information about an IP address.
- **Threats(ip string)**: Get threat intelligence related to an IP address.
- **BulkLookup(ips []string, params []string, lang ...string)**: Get geolocation information for multiple IP addresses.
- **Country(countryCode string, params []string, lang ...string)**: Get information about a country by its code.
- **Profanity(text string)**: Check if a given text contains profanity.
- **ASN(asn string)**: Get information about an ASN (Autonomous System Number).
- **Email(email string)**: Validate an email address.
- Phone(phone string, countryCode string): Validate or lookup a phone number.
- **IBAN(iban string)**: Validate or lookup an IBAN number.
- **Payment(data map[string]interface{})**: Check if a payment transaction is fraudulent.

## Example of Method Usage

```go
// Lookup country information
countryInfo, err := greipInstance.Country("US")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(countryInfo.CountryName, countryInfo.Population)
```

## Development Mode

If you need to test the integration without affecting your subscription usage, you can set the test attribute to true when initializing the Greip instance:

```go
greipInstance := greip.NewGreip("YOUR_API_TOKEN", true)
```

> [!WARNING]
> Enabling the test mode returns fake data. **Do not use it in production**.

## Error Handling

The library returns error for invalid parameters and request-related issues. Here’s an example of handling errors:

```go
response, err := greipInstance.Lookup("INVALID_IP")
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(response)
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any improvements or bugs.

## License

This project is licensed under the Apache 2.0 License - see the LICENSE file for details.
