package framework

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	serverCert = `-----BEGIN CERTIFICATE-----
MIIFvTCCA6WgAwIBAgIUBpS47ArkUC0MXYK3LvXU3eRh/CowDQYJKoZIhvcNAQEL
BQAwUjELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDELMAkGA1UEAwwCY2EwHhcNMjQxMTE1
MTcxNjI1WhcNMjcxMTE1MTcxNjI1WjByMQswCQYDVQQGEwJVUzELMAkGA1UECAwC
UEExFTATBgNVBAcMDFBoaWxhZGVscGhpYTETMBEGA1UECgwKTGlub2RlIExMQzEU
MBIGA1UECwwLTGlub2RlIFRlc3QxFDASBgNVBAMMC2xpbm9kZS50ZXN0MIICIjAN
BgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA27JxXbiT+0aptSvE2uFakJQf+qwT
5mGFCNaQFRBDdxxLTUF6UyglZflT9KdVtJ9kmsyZj9vhFnxszWnoUK1Y/runOYTf
JlNBVp28fO43HrUtaHFCZncvu0C4Tdc09721p2pP5nhgXv8BtZeDAVY/hjSIGgP1
1WNLSWP2eZn4+q4hr7iUqVqLRYVz5e489b1sEXpCiSDWuq6GWRzvEBquHX0F82mW
84DMfa2TrcG4bw0i0r4nKWcgB3at7sR32DvEPFsFiEreFgNsx7b1KcG+ngzA3ZKL
9bviQKSLjjn48VPoV/w5lT3PYGIAjwu2tbNY8J6dUcni4aHnIwhwBFVb5299eIEC
nccueVExw8LtXBYOUKT4A8doKy3ZBq4B+WY8N0QhE6H8tuLrAl6IUh8rduuvJc38
+QIDD6IKr58zuest6q0/lNvjruOfUMa+EsBPX795wyDuqL4tUyfySyUyYNXcQ4ip
2nFTBYXoB75jLsXHULhOC+7AbxzWeM76mjeNgKzUJaz+1EUMLYOSsfiYFMlWfoiL
ilf7WMdR3bLHccFAA/Qg3CZETU/B20amYDI/+0TvY1td01gzoUx3UjDPB6mpntgr
DoTISDNAvZgPOt9ebs7AEM6/iHgIQtAnCQULTzQ48i3WZlpPYb2IeWOsNCXiOZPN
+STXedL5M3IUwUcCAwEAAaNrMGkwJwYDVR0RBCAwHoILbGlub2RlLnRlc3SCD3d3
dy5saW5vZGUudGVzdDAdBgNVHQ4EFgQUgNqzhL/JpxllvFu18qvlg/usDrEwHwYD
VR0jBBgwFoAUC2AMOf90/zpuQ588rPLfe7EukIUwDQYJKoZIhvcNAQELBQADggIB
AL38v8A0Yfi3Qcr7JtMJ+EOgiHo+W1PW05CAKrswqIZGb9pLwcc46N1ICX4/wItH
DfOmiLHEJ+eEaf07XWy1G+orvqsz6FLh2lfr1cne2DH1udiBXw2VyHDeaighgqTX
rHPcV9lLPcRgQgE8AC2WSn3Rmjd4eU+twlqYcJTLt3cy+TulwXxGBjn7CSmRamRA
AaURnVpsMhw9baINrN6+3zbjw1LKpMO3JfPx9NPw0iUYYbUWFMli2RTEwdR0o9Fu
Om6ogyYHHLTUDv2+cHYY4TKJ0LGz9PGB3iwdGbSSpLadjV7xkFERio5B4o/FedLB
CuECSIoWqjScSrVWjpIpG6b7LVkuDI7ZrZ6Rvkwcv4Zezx5TkynQUw9EezEgGRQf
RiBSKoPGKJfRGiYGNXDjqENX3kxqt5cuVe/Z0czrb+2zOMfaTZwJtp2rrJqckxBh
CK4CXQz2nsfGRW/lyJ1Jyc+ul0obXXhynDBA9dE5woCIwgTCRL9M0ZOHjoQi1tDh
27i0j4YzIvlIDIi6iex/XVZi9mhuRvDR7f7c5RVpHsu38znCLyQetFnwOQOmIVZI
lEUQvU1Jnk+e5+RqvOcZ0ZcLppBa71XjUdYm56mzY1ph04n1VUO4rmaI3wNBETGd
jJ3K7XuBBL/YT+02AzsZR/0fiHLdA9DbLUdhtRs0mb5u
-----END CERTIFICATE-----`
	serverKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
)

func (i *lbInvocation) CreateTLSSecret(secretName string) (err error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},
		Data: map[string][]byte{
			corev1.TLSCertKey:       []byte(serverCert),
			corev1.TLSPrivateKeyKey: []byte(serverKey),
		},
		Type: corev1.SecretTypeTLS,
	}

	_, err = i.kubeClient.CoreV1().Secrets(i.Namespace()).Create(context.TODO(), secret, metav1.CreateOptions{})

	return err
}

func (i *lbInvocation) DeleteSecret(name string) error {
	err := i.kubeClient.CoreV1().Secrets(i.Namespace()).Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}
