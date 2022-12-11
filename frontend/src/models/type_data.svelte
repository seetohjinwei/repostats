<script lang="ts" context="module">
	export interface TypeData {
		language: string;
		file_count: number;
		bytes: number;
	}

	export function TypeDataPrettify(td: TypeData): string {
		const language: string = td.language === " " ? "no extension" : td.language;

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

	export function TypeDataCompareFn(a: TypeData, b: TypeData): number {
		if (a.bytes !== b.bytes) {
			return a.bytes - b.bytes;
		} else if (a.file_count !== b.file_count) {
			return a.file_count - b.file_count;
		}

		return a.language.localeCompare(b.language);
	}

	export function TypeDataReverseCompareFn(a: TypeData, b: TypeData): number {
		return -1 * TypeDataCompareFn(a, b);
	}
</script>
