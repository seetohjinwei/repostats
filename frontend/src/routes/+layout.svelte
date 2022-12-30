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
		margin-top: 2em;

		@media (min-width: 1000px) {
			margin-top: 0;
			height: 90vh;
		}
	}

	.container {
		@media (min-width: 1000px) {
			position: relative;
			top: 50%;
			-webkit-transform: translateY(-50%);
			-ms-transform: translateY(-50%);
			transform: translateY(-50%);
		}

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
