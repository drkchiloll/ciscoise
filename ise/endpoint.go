package ise

import "encoding/json"

// Endpoint ISE endpoint
type Endpoint struct {
	Object
	MacAddr                 string `json:"mac,omitempty"`
	ProfileID               string `json:"profileId,omitempty"`
	StaticProfileAssignment bool   `json:"staticProfileAssignment,omitempty"`
	GroupID                 string `json:"groupId,omitempty"`
	StaticGroupAssignment   bool   `json:"staticGroupAssignment,omitempty"`
	PortalUser              string `json:"portalUser,omitempty"`
	IdentityStore           string `json:"identityStore,omitempty"`
	IdentityStoreID         string `json:"identityStoreId,omitempty"`
	CustomAttrs             struct {
		CustomAttrs map[string]string `json:"customAttributes,omitempty"`
	} `json:"customAttributes,omitempty"`
	ObjRef
}

// GetEndpoints retrieves a paginated list of EPs
func (c *Client) GetEndpoints(r ReqParams) ([]Endpoint, error) {
	res, err := c.MakeReq("/endpoint", "GET", r, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type respData struct {
		SearchResult struct {
			Total     int64      `json:"total"`
			Resources []Endpoint `json:"resources"`
			ObjRef
		} `json:"SearchResult"`
	}
	var epRes respData
	json.NewDecoder(res.Body).Decode(&epRes)
	return epRes.SearchResult.Resources, nil
}

// GetEndpoint retrieves ep by ep ID
func (c *Client) GetEndpoint(id string) (Endpoint, error) {
	res, err := c.MakeReq("/endpoint/"+id, "GET", ReqParams{}, nil)
	if err != nil {
		return Endpoint{}, err
	}
	defer res.Body.Close()
	type epResp struct {
		ERSEndpoint Endpoint `json:"ERSEndpoint"`
	}
	var ep epResp
	json.NewDecoder(res.Body).Decode(&ep)
	return ep.ERSEndpoint, nil
}
