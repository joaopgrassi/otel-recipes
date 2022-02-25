<script lang="ts">
	import { onMount } from 'svelte';
	import Highlight from 'svelte-highlight';
	import 'svelte-highlight/src/styles/github-dark-dimmed.css';
	import type { HljsLanguage } from 'svelte-highlight/src/languages';
	import { copyToClipboard } from '../_utils/utils';

	import { Language, Languages, Sample, Samples, Step } from '$lib/common/types';

	export let step: Step;
	export let sample: Sample;
	export let language: Language;

	let highlightJsLang: HljsLanguage;
	let code = '';

	onMount(async () => {
		if (!sample || sample.id === Samples.none.id) {
			return;
		}
		highlightJsLang = await getHighlightJsLang();
		const res = await fetch(step.source);
		code = await res.text();
	});

	/**
	 * Loads the HighlightJs language to use in the code snippet.
	 * This circumvents dynamic imports limitations:
	 * https://github.com/rollup/plugins/tree/master/packages/dynamic-import-vars#limitations
	 */
	async function getHighlightJsLang(): Promise<HljsLanguage> {
		switch (language.id) {
			case Languages.csharp.id:
				return (await import('svelte-highlight/src/languages/csharp')).default;
			case Languages.go.id:
				return (await import('svelte-highlight/src/languages/go')).default;
			case Languages.java.id:
				return (await import('svelte-highlight/src/languages/java')).default;
			case Languages.js.id:
				return (await import('svelte-highlight/src/languages/javascript')).default;
		}
	}
</script>

{#if code}
	<nav class="level mt-2">
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
