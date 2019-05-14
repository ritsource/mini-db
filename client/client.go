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

// Client contains Get, Set, and Delete methods
// necessary for interacting with the corrosponding server and manipulating data
type Client struct {
	Network string // Network type (TCP)
	Address string // Server address
}

// Get creates a TCP-connection to the corrosponding server
// and queries data from it given a valid key
// Example: resp, err := mdb.Get("myname")
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

// Set creates a TCP-connection to the server
// and inserts data into it
// Example: resp, err := mdb.Set("myname", "Ritwik Saha", "str")
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

// Delete creates a TCP-connection to the server
// and deletes data from it
// Example: resp, err := mdb.Delete("myname")
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
