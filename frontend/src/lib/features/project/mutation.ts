import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { createMutation } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';
import { getProjectsQueryKey } from './query';
import type { ApiRes, CreateProjectPayload, DeleteProjectPayload, Project } from './types';

// Project mutations keep the list cache in sync so page state updates immediately.
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
		onError: (error) => axiosErr(error, 'Failed to create project')
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
		onError: (error) => axiosErr(error, 'Failed to delete project'),
		onSettled
	}));
}
