<script lang="ts">
	import { onMount } from 'svelte';
	import Highlight from 'svelte-highlight';
	import 'svelte-highlight/src/styles/github-dark-dimmed.css';
	import { shell } from 'svelte-highlight/src/languages';

	import { Dependency, Language, Languages, Sample } from '$lib/common/types';

	export let sample: Sample;
	export let language: Language;

	let code = '';

	function getInstallText(deps: Dependency[], supplier: (dep: Dependency) => string): string {
		let lines: string[] = [];
		deps.forEach((d: Dependency) => lines.push(supplier(d)));
		return lines.join('\n');
	}

	onMount(() => {
		switch (language.id) {
			case Languages.csharp.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `dotnet add package ${d.id} --version ${d.version}`
				);
				return;
			case Languages.go.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `go get ${d.id}@${d.version}`
				);
				return;
			case Languages.java.id:
				code = getInstallText(sample.dependencies, (d: Dependency) => `${d.id} @ ${d.version}`);
				return;
			case Languages.js.id:
				code = getInstallText(
					sample.dependencies,
					(d: Dependency) => `npm install${d.id}@${d.version}`
				);
				return;
		}
	});
</script>

<div>
	<Highlight language={shell} {code} />
	<nav class="buttons is-right">
		<button class="button is-small is-link is-right">Copy</button>
	</nav>
</div>
