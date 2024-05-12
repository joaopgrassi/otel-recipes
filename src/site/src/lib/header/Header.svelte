<script lang="ts">
	import { resetSearch } from '$lib/store/store';
	let showNav = false;
</script>

<nav class="navbar py-3 has-background-grey-darker" aria-label="Menu navigation bar">
	<div class="container">
		<div class="navbar-brand">
			<a class="navbar-item" href="/">
				<div class="mr-2">
					<img
						src="logotype-white.svg"
						alt="OTel recipes logo"
						style="min-height: 50px;"
						width="auto"
					/>
				</div>
			</a>
			<!-- svelte-ignore a11y-missing-attribute -->
			<a
				role="button"
				tabindex="0"
				class="navbar-burger"
				aria-label="menu"
				aria-expanded="false"
				data-target="navMenu"
				class:is-active={showNav}
				on:click={(e) => (showNav = !showNav)}
				on:keydown={(e) => (showNav = !showNav)}
			>
				<span aria-hidden="true" />
				<span aria-hidden="true" />
				<span aria-hidden="true" />
			</a>
		</div>
		<div class="navbar-menu has-background-grey-darker" class:is-active={showNav} id="navMenu">
			<div class="navbar-end">
				<a href="/" class="navbar-item">Home</a>
				<a href="/recipes" class="navbar-item" on:click={() => resetSearch()}> Recipes </a>
				<a href="https://github.com/joaopgrassi/otel-recipes" class="navbar-item"> <img src="github-mark-white.svg" alt="GitHub"/> </a>
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
