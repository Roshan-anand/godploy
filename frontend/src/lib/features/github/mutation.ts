import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { createMutation } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';
import { getGithubAppsQueryKey } from './query';
import type { GetRepoResult, GitProviderOption, GithubApp, GithubRepo } from './types';

type ApiMessageRes = {
	message: string;
};

// GitHub mutations cover app deletion and repo discovery for the service creation flow.
export function createDeleteGithubAppMutation() {
	return createMutation(() => ({
		mutationFn: (payload: { app_id: number }) =>
			api.delete('/provider/github/app', { data: payload }).then((res) => res.data),
		onSuccess: (_res, payload) => {
			queryClient.setQueryData(
				getGithubAppsQueryKey(),
				(cachedApps: GithubApp[] | null | undefined) => {
					if (!cachedApps) return null;

					const remainingApps = cachedApps.filter((app) => app.app_id !== payload.app_id);
					return remainingApps.length > 0 ? remainingApps : null;
				}
			);

			toast.success('GitHub app deleted successfully');
		},
		onError: (error) => axiosErr(error, 'Failed to delete github app')
	}));
}

export function createGetReposMutation(setGithubRepos: (repos: GithubRepo[]) => void) {
	return createMutation(() => ({
		mutationFn: async ({
			provider,
			appId
		}: {
			provider: GitProviderOption;
			appId: number;
		}): Promise<GetRepoResult> => {
			const response = await api.get<GithubRepo[] | ApiMessageRes>(provider.api, {
				params: { app_id: appId },
				validateStatus: (status) => status === 200 || status === 204 || status === 409
			});

			return {
				status: response.status,
				repos: response.status === 200 ? (response.data as GithubRepo[]) : [],
				message: response.status === 409 ? ((response.data as ApiMessageRes)?.message ?? '') : '',
				provider: provider.key
			};
		},
		onSuccess: (result) => {
			setGithubRepos(result.repos);

			if (result.status === 409) {
				toast.error(result.message || 'No github connected');
			}
		},
		onError: (error) => {
			setGithubRepos([]);
			axiosErr(error, 'Failed to fetch repositories');
		}
	}));
}
