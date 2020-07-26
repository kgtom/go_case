
1.工厂模式模式 newXXX

~~~
func addHeaders(r *http.Request) {
	req.Header.Add("x-api-key", "my-secret-token")
}

// somewhere in your code

req, err := http.NewRequest(/* ... */)
if err != nil {
	return nil
}

addHeaders(req)

// do your job here
~~~
2.包装器模式

~~~
type httpClient struct {
	c        http.Client
	apiToken string
}

func (c *httpClient) Get(url string) (resp *Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err

	}

	return c.Do(req)

}

func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)

}

func (c *httpClient) Do(req *Request) (*Response, error) {
	req.Header.Add("x-api-key", c.apiToken)
	return c.c.Do(req)

}
~~~
https://developer20.com/add-header-to-every-request-in-go/
