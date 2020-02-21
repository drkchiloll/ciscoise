package ise

// Node ISE Deployment Server
type Node struct {
	Object
	Gateway        string   `json:"gateway"`
	DisplayName    string   `json:"displayName"`
	InDeployment   bool     `json:"inDeployment"`
	OtherPapFqdn   string   `json:"otherPapFqdn,omitempty"`
	IPAddrs        []string `json:"ipAddresses"`
	IPAddr         string   `json:"ipAddress"`
	NodeTypes      string   `json:"nodeServiceTypes"`
	FQDN           string   `json:"fqdn"`
	PrimaryPapNode bool     `json:"primaryPapNode"`
	PxGridNode     bool     `json:"pxGridNode"`
	ObjRef
}

// GetNodes retrieves list of all ISE Nodes in a Deployment
func (c *Client) GetNodes() {
	res, err := c.MakeReq("/node", "GET", ReqParams{}, nil)
	if err != nil {
	}
	defer res.Body.Close()
	type respData struct {
		SearchResult struct {
			Total     int64  `json:"total"`
			Resources []Node `json:"resources"`
			ObjRef
		} `json:"SearchResult"`
	}
}
