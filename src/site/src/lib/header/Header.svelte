<script lang="ts">
	import logo from './svelte-logo.svg';
	import { page } from '$app/stores';
	let showNav = false;
</script>

<nav class="navbar py-3 has-background-light">
	<div class="container">
		<div class="navbar-brand">
			<a class="navbar-item" href="/">
				<img src={logo} alt="" style="min-height: 40px;" width="auto" />
			</a>
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

	a.navbar-item.is-active:not(:focus):not(:hover),
	.navbar-link.is-active:not(:focus):not(:hover) {
		color: $grey-dark;
	}

	a.navbar-item:focus,
	a.navbar-item:focus-within,
	a.navbar-item:hover,
	a.navbar-item.is-active,
	.navbar-link:focus,
	.navbar-link:focus-within,
	.navbar-link:hover,
	.navbar-link.is-active {
		background-color: transparent;
		color: $link;
	}
</style>
