<script lang="ts">
	import { page } from '$app/stores';
	let showNav = false;
</script>

<nav class="navbar py-3 has-background-light" aria-label="Menu navigation bar">
	<div class="container">
		<div class="navbar-brand">
			<a class="navbar-item" href="/">
				<div class="mr-2">
					<img src="logo.svg" alt="OTel recipes logo" style="min-height: 50px;" width="auto" />
				</div>
				<h1 class="title is-5 mb-3">OTel Recipes</h1>
			</a>
			<!-- svelte-ignore a11y-missing-attribute -->
			<a
				role="button"
				class="navbar-burger"
				aria-label="menu"
				aria-expanded="false"
				data-target="navMenu"
				class:is-active={showNav}
				on:click={(e) => (showNav = !showNav)}
			>
				<span aria-hidden="true" />
				<span aria-hidden="true" />
				<span aria-hidden="true" />
			</a>
		</div>
		<div class="navbar-menu" class:is-active={showNav} id="navMenu">
			<div class="navbar-end">
				<a href="/" class="navbar-item" class:is-active={$page.url.pathname === '/'}>Home</a>
				<a href="/recipes" class="navbar-item" class:is-active={$page.url.pathname === '/recipes'}>
					Recipes
				</a>
			</div>
		</div>
	</div>
</nav>

<style lang="scss">
	@import 'bulma/sass/utilities/initial-variables.sass';
	.navbar-menu .navbar-item {
		margin-left: 1.2rem;
		font-size: 1.3rem;
		padding: 0;
	}
</style>
