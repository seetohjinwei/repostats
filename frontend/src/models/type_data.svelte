<script lang="ts" context="module">
	export interface TypeData {
		language: string;
		file_count: number;
		bytes: number;
	}

	function prettifyLanguage(language: string): string {
		if (language === " ") {
			return "no extension";
		}

		return language;
	}

	export function TypeDataPrettify(td: TypeData): string {
		const language: string = prettifyLanguage(td.language);

		const files: string = td.file_count === 1 ? "file" : "files";

		const units: string[] = ["B", "kB", "MB", "GB", "TB"];
		let bytes: number = td.bytes;
		let index: number = 0;
		while (bytes >= 1000) {
			bytes = Math.floor(bytes / 1000);
			index++;
		}
		const unit: string = units[index];

		return `${language}: ${td.file_count} ${files} (${bytes}${unit})`;
	}

	export function TypeDataToChartJS(tds: TypeData[], limit: number) {
		const truncated = tds.slice(0, limit);

		const labels = [...truncated.map((td) => prettifyLanguage(td.language)), "others"];

		const rest = tds.slice(limit);
		const others = rest.reduce(
			(acc, curr) => {
				return {
					...acc,
					file_count: acc.file_count + curr.file_count,
					bytes: acc.bytes + curr.bytes,
				};
			},
			{ language: "others", file_count: 0, bytes: 0 },
		);

		// Colour scheme from: https://www.schemecolor.com/orange-green-blue-pie-chart.php
		const colors = ["#F47A1F", "#FDBB2F", "#377B2B", "#7AC142", "#007CC3", "#00529B"];
		const labelColors = ["#6C0B89", "#3A1B2F"];

		const datasets = [
			{
				label: "bytes",
				data: [...truncated.map((td) => td.bytes), others.bytes],
				backgroundColor: colors,
				hoverOffset: 4,
			},
			{
				label: "file count",
				data: [...truncated.map((td) => td.file_count), others.file_count],
				backgroundColor: colors,
				hoverOffset: 4,
			},
		];

		const legendButtons = ["bytes", "file count"];

		const data = { labels, datasets };
		const options = {
			responsive: true,
			plugins: {
				legend: {
					onClick: function (e: any, legendItem: any, legend: any) {
						const legendIndex: number = legendButtons.findIndex((x) => x === legendItem.text);
						const ci = legend.chart;
						ci.getDatasetMeta(legendIndex).hidden = !ci.getDatasetMeta(legendIndex).hidden;
						ci.update();
					},
					display: true,
					position: "bottom",
					labels: {
						generateLabels: (chart: any) => {
							return chart.data.labels.slice(0, legendButtons.length).map((l: any, i: number) => ({
								fontColor: "#f3efe0",
								datasetIndex: i,
								text: legendButtons[i],
								fillStyle: labelColors[i],
								strokeStyle: labelColors[i],
								hidden: chart.getDatasetMeta(i).hidden,
							}));
						},
					},
				},
			},
		};

		return [data, options];
	}

	export function TypeDataCompareFn(a: TypeData, b: TypeData): number {
		if (a.file_count !== b.file_count) {
			return a.file_count - b.file_count;
		} else if (a.bytes !== b.bytes) {
			return a.bytes - b.bytes;
		}

		return a.language.localeCompare(b.language);
	}

	export function TypeDataReverseCompareFn(a: TypeData, b: TypeData): number {
		return -1 * TypeDataCompareFn(a, b);
	}
</script>
