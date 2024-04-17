<footer class="footer has-background-grey-darker">
	<div class="container">
		<div class="columns">
			<div class="column is-5 is-4-widescreen">
				<a class="navbar-item" href="/">
					<div class="mr-2">
						<img
							src="logotype-white.svg"
							alt="OTel recipes logo"
							style="min-height: 60px;"
							width="auto"
						/>
					</div>
				</a>
			</div>
			<div class="column is-7 is-6-widescreen is-offset-2-widescreen">
				<div class="columns">
					<div class="column is-4">
						<div class="menu">
							<p class="menu-label">The project</p>
							<ul class="menu-list">
								<li><a href="/about">About</a></li>
								<li><a href="/privacy">Privacy</a></li>
								<li><a href="https://github.com/joaopgrassi/otel-recipes">GitHub repo</a></li>
							</ul>
						</div>
					</div>
					<div class="column is-4">
						<div class="menu">
							<p class="menu-label">OTel Resources</p>
							<ul class="menu-list">
								<li><a href="https://opentelemetry.io">Official website</a></li>
								<li>
									<a
										href="https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/overview.md"
									>
										Specification
									</a>
								</li>
							</ul>
						</div>
					</div>
					<div class="column is-4">
						<div class="menu">
							<p class="menu-label">Social</p>
							<ul class="menu-list">
								<li>
									<a
										href="https://twitter.com/otelrecipes"
										target="_blank"
										rel="noopener noreferrer"
									>
										Twitter
									</a>
								</li>
							</ul>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</footer>

<style lang="scss">
	.menu {
		font-size: 1rem;
	}

	.menu-label {
		font-size: 1.3rem;
		font-weight: 500;
		text-transform: none;
		letter-spacing: 0;
		color: $cultured;
	}

	.menu-list {
		line-height: 1.25;
		margin-left: -0.75rem;
	}

	.menu-list a {
		color: $cultured;
	}

	.menu-list a:hover {
		background-color: transparent;
		color: $link;
	}
</style>
