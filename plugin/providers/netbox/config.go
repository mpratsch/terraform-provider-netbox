package netbox

import (
	"log"

	api "github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client"
)

// Config provides the configuration for the NETBOX provider.
type Config struct {
	// The application ID required for API requests. This needs to be created
	// in the NETBOX console. It can also be supplied via the NETBOX_APP_ID
	// environment variable.
	AppID string

	// The API endpoint. This defaults to http://localhost/api, and can also be
	// supplied via the NETBOX_ENDPOINT_ADDR environment variable.
	Endpoint string

	// If no Vlan is assigned to a prefixes you need to set the value to 0
	UseVlan int
}

type ProviderNetboxClient struct {
	client        *client.NetBox
	configuration Config
}

// ProviderNetboxClient is a structure that contains the client connections
// necessary to interface with the Go-Netbox API
//type ProviderNetboxClient struct {
//		client *Client
//}

func (c *Config) Client() (interface{}, error) {
	log.Printf("[DEBUG] config.go Client() AppID: %s", c.AppID)
	log.Printf("[DEBUG] config.go Client() Endpoint: %s", c.Endpoint)
	log.Printf("[DEBUG] config.go Client() UseVlan: %s", c.UseVlan)
	cfg := Config{
		AppID:    c.AppID,
		Endpoint: c.Endpoint,
		UseVlan:  c.UseVlan,
	}
	log.Printf("[DEBUG] config.go Initializing Netbox controllers")
	// sess := session.NewSession(cfg)
	// Create the Client
	cli := api.NewNetboxWithAPIKey(cfg.Endpoint, cfg.AppID)

	// Validate that our connection is okay
	if err := c.ValidateConnection(cli); err != nil {
		log.Printf("[DEBUG] config.go Client() Error %s", err)
		return nil, err
	}
	cs := ProviderNetboxClient{
		client:        cli,
		configuration: cfg,
	}
	return &cs, nil
}

// ValidateConnection ensures that we can connect to Netbox early, so that we
// do not fail in the middle of a TF run if it can be prevented.
func (c *Config) ValidateConnection(sc *client.NetBox) error {
	log.Printf("[DEBUG] config.go ValidateConnection() valitation ")
	rs, err := sc.Dcim.DcimRacksList(nil, nil)
	log.Println(rs)
	return err
}
