package client

import (
	"io/ioutil"
	"net"

	"github.com/ritwik310/mini-db/shell"
	"github.com/ritwik310/mini-db/src"
)

// New returns a new Client, that interacts with the server
func New(network, address string) Client {
	return Client{
		Network: network,
		Address: address,
	}
}

// Client contains CRUD methods
type Client struct {
	Network string
	Address string
}

// Get queries data from server
func (c *Client) Get(key string) (map[string]interface{}, error) {
	// New TCP-connection to the server
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	// Writing GET message/command to server
	fstr := shell.FormatCmd("GET " + key) // Format data according to teh protocol
	conn.Write([]byte(fstr))

	return handleResponse(&conn) // Returning server response
}

// Set inserts data in server
func (c *Client) Set(key, val, typ string) (map[string]interface{}, error) {
	// New TCP-connection to the server
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	// Writing SET message/command to server
	fstr := shell.FormatCmd("SET " + key + " " + val + " --" + typ)
	// fmt.Println("fstr", fstr)
	conn.Write([]byte(fstr))

	return handleResponse(&conn) // Returning server response
}

// Delete deletes key from the server
func (c *Client) Delete(key string) (map[string]interface{}, error) {
	// New TCP-connection to the server
	conn, err := net.Dial(c.Network, c.Address)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer conn.Close()

	// Writing DELETE message/command to server
	fstr := shell.FormatCmd("DELETE	 " + key)
	conn.Write([]byte(fstr))

	return handleResponse(&conn) // Returning server response
}

// handleResponse recieves the response message from connection
// and returns the data as map[string]interface{}
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
