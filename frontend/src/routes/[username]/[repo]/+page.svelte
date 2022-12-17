<script lang="ts">
	import {
		TypeDataPrettify,
		TypeDataReverseCompareFn,
		TypeDataToChartJS,
	} from "../../../models/type_data.svelte";
	import type { PageData } from "./$types";
	import { Doughnut } from "svelte-chartjs";
	import { Chart as ChartJS, ArcElement, Legend, Tooltip } from "chart.js";
	ChartJS.register(ArcElement, Legend, Tooltip);

	const LIMIT = 5;

	export let data: PageData;

	const githubLink: string = `https://github.com/${data.username}/${data.repo}/`;

	$: sorted = Object.values(data.typeData).sort(TypeDataReverseCompareFn);
	$: languages = sorted.map(TypeDataPrettify);
	$: truncatedLanguages = languages.slice(0, LIMIT);
	$: [doughnutData, doughnutOptions] = TypeDataToChartJS(sorted, LIMIT);

	async function handleForceRefresh() {
		const res = await fetch(
			`https://repostats.jinwei.dev/api/repo_force?username=${data.username}&repo=${data.repo}`,
		);

		const json = await res.json();
		if (json.data === undefined) {
			return;
		}

		data.typeData = json.data;
	}
</script>

<h1 class="title">{data.username}'s {data.repo}</h1>

{#if !data.ok}
	<p>Repository not found! <a href="/">Try another one?</a></p>
{:else}
	<a target="_blank" rel="noopener noreferrer" href={githubLink}>{githubLink}</a>

	<div class="content">
		<div class="list">
			<b>Top 5</b>
			{#each truncatedLanguages as td}
				<p>{td}</p>
			{/each}
		</div>

		<div class="chart">
			<Doughnut data={doughnutData} options={doughnutOptions} />
		</div>
	</div>

	<button on:click={handleForceRefresh} title="Please don't overuse this!">Refresh Data!</button>

	<p class="links">
		<a class="back" href="/">Home</a>
		|
		<a class="back" href="/{data.username}">Repos</a>
		| Made by
		<a class="author" href="https://jinwei.dev" target="_blank" rel="noopener noreferrer">Jin Wei</a
		>
	</p>
{/if}

<style lang="scss">
	@import "../../../styles/colours.scss";

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

	.content {
		margin: 2em 0;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column-reverse;

		@media (min-width: 1000px) {
			flex-direction: row;
			gap: 50px;
		}

		.chart {
			height: 200px;
		}
	}

	button {
		border-radius: 6px;
		padding: 0.3ch 0.7ch;
	}
</style>
