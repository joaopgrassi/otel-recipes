<script lang="ts">
	export let title: string, description: string, source: string;
	export let tags: string[] = [];
</script>

<div class="column is-4-desktop is-6-tablet">
	<div class="card p-3">
		<div class="card-content">
			<!-- Card title -->
			<div class="media">
				<div class="media-left">
					<figure class="image is-48x48">
						<img src="langs/csharp-plain.svg" alt="Language logo" />
					</figure>
				</div>
				<div class="media-content">
					<p class="title is-4">{title}</p>
					<p class="subtitle">
						{#each tags as tag}
							<span class="tag is-primary is-light">{tag}</span>
						{/each}
					</p>
				</div>
			</div>
			<!-- Card content -->
			<div class="content">
				{description}
			</div>
		</div>
		<!-- Card footer -->
		<footer class="card-footer">
			<div class="card-footer-item">
				<div class="is-flex is-justify-content-space-between" style="width: 100%;">
					<a class="button is-primary ml-auto" href={source}>See recipe</a>
				</div>
			</div>
		</footer>
	</div>
</div>
