<script lang="ts">
	import { page } from '$app/stores';
	import RecipeSelector from '$lib/_components/RecipeSelector.svelte';
	import { selectedSampleId, setSelectedLanguage, setSelectedSignal } from '$lib/store/store';

	const lang = $page.params.lang;
	setSelectedLanguage(lang);

	const signal = $page.params.signal;
	setSelectedSignal(signal);

	const sample = $page.params.sample;
	$: selectedSampleId.set(sample);
</script>

<RecipeSelector />
