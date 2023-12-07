package main

import (
	"context"
	"flea/flea"
	"flea/generator"
	"flea/k8s"
	"flea/store"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Info("Start kubeconfig generation...")
	fleaConfig := new(flea.FleaConfig)
	flea.LoadConfig(fleaConfig)
	ctx := context.Background()

	vaultClient, err := store.InitVaultConnect(ctx, fleaConfig)
	if err != nil {
		log.Fatal(err)
	}

	k, err := k8s.InitK8sConnect(fleaConfig.KubeconfigCurrentContext, flea.SetKubeconfigPath(fleaConfig))
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

	log.Infof("Get kubeconfig: %s/ui/vault/secrets/kv/list/common/kubeconfig/", viper.GetString("VAULT_ADDR"))
	log.Info("Kubeconfig generation is done.")
}
