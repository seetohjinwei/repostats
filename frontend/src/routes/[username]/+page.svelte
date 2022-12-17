<script lang="ts">
	import type { PageData } from "./$types";

	export let data: PageData;
</script>

<h1 class="title">{data.username}'s repos</h1>

{#if !data.ok}
	<p>User not found! <a href="/">Try another one?</a></p>
{:else}
	<div class="repos">
		{#each data.repos as repo}
			<p><a href="{data.username}/{repo.repo}">{repo.repo}</a></p>
		{/each}
	</div>
	{#if data.repos.length == 30}
		<!-- because GitHub API only returns first 30 -->
		<p class="hint">(only first 30 repos are shown!)</p>
	{/if}
	<p class="links">
		<a class="back" href="/">Home</a>
		| Made by
		<a class="author" href="https://jinwei.dev" target="_blank" rel="noopener noreferrer">Jin Wei</a
		>
	</p>
{/if}

<style lang="scss">
	@import "../../styles/colours.scss";

	.title {
		margin: 0;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	div.repos {
		max-height: 70vh;
		overflow-y: auto;
		p {
			margin: 1ch 0;
		}

		@media (min-width: 700px) {
			display: flex;
			margin: 1em auto 0 auto;
			justify-content: space-between;
			flex-wrap: wrap;
			width: 40%;
			gap: 10px 20px;

			p {
				display: inline;
				margin: 0;
			}
		}
	}

	.hint {
		opacity: 0.5;
	}

	p.links {
		opacity: 0.7;

		a {
			$trans: 0.5s cubic-bezier(0.06, 0.53, 0.56, 0.34);
			color: $link-text;
			box-shadow: inset 0 0 0 0 $page-bg;
			transition: color $trans, box-shadow $trans;

			&:hover,
			&:active {
				color: $page-bg;
			}
		}

		.back {
			&:hover,
			&:active {
				box-shadow: inset 5ch 0 0 0 $link-text;
			}
		}

		.author {
			&:hover,
			&:active {
				box-shadow: inset 6ch 0 0 0 $link-text;
			}
		}
	}
</style>
