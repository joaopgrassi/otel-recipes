<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	const texts = [
		'xLorem Ipsum is simply dummy text of\n the printing and typesetting industry. Lorem Ipsum has',
		'Lorem Ipsum is simply dummy text of\n the printing and typesetting industry. Lorem Ipsum has'
	];
	let textIndex = 0;
	let code = '';

	onMount(() => {
		code = texts[0];

		const interval = setInterval(() => {
			if (textIndex >= texts.length) {
				textIndex = 0;
			}
			code = texts[textIndex];
			textIndex++;
		}, 5000);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<!-- intro -->
<section class="section is-medium has-background-light">
	<div class="container">
		<div class="columns" />
		<div class="columns is-vcentered">
			<div class="column is-half is-5-fullhd is-offset-1-fullhd">
				<div class="section-header">
					<h1 class="title is-spaced is-1">
						OpenTelemetry
						<span class="is-hidden-touch"><br /></span>
						Recipes
					</h1>

					<div class="block">
						<span class="subtitle is-4">
							The place to get started with OpenTelemetry <span class="is-hidden-touch">
								<br />
							</span>
							<span class="is-underlined">the easy way.</span>
						</span>
					</div>

					<div class="block">
						<span class="subtitle is-4">
							Explore the collection of bite-sized samples, specially curated to get you started
							quickly.
						</span>
					</div>

					<div class="field is-grouped">
						<div class="control">
							<a class="button is-link is-medium" href="#">🌮 Check the recipes</a>
						</div>
					</div>
				</div>
			</div>
			<div class="column is-7">
				<img src="vision.svg" alt="Girl on laptop chilling with cute dogue" />
			</div>
		</div>
	</div>
</section>

<!-- How it works section -->
<section class="section is-medium">
	<div class="container">
		<div class="column has-text-centered mb-3">
			<h1 class="title is-3">Here's how it works:</h1>
		</div>
		<div class="columns">
			<div class="column is-4 has-text-centered">
				<img src="language-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">1. Programming language</h3>
				<p class="subtitle is-5">Select one of the available programming languages.</p>
			</div>
			<div class="column is-4 has-text-centered">
				<img src="signal-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">2. Signal</h3>
				<p class="subtitle is-5">
					Then, the <a href="https://github.com/open-telemetry/opentelemetry-specification/blob/v1.8.0/specification/glossary.md#signals">Signal</a> you want to add to your app: Trace, Metric or Log.
				</p>
			</div>
			<div class="column is-4 has-text-centered">
				<img src="sample-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">3. Sample app</h3>
				<p class="subtitle is-5">Lastly, select the sample app. The sample app will demonstrate in a concise way what you need to do to configure OTel in <b>your</b> apps.</p>
			</div>
		</div>
	</div>
</section>

<!-- Code slideshow section -->
<section class="section is-medium code-slideshow">
	<div class="container">
		<div class="columns is-5 is-variable is-vcentered">
			<div class="column is-half">
				<div class="browser-mockup">
					<svg xmlns="http://www.w3.org/2000/svg" width="54" height="14" viewBox="0 0 54 14">
						<g fill="none" fill-rule="evenodd" transform="translate(1 1)">
							<circle cx="6" cy="6" r="6" fill="#FF5F56" stroke="#E0443E" stroke-width=".5" />
							<circle cx="26" cy="6" r="6" fill="#FFBD2E" stroke="#DEA123" stroke-width=".5" />
							<circle cx="46" cy="6" r="6" fill="#27C93F" stroke="#1AAB29" stroke-width=".5" />
						</g>
					</svg>
					<div class="container mt-2">
						{#key code}
							<pre in:fade>{code}</pre>
						{/key}
					</div>
				</div>
			</div>
			<div class="column">
				<div class="content">
					<h1 class="title is-3 is-spaced has-text-light">
						Learn how to configure and use OTel in your projects.
					</h1>
					<hr class="is-line-learn" />
					<p class="subtitle is-5 has-text-light">
						The OpenTelemetry project can greatly improve the observability of your your apps, but
						it also brings a bit of a baggage with it.
					</p>
					<p class="subtitle is-5 has-text-light">
						If you are new to the concepts that OpenTelemetry uses, it can be a bit overwhelming.
						But don't worry, <b>we're here for you.</b>
					</p>

					<p class="subtitle is-5 has-text-light">
						With the help of <i>small, to-the-point samples,</i>
						you can learn how to quickly get started and add OTel to your projects.
					</p>

					<p class="subtitle is-5 has-text-light">
						Once you are ready, you can dive into more advance topics, over on
						<a href="https://opentelemetry.io" class="has-text-light islink-underline">
							opentelemetry.io
						</a>
						and on
						<a
							href="https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification"
							class="has-text-light islink-underline"
						>
							GitHub.
						</a>
					</p>
				</div>
			</div>
		</div>
	</div>
</section>

<!-- Get involved section -->
<section class="section is-medium">
	<div class="container">
		<div class="columns is-5 is-variable is-vcentered">
			<div class="column is-half">
				<div class="content">
					<h1 class="title is-2 is-spaced">Get involved!</h1>
					<hr class="is-line-involved" />
					<p class="subtitle is-4">
						You can contribute to OTel Recipes in several ways. Helping maintaining existing
						samples, adding new ones or improving our docs and website. All contributions are
						welcome!
					</p>
					<p class="subtitle is-4">
						We foster an inclusive and welcoming community. It doesn't matter whether you are already familiar with
						OpenTelemetry or just getting started.
					</p>
					<div class="field is-grouped">
						<div class="control">
							<a class="button is-dark is-medium" href="https:">View the project on GitHub</a>
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<img src="welcome.svg" alt="Girl on laptop chilling with cute dogue" />
			</div>
		</div>
	</div>
</section>

<style>
	.spacer {
		display: block;
		margin: 0;
		padding: 0;
		height: 0;
		background: none;
		border: none;
	}

	.spacer.is-1 {
		height: 1rem;
	}

	/* https://gist.github.com/jarthod/8719db9fef8deb937f4f */
	.browser-mockup {
		box-shadow: 0 0.1em 1em 0 rgba(0, 0, 0, 0.4);
		border-radius: 5px;
		background-color: #2b2b2b;
		min-height: 300px;
		padding: 18px;
	}

	.code-slideshow {
		background: #e96443; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to bottom,
			#904e95,
			#e96443
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to bottom,
			#904e95,
			#e96443
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
	}

	.islink-underline {
		border-bottom: 1px solid currentColor;
	}

	hr {
		height: 1px;
	}

	.is-line-learn {
		background: linear-gradient(to right, #fff 10%, rgba(1, 254, 228, 0) 100%);
	}

	.is-line-involved {
		background: linear-gradient(to right, #e96443 10%, rgba(1, 254, 228, 0) 100%);
	}
</style>
