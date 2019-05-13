package client

import (
	"io/ioutil"
	"net"

	"github.com/ritwik310/mini-db/shell"
	"github.com/ritwik310/mini-db/src"
)

func New(network, address string) Client {
	return Client{
		Network: network,
		Address: address,
	}
}

type Client struct {
	Network string
	Address string
}

func (c *Client) Get(key string) (map[string]interface{}, error) {
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	fstr := shell.FormatCmd("GET " + key)
	conn.Write([]byte(fstr))

	return handleResponse(&conn)
}

func (c *Client) Set(key, val, typ string) (map[string]interface{}, error) {
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	fstr := shell.FormatCmd("SET " + key + " " + val + " --" + typ)
	// fmt.Println("fstr", fstr)
	conn.Write([]byte(fstr))

	return handleResponse(&conn)
}

func (c *Client) Delete(key string) (map[string]interface{}, error) {
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	fstr := shell.FormatCmd("DELETE	 " + key)
	conn.Write([]byte(fstr))

	return handleResponse(&conn)
}

func handleResponse(conn *net.Conn) (map[string]interface{}, error) {
	bs, err := ioutil.ReadAll(*conn)
	if err != nil {
		return map[string]interface{}{}, err
	}

	d, err := src.UnmarshalData(bs)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return d, nil
}
