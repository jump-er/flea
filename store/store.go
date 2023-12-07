package store

import (
	"context"
	"flea/flea"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	log "github.com/sirupsen/logrus"
)

func InitVaultConnect(ctx context.Context, fleaconfig *flea.FleaConfig) (*vault.Client, error) {
	vaultClient, err := vault.New(
		vault.WithAddress(fleaconfig.VaultAddr),
		vault.WithRequestTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	if err := vaultClient.SetToken(fleaconfig.VaultToken); err != nil {
		return nil, err
	}

	log.Info("Connection to the Vault successful.")
	return vaultClient, nil
}

func PutToVault(ctx context.Context, fleaconfig *flea.FleaConfig, kubeconfig string, client *vault.Client) error {
	configIdentity := "developer-" + fleaconfig.KubeconfigClusterName
	_, err := client.Secrets.KvV2Write(ctx, "common/kubeconfig/"+configIdentity, schema.KvV2WriteRequest{
		Data: map[string]any{
			configIdentity: kubeconfig,
		}},
		vault.WithMountPath("kv"),
	)
	if err != nil {
		return err
	}

	log.Info(configIdentity, " config was written to the Vault successfully.")
	return nil
}
