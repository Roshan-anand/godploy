import { createForm } from '@tanstack/svelte-form';
import { z } from 'zod';
import type { LoginMutation } from '../auth.api';

// Keeps login form contract and submit mapping beside login route.
export interface LoginFormValues {
	email: string;
	password: string;
	rememberMe: boolean;
}

export function createLoginForm(login: LoginMutation) {
	const defaultValues: LoginFormValues = {
		email: '',
		password: '',
		rememberMe: false
	};

	return createForm(() => ({
		defaultValues,
		onSubmit: async ({ value }) => {
			login.mutate({ email: value.email, password: value.password });
		}
	}));
}

export const loginFieldValidators = {
	email: {
		onChange: z.email('Please enter a valid email')
	},
	password: {
		onChange: z.string().min(8, 'Password must be at least 8 characters')
	}
} as const;
