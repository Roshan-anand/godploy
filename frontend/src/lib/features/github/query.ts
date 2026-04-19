import { api } from '@/axios';
import { userState } from '@/store/userState.svelte';
import { createQuery } from '@tanstack/svelte-query';
import type { GithubApp } from './types';

type GithubAppsQueryOptions = {
	enabled?: boolean;
	onSuccess?: (apps: GithubApp[]) => void;
	onError?: (error: unknown) => void;
};

// GitHub app listings are shared between the integrations page and the service picker dialog.
export function getGithubAppsQueryKey() {
	return ['github-apps', userState.currentOrg.id] as const;
}

export function createGithubAppsQuery({
	enabled = userState.currentOrg.id !== '',
	onSuccess,
	onError
}: GithubAppsQueryOptions = {}) {
	return createQuery(() => ({
		queryKey: getGithubAppsQueryKey(),
		queryFn: () =>
			api.get<GithubApp[] | null>('/provider/github/app/list').then((res) => res.data ?? []),
		enabled,
		onSuccess,
		onError
	}));
}
