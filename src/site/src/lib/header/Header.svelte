<script lang="ts">
	import { page } from '$app/stores';
	import { loop_guard, onMount } from 'svelte/internal';
	import { fade, fly } from 'svelte/transition';

	import logo from './svelte-logo.svg';
	let toggle = false;

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

<nav class="navbar is-relative is-transparent py-3 has-background-light">
	<div class="container">
		<div class="navbar-brand">
			<a class="navbar-item" href="#">
				<img src={logo} alt="" style="min-height: 40px;" width="auto" />
			</a>
		</div>
		<div class="navbar-menu">
			<div class="navbar-end">
				<div class="navbar-item">
					<div class="buttons">
						<a class="button is-primary is-outlined" href="#">Log In</a>
						<a class="button is-primary" href="#">Sign Up</a>
					</div>
				</div>
			</div>
		</div>
	</div>
</nav>

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
				<img src="wfh_1.svg" alt="Girl on laptop chilling with cute dogue" />
			</div>
		</div>
	</div>
</section>

<section class="section is-medium">
	<div class="container">
		<div class="column has-text-centered mb-3">
			<h1 class="title is-3">Here's how it works:</h1>
		</div>
		<div class="columns">
			<div class="column is-4 has-text-centered">
				<img src="language-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">1. Language</h3>
				<p class="subtitle is-5">Select one of the available languages.</p>
			</div>
			<div class="column is-4 has-text-centered">
				<img src="signal-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">2. Telemetry Signal</h3>
				<p class="subtitle is-5">
					Then, the telemetry you want to add to your app: Trace, Metric or Log.
				</p>
			</div>
			<div class="column is-4 has-text-centered">
				<img src="sample-icon.svg" alt="language" width="150" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">3. Sample</h3>
				<p class="subtitle is-5">Lastly, select the sample type.</p>
			</div>
		</div>
	</div>
</section>

<section class="section is-medium demo">
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
				<h1 class="title is-4 has-text-light is-spaced">Easily find how to configure SDKs</h1>
				<p class="subtitle has-text-light">
					Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has
					been the industry's standard dummy text ever since the 1500s, when an unknown printer took
					a galley of type and scrambled it to make a type specimen book.
				</p>
			</div>
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<div class="columns is-5 is-variable is-vcentered">
			<div class="column is-half">
				<h1 class="title is-3 is-spaced">Get involved!</h1>
				<p class="subtitle">
					Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has
					been the industry's standard dummy text ever since the 1500s, when an unknown printer took
					a galley of type and scrambled it to make a type specimen book.
				</p>

				<div class="field is-grouped">
					<div class="control">
						<a class="button is-dark is-medium" href="#">Visit the repo on GitHub</a>
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

	.demo {
		/* background: #CB356B;
background: -webkit-linear-gradient(to bottom, #BD3F32, #CB356B);
background: linear-gradient(to bottom, #BD3F32, #CB356B); */

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
</style>
