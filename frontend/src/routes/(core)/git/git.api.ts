import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { userState } from '@/store/userState.svelte';
import { createMutation, createQuery } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';

// Route-local provider/query/mutation contracts for git integrations page.
export interface GithubApp {
	name: string;
	app_id: number;
	created_at: string;
}

export interface GitProvider {
	name: string;
	icon: string;
	redirect: string;
}

export const providers: GitProvider[] = [
	{
		name: 'Github',
		icon: 'meteor-icons:github',
		redirect: '/api/provider/github/app/create'
	},
	{
		name: 'GitLab',
		icon: 'material-icon-theme:gitlab',
		redirect: ''
	},
	{
		name: 'BitBucket',
		icon: 'material-icon-theme:bitbucket',
		redirect: ''
	}
];

export function getGithubAppsQueryKey() {
	return ['github-apps', userState.currentOrg.id] as const;
}

export function createGithubAppsQuery() {
	return createQuery(() => ({
		queryKey: getGithubAppsQueryKey(),
		queryFn: () => api.get<GithubApp[] | null>('/provider/github/app/list').then((res) => res.data),
		enabled: userState.currentOrg.id !== ''
	}));
}

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

			toast.success('Github app deleted successfully');
		},
		onError: (error) => axiosErr(error, 'Failed to delete github app')
	}));
}

export function formatCreatedAt(createdAt: string) {
	const parsedDate = new Date(createdAt);
	if (Number.isNaN(parsedDate.getTime())) return createdAt;
	return parsedDate.toLocaleString();
}
