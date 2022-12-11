<script lang="ts">
	import { TypeDataPrettify, TypeDataReverseCompareFn } from "../../../models/type_data.svelte";
	import type { PageData } from "./$types";

	export let data: PageData;

	const githubLink: string = `https://github.com/${data.username}/${data.repo}/`;

	const languages = Object.values(data.typeData)
		.sort(TypeDataReverseCompareFn)
		.map(TypeDataPrettify);

	const truncatedLanguages = languages.slice(0, 5);
</script>

<div class="wrapper">
	<div class="container">
		<h1 class="title">{data.username}'s {data.repo}</h1>

		{#if !data.ok}
			<p>Repository not found! <a href="..">Try another one?</a></p>
		{:else}
			<a target="_blank" rel="noopener noreferrer" href={githubLink}>{githubLink}</a>

			{#each truncatedLanguages as td}
				<p>{td}</p>
			{/each}

			<!-- TODO: pie chart or other visualisation? -->

			<p class="links">
				<a class="more" href="..">Check out another!</a>
				| Made by
				<a class="author" href="https://jinwei.dev" target="_blank" rel="noopener noreferrer"
					>Jin Wei</a
				>
			</p>
		{/if}
	</div>
</div>

<style lang="scss">
	@import "../../../styles/colours.scss";

	.wrapper {
		height: 90vh;
	}

	.container {
		position: relative;
		top: 50%;
		-webkit-transform: translateY(-50%);
		-ms-transform: translateY(-50%);
		transform: translateY(-50%);

		text-align: center;

		a {
			color: $link-text;
			&:hover,
			&:active {
				color: $link-text-focus;
			}
		}
	}

	.title {
		margin: 0;
		overflow: hidden;
		text-overflow: ellipsis;
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

		.more {
			&:hover,
			&:active {
				box-shadow: inset 18ch 0 0 0 $link-text;
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
