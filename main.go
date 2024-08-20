package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	Id          string `json:"id"`
	StatusCode  string `json:"statusCode"`
	Description string `json:"description"`
}

func main() {
	for i := 0; i < 100; i++ {
		go callAu()
		go callAu()
		go callAu()
		time.Sleep(1 * time.Second)
	}
}

func callAu() {
	posturl := "https://auspark.au.edu/Planner/SubmitRegistration"

	body := []byte(`{
		"SectionIds": [
		  "290073",
		  "290077",
		  "290109",
		  "290116",
		  "290871",
		  "291132"
		],
		"Channel": "W"
	  }`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	cookie := "_ga=GA1.2.128613498.1715101706; Identity.External=CfDJ8MJ3Ss7oxyFKjxmon9SVuR9KmPZ3HAVWQP3pgY2GtW4WQf6Mu0f7yb7ORAnjqoof3zcilM5OHTyDZj8FU2jFap2tmsaeWWjonHugDllB62sY9YloQE0fSO0qWD_jdKWFfYS2LMNeFIzLeD2BGBM8DTb5_hTMpJ1hFfrK9EuIDm-yFBX35k3DHWU8mZl_Wt0fJkR2cyej1_vNPJVIoYvfvaeJcKz5EsEeZZcIMtnfus4MdorLjJMJ46PcQ9MkXGfHOzQ0tmy5mnQB7I_asizGERXucjjOXpd8s6R-_FIPm-QPSnVJdPbK38ltXiUPqAuya7Xw6V12QOW1y5Q7Wan3FIsIJckyQDgfGMYI076CHTIR8gm1v3jz3j7bhuNBeDzjXWv9I5U4zxkQaUYixfJ91gM3DMHNGjqTKSOLZRgssd8y11sW5iluKmKGD5yNHukP9A2v2LL54j180dmn9Sx0MkkiRbFuLRvvXr3czY7aRmNz1wJyP5xJ25ccg_YXp1GHHu2KS4UflVEhaWASFZxUsHpTB70_BQeWKfeRDF2z2TgsLS_kSmhi0L-pxaC0JFi2FXPG9W71Dwp964Dkx1mFbQAiBcUjnvyao5fEHGS1upgfvoT7QyX-Ye7iy_pTvqX6Txl9sDVWShj3-YhrmlC2Dfe_WJugXFzSl9uUD3syNd32fbvUvSy3SAmzzouflR-M1SUljvtMZesl2pS3mpRSblSnm1jh15-uWaLKzCmnUl62YOTeUwSvk-xb_gW95Qg2hzZ0IBBKij_IuwhyQcL6VSLl7gi6dIYBxmUxSSAIDPqiRrGkmMch-6DrbIQrxjW8Vio8zSpvXmPMB6B6WeagEM8QzN4fszy7uh2u7Ms_5LpQwnSyEi5vN1g5uG7QSiRRKU8PYsPYdnYcLqqJRdNUFch1oaFiHBGcVIDxJAI1QX0c; .AspNetCore.Antiforgery.9fXoN5jHCXs=CfDJ8MJ3Ss7oxyFKjxmon9SVuR_8J5wCWB-Zt_u9UBsB-1pDNksTO136-jcnpZhJontk_kZWd0l27-bxHK8g1XtVdaQ1pIp4V4w4XkWUnYbiwcYeUwEVoblGDOhK3mM4ou6CcPkijRbi1XVJcln4c0cUMxQ; ARRAffinity=012150c154c6bea60792ab57143f1b8e58d3e5118bedc36a64886b13b86a6f8e; ARRAffinitySameSite=012150c154c6bea60792ab57143f1b8e58d3e5118bedc36a64886b13b86a6f8e; .AspNetCore.Identity.Application=CfDJ8MJ3Ss7oxyFKjxmon9SVuR8HSlOuvNTpJM0OJHuN0h8G5z-BYgFtEIxSfSC3wrxi7jIEn3NagFguOEHg9XXP4ca-xlzyKzLAf3V7dOvr59ULsxOPFOhPh14_-s8EMOqbtaNo8KCBK2fUcQmxvDhRWeEPwCRWWLBJUKHNCaWXvnqcQry7joJAav_VLG0aADRjkP7pUjH5VAs_7U3hTpDBb5QbiJgoIJWUis8waNmB2Ws-w5NsKOUqduymjRiv5e3V-0jl5WYxjT498L3jZsW71xtAQ41j9v6jsojJ-tTfzStWVyurwSYiZFTPV-7t_qU6MqORo2TmOkZvzO71jB8Sk-1gQYR3tgJq9y9MZrj8vafnDobTj8A5WUgtPLVsbpCdLPKDDU2xE_c-JiL4ZCZxLBUOV6OuftLOgeQKXZGtjo9rOxVUWhQavUeHflv4be47W0YAbOkzh-9sUX6oVxQSCUcCbJYZkzt_Yz8Qp_PjOFtHto89UODia8tp-S4sxFAVBlyal-g_KhkVOt3JF2VXiBKAx4i2cO1CE5JzPVMaTd0M10KixJc-BK4PlfOKogn07tXnLNeo575j9XdEzzqYjt2t8h30WvEFxohGlHN-qeu5dUkZzY6jQVbfgZDK_NIKpH5jNT4JjipTyJjLzoMMrRsWBjuKUzVzn6AXbu-bDqFfW3FSjo-gqp9oHLo5J-3_mA"
	xsrf_token := "CfDJ8MJ3Ss7oxyFKjxmon9SVuR-5QFIkF-BRCoCz_ToYPEFw6aAx6CCa15JONNf6yWaT9V8hIsenKqKdscNc8iNZmsqc9TqisjLB9ATXaA7l4suS7wFqJiW066TGbV9VnoOoU2N3VN4MwfAwjeHvA_BjBiMikgP9J1aydaFKUkkkgLaD5L4kKtWE3IcGJFX29zhYog"

	r.Header.Add("Accept", "*/*")
	r.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	r.Header.Add("Accept-Language", "th-TH,th;q=0.9,en;q=0.8")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Content-Length", "98")
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("Origin", "https://auspark.au.edu")
	r.Header.Add("Referer", "https://auspark.au.edu/Planner")
	r.Header.Add("Sec-Ch-Ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
	r.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	r.Header.Add("Sec-Ch-Ua-Platform", `"macOS"`)
	r.Header.Add("Sec-Fetch-Dest", "empty")
	r.Header.Add("Sec-Fetch-Mode", "cors")
	r.Header.Add("Sec-Fetch-Site", "same-origin")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	r.Header.Add("X-Requested-With", "XMLHttpRequest")
	r.Header.Add("Cookie", cookie)
	r.Header.Add("X-Xsrf-Token", xsrf_token)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))

	// post := &Post{}
	// derr := json.NewDecoder(res.Body).Decode(post)
	// if derr != nil {
	// 	fmt.Println("Error decoding response:", derr)
	// 	return
	// }

	// if post.StatusCode == "200" || post.StatusCode == "201" {
	// 	panic("DONE")
	// }

	// fmt.Println("Id:", post.Id)
	// fmt.Println("StatusCode:", post.StatusCode)
	// fmt.Println("Description:", post.Description)

	return
}
