<script lang="ts">
	import { page } from '$app/stores';
	import logo from './svelte-logo.svg';
</script>

<nav class="navbar is-relative is-transparent py-3" aria-label="main navigation">
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
						<a class="button is-primary is-outlined" href="#">Log In</a><a
							class="button is-primary"
							href="#">Sign Up</a
						>
					</div>
				</div>
			</div>
		</div>
	</div>
</nav>

<section class="section hero-image">
	<div class="container">
		<div class="columns is-vcentered">
			<div class="column is-half is-5-fullhd is-offset-1-fullhd">
				<div class="section-header">
					<h1 class="title is-spaced is-1">
						OpenTelemetry<span class="is-hidden-touch"><br /></span> Recipes
					</h1>

					<div class="block">
						<span class="subtitle is-4">
							The place to get started with OpenTelemetry <span class="is-hidden-touch"><br /></span
							>
							<span class="is-underlined">the easy way.</span></span
						>
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
			<div
				class="column is-full-mobile is-three-fifths-tablet is-half-desktop is-one-third-widescreen is-half-fullhd"
			>
				<div class="section-media">
					<img src="wfh_1.svg" alt="Girl on laptop chilling with cute dogue" />
				</div>
			</div>
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
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
				<p class="subtitle is-5">Lastly, select the sample type</p>
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

	.hero-image {
		position: relative;
	}

	.hero-image:before {
		content: ' ';
		display: block;
		position: absolute;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		opacity: 0.2;
		background-image: url('pattern5.svg');
		background-repeat: no-repeat;
		background-position: 20% 93%;
		background-size: cover;
	}
</style>
