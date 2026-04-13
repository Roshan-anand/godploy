<script lang="ts">
	import type { PsqlService, ServiceBase } from '@/types';

	let { service } = $props<{ service: ServiceBase & PsqlService }>();
</script>

<div class="rounded-lg border bg-card text-card-foreground shadow-sm p-5 space-y-3">
	<div>
		<h1 class="text-xl font-semibold">{service.name}</h1>
		<p class="text-sm uppercase text-muted-foreground">{service.type}</p>
	</div>

	<p class="text-sm text-muted-foreground">{service.description || 'No description'}</p>

	<div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
		<div>
			<p class="text-muted-foreground">Service ID</p>
			<p class="font-medium break-all">{service.id}</p>
		</div>
		<div>
			<p class="text-muted-foreground">App Name</p>
			<p class="font-medium">{service.app_name}</p>
		</div>
		<div>
			<p class="text-muted-foreground">Database Name</p>
			<p class="font-medium">{service.db_name}</p>
		</div>
		<div>
			<p class="text-muted-foreground">Database User</p>
			<p class="font-medium">{service.db_user}</p>
		</div>
		<div>
			<p class="text-muted-foreground">Image</p>
			<p class="font-medium break-all">{service.image}</p>
		</div>
		<div>
			<p class="text-muted-foreground">Internal URL</p>
			<p class="font-medium break-all">{service.internal_url || 'N/A'}</p>
		</div>
	</div>
</div>
