<script lang="ts">
	import { api } from '@/axios';
	import { Skeleton } from '@/components/ui/skeleton';
	import { createQuery } from '@tanstack/svelte-query';
	import { page } from '$app/state';

	type ServiceType = 'psql' | 'app';

	interface ServiceBase {
		id: string;
		project_id: string;
		type: ServiceType;
		service_id: string;
		name: string;
		app_name: string;
		description: string;
		created_at: string;
	}

	interface PsqlService extends ServiceBase {
		type: 'psql';
		db_name: string;
		db_user: string;
		db_password: string;
		image: string;
		internal_url: string;
	}

	interface AppService extends ServiceBase {
		type: 'app';
	}

	type ServiceDetails = PsqlService | AppService;

	const serviceId = $derived(page.params.id ?? '');
	const serviceType = $derived(page.url.searchParams.get('type') as ServiceType | null);

	// Discriminated union keeps rendering type-safe while selecting endpoint by service type.
	const serviceQuery = createQuery(() => ({
		queryKey: ['service-details', serviceId, serviceType],
		queryFn: async (): Promise<ServiceDetails> => {
			const url = serviceType === 'app' ? `/service/app/${serviceId}` : `/service/psql/${serviceId}`;
			return api.get<ServiceDetails>(url).then((res) => res.data);
		},
		enabled: serviceId !== '' && (serviceType === 'psql' || serviceType === 'app')
	}));
</script>

<section class="p-2 flex-1">
	{#if !serviceType}
		<p class="text-muted-foreground">Missing service type in URL</p>
	{:else if serviceQuery.isPending}
		<div class="rounded-lg border bg-card text-card-foreground shadow-sm p-4 space-y-3">
			<Skeleton class="h-6 w-1/3" />
			<Skeleton class="h-4 w-2/3" />
			<Skeleton class="h-4 w-1/2" />
		</div>
	{:else if serviceQuery.isError || !serviceQuery.data}
		<p class="text-red-500">Failed to load service details</p>
	{:else}
		<div class="rounded-lg border bg-card text-card-foreground shadow-sm p-5 space-y-3">
			<div>
				<h1 class="text-xl font-semibold">{serviceQuery.data.name}</h1>
				<p class="text-sm uppercase text-muted-foreground">{serviceQuery.data.type}</p>
			</div>

			<p class="text-sm text-muted-foreground">{serviceQuery.data.description || 'No description'}</p>

			<div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
				<div>
					<p class="text-muted-foreground">Service ID</p>
					<p class="font-medium break-all">{serviceQuery.data.id}</p>
				</div>
				<div>
					<p class="text-muted-foreground">App Name</p>
					<p class="font-medium">{serviceQuery.data.app_name}</p>
				</div>
			</div>

			{#if serviceQuery.data.type === 'psql'}
				<div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
					<div>
						<p class="text-muted-foreground">Database Name</p>
						<p class="font-medium">{serviceQuery.data.db_name}</p>
					</div>
					<div>
						<p class="text-muted-foreground">Database User</p>
						<p class="font-medium">{serviceQuery.data.db_user}</p>
					</div>
					<div>
						<p class="text-muted-foreground">Image</p>
						<p class="font-medium break-all">{serviceQuery.data.image}</p>
					</div>
					<div>
						<p class="text-muted-foreground">Internal URL</p>
						<p class="font-medium break-all">{serviceQuery.data.internal_url || 'N/A'}</p>
					</div>
				</div>
			{/if}
		</div>
	{/if}
</section>
