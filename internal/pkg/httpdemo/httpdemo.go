package httpdemo

// import (
// 	"net/http"
// 	"fmt"
// 	"http"
// )

// func HttpClientDemo() {
// 	client := http.Client{
// 		CheckRedirect: func(req *http.Request, via []*http.ReadRequest) error {
// 			return http.ErrUseLastResponse
// 		}
// 	}

// 	req, reqErr := http.NewRequest(requestMethod, requestUrl, nil)
// 	resp, clientErr := client.Do(req)
// }
