package tambahan

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParsingUrl(t *testing.T) {
	urlString := "https://www.tokopedia.com/globallaptopthinkpad/laptop-lenovo-thinkpad-t450-t450s-ts-intel-i5-i7-gen-5-ram-4-8-gb-ssd-t450-i5gen5-4gb-ssd-256gb?extParam=whid%3D13851530"

	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	fmt.Println("HOST:", u.Host)
	fmt.Println("SCHEME:", u.Scheme)
	fmt.Println("FRAGMENT:", u.Fragment)
	fmt.Println("PATH:", u.Path)
	fmt.Println("QUERY:", u.Query())
}

/*
=== RUN   TestParsingUrl
HOST: www.tokopedia.com
SCHEME: https
FRAGMENT:
PATH: /globallaptopthinkpad/laptop-lenovo-thinkpad-t450-t450s-ts-intel-i5-i7-gen-5-ram-4-8-gb-ssd-t450-i5gen5-4gb-ssd-256gb
QUERY: map[extParam:[whid=13851530]]
--- PASS: TestParsingUrl (0.00s)
PASS
ok      learn-go-ii/tambahan    0.390s
*/
