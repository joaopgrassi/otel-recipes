<script context="module" lang="ts">
	export const prerender = true;
</script>

<svelte:head>
	<title>OTel Recipes - About</title>
</svelte:head>

<section class="section">
	<div class="container">
		<div class="columns">
			<div class="column is-8">
				<h3 class="title is-3">Introducing: OTel Recipes</h3>
				<div class="block">
					OTel Recipes is an open source project composed of two parts: a collection of sample
					applications that focus on achieving a specific goal with OpenTelemetry and this website,
					that allows you to browse the samples in a simple and structured way.
				</div>

				<div class="block">
					The samples are curated by OpenTelemetry enthusiasts and are focused on teaching you how
					to configure and use OpenTelemetry for various tasks <strong>the easy way</strong>
					. The samples help you answer questions such as:
				</div>

				<div class="content">
					<ul>
						<li>How can I configure my Java application to emit OpenTelemetry traces?</li>
						<li>How can I record metrics in my .NET service?</li>
						<li>How can I enrich my application logs with trace information?</li>
					</ul>
				</div>

				<div class="block">
					The sample applications follow a structure and are <strong>broken down into steps</strong>
					 (just like a recipe! 🥙). The website then uses this structure to present the steps needed
					to successfully configure and use OpenTelemetry in your projects. 🚀 🔭
				</div>

				<div class="block">
					Since the project is open source, anyone can contribute to OTel recipes. You can help
					maintain and update the existing samples, introduce new samples for your favorite
					programming language(s), help with documentation and many more. Take a look at our <a
						href="https://github.com/joaopgrassi/otel-recipes/blob/main/CONTRIBUTING.md"
					>
						contributing guidelines
					</a>
					 to learn how to get started.
				</div>

				<h1 class="title is-4">Why OTel recipes?</h1>

				<div class="block">
					There's no doubt that OpenTelemetry is here to stay, and will eventually become the
					de-facto solution for all things observability.
				</div>

				<div class="block">
					As OpenTelemetry adoption grows, more and more people will start to hear about it and want
					to see what it is about. The truth is that, for many, it will be the first time they are
					exposed to the many concepts that the OpenTelemetry project brings.
				</div>

				<div class="block">
					There is a lot of effort put in by the the OpenTelemetry community in making the barrier
					to start low, but this is not an easy task. On top of that, the project is advancing
					quickly, (which is great) but it can be tricky to keep up with things.
				</div>

				<div class="block">
					After actively collaborating to OpenTelemetry and using it in my projects I, noticed the
					following pattern emerge:
					<div class="content">
						<ul>
							<li>I need to instrument some code/application</li>
							<li>
								I go to either the official documentation or to the GitHub SDK repository to see how
								to use it
							</li>
							<li>
								The documentation is either lacking or outdated (remember, the project is evolving
								so this is somewhat expected)
							</li>
							<li>
								I end up looking at the SDK tests to find out how to use it and interact with the
								SDK
							</li>
						</ul>
					</div>
				</div>

				<div class="block">
					The "cycle" above works but it's a bit cumbersome. Of course, things will improve as the
					project matures, but I realized that many will probably face the same challenges.
				</div>

				<div>
					So I thought: Wouldn't it be nice if there was a site where I could "ask" questions on how
					to achieve XYZ with OpenTelemetry and the site would "reply" me with a list of steps that
					I could simply follow do achieve what I wanted? The rest, as they say, is history 😅.
				</div>
			</div>
		</div>
	</div>
</section>
