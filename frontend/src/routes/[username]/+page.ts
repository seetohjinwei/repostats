import type Repository from "src/models/repository.svelte";
import type { PageLoad } from "./$types";

export const load = (async ({ params }) => {
	const { username } = params;

	const res = await fetch(`https://repostats.jinwei.dev/api/user?username=${username}`);

	const json = await res.json();
	const repos: Repository[] | undefined = json.data;
	const reposSorted: Repository[] = (repos ?? []).sort((a, b) => a.repo.localeCompare(b.repo));
	const error: string | undefined = json.error;

	return {
		ok: res.ok,
		error,
		repos: reposSorted,
		username,
	};
}) satisfies PageLoad;
