<script lang="ts">
	import ServiceList from '@/components/services/service-list.svelte';
	import { page } from '$app/state';

	const projectId = $derived(page.params.id ?? '');
</script>

<ServiceList scopeType="project" scopeId={projectId} />
