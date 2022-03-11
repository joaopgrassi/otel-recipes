<script lang="ts">
	import { slide } from 'svelte/transition';
	import type { Language, Sample } from '$lib/common/types';
	export let sample: Sample;
	export let language: Language;

	let metadataIsVisible = false;
	function toggleMetadata() {
		metadataIsVisible = metadataIsVisible ? false : true;
	}
</script>

<nav class="level mb-2">
	<div class="level-left">
		<div class="level-item">
			<h1 class="title is-5 has-text-grey-lighter">Sample metadata</h1>
		</div>
	</div>
	<div class="level-right">
		<input
			type="button"
			class="button is-small is-link is-right"
			on:click={() => toggleMetadata()}
			value={metadataIsVisible ? 'Hide' : 'Show'}
		/>
	</div>
</nav>
{#if metadataIsVisible}
	<div class="metadata" transition:slide|local>
		<div class="metadata-line p-1 mb-1">
			<p>
				<strong>Language:</strong>
				{language.displayName}
			</p>
		</div>
		{#if sample.description}
			<div class="metadata-line p-1 mb-1">
				<p>
					<strong>Description:</strong>
					{sample.description}
				</p>
			</div>
		{/if}
		<div class="p-1">
			<p>
				<strong>Sample source:</strong>
				<a href={sample.sourceRoot}>{sample.sourceRoot}</a>
			</p>
		</div>
	</div>
{/if}

<style lang="scss">
	.level {
		margin-bottom: 0;
	}

	.metadata,
	strong {
		color: $white-ter;
	}

	.metadata a:hover {
		color: $link;
		text-decoration: underline;
	}

	.metadata-line {
		border-bottom: 1px solid rgba(255, 255, 255, 0.3);
	}
</style>
