<script lang="ts">
	import { page } from '$app/stores';
	import RecipeSelector from '$lib/components/recipe-selector/RecipeSelector.svelte';
	import { selectedLanguage, selectedSignal, setSelectedLanguage } from '$lib/store/store';

	const lang = $page.params.lang;
	setSelectedLanguage(lang);

	const signal = $page.params.signal;
	$: selectedSignal.set(signal);
</script>

<RecipeSelector />
