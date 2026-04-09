import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { api } from '@/axios';
import { userState } from '@/store/user-state.svelte';
import { createMutation } from '@tanstack/svelte-query';

export interface LoginPayload {
	email: string;
	password: string;
}

export interface RegisterPayload {
	name: string;
	email: string;
	password: string;
	org_name: string;
}

export interface Organization {
	id: string;
	name: string;
}

export interface AuthResponse {
	message: string;
	name: string;
	email: string;
	org_id: string;
	org_name: string;
}

export function useLoginMutation() {
	return createMutation(() => ({
		mutationFn: (payload: LoginPayload) =>
			api.post<AuthResponse>('/auth/login', payload).then((res) => res.data),
		onSuccess: (data) => {
			userState.name = data.name;
			userState.email = data.email;
			goto(resolve('/'));
		}
	}));
}

export function useRegisterMutation() {
	return createMutation(() => ({
		mutationFn: (payload: RegisterPayload) =>
			api.post<AuthResponse>('/auth/register', payload).then((res) => res.data),
		onSuccess: (data) => {
			userState.name = data.name;
			userState.email = data.email;
			goto(resolve('/'));
		}
	}));
}
