package k8s

import (
	"context"
	"encoding/base64"
	"flea/flea"
	"fmt"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetUserTokenData(fleaconfig *flea.FleaConfig, tokenName string, client kubernetes.Interface) (string, error) {
	userToken, err := client.CoreV1().Secrets(fleaconfig.FleaK8sUserTokenNamespace).Get(context.Background(),
		tokenName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("Error getting certificate-authority-data: %v", err)
	}

	log.Info("K8s user token data recieved.")
	return string(userToken.Data["token"]), nil
}

func GetClusterCA(fleaconfig *flea.FleaConfig, client kubernetes.Interface) (string, error) {
	CA, err := client.CoreV1().ConfigMaps(fleaconfig.FleaK8sCANamespace).Get(context.Background(),
		fleaconfig.FleaK8sCAName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("Error getting CA: %v", err)
	}

	return base64.StdEncoding.EncodeToString([]byte(CA.Data["ca.crt"])), nil
}

func InitK8sConnect(context, kubeconfigPath string) (*kubernetes.Clientset, error) {
	k, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("error init k8s connect: %v", err)
	}

	return kubernetes.NewForConfig(k)
}
