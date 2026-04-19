<script lang="ts">
	import { createRegisterMutation } from '@/features/auth/mutation';
	import { toast } from 'svelte-sonner';
	import AuthBranding from '@/components/auth-branding.svelte';
	import { Button } from '@/components/ui/button';
	import { Input } from '@/components/ui/input';
	import { Checkbox } from '@/components/ui/checkbox';
	import { Label } from '@/components/ui/label';
	import { createForm } from '@tanstack/svelte-form';
	import { resolve } from '$app/paths';
	import { z } from 'zod';
	import { getFieldErrMsg } from '@/utils';

	interface RegisterFormValues {
		name: string;
		email: string;
		password: string;
		organisation: string;
		rememberMe: boolean;
	}

	const defaultValues: RegisterFormValues = {
		name: '',
		email: '',
		password: '',
		organisation: '',
		rememberMe: false
	};

	const register = createRegisterMutation();
	const form = createForm(() => ({
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

	const registerFieldValidators = {
		name: {
			onBlur: z.string().min(2, 'Name must be at least 2 characters')
		},
		email: {
			onBlur: z.email('Please enter a valid email')
		},
		password: {
			onBlur: z.string().min(8, 'Password must be at least 8 characters')
		},
		organisation: {
			onBlur: z.string().min(2, 'Organisation must be at least 2 characters')
		}
	} as const;

	$effect(() => {
		if (register.isError) {
			toast.error(
				register.error?.message ?? 'An error occurred while registering. Please try again.'
			);
		}
	});
</script>

<div class="grid min-h-svh lg:grid-cols-2">
	<AuthBranding />
	<div class="flex items-center justify-center p-8">
		<div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-100">
			<div class="flex flex-col space-y-2 text-center">
				<h1 class="text-2xl font-semibold tracking-tight">Create an account</h1>
				<p class="text-sm text-muted-foreground">Enter your details below to create your account</p>
			</div>

			<form
				onsubmit={(e) => {
					e.preventDefault();
					e.stopPropagation();
					form.handleSubmit();
				}}
			>
				<div class="grid gap-4">
					<form.Field name="name" validators={registerFieldValidators.name}>
						{#snippet children(field)}
							<div class="grid gap-2">
								<Label for={field.name}>Name</Label>
								<Input
									id={field.name}
									name={field.name}
									type="text"
									placeholder="John Doe"
									value={field.state.value}
									onblur={field.handleBlur}
									oninput={(e) => field.handleChange(e.currentTarget.value)}
								/>
								{#if field.state.meta.errors.length}
									<p class="text-sm font-medium text-destructive">
										{getFieldErrMsg(field.state.meta.errors[0])}
									</p>
								{/if}
							</div>
						{/snippet}
					</form.Field>

					<form.Field name="email" validators={registerFieldValidators.email}>
						{#snippet children(field)}
							<div class="grid gap-2">
								<Label for={field.name}>Email</Label>
								<Input
									id={field.name}
									name={field.name}
									type="email"
									placeholder="name@example.com"
									value={field.state.value}
									onblur={field.handleBlur}
									oninput={(e) => field.handleChange(e.currentTarget.value)}
								/>
								{#if field.state.meta.errors.length}
									<p class="text-sm font-medium text-destructive">
										{getFieldErrMsg(field.state.meta.errors[0])}
									</p>
								{/if}
							</div>
						{/snippet}
					</form.Field>

					<form.Field name="password" validators={registerFieldValidators.password}>
						{#snippet children(field)}
							<div class="grid gap-2">
								<Label for={field.name}>Password</Label>
								<Input
									id={field.name}
									name={field.name}
									type="password"
									value={field.state.value}
									onblur={field.handleBlur}
									oninput={(e) => field.handleChange(e.currentTarget.value)}
								/>
								{#if field.state.meta.errors.length}
									<p class="text-sm font-medium text-destructive">
										{getFieldErrMsg(field.state.meta.errors[0])}
									</p>
								{/if}
							</div>
						{/snippet}
					</form.Field>

					<form.Field name="organisation" validators={registerFieldValidators.organisation}>
						{#snippet children(field)}
							<div class="grid gap-2">
								<Label for={field.name}>Organisation</Label>
								<Input
									id={field.name}
									name={field.name}
									type="text"
									placeholder="Acme Inc."
									value={field.state.value}
									onblur={field.handleBlur}
									oninput={(e) => field.handleChange(e.currentTarget.value)}
								/>
								{#if field.state.meta.errors.length}
									<p class="text-sm font-medium text-destructive">
										{getFieldErrMsg(field.state.meta.errors[0])}
									</p>
								{/if}
							</div>
						{/snippet}
					</form.Field>

					<form.Field name="rememberMe">
						{#snippet children(field)}
							<div class="flex items-center space-x-2 mt-2">
								<Checkbox
									id={field.name}
									checked={field.state.value}
									onchange={() => field.handleChange(!field.state.value)}
								/>
								<Label
									for={field.name}
									class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
								>
									Remember me
								</Label>
							</div>
						{/snippet}
					</form.Field>

					<form.Subscribe
						selector={(state) => ({
							canSubmit: state.canSubmit,
							isSubmitting: state.isSubmitting
						})}
					>
						{#snippet children(state)}
							<Button
								type="submit"
								class="w-full mt-2"
								disabled={!state.canSubmit || register.isPending}
							>
								{state.isSubmitting || register.isPending
									? 'Creating account...'
									: 'Create account'}
							</Button>
						{/snippet}
					</form.Subscribe>
				</div>
			</form>

			<p class="px-8 text-center text-sm text-muted-foreground">
				Already have an account?
				<Button variant="link" href={resolve('/login')}>Log in</Button>
			</p>
		</div>
	</div>
</div>
