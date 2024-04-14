package main

import (
	"context"
	"flea/flea"
	"flea/generator"
	"flea/k8s"
	"flea/store"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var version = "1.1.0"

func main() {
	log.Infof("Flea version: %s", version)
	log.Info("Start kubeconfig generation...")

	fleaConfig := new(flea.FleaConfig)
	flea.LoadConfig(fleaConfig)
	ctx := context.Background()

	vaultClient, err := store.InitVaultConnect(ctx, fleaConfig)
	if err != nil {
		log.Fatal(err)
	}

	if fleaConfig.FleaKubeconfigVault {
		d, err := store.GetFromVault(ctx, fleaConfig.Env+"/.kube/config", "value", vaultClient)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("./kubeconfig", []byte(d), 0600)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Kubeconfig taken from Vault.")
	}

	k, err := k8s.InitK8sConnect(fleaConfig.KubeconfigCurrentContext, flea.SetKubeconfigPath(fleaConfig))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("./kubeconfig")
	if err != nil {
		log.Fatal(err)
	}

	kubeconfig, err := generator.GenerateKubeconfig(generator.KubeconfigTemplate, fleaConfig, k)
	if err != nil {
		log.Fatal(err)
	}

	err = store.PutToVault(ctx, fleaConfig, kubeconfig, vaultClient)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Link to get the kubeconfig: %s/ui/vault/secrets/kv/list/common/kubeconfig/", viper.GetString("VAULT_ADDR"))
	log.Info("Kubeconfig generation is done.")
}
