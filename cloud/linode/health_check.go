package linode

import (
	"context"
	"time"

	"github.com/linode/linode-cloud-controller-manager/cloud/linode/client"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type healthChecker struct {
	period       time.Duration
	linodeClient client.Client
	stopCh       chan<- struct{}
}

func newHealthChecker(apiToken string, timeout time.Duration, period time.Duration, stopCh chan<- struct{}) (*healthChecker, error) {
	client, err := client.New(apiToken, timeout)
	if err != nil {
		return nil, err
	}

	return &healthChecker{
		period:       period,
		linodeClient: client,
		stopCh:       stopCh,
	}, nil
}

func (r *healthChecker) Run(stopCh <-chan struct{}) {
	ctx := wait.ContextForChannel(stopCh)
	wait.Until(r.worker(ctx), r.period, stopCh)
}

func (r *healthChecker) worker(ctx context.Context) func() {
	return func() {
		r.do(ctx)
	}
}

func (r *healthChecker) do(ctx context.Context) {
	if r.stopCh == nil {
		klog.Errorf("stop signal already fired. nothing to do")
		return
	}

	authenticated, err := client.CheckClientAuthenticated(ctx, r.linodeClient)
	if err != nil {
		klog.Warningf("unable to determine linode client authentication status: %s", err.Error())
		return
	}

	if !authenticated {
		klog.Error("detected invalid linode api token: stopping controllers")

		close(r.stopCh)
		r.stopCh = nil
		return
	}

	klog.Info("linode api token is healthy")
}