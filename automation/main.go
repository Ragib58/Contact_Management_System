package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// URL and headers
	url := "https://training.mygov.bd/application/apply?id=BDGS-1747039917"

	// Initialize CookieJar to handle cookies
	jar, _ := cookiejar.New(nil)

	// Create HTTP client with CookieJar
	client := &http.Client{
		Jar: jar,
	}

	// Headers
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("your form data here")))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://training.mygov.bd")
	req.Header.Set("Referer", "https://training.mygov.bd/services/form?id=BDGS-1747039917")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36")

	// Set cookies (use actual cookie values from your browser)
	cookies := []*http.Cookie{
		&http.Cookie{Name: "cf_clearance", Value: "b9gPuHam...<cut>"},
		&http.Cookie{Name: "XSRF-TOKEN", Value: "eyJpdiI6...<cut>"},
		&http.Cookie{Name: "mygov_v2_session", Value: "ANu1TTVCP4...<cut>"},
	}

	// Add cookies to the request
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// Form data (make sure to replace with actual data, similar to your cURL data)
	formData := url.Values{}
	formData.Set("_token", "IIKljfEdTRGA0nD0xJcRe0vP0LuCL2BAnTskeia6")
	formData.Set("office_layer", "3")
	formData.Set("office_origin", "17")
	formData.Set("office_id", "65")
	formData.Set("is_complete", "1")
	formData.Set("text-27e5719252f9", "Rajshahi")
	formData.Set("applicant_name", "মোঃ রাগীব শাহরিয়ার")
	formData.Set("text-4c8546b6d034e8", "321312312")
	formData.Set("text-c04af5f38aa0b8", "2433421")
	formData.Set("text-44cfd6fcc4cc4", "Rajshahi")
	formData.Set("text-004670af4966", "56344343")
	formData.Set("text-66073222b57868", "13-05-2025")
	formData.Set("usignature1", "data:image/png;base64,iVBORw0KGgoAAAANS...<cut>")
	formData.Set("text-04519375b5b15", "মোঃ রাগীব শাহরিয়ার")
	formData.Set("text-0cbba2330c4da", "01774896515")
	formData.Set("attachment[profile_photo]", "আবেদনকারীর ছবি.jpg")
	formData.Set("attachments[]", "b9addb6a1d9dd8504e3850976cc812f9")
	formData.Set("attachment[nid]", "জাতীয় পরিচয়পত্রের কপি.jpg")

	// Set form data as the request body
	req.Body = ioutil.NopCloser(bytes.NewBufferString(formData.Encode()))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print response
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
}
