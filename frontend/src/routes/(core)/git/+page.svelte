<script lang="ts">
	import { api, axiosErr } from '@/axios';
	import { queryClient } from '@/query';
	import { userState } from '@/store/user-state.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import Icon from '@iconify/svelte';
	import { createMutation, createQuery } from '@tanstack/svelte-query';
	import { toast } from 'svelte-sonner';

	interface GithubApp {
		name: string;
		app_id: number;
		created_at: string;
	}

	const providers = [
		{
			name: 'Github',
			icon: 'meteor-icons:github',
			redirect: '/api/provider/github/app/create'
		},
		{
			name: 'GitLab',
			icon: 'material-icon-theme:gitlab',
			redirect: ''
		},
		{
			name: 'BitBucket',
			icon: 'material-icon-theme:bitbucket',
			redirect: ''
		}
	];

	const getGithubAppsQueryKey = () => ['github-apps', userState.currentOrg.id] as const;

	const getGithubAppsQuery = createQuery(() => ({
		queryKey: getGithubAppsQueryKey(),
		queryFn: () => api.get<GithubApp[] | null>('/provider/github/app/list').then((res) => res.data),
		enabled: userState.currentOrg.id !== ''
	}));

	const deleteGithubAppMutation = createMutation(() => ({
		mutationFn: (payload: { app_id: number }) =>
			api.delete('/provider/github/app', { data: payload }).then((res) => res.data),
		onSuccess: (_res, payload) => {
			queryClient.setQueryData(
				getGithubAppsQueryKey(),
				(cachedApps: GithubApp[] | null | undefined) => {
					if (!cachedApps) return null;

					const remainingApps = cachedApps.filter((app) => app.app_id !== payload.app_id);
					return remainingApps.length > 0 ? remainingApps : null;
				}
			);

			toast.success('Github app deleted successfully');
		},
		onError: (error) => axiosErr(error, 'Failed to delete github app')
	}));

	const providerRedirect = (loc: string) => (window.location.href = loc);

	const formatCreatedAt = (createdAt: string) => {
		const parsedDate = new Date(createdAt);
		if (Number.isNaN(parsedDate.getTime())) return createdAt;
		return parsedDate.toLocaleString();
	};
</script>

<section class="p-2">
	<h1 class="my-2">Connect any git provider</h1>

	<section class="flex items-center gap-4 w-full">
		{#each providers as p (p)}
			<Button
				id={p.name}
				variant="outline"
				disabled={p.redirect == ''}
				onclick={() => providerRedirect(p.redirect)}
				class="flex-1"
			>
				<Icon icon={p.icon} width="24" height="24" />
				<p>{p.name}</p>
			</Button>
		{/each}
	</section>
</section>

<hr class="my-3" />

<section class="flex-1">
	{#if getGithubAppsQuery.isPending && userState.currentOrg.id !== ''}
		<p class="text-muted-foreground">Loading provider details...</p>
	{:else if getGithubAppsQuery.isError}
		<p class="text-destructive">Failed to load provider details.</p>
	{:else if !getGithubAppsQuery.data || getGithubAppsQuery.data.length === 0}
		<div class="flex items-center gap-2 text-muted-foreground size-full justify-center">
			<Icon icon="material-icon-theme:git" width="24" height="24" />
			<p>No provider connected</p>
		</div>
	{:else}
		<div class="space-y-3">
			{#each getGithubAppsQuery.data as app (app.app_id)}
				<div class="rounded-lg border bg-card text-card-foreground shadow-sm p-4 space-y-3">
					<div class="flex items-start justify-between gap-3">
						<div class="flex items-center gap-2">
							<Icon icon="meteor-icons:github" width="24" height="24" />
							<div>
								<h2 class="font-semibold">GitHub</h2>
								<p class="text-sm text-muted-foreground">{app.name}</p>
							</div>
						</div>

						<Button
							variant="destructive"
							size="sm"
							onclick={() => deleteGithubAppMutation.mutate({ app_id: app.app_id })}
							disabled={deleteGithubAppMutation.isPending}
						>
							{deleteGithubAppMutation.isPending ? 'Deleting...' : 'Delete'}
						</Button>
					</div>

					<div class="text-sm text-muted-foreground space-y-1">
						<span class="font-medium text-foreground">Created:</span>
						{formatCreatedAt(app.created_at)}
					</div>
				</div>
			{/each}
		</div>
	{/if}
</section>
