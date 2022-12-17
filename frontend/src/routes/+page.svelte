<script lang="ts">
	import Icon from "../components/Icon.svelte";

	let username: string = "";
	let repository: string = "";
	let error: string = "";
	const resetError = () => (error = "");

	// Regex from: https://github.com/shinnn/github-username-regex
	const usernameTest = /^[a-z\d](?:[a-z\d]|-(?=[a-z\d])){0,38}$/i;

	// StackOverflow answer that talks about is a valid GitHub repo name: https://stackoverflow.com/a/64147124/
	// I can't find a lot of information surrounding this, but I believe it should be largely true.
	// Furthermore, the regex is intentionally less restrictive (read: doesn't check the length).
	const repositoryTest = /^[a-zA-Z\d._-]+$/i;

	function handleSubmit() {
		// `api` is reserved on both GitHub and RepoStats
		if (username === "api" || username === "" || !usernameTest.test(username)) {
			error = `"${username}" is not a valid username!`;
			return;
		}

		if (repository === "") {
			window.location.href += `${username}`;
			return;
		}

		if (!repositoryTest.test(repository)) {
			error = `"${repository}" is not a valid repository name!`;
			return;
		}

		window.location.href += `${username}/${repository}`;
	}
</script>

<h1 class="title">RepoStats</h1>
<p class="subtitle">Statistics on your Repositories</p>

<form on:submit={handleSubmit}>
	<div class="fields">
		<input type="text" bind:value={username} placeholder="username" required />
		<input type="text" bind:value={repository} placeholder="repository" />
	</div>

	<div id="error" on:click={resetError} on:keypress={resetError}>{error}</div>
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
	<a class="author" href="https://jinwei.dev" target="_blank" rel="noopener noreferrer">Jin Wei</a>
</p>

<style lang="scss">
	@import "../styles/colours.scss";

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

	#error {
		cursor: pointer;
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
