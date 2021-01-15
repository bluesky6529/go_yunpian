package Money

import (
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func Money(Apikey string) float64 {
	c := resty.New().SetHostURL("https://sms.yunpian.com")
	//c := resty.New().SetHostURL("https://sms.yunpian.com").EnableTrace().SetDebug(true)
	resp, err := c.R().
		SetHeaders(map[string]string{
			"Accept": "application/json;charset=utf-8",
		}).
		SetFormData(map[string]string{
			"apikey": Apikey,
		}).
		Post("/v2/user/get.json")
	if err != nil {
		panic(err)
	}

	/*fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	//fmt.Println("Request Trace Info:")
	//ti := resp.Request.TraceInfo()
	//fmt.Println("  DNSLookup     :", ti.DNSLookup)
	//fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
	*/
	//log.Printf("%f", gjson.GetBytes(resp.Body(), "balance").Float())
	return gjson.GetBytes(resp.Body(), "balance").Float()
}
