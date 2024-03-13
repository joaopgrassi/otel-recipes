<script lang="ts">
	import type { Recipe } from '$lib/common/types';
	import { selectedRecipeId } from '$lib/store/store';
	import { fly } from 'svelte/transition';
	export let sample: Recipe;
</script>

<div class="column is-4-desktop is-6-tablet" in:fly={{ x: 100, duration: 300 }}>
	<div class="card p-3 card-equal-height">
		<div class="card-content pr-4">
			<!-- Card title -->
			<div class="media">
				<div class="media-left">
					<figure class="image is-48x48">
						<img src="langs/{sample.languageId}-plain.svg" alt="Language logo" />
					</figure>
				</div>
				<div class="media-content">
					<p class="title is-4">{sample.displayName}</p>
				</div>
				<div class="media-right">
					<span class="tag is-info is-light">
						<span class="mr-1 mt-1">
							<img
								src="icons/{sample.signal}-icon-card.svg"
								width="15"
								alt="{sample.signal}-icon"
							/>
						</span>
						<span class="is-size-7"
							>{sample.signal[0].toUpperCase() + sample.signal.slice(1).toLowerCase()}</span
						>
					</span>
				</div>
			</div>
			<!-- Card content -->
			<div class="content is-text-overflow card-test">
				<div class="block">
					{sample.description}
				</div>
				<div class="block">
					{#each sample.tags || [] as tag}
						<span class="tag is-primary is-light mr-2">{tag}</span>
					{/each}
				</div>
			</div>
		</div>
		<!-- Card footer -->
		<footer class="card-footer">
			<div class="card-footer-item">
				<div class="is-flex is-justify-content-space-between" style="width: 100%;">
					<button class="button is-primary ml-auto" on:click={() => selectedRecipeId.set(sample.id)}
						>See recipe</button
					>
				</div>
			</div>
		</footer>
	</div>
</div>
