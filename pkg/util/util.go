package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"hash/fnv"
	"net/http"
	"net/url"
	"time"

	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/sirupsen/logrus"
)

// ComputeHash - get the hash of the byte array sent in
func ComputeHash(data interface{}) (uint64, error) {
	dataB, err := json.Marshal(data)
	if err != nil {
		return 0, fmt.Errorf("could not marshal data to bytes")
	}

	h := fnv.New64a()
	h.Write(dataB)
	return h.Sum64(), nil
}

// MaskValue - mask sensitive information with * (asterisk).  Length of sensitiveData to match returning maskedValue
func MaskValue(sensitiveData string) string {
	var maskedValue string
	for i := 0; i < len(sensitiveData); i++ {
		maskedValue += "*"
	}
	return maskedValue
}

// PrintDataInterface - prints contents of the interface only if in debug mode
func PrintDataInterface(data interface{}) {
	if log.GetLevel() == logrus.DebugLevel {
		PrettyPrint(data)
	}
}

// PrettyPrint - print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// GetProxyURL - need to provide my own function (instead of http.ProxyURL()) to handle empty url. Returning nil
// means "no proxy"
func GetProxyURL(fixedURL *url.URL) func(*http.Request) (*url.URL, error) {
	return func(*http.Request) (*url.URL, error) {
		if fixedURL == nil || fixedURL.Host == "" {
			return nil, nil
		}
		return fixedURL, nil
	}
}

// GetURLHostName - return the host name of the passed in URL
func GetURLHostName(urlString string) string {
	host, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return host.Hostname()
}

// StringSliceContains - does the given string slice contain the specified string?
func StringSliceContains(items []string, s string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveDuplicateValuesFromStringSlice - remove duplicate values from a string slice
func RemoveDuplicateValuesFromStringSlice(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ConvertTimeToMillis - convert to milliseconds
func ConvertTimeToMillis(tm time.Time) int64 {
	return tm.UnixNano() / 1e6
}

// IsNotTest determines if a test is running or not
func IsNotTest() bool {
	return flag.Lookup("test.v") == nil
}
