package httpdemo

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BojanKomazec/go-demo/internal/pkg/osdemo"
)

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

func downloadFile(filepath string, url string, wg *sync.WaitGroup) (err error) {

	wg.Add(1)
	fmt.Println("Writing to file: ", filepath)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Writing to file completed!: ", filepath)
	defer wg.Done()
	return nil
}

// https://www.adam.com.au/support/blank-test-files
// http://xcal1.vodafone.co.uk/
// https://speed.hetzner.de/
// https://www.thinkbroadband.com/download
func downloadInParallelDemo(outputDir string) error {
	urls := [...]string{
		"http://mirror.filearena.net/pub/speed/SpeedTest_16MB.dat?_ga=2.196721836.1613271072.1560530397-1086101020.1560530397",
		"http://mirror.filearena.net/pub/speed/SpeedTest_32MB.dat?_ga=2.196721836.1613271072.1560530397-1086101020.1560530397",
		"http://mirror.filearena.net/pub/speed/SpeedTest_64MB.dat?_ga=2.196721836.1613271072.1560530397-1086101020.1560530397",
		"https://speed.hetzner.de/100MB.bin",
		"http://212.183.159.230/5MB.zip",
		// "http://212.183.159.230:81/5MB.zip",
		"http://212.183.159.230/10MB.zip",
		// "http://212.183.159.230:81/10MB.zip",
		"http://212.183.159.230/20MB.zip",
		// "http://212.183.159.230:81/20MB.zip",
		"http://ipv4.download.thinkbroadband.com/5MB.zip",
		"http://ipv4.download.thinkbroadband.com:81/5MB.zip",
		"http://ipv4.download.thinkbroadband.com/10MB.zip",
		"http://ipv4.download.thinkbroadband.com:81/10MB.zip",
		"http://ipv4.download.thinkbroadband.com/20MB.zip",
		"http://ipv4.download.thinkbroadband.com:81/20MB.zip",
	}

	var wg sync.WaitGroup
	downloadsDirName := filepath.Join(outputDir, "downloads")
	err := osdemo.CreateDirIfNotExist(downloadsDirName)
	if err != nil {
		return err
	}

	for _, rawURL := range urls {
		fmt.Println("url:", rawURL)
		parts, err := splitURL(rawURL)
		if err != nil {
			return err
		}
		fmt.Println("parts =", parts)
		fileName := filepath.Join(downloadsDirName, strings.Join(parts, "_"))
		fmt.Println("*********** fileName =", fileName)
		go downloadFile(fileName, rawURL, &wg)
	}

	fmt.Println("All goroutines initiated.")
	wg.Wait()
	fmt.Println("All goroutines completed.")
	return nil
}

func splitURL(rawURL string) (parts []string, err error) {
	parts = make([]string, 0)

	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	fmt.Println("url =", url)

	path := url.Path
	fmt.Println("path =", path)

	host := url.Host
	fmt.Println("host (may include port) =", host)

	if strings.Index(host, ":") > -1 {
		var port string
		host, port, _ = net.SplitHostPort(host)
		fmt.Println("host =", host)
		fmt.Println("port =", port)
		host = host + "_port_" + port
	}

	parts = strings.Split(host, ".")

	fragment := url.Fragment
	fmt.Println("fragment =", fragment)

	uriSegments := strings.Split(path, "/")
	fmt.Println("uriSegments (might contain empty string element) =", uriSegments)

	splitFn := func(c rune) bool {
		return c == '/'
	}
	uriSegments = strings.FieldsFunc(path, splitFn)
	fmt.Println("uriSegments (without empty string elements) =", uriSegments)

	fmt.Println("last uriSegment (resource name) =", uriSegments[len(uriSegments)-1])

	parts = append(parts, uriSegments...)
	return parts, err
}

// ShowDemo func
func ShowDemo(outputDir string) {
	fmt.Println("outputDir =", outputDir)
	// downloadInParallelDemo(outputDir)
}
