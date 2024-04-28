<svelte:head>
	<title>OTel Recipes - Home</title>
</svelte:head>

<!-- intro -->
<section class="hero is-fullheight has-background-grey-darker">
	<div class="hero-body">
		<div class="container is-fluid">
			<div class="columns is-vcentered">
				<div class="column is-7">
					<div class="section-header">
						<h1 class="title is-size-1 is-size-3-mobile has-text-white-bis mb-6">
							Navigate the world of OpenTelemetry
						</h1>
						<h2 class="subtitle is-size-4 is-size-6-mobile has-text-white-bis">
							With the help of our sample applications and well-defined steps (just like a recipe!)
							you will learn how instrument your applications and become an OpenTelemetry chef in no time 😎
						</h2>
						<div class="field is-grouped">
							<div class="control">
								<a class="button is-primary is-medium" href="/recipes">Check the recipes</a>
							</div>
						</div>
					</div>
				</div>
				<div class="column is-offset-1 mt-6">
					<img src="chef-amico.svg" alt="Girl on laptop chilling with cute dogue" />
				</div>
			</div>
		</div>
	</div>
</section>

<!-- How it works section -->
<section class="section is-medium">
	<div class="container">
		<div class="column has-text-centered mb-3">
			<h1 class="title is-2">Here's how it works:</h1>
		</div>
		<div class="columns">
			<div class="column is-4 has-text-centered">
				<img src="searching-bro.svg" alt="explore apps" width="300" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">1. Explore the samples</h3>
				<p class="subtitle is-5">Explore the available samples. You can text search or use the filters.</p>
			</div>

			<div class="column is-4 has-text-centered">
				<img src="choice-bro.svg" alt="sample app" width="300" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">2. Pick a sample</h3>
				<p class="subtitle is-5">
					Click on the sample to view its details
				</p>
			</div>

			<div class="column is-4 has-text-centered">
				<img src="steps.svg" alt="language" width="300" />
				<hr class="spacer is-1" />
				<h3 class="title is-4 is-spaced">3. Follow the steps</h3>
				<p class="subtitle is-5">
					Follow the steps given by the sample to configure OTel in your application.
				</p>
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
					<img src="browser-buttons.svg" alt="browser top-bar icons" />
					<div class="container mt-2 p-4">
						<div class="ph-row">
							<div class="ph-col-6 mb-1" />
							<div class="ph-col-4 empty" />
							<div class="ph-col-10 mb-1" />
							<div class="ph-col-2 empty" />
							<div class="ph-col-2 mb-1" />
							<div class="ph-col-10 empty" />
							<div class="ph-col-4 mb-1" />
							<div class="ph-col-8 empty" />
							<div class="ph-col-12 big mb-1" />
							<div class="ph-col-6 mb-1" />
							<div class="ph-col-6 empty" />
							<div class="ph-col-10 mb-1" />
							<div class="ph-col-2 empty" />
							<div class="ph-col-2 mb-1" />
							<div class="ph-col-10 empty mb-1" />
							<div class="ph-col-4 mb-1" />
							<div class="ph-col-8 empty mb-1" />
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<div class="content">
					<h1 class="title is-3 is-spaced has-text-light">
						Learn and apply OpenTelemetry concepts the practical way.
					</h1>
					<hr class="is-line-learn" />
					<p class="subtitle is-5 has-text-light">
						For many, observability is a whole new world and getting started can be a bit overwhelming in the beginning.
						But don't worry, OTel recipes is here to help!
					</p>
					<p class="subtitle is-5 has-text-light">
						With the help of <i>small, to-the-point samples,</i>
						and well-defined steps, you can learn how to quickly get started and add OTel to your projects.
					</p>
					<p class="subtitle is-5 has-text-light">
						The samples are end-to-end tested, to guarantee a frustration-free experience when trying it out on your own. 🧘‍♂️
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
						We foster an inclusive and welcoming community. It doesn't matter whether you are
						already familiar with OpenTelemetry or just getting started.
					</p>
					<div class="field is-grouped">
						<div class="control">
							<a class="button is-dark is-medium" href="https://github.com/joaopgrassi/otel-recipes"
								>View the project on GitHub</a
							>
						</div>
					</div>
				</div>
			</div>
			<div class="column">
				<img src="welcome.svg" alt="Girl on laptop chilling with cute cat" />
			</div>
		</div>
	</div>
</section>

<style lang="scss">
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
