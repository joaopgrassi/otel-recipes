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

<section class="section">
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
			<div class="column is-6">
				<div class="section-media">
					<img role="presentation" src="wfh_1.svg" alt="Girl on laptop chilling with cute dogue" />
				</div>
			</div>
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<div class="columns">
			<div class="column is-4">
				<div class="card">
					<figure class="image is-128x128">
						<img src="language-icon.svg" alt="language" />
					</figure>
					<hr class="spacer is-1" />
					<h3 class="title is-4 is-spaced">Pick a language</h3>
					<p class="subtitle is-5">Select one of the available languages.</p>
				</div>
			</div>
			<div class="column is-4">
				<div class="card is-usp">
					<figure class="image is-128x128">
						<img src="signal-icon.svg" alt="language" />
					</figure>
					<hr class="spacer is-1" />
					<h3 class="title is-4 is-spaced">Then a Signal</h3>
					<p class="subtitle is-5">
						Select the telemetry you want to add to your app: Trace, Metric or Log.
					</p>
				</div>
			</div>
			<div class="column is-4">
				<div class="card is-usp">
					<figure class="image is-128x128">
						<img src="sample-icon.svg" alt="language" />
					</figure>
					<hr class="spacer is-1" />
					<h3 class="title is-4 is-spaced">Lastly, the sample type</h3>
					<p class="subtitle is-5">Select the type of app.</p>
				</div>
			</div>
		</div>
	</div>
</section>


<style>
	.step {
		width: 40px;
		height: 40px;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		border-radius: 50%;
		border: none;
		text-decoration: none;
	}

	.numberCircle {
		width: 40px;
		height: 40px;
		text-align: center;
		font-size: 22px;
		border-radius: 50%;
		border: 2px solid #3498db;
		box-shadow: 0 0 0 8px #fff;
		color: #000;
		background: #fff;
		margin: 0 auto;
		z-index: 1;
		transform: translateY(-90%);
		line-height: 33px;
	}

	.card {
		background-color: transparent;
		box-shadow: none;
		color: #998da0;
		max-width: 100%;
		position: relative;
	}

	.card.is-usp {
		max-width: 320px;
		margin: 0 auto;
	}

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
</style>
