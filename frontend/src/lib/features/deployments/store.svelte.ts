// AI summary: Class-based context store for deployment tab UI state consumed by deployment feature hooks.
import { getContext, setContext } from 'svelte';

interface DeploymentsFeatureState {
	deletingDeploymentId: string;
}

class DeploymentsFeatureStateClass implements DeploymentsFeatureState {
	deletingDeploymentId = $state('');
}

const DEFAULT_KEY = 'deployments:feature:state';

export const getDeploymentsFeatureState = (key: string = DEFAULT_KEY) => {
	return getContext<DeploymentsFeatureState>(key);
};

export const setDeploymentsFeatureState = (key: string = DEFAULT_KEY) => {
	return setContext(key, new DeploymentsFeatureStateClass());
};
