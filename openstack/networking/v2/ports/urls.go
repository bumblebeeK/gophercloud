package ports

import "github.com/gophercloud/gophercloud/v2"

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
