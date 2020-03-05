package httpdemo

import (
	"fmt"
	"io"
	"log"
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

// For url https://danube.example.com:8080/?cid={customerId}&country=rs
// it returns [danube example com_port_8080].
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

	if len(uriSegments) > 0 {
		fmt.Println("last uriSegment (resource name) =", uriSegments[len(uriSegments)-1])
	}

	parts = append(parts, uriSegments...)
	return parts, err
}

// "https://danube.example.com/?cid={customerId}&country=rs
// output: danube.com
func extractDomain(rawURL string) (string, error) {

	url, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	fmt.Println("url =", url)

	host := url.Host
	fmt.Println("host (may include port) =", host)

	host = strings.Split(host, ":")[0]
	fmt.Println("host (must not include port) =", host)

	parts := strings.Split(host, ".")
	fmt.Println("parts =", parts)

	return parts[len(parts)-2] + "." + parts[len(parts)-1], nil
}

func extractDomainDemo() {
	url := "https://danube.example.com/?cid={customerId}&country=rs"
	domain, err := extractDomain(url)
	if err != nil {
		return
	}
	log.Println("Extracted domain:", domain)

	url2 := "https://danube.example.com:8080/?cid={customerId}&country=rs"
	domain, err = extractDomain(url2)
	if err != nil {
		return
	}
	log.Println("Extracted domain:", domain)
}

// function name based on https://github.com/google/guava/wiki/InternetDomainNameExplained
// @todo: find out how to extract domain segments when it contains top public suffixes e.g. org.uk
func ensureTopDomainUnderRegistrySuffix(rawURL string, topDomainUnderRegistrySuffix string) (string, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	fmt.Println("url =", url)
	fmt.Println("url.Scheme =", url.Scheme)
	fmt.Println("url.Opaque =", url.Opaque)
	fmt.Println("url.User =", url.User)
	fmt.Println("url.Host =", url.Host)
	fmt.Println("url.Path =", url.Path)
	fmt.Println("url.RawPath =", url.RawPath)
	fmt.Println("url.ForceQuery =", url.ForceQuery)
	fmt.Println("url.RawQuery =", url.RawQuery)
	fmt.Println("url.Fragment =", url.Fragment)

	topDomainUnderRegistrySuffixAndPort := strings.Split(url.Host, ":")
	currentTopDomainUnderRegistrySuffix := topDomainUnderRegistrySuffixAndPort[0]
	fmt.Println("currentTopDomainUnderRegistrySuffix =", currentTopDomainUnderRegistrySuffix)
	if len(topDomainUnderRegistrySuffixAndPort) > 1 {
		port := topDomainUnderRegistrySuffixAndPort[1]
		fmt.Println("port =", port)
		url.Host = strings.Join([]string{topDomainUnderRegistrySuffix, port}, ":")
	} else {
		url.Host = topDomainUnderRegistrySuffix
	}

	return url.String(), nil
}

func ensureTopDomainUnderRegistrySuffixDemo() {
	url := "https://danube.example.com/?cid={customerId}&country=rs"
	resURL, err := ensureTopDomainUnderRegistrySuffix(url, "example2.rs")
	if err != nil {
		return
	}
	log.Println("Output url:", resURL)

	url2 := "https://danube.example.com:8080/?cid={customerId}&country=rs"
	resURL, err = ensureTopDomainUnderRegistrySuffix(url2, "dev.example2.rs")
	if err != nil {
		return
	}
	log.Println("Output url:", resURL)
}

// topDomain is e.g. example.com or example.org.uk
func replaceDomainSegment(rawURL string, oldDomainSegment string, newDomainSegment string) (string, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	domainAndPort := strings.Split(url.Host, ":")
	domain := domainAndPort[0]
	domain = strings.Replace(domain, oldDomainSegment, newDomainSegment, 1)

	if len(domainAndPort) > 1 {
		url.Host = strings.Join([]string{domain, domainAndPort[1]}, ":")
	} else {
		url.Host = domain
	}

	return url.String(), nil
}

func replaceDomainSegmentDemo() {
	log.Println("replaceDomainSegment()")

	url := "https://danube.example.com/?cid={customerId}&country=rs"
	resURL, err := replaceDomainSegment(url, "example.com", "example2.rs")
	if err != nil {
		return
	}
	log.Println("Output url:", resURL)

	url2 := "https://danube.example.org.uk:8080/?cid={customerId}&country=rs"
	resURL, err = replaceDomainSegment(url2, "danube.example.org.uk", "dev.example2.rs")
	if err != nil {
		return
	}
	log.Println("Output url:", resURL)
}

// ShowDemo func
func ShowDemo(outputDir string) {
	// fmt.Println("outputDir =", outputDir)
	// downloadInParallelDemo(outputDir)
	extractDomainDemo()
	ensureTopDomainUnderRegistrySuffixDemo()
	replaceDomainSegmentDemo()
}
