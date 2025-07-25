package kubeconfigstrategy

import "context"

// Option specifies which strategy will be applied
type Option struct {
	// Secret specifies option for the kubeconfig strategy based on a Secret.
	Secret SecretStrategyOption
}

// New creates a new kubeconfig strategy based on the provided options.
func New(ctx context.Context, option Option) (Interface, error) {
	return newSecretKubeConfigStrategy(ctx, option.Secret)
}
