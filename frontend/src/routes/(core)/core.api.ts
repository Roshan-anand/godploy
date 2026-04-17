import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { userState } from '@/store/userState.svelte';
import { createMutation, createQuery } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';

// Route-local query/mutation module for the core projects page.
export interface Project {
	id: string;
	name: string;
	description: string;
}

interface CreateProjectPayload {
	project_name: string;
	description: string;
}

interface DeleteProjectPayload {
	id: string;
}

interface ApiRes {
	message: string;
}

export function getProjectsQueryKey() {
	return ['projects', userState.currentOrg.id] as const;
}

export function createProjectsQuery() {
	return createQuery(() => ({
		queryKey: getProjectsQueryKey(),
		queryFn: () => api.get<Project[]>('/project/all').then((res) => res.data),
		enabled: userState.currentOrg.id !== ''
	}));
}

export function createProjectCreateMutation(onCreate: () => void) {
	return createMutation(() => ({
		mutationFn: (payload: CreateProjectPayload) =>
			api.post<Project>('/project', payload).then((res) => res.data),
		onSuccess: (createdProject) => {
			queryClient.setQueryData(getProjectsQueryKey(), (cachedProjects: Project[] | undefined) => {
				if (!cachedProjects) return [createdProject];
				return [createdProject, ...cachedProjects];
			});

			onCreate();
			toast.success('Project created successfully');
		},
		onError: (error) => axiosErr(error, 'Faild to create project')
	}));
}

export function createProjectDeleteMutation(
	onMutateProjectId: (projectId: string) => void,
	onSettled: () => void
) {
	return createMutation(() => ({
		mutationFn: (payload: DeleteProjectPayload) =>
			api.delete<ApiRes>('/project', { data: payload }).then((res) => res.data),
		onMutate: (payload) => {
			onMutateProjectId(payload.id);
		},
		onSuccess: (res, payload) => {
			queryClient.setQueryData(getProjectsQueryKey(), (cachedProjects: Project[] | undefined) => {
				if (!cachedProjects) return [];
				return cachedProjects.filter((project) => project.id !== payload.id);
			});

			toast.success(res.message || 'Project deleted successfully');
		},
		onError: (error) => axiosErr(error, 'Faild to delete project'),
		onSettled
	}));
}
