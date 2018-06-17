package testing

import (
	"strings"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/nodes"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateNode(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateSuccessfully(t)

	createOpts := nodes.CreateOpts{
		ClusterID: "e395be1e-8d8e-43bb-bd6c-943eccf76a6d",
		Metadata: map[string]interface{}{
			"foo": "bar",
			"test": map[string]interface{}{
				"nil_interface": interface{}(nil),
				"float_value":   float64(123.3),
				"string_value":  "test_string",
				"bool_value":    false,
			},
		},
		Name:      "node-e395be1e-002",
		ProfileID: "d8a48377-f6a3-4af4-bbbb-6e8bcaa0cbc0",
		Role:      "",
	}

	res := nodes.Create(fake.ServiceClient(), createOpts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-Openstack-Request-Id")
	th.AssertEquals(t, "req-3791a089-9d46-4671-a3f9-55e95e55d2b4", requestID)

	location := res.Header.Get("Location")
	th.AssertEquals(t, "http://senlin.cloud.blizzard.net:8778/v1/actions/ffd94dd8-6266-4887-9a8c-5b78b72136da", location)

	locationFields := strings.Split(location, "actions/")
	actionID := locationFields[1]
	th.AssertEquals(t, "ffd94dd8-6266-4887-9a8c-5b78b72136da", actionID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedCreate, *actual)
}
