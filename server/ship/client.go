package ship

import (
	"github.com/dcrodman/archon/data"
	"github.com/dcrodman/archon/encryption"
	"github.com/dcrodman/archon/packets"
	"github.com/dcrodman/archon/server"
	"github.com/dcrodman/archon/server/internal"
	"github.com/dcrodman/archon/server/internal/relay"
)

// Client2 implementation for the SHIP server.
type Client struct {
	cs *server.ConnectionState

	clientCrypt *encryption.PSOCrypt
	serverCrypt *encryption.PSOCrypt

	server.CommonClient

	account *data.Account
}

func (c Client) ConnectionState() *server.ConnectionState { return c.cs }

func (c *Client) clientVector() []uint8 { return c.clientCrypt.Vector }
func (c *Client) serverVector() []uint8 { return c.serverCrypt.Vector }

func (c Client) Encrypt(bytes []byte, length uint32) {
	c.serverCrypt.Encrypt(bytes, length)
}

func (c Client) Decrypt(bytes []byte, length uint32) {
	c.clientCrypt.Decrypt(bytes, length)
}

func (c Client) DebugInfo() map[string]interface{} {
	return map[string]interface{}{
		"server_type": "ship",
	}
}

func (c *Client) send(packet interface{}) error {
	return relay.SendPacket(c, packet, packets.BBHeaderSize)
}

func (c *Client) sendRaw(packet interface{}) error {
	bytes, size := internal.BytesFromStruct(packet)
	return relay.SendRaw(c, bytes, uint16(size))
}
