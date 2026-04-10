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
		created_at: string;
		github_app_id: string;
	}

	interface DeleteGithubAppPayload {
		org_id: string;
	}

	interface ApiRes {
		message: string;
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

	const getGithubAppQueryKey = () => ['github-app', userState.currentOrg.id] as const;

	const getGithubAppQuery = createQuery(() => ({
		queryKey: getGithubAppQueryKey(),
		queryFn: () =>
			api
				.get<GithubApp | null>('/provider/github/app', {
					params: { org_id: userState.currentOrg.id }
				})
				.then((res) => res.data),
		enabled: userState.currentOrg.id !== ''
	}));

	const deleteGithubAppMutation = createMutation(() => ({
		mutationFn: (payload: DeleteGithubAppPayload) =>
			api.delete<ApiRes>('/provider/github/app', { data: payload }).then((res) => res.data),
		onSuccess: (res) => {
			queryClient.setQueryData(getGithubAppQueryKey(), null);
			toast.success(res.message || 'Github app deleted successfully');
		},
		onError: (error) => axiosErr(error, 'Failed to delete github app')
	}));

	const providerRedirect = (loc: string) => (window.location.href = loc);

	const deleteGithubApp = () => {
		if (deleteGithubAppMutation.isPending || userState.currentOrg.id === '') return;

		deleteGithubAppMutation.mutate({ org_id: userState.currentOrg.id });
	};

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
	{#if getGithubAppQuery.isPending && userState.currentOrg.id !== ''}
		<p class="text-muted-foreground">Loading provider details...</p>
	{:else if getGithubAppQuery.isError || !getGithubAppQuery.data}
		<div class="flex items-center gap-2 text-muted-foreground size-full justify-center">
			<Icon icon="material-icon-theme:git" width="24" height="24" />
			<p>No provider connected</p>
		</div>
	{:else}
		<div class="rounded-lg border bg-card text-card-foreground shadow-sm p-4 space-y-3">
			<div class="flex items-start justify-between gap-3">
				<div class="flex items-center gap-2">
					<Icon icon="meteor-icons:github" width="24" height="24" />
					<div>
						<h2 class="font-semibold">GitHub</h2>
						<p class="text-sm text-muted-foreground">{getGithubAppQuery.data.name}</p>
					</div>
				</div>

				<Button
					variant="destructive"
					size="sm"
					onclick={deleteGithubApp}
					disabled={deleteGithubAppMutation.isPending}
				>
					{deleteGithubAppMutation.isPending ? 'Deleting...' : 'Delete'}
				</Button>
			</div>

			<div class="text-sm text-muted-foreground space-y-1">
				<span class="font-medium text-foreground">Created:</span>
				{formatCreatedAt(getGithubAppQuery.data.created_at)}
			</div>
		</div>
	{/if}
</section>
