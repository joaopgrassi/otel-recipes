<script lang="ts">
	import { selectedLanguage, allLanguages, allSignals, selectedSignal } from '$lib/store/store';
	import RecipeCard from './RecipeCard.svelte';
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

<section class="section">
  <div class="container">
    <div class="columns is-multiline">
      <div class="column is-4-desktop is-6-tablet">

				<RecipeCard></RecipeCard>
      </div>
</section>
