<script lang="ts">
	import { page } from '$app/stores';
	import RecipeSelector from '$lib/components/RecipeSelector.svelte';
	import { setSelectedLanguage, setSelectedSignal } from '$lib/store/store';
	import { onMount } from 'svelte';

	onMount(() => {
		const lang = $page.params.lang;
		setSelectedLanguage(lang);

		const signal = $page.params.signal;
		setSelectedSignal(signal);
	});
</script>

<RecipeSelector />
