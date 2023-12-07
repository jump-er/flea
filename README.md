# Flea

Flea предназначен для генерации готового к использованию kubeconfig (конфиг для kubectl).

Предполагается, что уже имеется ранее созданный конфиг с правами администратора.

Flea работает на основании уже добавленных в k8s кластер RBAC.

После генерации kubeconfig-а он складывается в Vault, откуда любой сможет его себе забрать (в дальнейшем планируется реализовать интерфейс, чтобы была возможность реализации своих хранилищ).

## Конфигурация

Приложение конфигурируется переменными окружения.

`ENV` - окружение, для которого запускается flea

`FLEA_KUBECONFIG_PATH` - путь к kubeconfig с административными правами, необходим для выполнения нужных операций в кластере (по умолчанию путь из vault)

`FLEA_K8S_USER_TOKEN_NAMESPACE` - неймспейс, в котором искать токен пользователя

`FLEA_K8S_USER_TOKEN_NAME` - название токена пользователя

`FLEA_K8S_CA_NAMESPACE` - неймспейс, в котором искать CA кластера

`FLEA_K8S_CA_NAME` - название CA кластера

`KUBECONFIG_CURRENT_CONTEXT` - значение `current-context`, которое подставится в kubeconfig

`KUBECONFIG_USER_NAME` - значение `users.name`, которое подставится в kubeconfig

`KUBECONFIG_CLUSTER_ENDPOINT` - значение `clusters.name.cluster.server`, которое подставится в kubeconfig

`KUBECONFIG_CLUSTER_NAME` - значение `clusters.name`, которое подставится в kubeconfig

`VAULT_ADDR` - HTTP адрес Vault

`VAULT_TOKEN` - токен для подключения в Vault
