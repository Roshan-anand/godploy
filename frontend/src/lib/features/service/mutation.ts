import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { api, axiosErr } from '@/axios';
import { queryClient } from '@/query';
import { createMutation } from '@tanstack/svelte-query';
import { toast } from 'svelte-sonner';
import type { CreateServicePayload, CreateServiceResponse } from './types';

// Service creation invalidates the service list and routes into the newly created service.
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
