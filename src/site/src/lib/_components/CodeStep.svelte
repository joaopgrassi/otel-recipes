<script lang="ts">
	import { onMount } from 'svelte';
	import Highlight from 'svelte-highlight';
	import 'svelte-highlight/styles/github-dark-dimmed.css';
	import type { javascript } from 'svelte-highlight/languages/javascript';
	import { copyToClipboard } from '../_utils/utils';

	import { LanguageDropDown, Languages, Recipe, Recipes, Step } from '$lib/common/types';
	import LoadingStep from './LoadingStep.svelte';
	import type { LanguageType } from 'svelte-highlight/languages';

	export let step: Step;
	export let sample: Recipe;
	export let language: LanguageDropDown;

	let highlightJsLang: LanguageType<string>;
	let code = '';
	let isLoading = false;

	onMount(async () => {
		isLoading = true;
		if (!sample || sample.id === Recipes.none.id) {
			return;
		}
		highlightJsLang = await getHighlightJsLang();
		const res = await fetch(step.source);
		code = await res.text();
		isLoading = false;
	});

	/**
	 * Loads the HighlightJs language to use in the code snippet.
	 * This circumvents dynamic imports limitations:
	 * https://github.com/rollup/plugins/tree/master/packages/dynamic-import-vars#limitations
	 */
	async function getHighlightJsLang(): Promise<LanguageType<string>> {
		switch (language.id) {
			case Languages.csharp.id:
				return (await import('svelte-highlight/languages/csharp')).default;
			case Languages.go.id:
				return (await import('svelte-highlight/languages/go')).default;
			case Languages.java.id:
				return (await import('svelte-highlight/languages/java')).default;
			case Languages.js.id:
				return (await import('svelte-highlight/languages/javascript')).default;
			case Languages.python.id:
				return (await import('svelte-highlight/languages/python')).default;
		}
	}
</script>

{#if isLoading}
	<div class="block has-text-centered">
		<LoadingStep />
	</div>
{/if}

{#if code && !isLoading}
	<nav class="level mt-2" aria-label="Sample step navigation bar">
		<div class="level-left">
			<div class="level-item">
				<h1 class="title is-5 has-text-grey-lighter">{step.displayName}</h1>
			</div>
		</div>
		<div class="level-right">
			<button class="button is-small is-link is-right" on:click={() => copyToClipboard(code)}>
				Copy
			</button>
		</div>
	</nav>
	<Highlight language={highlightJsLang} {code} />
{/if}
