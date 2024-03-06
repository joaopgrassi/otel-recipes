<script lang="ts">
	import { Recipes } from '$lib/common/types';
	import {
		selectedLanguage,
		allLanguages,
		allSignals,
		selectedSignal,
		filteredSamples,
		selectedSampleId
	} from '$lib/store/store';
	import RecipeCard from './RecipeCard.svelte';
	import { fly } from 'svelte/transition';
</script>

<section class="section">
	<div class="container">
		<div class="columns has-text-centered">
			<div class="column">
				<div class="field has-addons">
					<!-- Input search -->
					<div class="control is-expanded">
						<input class="input is-medium" type="text" placeholder="Search for sample apps" />
					</div>

					<!-- Language select -->
					<div class="control">
						<div class="select is-medium">
							<select id="language" name="language" bind:value={$selectedLanguage}>
								{#each $allLanguages as lang}
									<option value={lang}>
										{lang.displayName}
									</option>
								{/each}
							</select>
						</div>
					</div>

					<!-- Signal select -->
					<div class="control">
						<div class="select is-medium">
							<select class="" id="signal" name="signal" bind:value={$selectedSignal}>
								{#each $allSignals as signal}
									<option value={signal}>
										{signal.displayName}
									</option>
								{/each}
							</select>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>

{#if $filteredSamples.length > 0 && $selectedSampleId === Recipes.none.id}
	<section class="section" in:fly={{ x: 100, duration: 300 }}>
		<div class="container">
			<div class="columns is-multiline">
				{#each $filteredSamples as sample}
					<span />
					<RecipeCard {sample} />
				{/each}
			</div>
		</div>
	</section>
{/if}
