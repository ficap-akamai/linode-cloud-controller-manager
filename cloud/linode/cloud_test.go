package linode

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/linode/linodego"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudRouteControllerDisabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Setenv("LINODE_API_TOKEN", "dummyapitoken")
	t.Setenv("LINODE_REGION", "us-east")

	fake := newFake(t)
	ts := httptest.NewServer(fake)
	defer ts.Close()

	clientNewBak := clientNew

	clientNew = func(token string, _ time.Duration) (*linodego.Client, error) {
		userAgent := fmt.Sprintf("linode-cloud-controller-manager %s", linodego.DefaultUserAgent)

		linodeClient := linodego.NewClient(http.DefaultClient)
		linodeClient.SetBaseURL(ts.URL)

		linodeClient.SetUserAgent(userAgent)
		linodeClient.SetToken(token)

		return &linodeClient, nil
	}

	defer func() { clientNew = clientNewBak }()

	t.Run("should not fail if vpc is empty and routecontroller is disabled", func(t *testing.T) {
		Options.VPCName = ""
		Options.EnableRouteController = false
		_, err := newCloud()
		assert.NoError(t, err)
	})

	t.Run("fail if vpcname is empty and routecontroller is enabled", func(t *testing.T) {
		Options.VPCName = ""
		Options.EnableRouteController = true
		_, err := newCloud()
		assert.Error(t, err)
	})
}
