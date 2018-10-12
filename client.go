package secure

import (
	"./crypt"
	//"github.com/stuarthu/secureserver/crypt"
	"net/http"
	"strconv"
)

type Client struct {
	crypt.Decryptor
	*http.Client
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("secure-type", c.Type())
	req.Header.Add("secure-publickey", c.PublicKey())
	req.Header.Add("secure-messagesize", strconv.Itoa(c.MessageSize()))
	resp, e := c.Client.Do(req)
	if e != nil {
		return nil, e
	}
	t := resp.Header.Get("secure-type")
	if t == c.Type() {
		resp.Body = c.Decryptor.NewDecryptedReader(resp.Body)
	}
	return resp, nil
}
