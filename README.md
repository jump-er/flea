# Flea

Flea предназначен для генерации готового к использованию kubeconfig (конфиг для kubectl).

Предполагается, что преимущественно генерироваться будут конфиги для команд разработки, потому что административные конфиги создаются на этапе создания кластера.

Flea работает на основании уже добавленных в k8s кдастер RBAC.

После генерации kubeconfig-а он складывается в Vault, откуда любой сможет его себе забрать.

## Конфигурация

Приложение конфигурируется переменными окружения.

`ENV` - окружение, для которого запускается flea

`FLEA_KUBECONFIG_VAULT` - включает использование Hashicorp Vault для получения kubeconfig, если false, то kubeconfig будет браться по локальному пути из FLEA_KUBECONFIG_PATH

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
