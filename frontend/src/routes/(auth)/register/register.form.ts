import { createForm } from '@tanstack/svelte-form';
import { z } from 'zod';
import type { RegisterMutation } from '../auth.api';

// Keeps register form contract and submit mapping beside register route.
export interface RegisterFormValues {
	name: string;
	email: string;
	password: string;
	organisation: string;
	rememberMe: boolean;
}

export function createRegisterForm(register: RegisterMutation) {
	const defaultValues: RegisterFormValues = {
		name: '',
		email: '',
		password: '',
		organisation: '',
		rememberMe: false
	};

	return createForm(() => ({
		defaultValues,
		onSubmit: async ({ value }) => {
			register.mutate({
				name: value.name,
				email: value.email,
				password: value.password,
				org_name: value.organisation
			});
		}
	}));
}

export const registerFieldValidators = {
	name: {
		onChange: z.string().min(2, 'Name must be at least 2 characters')
	},
	email: {
		onChange: z.email('Please enter a valid email')
	},
	password: {
		onChange: z.string().min(8, 'Password must be at least 8 characters')
	},
	organisation: {
		onChange: z.string().min(2, 'Organisation must be at least 2 characters')
	}
} as const;
