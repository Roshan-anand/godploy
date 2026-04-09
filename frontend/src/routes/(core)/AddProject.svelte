<script lang="ts">
	import { Button } from '@/components/ui/button';
	import { Plus } from '@lucide/svelte';

	let { onclick, disabled = false }: { onclick?: () => void; disabled?: boolean } = $props();
</script>

<Button {disabled} {onclick}>
	<Plus />
	<p>Create</p>
</Button>
