<script>
	import { page } from '$app/stores';
</script>

<section class="section is-medium">
	<div class="container">
		<div class="columns is-variable is-vcentered">
			<div class="column is-8">
				<img src="error.png" alt="Astronaut falling" />
			</div>
			<div class="column has-text-centered">
				<h1 class="title is-size-1">You explored too hard.</h1>
				<h2 class="subtitle is-size-4">
					It's OK though, just go back <a href="/">Home</a>.
				</h2>
				<div class="block">
					<p class="is-size-6">But if you are curious, here's what went wrong:</p>
					<code>
						{$page.status}:
						{$page.error.message}
					</code>
				</div>
			</div>
		</div>
	</div>
</section>
