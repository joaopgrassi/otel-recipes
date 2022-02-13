<script lang="ts">
	import { Signals, NoneSample, Samples } from '$lib/common/types';

	import {
		selectedLanguage,
		languages,
		filteredSignals,
		selectedSignal,
		selectedSample,
		filteredSamples
	} from '$lib/store/store';

	function languageChanged() {
		selectedSignal.set(Signals.none);
		selectedSample.set(Samples.none);
	}

	function signalChanged() {
		selectedSample.set(Samples.none);
	}
</script>

<section class="section is-medium">
	<div class="container">
		<div class="columns has-text-centered">
			<div class="column">
				<div class="field">
					<label class="label" for="lang">Programming language</label>
					<div class="control">
						<div class="select is-rounded is-medium">
							<select name="lang" bind:value={$selectedLanguage} on:change={languageChanged}>
								{#each $languages as lang}
									<option value={lang.id}>
										{lang.displayName}
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
							<select name="signal" bind:value={$selectedSignal} on:change={signalChanged}>
								{#each $filteredSignals as signal}
									<option value={signal.id}>
										{signal.displayName}
									</option>
								{/each}
							</select>
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<div class="field">
					<label class="label" for="sample">Sample app</label>
					<div class="control">
						<div class="select is-rounded is-medium">
							<select name="sample" bind:value={$selectedSample}>
								{#each $filteredSamples as sample}
									<option value={sample.id}>
										{sample.displayName}
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
