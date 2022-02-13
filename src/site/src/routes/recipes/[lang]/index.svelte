<script lang="ts">
	import { page } from '$app/stores';
	import RecipeSelector from '$lib/components/recipe-selector/RecipeSelector.svelte';
	import { selectedLanguage } from '$lib/store/store';

	const lang = $page.params.lang;
	$: selectedLanguage.set(lang);
</script>

<RecipeSelector />
