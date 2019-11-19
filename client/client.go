/*
MiniDB-Client contains methods to programmatically interacts with the MiniDB-Server.

*/

package client

import (
	"io/ioutil"
	"net"

	"github.com/ritsource/mini-db/shell"
	"github.com/ritsource/mini-db/src"
)

// New returns a new Client, that contains Get, Set, and Delete function
// to interact with the MiniDB-Server running on the specified address.
func New(network, address string) Client {
	return Client{
		Network: network,
		Address: address,
	}
}

// Client contains Get, Set, and Delete methods necessary
// for interacting with the corrosponding MiniDB-Server and manipulating data
type Client struct {
	Network string // Network type (TCP)
	Address string // Server address
}

// Get creates a TCP-connection to the corrosponding MiniDB-Server
// and queries data from it given a valid key
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

// Set creates a TCP-connection to the MiniDB-Server and inserts data into it
// This method has three arguements, key, value and type
// typ indicates the data type of value. By default its a string,
// but it also supports integer and binary,
// Read more about data types, https://github.com/ritsource/mini-db#data-types
// Read more about type declaration in client, https://github.com/ritsource/mini-db/blob/master/client/README.md#data-type-declaration
func (c *Client) Set(key, val, typ string) (map[string]interface{}, error) {
	// typ = 'str' for string (default, if provided "")
	// typ = 'int' for integer
	// typ = 'bin' for binary

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

// Delete creates a TCP-connection to the server and deletes data from it
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
