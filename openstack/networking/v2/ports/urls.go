package ports

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
)

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("ports", id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("ports")
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func addAllowedAddressPairURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id) + "/add_allowed_address_pairs"
}

func removeAllowedAddressPairURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id) + "/remove_allowed_address_pairs"
}

// AddAllowedAddressPair accepts a UpdateOpts struct and updates an existing port using the
// values provided.
func AddAllowedAddressPair(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		r.Err = err
		return
	}
	for k := range h {
		if k == "If-Match" {
			h[k] = fmt.Sprintf("revision_number=%s", h[k])
		}
	}
	resp, err := c.Put(addAllowedAddressPairURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: h,
		OkCodes:     []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// RemoveAllowedAddressPair accepts a UpdateOpts struct and updates an existing port using the
// values provided.
func RemoveAllowedAddressPair(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	h, err := gophercloud.BuildHeaders(opts)
	if err != nil {
		r.Err = err
		return
	}
	for k := range h {
		if k == "If-Match" {
			h[k] = fmt.Sprintf("revision_number=%s", h[k])
		}
	}
	resp, err := c.Put(removeAllowedAddressPairURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		MoreHeaders: h,
		OkCodes:     []int{200, 201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
