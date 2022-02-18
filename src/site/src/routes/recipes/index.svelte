<script lang="ts">
	import { Samples } from '$lib/common/types';
	import { selectedSample, selectedSignal } from '$lib/store/store';
	import RecipeSelector from '$lib/components/recipe-selector/RecipeSelector.svelte';
</script>

<RecipeSelector />

{#if $selectedSample.id !== Samples.none.id}
	<section>
		<div class="container has-text-centered">
			<h1 class="title is-4">Steps</h1>
			<h2 class="subtitle is-6">
				Follow the steps below to configure OpenTelemetry in your project 🔭
			</h2>
		</div>
	</section>
	<div class="section">
		<div class="content">
			<div class="columns is-centered">
				<div class="column is-four-fifths">
					<div class="steps">
						<div class="center-line">
							<a href="#" class="scroll-icon">
								<img class="caret-up" alt="^" src="caret-up.svg" />
							</a>
						</div>

						<div class="step indicator">
							<section>
								<small class="icon">0</small>
								<div class="block">
									<h1 class="title is-5 is-spaced">Install the packages</h1>
								</div>
								<div class="block">
									{#each $selectedSample.dependencies as dep}
										<p>{dep.id}@{dep.version}</p>
									{/each}
								</div>
							</section>
						</div>
						{#each $selectedSample.steps as step}
							<div class="step indicator">
								<section>
									<small class="icon">{step.order}</small>
									<div class="block">
										<h1 class="title is-5 is-spaced">{step.displayName}</h1>
									</div>
									<div class="block">
										<pre>{step.source}</pre>
									</div>
								</section>
							</div>
						{/each}
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	@import 'bulma/sass/utilities/initial-variables.sass';
	.steps {
		/* max-width: 1080px; */
		padding: 0 20px;
		position: relative;
	}

	.steps .center-line {
		position: absolute;
		height: 100%;
		width: 4px;
		background: #fff;
		left: 40px;
		top: 20px;
		transform: translateX(-50%);
	}

	.steps .step {
		margin: 30px 0 3px 60px;
	}

	.steps .step section {
		background: #fff;
		border-radius: 5px;
		width: 100%;
		padding: 20px;
		position: relative;
		box-shadow: 0 0.5em 1em -0.125em rgba($black, 0.1), 0 0px 0 1px rgba($black, 0.02);
	}

	.steps .step section::before {
		position: absolute;
		content: '';
		height: 15px;
		width: 15px;
		background: #fff;
		top: 28px;
		z-index: -1;
		transform: rotate(45deg);
	}

	.indicator section::before {
		left: -7px;
	}

	.indicator section .icon {
		left: -60px;
		top: 15px;
		right: -60px;
	}

	.step section .icon,
	.center-line .scroll-icon {
		position: absolute;
		background: $link;
		height: 40px;
		width: 40px;
		font-weight: 700;
		text-align: center;
		line-height: 40px;
		border-radius: 50%;
		color: $white-ter;
		font-size: 17px;
	}

	.center-line .scroll-icon {
		bottom: 0px;
		left: 50%;
		font-size: 25px;
		transform: translateX(-50%);
	}

	.caret-up {
		margin: 3px 0 0 1px;
		width: 30px;
	}
</style>
