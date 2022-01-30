<script lang="ts">
	import {
		selectedLanguage,
		languages,
		filteredSignals,
		selectedSignal,
		selectedSample,
		filteredSamples
	} from '$lib/store/store';
</script>

<section class="section is-medium">

	<div class="container">

		<div class="columns has-text-centered">
			<div class="column">
				<div class="field">
					<label class="label" for="lang">Programming language</label>
					<div class="control">
						<div class="select is-rounded is-medium">
							<select name="lang" bind:value={$selectedLanguage}>
								{#each $languages as lang}
									<option value={lang}>
										{lang}
									</option>
								{/each}
							</select>
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<div class="field">
					<label class="label" for="signal">Signal</label>
					<div class="control">
						<div class="select is-rounded is-medium">
							<select name="signal" bind:value={$selectedSignal}>
								{#each $filteredSignals as signal}
									<option value={signal}>
										{signal}
									</option>
								{/each}
							</select>
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<div class="field">
					<label class="label" for="sample">Sample type</label>
					<div class="control">
						<div class="select is-rounded is-medium">
							<select name="sample" bind:value={$selectedSample}>
								{#each $filteredSamples as sample}
									<option value={sample}>
										{sample.type}
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
