package generator

import (
	"bytes"
	"flea/flea"
	"flea/k8s"
	"html/template"

	"k8s.io/client-go/kubernetes"
)

const KubeconfigTemplate = `apiVersion: v1
kind: Config
current-context: {{.CurrentContext}}
contexts:
- name: {{.ClusterName}}
  context:
    cluster: {{.ClusterName}}
    user: {{.UserName}}
clusters:
- name: {{.ClusterName}}
  cluster:
    certificate-authority-data: {{.CertificateAuthorityData}}
    server: {{.ClusterEndpoint}}
users:
- name: {{.UserName}}
  user:
    token: {{.UserToken}}
`

type fleaKubeconfigTemplateParams struct {
	CurrentContext           string
	ClusterName              string
	CertificateAuthorityData string
	ClusterEndpoint          string
	UserName                 string
	UserToken                string
}

func GenerateKubeconfig(kubeconfigTemplate string, fleaconfig *flea.FleaConfig, client kubernetes.Interface) (string, error) {
	buf := new(bytes.Buffer)

	userTokenData, err := k8s.GetUserTokenData(fleaconfig, fleaconfig.FleaK8sUserTokenName, client)
	if err != nil {
		return "", err
	}

	clusterCAData, err := k8s.GetClusterCA(fleaconfig, client)
	if err != nil {
		return "", err
	}

	c := fleaKubeconfigTemplateParams{
		CurrentContext:           fleaconfig.KubeconfigCurrentContext,
		ClusterName:              fleaconfig.KubeconfigClusterName,
		CertificateAuthorityData: clusterCAData,
		ClusterEndpoint:          fleaconfig.KubeconfigClusterEndpoint,
		UserName:                 fleaconfig.KubeconfigUserName,
		UserToken:                userTokenData,
	}

	kubeconfigT := template.Must(template.New("kubeconfig").Parse(kubeconfigTemplate))
	err = kubeconfigT.Execute(buf, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
