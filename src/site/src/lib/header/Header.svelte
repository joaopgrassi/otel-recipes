<script lang="ts">
	import { page } from '$app/stores';
	import logo from './svelte-logo.svg';
</script>

<section class="is-background-contain is-background-no-repeat">
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
	<div class="section">
		<div class="container">
			<div class="is-max-w-lg mx-auto has-text-centered mb-6">
				<h2 class="title is-spaced">
					<span>OpenTelemetry Recipes</span>
				</h2>
				<p class="subtitle" data-removed="true">
					Learn how to get started with OpenTelemetry without all the fuss, with the help of small,
					bite size recipes.
				</p>
			</div>
		</div>
	</div>
</section>

<secion>
	<div class="container">
		<div class="columns">

			<div class="column">
				<div class="box how">
					Step 1
				</div>
			</div>

			<div class="column">
				<div class="box">Step 2</div>
			</div>

			<div class="column">
				<div class="box">Step 3</div>
			</div>

			<div class="column">
				<div class="box">Step 4</div>
			</div>
		</div>
	</div>
</secion>
