// AI summary: Class-based context store for service-create flow state and callbacks used by service feature hooks.
import { getContext, setContext } from 'svelte';
import type { CreateServiceResponse, GithubApp, GithubRepo } from './type';

interface ServicesFeatureStore {
	githubApps: GithubApp[];
	githubRepos: GithubRepo[];
	afterCreateSuccess: (response: CreateServiceResponse) => Promise<void>;
	setAfterCreateSuccess: (fn: (response: CreateServiceResponse) => Promise<void>) => void;
}

class ServicesFeatureStoreClass implements ServicesFeatureStore {
	githubApps = $state<GithubApp[]>([]);
	githubRepos = $state<GithubRepo[]>([]);
	afterCreateSuccess = $state<(response: CreateServiceResponse) => Promise<void>>(async () => {});

	setAfterCreateSuccess(fn: (response: CreateServiceResponse) => Promise<void>) {
		this.afterCreateSuccess = fn;
	}
}

const DEFAULT_KEY = 'services:feature:state';

export const getServicesFeatureState = (key: string = DEFAULT_KEY) => {
	return getContext<ServicesFeatureStore>(key);
};

export const setServicesFeatureState = (key: string = DEFAULT_KEY) => {
	return setContext(key, new ServicesFeatureStoreClass());
};
