import type { TypeData } from "../../../models/type_data.svelte";
import type { PageLoad } from "./$types";

export const load = (async ({ params }) => {
	const { username, repo } = params;

	const res = await fetch(
		`https://repostats.jinwei.dev/api/repo?username=${username}&repo=${repo}`,
	);

	const json = await res.json();
	const typeData: { [key: string]: TypeData } | undefined = json.data;
	const error: string | undefined = json.error;

	return {
		ok: res.ok,
		error,
		typeData: typeData ?? {},
		username,
		repo,
	};
}) satisfies PageLoad;
