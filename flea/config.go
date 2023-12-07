package flea

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var mainEnvs = []string{
	"ENV",
	"KUBECONFIG_CLUSTER_ENDPOINT",
	"KUBECONFIG_CLUSTER_NAME",
	"KUBECONFIG_CURRENT_CONTEXT",
	"VAULT_TOKEN",
}

type FleaConfig struct {
	Env                       string `mapstructure:"ENV"`
	FleaKubeconfigPath        string `mapstructure:"FLEA_KUBECONFIG_PATH"`
	FleaK8sUserTokenNamespace string `mapstructure:"FLEA_K8S_USER_TOKEN_NAMESPACE"`
	FleaK8sUserTokenName      string `mapstructure:"FLEA_K8S_USER_TOKEN_NAME"`
	FleaK8sCANamespace        string `mapstructure:"FLEA_K8S_CA_NAMESPACE"`
	FleaK8sCAName             string `mapstructure:"FLEA_K8S_CA_NAME"`
	KubeconfigCurrentContext  string `mapstructure:"KUBECONFIG_CURRENT_CONTEXT"`
	KubeconfigUserName        string `mapstructure:"KUBECONFIG_USER_NAME"`
	KubeconfigClusterEndpoint string `mapstructure:"KUBECONFIG_CLUSTER_ENDPOINT"`
	KubeconfigClusterName     string `mapstructure:"KUBECONFIG_CLUSTER_NAME"`
	VaultAddr                 string `mapstructure:"VAULT_ADDR"`
	VaultToken                string `mapstructure:"VAULT_TOKEN"`
}

func LoadConfig(config *FleaConfig) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	err := isEnvExist(mainEnvs)
	if err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()
	viper.SetDefault("FLEA_K8S_USER_TOKEN_NAMESPACE", "default")
	viper.SetDefault("FLEA_K8S_USER_TOKEN_NAME", "developer-token")
	viper.SetDefault("FLEA_K8S_CA_NAMESPACE", "kube-system")
	viper.SetDefault("FLEA_K8S_CA_NAME", "kube-root-ca.crt")
	viper.SetDefault("KUBECONFIG_CURRENT_CONTEXT", viper.GetString("KUBECONFIG_CLUSTER_NAME"))
	viper.SetDefault("KUBECONFIG_USER_NAME", "developer")
	viper.SetDefault("VAULT_ADDR", "https://vault.offline.shelopes.com")

	viper.BindEnv("ENV")
	viper.BindEnv("FLEA_KUBECONFIG_PATH")
	viper.BindEnv("FLEA_K8S_USER_TOKEN_NAMESPACE")
	viper.BindEnv("FLEA_K8S_USER_TOKEN_NAME")
	viper.BindEnv("FLEA_K8S_CA_NAMESPACE")
	viper.BindEnv("FLEA_K8S_CA_NAME")
	viper.BindEnv("KUBECONFIG_CURRENT_CONTEXT")
	viper.BindEnv("KUBECONFIG_USER_NAME")
	viper.BindEnv("KUBECONFIG_CLUSTER_ENDPOINT")
	viper.BindEnv("KUBECONFIG_CLUSTER_NAME")
	viper.BindEnv("VAULT_ADDR")
	viper.BindEnv("VAULT_TOKEN")

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func isEnvExist(envs []string) error {
	for _, env := range envs {
		if _, ok := os.LookupEnv(env); !ok {
			return fmt.Errorf("Env %v not set.", env)
		}
	}
	return nil
}

func SetKubeconfigPath(config *FleaConfig) string {
	if config.FleaKubeconfigPath != "" {
		return config.FleaKubeconfigPath
	}
	return "./vault/" + config.Env + "/.kube/config"
}
