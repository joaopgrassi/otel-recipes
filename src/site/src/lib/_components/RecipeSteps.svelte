<script lang="ts">
	import { fly } from 'svelte/transition';
	import { Recipes, Step } from '$lib/common/types';
	import { selectedLanguage, selectedSample } from '$lib/store/store';
	import CodeStep from './CodeStep.svelte';
	import MetadataStep from './MetadataStep.svelte';
	import PackageInstallStep from './PackageInstallStep.svelte';

	const getSortedSteps = (steps: Step[]) => {
		return steps.sort((left: Step, right: Step) => left.order - right.order);
	};
</script>

{#if $selectedSample.id !== Recipes.none.id}
	<div in:fly={{ x: 100, duration: 300 }}>
		<section>
			<div class="container has-text-centered">
				<h2 class="subtitle is-5">
					Follow the steps below to configure OpenTelemetry in your project 🔭
				</h2>
			</div>
		</section>
		<div class="section">
			<div class="content">
				<div class="columns is-centered">
					<div class="column is-11">
						<div class="steps">
							<div class="center-line">
								<!-- svelte-ignore a11y-invalid-attribute -->
								<a href="#" class="scroll-icon">
									<img class="caret-up" alt="^" src="caret-up.svg" />
								</a>
							</div>
							<div class="step step-metadata indicator">
								<section>
									<small class="icon">
										<img class="icon-info" alt="info" src="info.svg" />
									</small>
									<MetadataStep sample={$selectedSample} language={$selectedLanguage} />
								</section>
							</div>
							<div class="step indicator">
								<p class="step-language">
									<span class="step-language-tag bd-is-html">shell</span>
								</p>
								<section>
									<small class="icon">
										<img class="icon-package" alt="package" src="package.svg" />
									</small>
									<img src="browser-buttons.svg" alt="browser top-bar icons" />
									<PackageInstallStep sample={$selectedSample} language={$selectedLanguage} />
								</section>
							</div>
							{#each getSortedSteps($selectedSample.steps) as step}
								<div class="step indicator">
									<p class="step-language">
										<span class="step-language-tag bd-is-html">{$selectedLanguage.id}</span>
									</p>
									<section>
										<small class="icon">{step.order}</small>
										<img src="browser-buttons.svg" alt="browser top-bar icons" />
										<CodeStep {step} sample={$selectedSample} language={$selectedLanguage} />
									</section>
								</div>
							{/each}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	@import 'bulma/sass/utilities/initial-variables.sass';

	.step-language {
		margin-bottom: 0;
	}

	.step-language-tag.bd-is-html {
		background: $code-bg;
		border-bottom-left-radius: 0;
		border-bottom-right-radius: 0;
		color: $grey-lighter;
	}

	.step-language-tag {
		align-items: center;
		background: $code-bg;
		border-radius: 0.5em;
		border-bottom-right-radius: 0.5em;
		border-bottom-left-radius: 0.5em;
		color: $grey-lighter;
		display: inline-flex;
		flex-grow: 0;
		flex-shrink: 0;
		font-size: 0.9em;
		font-weight: 600;
		height: 1.5rem;
		padding: 0 1em;
		vertical-align: top;
	}

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
		background: #22272e;
		border-radius: 5px;
		border-top-left-radius: 0;
		width: 100%;
		padding: 20px;
		position: relative;
		box-shadow: 0 0.1em 1em 0 rgba($black, 0.4);
	}

	.steps .step-metadata section {
		border-top-left-radius: 5px;
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

	.icon-package {
		margin: 7px 0 0 1px;
		height: 25px;
	}

	.icon-info {
		margin: 8px 0 0 1px;
		height: 23px;
	}
</style>
