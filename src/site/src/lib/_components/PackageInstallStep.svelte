<script lang="ts">
	import Highlight from 'svelte-highlight';
	import 'svelte-highlight/src/styles/github-dark-dimmed.css';
	import { shell } from 'svelte-highlight/src/languages';

	import { copyToClipboard } from '../_utils/utils';
	import { Dependency, Language, Languages, Sample } from '$lib/common/types';

	export let sample: Sample;
	export let language: Language;

	let code = '';

	$:{
		switch (language.id) {
			case Languages.csharp.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `dotnet add package ${d.id} --version ${d.version}`
				);
				break;
			case Languages.go.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `go get ${d.id}@${d.version}`
				);
				break;
			case Languages.java.id:
				code = getInstallText(sample.dependencies, (d: Dependency) => `${d.id} @ ${d.version}`);
				break;
			case Languages.js.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `npm install ${d.id}@${d.version}`
				);
				break;
		}
	}

	function getInstallText(deps: Dependency[], supplier: (dep: Dependency) => string): string {
		let lines: string[] = [];
		deps.forEach((d: Dependency) => lines.push(supplier(d)));
		return lines.join('\n');
	}
</script>

{#if code}
	<nav class="level mt-2" aria-label="Sample package list navigation bar">
		<div class="level-left">
			<div class="level-item">
				<h1 class="title is-5 has-text-grey-lighter">Install the packages</h1>
			</div>
		</div>
		<div class="level-right">
			<button class="button is-small is-link is-right" on:click={() => copyToClipboard(code)}>
				Copy
			</button>
		</div>
	</nav>
	<Highlight language={shell} {code} />
{/if}
