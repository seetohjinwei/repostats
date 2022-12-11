<script lang="ts">
	import Icon from "../components/Icon.svelte";

	let username: string = "";
	let repository: string = "";

	function handleSubmit() {
		if (username === "" || repository === "") {
			return;
		}

		window.location.href += `${username}/${repository}`;
	}
</script>

<div class="wrapper">
	<div class="container">
		<h1 class="title">RepoStats</h1>
		<p class="subtitle">Statistics on your Repositories</p>

		<form on:submit={handleSubmit}>
			<div class="fields">
				<input type="text" bind:value={username} placeholder="username" required />
				<input type="text" bind:value={repository} placeholder="repository" required />
			</div>

			<button type="submit">Search <Icon name="right-arrow" /></button>
		</form>

		<p class="links">
			<a
				class="source"
				href="https://github.com/seetohjinwei/repostats"
				target="_blank"
				rel="noopener noreferrer">Source Code</a
			>
			| Made by
			<a class="author" href="https://jinwei.dev" target="_blank" rel="noopener noreferrer"
				>Jin Wei</a
			>
		</p>
	</div>
</div>

<style lang="scss">
	@import "../styles/colours.scss";

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
	}

	.title {
		font-size: 4em;
		margin: 0;
	}

	.subtitle {
		font-size: 1.3em;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

		div.fields {
			@media (min-width: 1000px) {
				flex-direction: row;
			}

			input {
				background-color: $input-bg;
				color: $input-text;
				font-size: medium;

				border-radius: 6px;
				border: 0px;
				padding: 1ch 1ch;
				/* top left-right bottom */
				margin: 0 6px 6px;

				::placeholder {
					color: $input-text;
					opacity: 0.7;
				}
			}
		}

		button {
			display: inline-flex;
			align-items: center;

			margin-top: 6px;
			background-color: $button-bg;
			border-radius: 6px;
			font-size: medium;
			padding: 0.7ch 1.5ch;
		}
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

		.source {
			&:hover,
			&:active {
				box-shadow: inset 10ch 0 0 0 $link-text;
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
