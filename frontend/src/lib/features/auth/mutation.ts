import { goto } from '$app/navigation';
import { resolve } from '$app/paths';
import { api } from '@/axios';
import { userState } from '@/store/userState.svelte';
import { createMutation } from '@tanstack/svelte-query';
import type { AuthResponse, LoginPayload, RegisterPayload } from './types';

// Login and register mutations both update user state and transition into the app shell.
export function createLoginMutation() {
	return createMutation(() => ({
		mutationFn: (payload: LoginPayload) =>
			api.post<AuthResponse>('/auth/login', payload).then((res) => res.data),
		onSuccess: (data) => {
			userState.setUser(data);
			goto(resolve('/'));
		}
	}));
}

export function createRegisterMutation() {
	return createMutation(() => ({
		mutationFn: (payload: RegisterPayload) =>
			api.post<AuthResponse>('/auth/register', payload).then((res) => res.data),
		onSuccess: (data) => {
			userState.setUser(data);
			goto(resolve('/'));
		}
	}));
}

export type LoginMutation = ReturnType<typeof createLoginMutation>;
export type RegisterMutation = ReturnType<typeof createRegisterMutation>;
