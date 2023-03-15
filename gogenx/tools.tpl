package utils

import (
	"fmt"

	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MakeTraceID(prefix, suffix string) string {
	t := time.Now()
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(999) + 1
	y := rand.Intn(999) + 1
	return strings.ToUpper(prefix) + t.Format("060102150405") +
		fmt.Sprintf("%03d", x) + fmt.Sprintf("%03d", y) +
		strings.ToUpper(suffix)
}

func GetCurrentWeek() string {
	year, week := time.Now().UTC().ISOWeek()
	return strconv.Itoa(year) + fmt.Sprintf("%02d",week)
}

func IPChecker(myip string, whitelist string) bool {
	if whitelist == "" {
		return false
	}
	for _, ip := range strings.Split(whitelist, ",") {
		if !strings.Contains(ip, "/") {
			ip = ip + "/32"
		}
		_, ipnetA, _ := net.ParseCIDR(ip)
		ipB := net.ParseIP(myip)

		if ipnetA.Contains(ipB) {
			return true
		}
	}
	return false
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}