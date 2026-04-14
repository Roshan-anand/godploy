<script lang="ts">
	import { api, axiosErr } from '@/axios';
	import { Button } from '@/components/ui/button';
	import * as Dialog from '@/components/ui/dialog';
	import { Input } from '@/components/ui/input';
	import { Label } from '@/components/ui/label';
	import * as Select from '@/components/ui/select';
	import { Textarea } from '@/components/ui/textarea';
	import { queryClient } from '@/query';
	import { createForm } from '@tanstack/svelte-form';
	import { createMutation, createQuery } from '@tanstack/svelte-query';
	import Icon from '@iconify/svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { toast } from 'svelte-sonner';
	import { z } from 'zod';
	import type { ServiceType } from '@/types';
	import { userState } from '@/store/user-state.svelte';
	import type { ServicePageUiState } from '@/components/services/context.svelte';

	const {
		pageUi
	}: {
		pageUi: ServicePageUiState;
	} = $props();

	interface Project {
		id: string;
		name: string;
	}

	interface CreateServiceResponse {
		id: string;
		type: ServiceType;
	}

	type GitProviderKey = 'github' | 'gitlab' | 'bitbucket';
	type GitRepoFetchStatus =
		| 'idle'
		| 'loading'
		| 'provider_not_connected'
		| 'no_repo_authorized'
		| 'loaded'
		| 'error';

	interface GitProviderOption {
		key: GitProviderKey;
		name: string;
		icon: string;
		api: string;
	}

	interface GithubRepo {
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
		provider: GitProviderKey;
	}

	const projectIdFromPath = $derived(page.params.id ?? '');
	const isProjectScoped = $derived(projectIdFromPath !== '');

	const serviceTypes = [
		{ value: 'app' as const, label: 'App Service' },
		{ value: 'psql' as const, label: 'PostgreSQL Service' }
	];

	const gitProviders: GitProviderOption[] = [
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

	const getProjectsQueryKey = () =>
		['projects', userState.currentOrg.id, 'service-create'] as const;

	const projectsQuery = createQuery(() => ({
		queryKey: getProjectsQueryKey(),
		queryFn: async () => {
			return api
				.get<Project[]>('/project/all', { params: { org_id: userState.currentOrg.id } })
				.then((res) => res.data);
		},
		enabled: !isProjectScoped && userState.currentOrg.id !== ''
	}));

	// Git selection is status-driven by backend response codes: 409 (provider missing), 204 (no repos), 200 (repo list).
	let selectedGitProvider = $state<GitProviderKey | ''>('');
	let selectedGitRepoId = $state('');
	let selectedGitRepoName = $state('');
	let gitRepoFetchStatus = $state<GitRepoFetchStatus>('idle');
	let gitRepoList = $state<GithubRepo[]>([]);

	const selectedGitProviderName = $derived(
		gitProviders.find((provider) => provider.key === selectedGitProvider)?.name ?? 'provider'
	);

	const getReposMutation = createMutation(() => ({
		mutationFn: async (provider: GitProviderOption): Promise<GetRepoResult> => {
			const response = await api.get<GithubRepo[]>(provider.api, {
				params: { org_id: userState.currentOrg.id },
				validateStatus: (status) => status === 200 || status === 204 || status === 409
			});

			return {
				status: response.status,
				repos: response.status === 200 ? response.data : [],
				provider: provider.key
			};
		},
		onSuccess: (result) => {
			if (result.status === 200 && result.repos.length > 0) {
				gitRepoList = result.repos;
				gitRepoFetchStatus = 'loaded';
				return;
			}

			gitRepoList = [];
			selectedGitRepoId = '';
			selectedGitRepoName = '';
			gitRepoFetchStatus =
				result.status === 409
					? 'provider_not_connected'
					: result.status === 204 || result.status === 200
						? 'no_repo_authorized'
						: 'error';
		},
		onError: (error) => {
			gitRepoList = [];
			selectedGitRepoId = '';
			selectedGitRepoName = '';
			gitRepoFetchStatus = 'error';
			axiosErr(error, 'Failed to fetch repositories');
		}
	}));

	const createServiceMutation = createMutation(() => ({
		mutationFn: async (payload: { type: ServiceType; body: Record<string, string> }) => {
			const url = payload.type === 'app' ? '/service/app' : '/service/psql';
			return api.post<CreateServiceResponse>(url, payload.body).then((res) => res.data);
		},
		onSuccess: async (createdService) => {
			await queryClient.invalidateQueries({ queryKey: ['services'] });
			pageUi.closeCreateDialog();
			resetGitRepoSelection();
			form.reset();

			toast.success('Service created successfully');
			goto(
				resolve(`/(core)/service/[service]?id=${createdService.id}`, {
					service: createdService.type
				})
			);
		},
		onError: (error) => axiosErr(error, 'Failed to create service')
	}));

	// TanStack Form handles one dynamic service form for both app and psql service creation.
	const form = createForm(() => ({
		defaultValues: {
			project_id: '',
			name: '',
			description: '',
			type: 'app',
			app_name: '',
			db_name: '',
			db_user: '',
			db_password: '',
			image: ''
		},
		onSubmit: ({ value }) => {
			const projectId = projectIdFromPath || value.project_id;
			if (projectId === '') {
				toast.error('Please select a project');
				return;
			}

			if (value.type === 'app') {
				if (selectedGitProvider === '') {
					toast.error('Please select a git provider');
					return;
				}

				if (selectedGitRepoId === '' || selectedGitRepoName === '') {
					toast.error('Please select a repository');
					return;
				}

				createServiceMutation.mutate({
					type: 'app',
					body: {
						project_id: projectId,
						name: value.name.trim(),
						description: value.description.trim(),
						app_name: value.app_name.trim(),
						git_provider: selectedGitProvider,
						git_repo_id: selectedGitRepoId,
						git_repo_name: selectedGitRepoName
					}
				});
				return;
			}

			createServiceMutation.mutate({
				type: 'psql',
				body: {
					project_id: projectId,
					name: value.name.trim(),
					description: value.description.trim(),
					app_name: value.app_name.trim(),
					db_name: value.db_name.trim(),
					db_user: value.db_user.trim(),
					db_password: value.db_password,
					image: value.image.trim()
				}
			});
		}
	}));

	function closeDialog() {
		if (createServiceMutation.isPending) return;
		pageUi.closeCreateDialog();
	}

	function resetGitRepoSelection() {
		selectedGitProvider = '';
		selectedGitRepoId = '';
		selectedGitRepoName = '';
		gitRepoFetchStatus = 'idle';
		gitRepoList = [];
	}

	function fetchGitRepos(provider: GitProviderOption) {
		if (provider.api === '' || userState.currentOrg.id === '' || createServiceMutation.isPending) return;

		selectedGitProvider = provider.key;
		selectedGitRepoId = '';
		selectedGitRepoName = '';
		gitRepoFetchStatus = 'loading';
		gitRepoList = [];
		getReposMutation.mutate(provider);
	}

	function onRepoSelect(repoId: string) {
		const repo = gitRepoList.find((r) => r.id.toString() === repoId);
		if (!repo) return;

		selectedGitRepoId = repoId;
		selectedGitRepoName = repo.full_name;
	}
</script>

<Dialog.Root bind:open={pageUi.createDialogOpen}>
	<Dialog.Portal>
		<Dialog.Overlay class="fixed inset-0 z-40 bg-black/40" />
		<Dialog.Content
			class="fixed z-50 top-1/2 left-1/2 w-[92vw] max-w-lg -translate-x-1/2 -translate-y-1/2 rounded-xl border bg-background p-5 shadow-lg"
		>
			<Dialog.Title class="text-lg font-semibold">Create Service</Dialog.Title>
			<Dialog.Description class="text-sm text-muted-foreground"
				>Create a new service for your project.</Dialog.Description
			>

			<form
				class="mt-4 space-y-4"
				onsubmit={(e) => {
					e.preventDefault();
					e.stopPropagation();
					form.handleSubmit();
				}}
			>
				{#if isProjectScoped}
					<input type="hidden" name="project_id" value={projectIdFromPath} />
				{/if}

				{#if !isProjectScoped}
					<form.Field
						name="project_id"
						validators={{ onChange: z.string().min(1, 'Project is required') }}
					>
						{#snippet children(field)}
							<div class="space-y-1.5">
								<Label for={field.name}>Project</Label>
								<Select.Root
									type="single"
									value={field.state.value}
									onValueChange={(value) => field.handleChange(value)}
								>
									<Select.Trigger class="w-full" id={field.name}>
										{field.state.value
											? projectsQuery.data?.find((project) => project.id === field.state.value)
													?.name
											: projectsQuery.isPending
												? 'Loading projects...'
												: 'Select project'}
									</Select.Trigger>
									<Select.Content>
										{#each projectsQuery.data ?? [] as project (project.id)}
											<Select.Item value={project.id} label={project.name} />
										{/each}
									</Select.Content>
								</Select.Root>
								{#if field.state.meta.errors.length}
									<p class="text-sm font-medium text-destructive">{field.state.meta.errors[0]}</p>
								{/if}
							</div>
						{/snippet}
					</form.Field>
				{/if}

				<form.Field
					name="name"
					validators={{ onChange: z.string().min(3, 'Service name must be at least 3 characters') }}
				>
					{#snippet children(field)}
						<div class="space-y-1.5">
							<Label for={field.name}>Service Name</Label>
							<Input
								id={field.name}
								placeholder="Payments Database"
								value={field.state.value}
								onblur={field.handleBlur}
								oninput={(e) => field.handleChange(e.currentTarget.value)}
								disabled={createServiceMutation.isPending}
							/>
							{#if field.state.meta.errors.length}
								<p class="text-sm font-medium text-destructive">{field.state.meta.errors[0]}</p>
							{/if}
						</div>
					{/snippet}
				</form.Field>

				<form.Field
					name="description"
					validators={{ onChange: z.string().min(1, 'Description is required') }}
				>
					{#snippet children(field)}
						<div class="space-y-1.5">
							<Label for={field.name}>Service Description</Label>
							<Textarea
								id={field.name}
								placeholder="What does this service do?"
								value={field.state.value}
								onblur={field.handleBlur}
								oninput={(e) => field.handleChange(e.currentTarget.value)}
								disabled={createServiceMutation.isPending}
							/>
							{#if field.state.meta.errors.length}
								<p class="text-sm font-medium text-destructive">{field.state.meta.errors[0]}</p>
							{/if}
						</div>
					{/snippet}
				</form.Field>

				<form.Field name="type">
					{#snippet children(field)}
						<div class="space-y-1.5">
							<Label for={field.name}>Service Type</Label>
							<Select.Root
								type="single"
								value={field.state.value}
								onValueChange={(value) => field.handleChange(value as ServiceType)}
							>
								<Select.Trigger class="w-full" id={field.name}>
									{serviceTypes.find((item) => item.value === field.state.value)?.label}
								</Select.Trigger>
								<Select.Content>
									{#each serviceTypes as item (item.value)}
										<Select.Item value={item.value} label={item.label} />
									{/each}
								</Select.Content>
							</Select.Root>
						</div>
					{/snippet}
				</form.Field>

				<form.Field
					name="app_name"
					validators={{ onChange: z.string().min(3, 'App name must be at least 3 characters') }}
				>
					{#snippet children(field)}
						<div class="space-y-1.5">
							<Label for={field.name}>App Name</Label>
							<Input
								id={field.name}
								placeholder="payments-db"
								value={field.state.value}
								onblur={field.handleBlur}
								oninput={(e) => field.handleChange(e.currentTarget.value)}
								disabled={createServiceMutation.isPending}
							/>
							{#if field.state.meta.errors.length}
								<p class="text-sm font-medium text-destructive">{field.state.meta.errors[0]}</p>
							{/if}
						</div>
					{/snippet}
				</form.Field>

				<form.Subscribe selector={(state) => state.values.type}>
					{#snippet children(currentType)}
						{#if currentType === 'app'}
							<div class="space-y-2">
								<Label>Git</Label>
								<div class="flex items-center gap-3 w-full">
									{#each gitProviders as provider (provider.key)}
										<Button
											type="button"
											variant="outline"
											disabled={provider.api === '' ||
												userState.currentOrg.id === '' ||
												getReposMutation.isPending ||
												createServiceMutation.isPending}
											onclick={() => fetchGitRepos(provider)}
											class="flex-1"
										>
											<Icon icon={provider.icon} width="20" height="20" />
											<p>{provider.name}</p>
										</Button>
									{/each}
								</div>

								{#if gitRepoFetchStatus === 'loading'}
									<p class="text-sm text-muted-foreground">Loading repositories...</p>
								{:else if gitRepoFetchStatus === 'provider_not_connected'}
									<p class="text-sm text-muted-foreground">No {selectedGitProviderName} connected</p>
								{:else if gitRepoFetchStatus === 'no_repo_authorized'}
									<p class="text-sm text-muted-foreground">No repo authorized.</p>
								{:else if gitRepoFetchStatus === 'loaded'}
									<div class="space-y-1.5">
										<Label for="git-repo-select">Repository</Label>
										<Select.Root
											type="single"
											value={selectedGitRepoId}
											onValueChange={onRepoSelect}
										>
											<Select.Trigger class="w-full" id="git-repo-select">
												{selectedGitRepoName || 'Select repository'}
											</Select.Trigger>
											<Select.Content>
												{#each gitRepoList as repo (repo.id)}
													<Select.Item value={repo.id.toString()} label={repo.full_name} />
												{/each}
											</Select.Content>
										</Select.Root>
									</div>
								{:else if gitRepoFetchStatus === 'error'}
									<p class="text-sm text-destructive">Failed to fetch repositories.</p>
								{/if}
							</div>
						{/if}

						{#if currentType === 'psql'}
							<form.Field
								name="db_name"
								validators={{ onChange: z.string().min(1, 'Database name is required') }}
							>
								{#snippet children(field)}
									<div class="space-y-1.5">
										<Label for={field.name}>Database Name</Label>
										<Input
											id={field.name}
											placeholder="payments"
											value={field.state.value}
											onblur={field.handleBlur}
											oninput={(e) => field.handleChange(e.currentTarget.value)}
											disabled={createServiceMutation.isPending}
										/>
										{#if field.state.meta.errors.length}
											<p class="text-sm font-medium text-destructive">
												{field.state.meta.errors[0]}
											</p>
										{/if}
									</div>
								{/snippet}
							</form.Field>

							<form.Field
								name="db_user"
								validators={{ onChange: z.string().min(1, 'Database user is required') }}
							>
								{#snippet children(field)}
									<div class="space-y-1.5">
										<Label for={field.name}>Database User</Label>
										<Input
											id={field.name}
											placeholder="postgres"
											value={field.state.value}
											onblur={field.handleBlur}
											oninput={(e) => field.handleChange(e.currentTarget.value)}
											disabled={createServiceMutation.isPending}
										/>
										{#if field.state.meta.errors.length}
											<p class="text-sm font-medium text-destructive">
												{field.state.meta.errors[0]}
											</p>
										{/if}
									</div>
								{/snippet}
							</form.Field>

							<form.Field
								name="db_password"
								validators={{ onChange: z.string().min(1, 'Database password is required') }}
							>
								{#snippet children(field)}
									<div class="space-y-1.5">
										<Label for={field.name}>Database Password</Label>
										<Input
											id={field.name}
											type="password"
											placeholder="********"
											value={field.state.value}
											onblur={field.handleBlur}
											oninput={(e) => field.handleChange(e.currentTarget.value)}
											disabled={createServiceMutation.isPending}
										/>
										{#if field.state.meta.errors.length}
											<p class="text-sm font-medium text-destructive">
												{field.state.meta.errors[0]}
											</p>
										{/if}
									</div>
								{/snippet}
							</form.Field>

							<form.Field
								name="image"
								validators={{ onChange: z.string().min(1, 'Image is required') }}
							>
								{#snippet children(field)}
									<div class="space-y-1.5">
										<Label for={field.name}>Image</Label>
										<Input
											id={field.name}
											placeholder="postgres:16"
											value={field.state.value}
											onblur={field.handleBlur}
											oninput={(e) => field.handleChange(e.currentTarget.value)}
											disabled={createServiceMutation.isPending}
										/>
										{#if field.state.meta.errors.length}
											<p class="text-sm font-medium text-destructive">
												{field.state.meta.errors[0]}
											</p>
										{/if}
									</div>
								{/snippet}
							</form.Field>
						{/if}
					{/snippet}
				</form.Subscribe>

				<form.Subscribe
					selector={(state) => ({ canSubmit: state.canSubmit, isSubmitting: state.isSubmitting })}
				>
					{#snippet children(state)}
						<div class="flex justify-end gap-2 pt-1">
							<Button
								variant="outline"
								type="button"
								onclick={closeDialog}
								disabled={createServiceMutation.isPending}
							>
								Cancel
							</Button>
							<Button
								type="submit"
								disabled={!state.canSubmit || state.isSubmitting || createServiceMutation.isPending}
							>
								{state.isSubmitting || createServiceMutation.isPending ? 'Creating...' : 'Create'}
							</Button>
						</div>
					{/snippet}
				</form.Subscribe>
			</form>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
