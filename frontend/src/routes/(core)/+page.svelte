<script>
	import Button from '@/components/ui/button/button.svelte';

	const githuRedirect = () => (window.location.href = '/api/provider/github/app/create');
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<Button onclick={githuRedirect} variant="secondary">connect</Button>
