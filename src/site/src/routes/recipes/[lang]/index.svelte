<script lang="ts">
	import { page } from '$app/stores';
	import RecipeSelector from '$lib/components/recipe-selector/RecipeSelector.svelte';
	import { setSelectedLanguage } from '$lib/store/store';

	const lang = $page.params.lang;
	setSelectedLanguage(lang);
</script>

<RecipeSelector />
