import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { userState } from '@/store/userState.svelte';
import { createMutation, createQuery } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';

// Route-local data module for create-service dialog query and mutation contracts.
export type ServiceType = 'app' | 'psql';
export type GitProviderKey = 'github' | 'gitlab' | 'bitbucket';

export interface Project {
	id: string;
	name: string;
}

export interface CreateServiceResponse {
	id: string;
	type: ServiceType;
}

export interface GitProviderOption {
	key: GitProviderKey;
	name: string;
	icon: string;
	api: string;
}

interface ApiMessageRes {
	message: string;
}

export interface GithubApp {
	name: string;
	app_id: number;
	created_at: string;
}

export interface GithubRepo {
	id: number;
	name: string;
	full_name: string;
	html_url: string;
	private: boolean;
	default_branch: string;
}

interface GetRepoResult {
	status: number;
	repos: GithubRepo[];
	message: string;
	provider: GitProviderKey;
}

interface CreateAppServiceBody {
	project_id: string;
	name: string;
	description: string;
	app_name: string;
	git_provider: GitProviderKey;
	git_repo_id: string;
	git_repo_name: string;
	git_branch: string;
	build_path: string;
}

interface CreatePsqlServiceBody {
	project_id: string;
	name: string;
	description: string;
	app_name: string;
	db_name: string;
	db_user: string;
	db_password: string;
	image: string;
}

type CreateServicePayload =
	| { type: 'app'; body: CreateAppServiceBody }
	| { type: 'psql'; body: CreatePsqlServiceBody };

export const serviceTypes = [
	{ value: 'app' as const, label: 'App Service' },
	{ value: 'psql' as const, label: 'PostgreSQL Service' }
];

export const gitProviders: GitProviderOption[] = [
	{
		key: 'github',
		name: 'Github',
		icon: 'meteor-icons:github',
		api: '/provider/github/repo/list'
	},
	{
		key: 'gitlab',
		name: 'GitLab',
		icon: 'material-icon-theme:gitlab',
		api: ''
	},
	{
		key: 'bitbucket',
		name: 'BitBucket',
		icon: 'material-icon-theme:bitbucket',
		api: ''
	}
];

export function getProjectsQueryKey() {
	return ['projects', userState.currentOrg.id, 'service-create'] as const;
}

export function createProjectsQuery(isProjectScoped: () => boolean) {
	return createQuery(() => ({
		queryKey: getProjectsQueryKey(),
		queryFn: () => api.get<Project[]>('/project/all').then((res) => res.data),
		enabled: !isProjectScoped() && userState.currentOrg.id !== ''
	}));
}

export function getGithubAppsQueryKey() {
	return ['github-apps', userState.currentOrg.id] as const;
}

export function createGithubAppsQuery(setGithubApps: (apps: GithubApp[]) => void) {
	return createQuery(() => ({
		queryKey: getGithubAppsQueryKey(),
		queryFn: async () => {
			try {
				const response = await api.get<GithubApp[] | null>('/provider/github/app/list');
				const apps = response.data ?? [];
				setGithubApps(apps);
				return apps;
			} catch (error) {
				const err = error instanceof Error ? error : new Error('Failed to load GitHub apps');
				setGithubApps([]);
				axiosErr(err, 'Failed to load GitHub apps');
				return [];
			}
		},
		enabled: false
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

export function createServiceCreateMutation(onCreated: () => void) {
	return createMutation(() => ({
		mutationFn: async (payload: CreateServicePayload) => {
			const url = payload.type === 'app' ? '/service/app' : '/service/psql';
			return api.post<CreateServiceResponse>(url, payload.body).then((res) => res.data);
		},
		onSuccess: async (createdService) => {
			await queryClient.invalidateQueries({ queryKey: ['services'] });
			onCreated();

			toast.success('Service created successfully');
			goto(
				resolve(`/(core)/service/[service]?id=${createdService.id}`, {
					service: createdService.type
				})
			);
		},
		onError: (error) => axiosErr(error, 'Failed to create service')
	}));
}
