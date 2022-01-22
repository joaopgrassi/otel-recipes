<script lang="ts">
	import logo from './svelte-logo.svg';
</script>

<nav class="navbar is-relative is-transparent py-3 has-background-light">
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
						<a class="button is-primary is-outlined" href="#">Log In</a>
						<a class="button is-primary" href="#">Sign Up</a>
					</div>
				</div>
			</div>
		</div>
	</div>
</nav>
