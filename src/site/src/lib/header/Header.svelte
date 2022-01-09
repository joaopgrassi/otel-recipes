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

<div class="container">
	<div class="is-max-w-lg mx-auto has-text-centered mb-6">
		<h2 class="title is-spaced">
			<span>OpenTelemetry Recipes</span>
		</h2>
		<p class="subtitle" data-removed="true">
			Learn how to get started with OpenTelemetry without all the fuss, with the help of small, bite
			size recipes.
		</p>
	</div>
</div>

<div class="content is-medium has-background-light">
	<div class="hero-body">
		<div class="columns is-vcentered">
			<div class="column is-three-fifths">
				<div class="content has-text-centered is-large">
					<h1>How it works</h1>

					<div class="tile is-ancestor">
						<div class="tile is-parent">
							<article class="tile is-child box">
								<p class="title">1</p>
								<p class="subtitle">Pick a language</p>
							</article>
						</div>
						<div class="tile is-parent">
							<article class="tile is-child box">
								<p class="title">2</p>
								<p class="subtitle">Then the type of instrumentation (Signal) </p>
							</article>
						</div>
						<div class="tile is-parent">
							<article class="tile is-child box">
								<p class="title">3</p>
								<p class="subtitle">And finally, the sample type</p>
							</article>
						</div>
					</div>

				</div>
			</div>

			<div class="column">
				<figure class="image">
					<img src="wfh_1.svg" />
				</figure>
			</div>
		</div>
		<!-- <div class="columns">
			<div class="column mb-3">
				<div class="box">
					<div class="numberCircle">1</div>
					<div class="block">Pick a programming language</div>
				</div>
			</div>

			<div class="column mb-3">
				<div class="box">
					<div class="numberCircle">2</div>
					<div class="block">Then the type of instrumentation (Signal)</div>
				</div>
			</div>

			<div class="column mb-3">
				<div class="box">
					<div class="numberCircle">3</div>
					<div class="block">And finally, the sample type</div>
				</div>
			</div>
		</div> -->
	</div>
</div>

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
</style>
