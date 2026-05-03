import { api, axiosErr } from '@/axios';
import { createQuery } from '@tanstack/svelte-query';
import { getServiceState } from './store.svelte';
import type { GithubApp, ServiceListResponse } from './type';
import { getUserState } from '../global/store.svelte';

export const getGithubAppsQueryKey = (orgId: string) => ['github-apps', orgId] as const;
export const getOrgServicesQueryKey = (orgId: string) => ['services', 'org', orgId] as const;

export function useGetServicesQuery() {
	const { currentOrg } = getUserState();
	return createQuery(() => ({
		queryKey: ['services', currentOrg],
		queryFn: async () => {
			return api.get<ServiceListResponse>('/service').then((res) => res.data.services);
		}
	}));
}

export function useGithubAppsQuery() {
	const featureState = getServiceState();
	const { currentOrg } = getUserState();

	return createQuery(() => ({
		queryKey: ['github-app', currentOrg],
		queryFn: async () => {
			try {
				const response = await api.get<GithubApp[] | null>('/provider/github/app/list');
				const apps = response.data ?? [];
				featureState.githubApps = apps;
				return apps;
			} catch (error) {
				featureState.githubApps = [];
				const err = error instanceof Error ? error : new Error('Failed to load GitHub apps');
				axiosErr(err, 'Failed to load GitHub apps');
				return [];
			}
		},
		enabled: false
	}));
}
