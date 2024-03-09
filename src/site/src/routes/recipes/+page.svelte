<script lang="ts">
	import { page } from '$app/stores';
	import { fly } from 'svelte/transition';
	import {
		LanguageDropDown,
		Languages,
		Recipe,
		Recipes,
		SignalDropDown,
		Signals
	} from '$lib/common/types';
	import {
		filteredSamples,
		setFromUrl,
		selectedRecipeId,
		resetSearch,
		selectedSample,
		selectedLanguage,
		selectedSignal
	} from '$lib/store/store';
	import RecipeSelector from '$lib/_components/RecipeSelector.svelte';
	import RecipeSteps from '$lib/_components/RecipeSteps.svelte';
	import { onDestroy, onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(() => {
		const language = $page.url.searchParams.get('language');
		const signal = $page.url.searchParams.get('signal');
		const sample = $page.url.searchParams.get('recipe');
		setFromUrl(language, signal, sample);
	});

	const selectedRecipeStore$ = selectedSample.subscribe((recipe: Recipe) => {
		if (recipe.id !== 'none') {
			$page.url.searchParams.set('recipe', recipe.id);
			goto(`?${$page.url.searchParams.toString()}`);
		}
	});

	const selectedLanguageStore$ = selectedLanguage.subscribe((lang: LanguageDropDown) => {
		if (lang.id !== Languages.none.id) {
			$page.url.searchParams.set('language', lang.id);
			goto(`?${$page.url.searchParams.toString()}`);
		}
	});

	const selectedSignalStore$ = selectedSignal.subscribe((signal: SignalDropDown) => {
		if (signal.id !== Signals.none.id) {
			$page.url.searchParams.set('signal', signal.id);
			goto(`?${$page.url.searchParams.toString()}`);
		}
	});

	onDestroy(() => {
		selectedRecipeStore$;
		selectedLanguageStore$;
		selectedSignalStore$;
		// resetSearch();
	});
</script>

<svelte:head>
	<title>OTel Recipes - Catalog</title>
</svelte:head>

<div class="container">
	<RecipeSelector />
	{#if $filteredSamples.length === 0 && $selectedRecipeId === Recipes.none.id}
		<section class="section" in:fly={{ x: 100, duration: 300 }}>
			<div class="container">
				<div class="columns is-5 is-variable is-vcentered ml-5">
					<div class="column is-half">
						<div class="content">
							<h1 class="title is-3 is-spaced">Explore away!</h1>
							<!-- <hr class="is-line-involved" /> -->
							<p class="subtitle is-5">Browse the collection of recipes using the filters above.</p>
							<p class="subtitle is-">Missing a recipe for a particular programming language?</p>
							<p class="subtitle is-5">
								You can either open a feature request over on <a
									href="https://github.com/joaopgrassi/otel-recipes/labels/feature-request"
								>
									GitHub
								</a>
								, or send your contributions directly. 🚀
							</p>
						</div>
					</div>
					<div class="column">
						<img src="space-discovery.svg" alt="Girl on laptop chilling with cute dogue" />
					</div>
				</div>
			</div>
		</section>
	{/if}
	<RecipeSteps />
</div>
