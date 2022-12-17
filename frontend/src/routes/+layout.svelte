<div class="wrapper">
	<div class="container">
		<slot />
	</div>
</div>

<style global lang="scss">
	/* Global SCSS */

	@import "../styles/colours.scss";

	html {
		background-color: $page-bg;
		color: $primary-text;
		font-family: $font-family;
	}

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
</style>
