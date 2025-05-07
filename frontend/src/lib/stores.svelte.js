import { get, writable } from 'svelte/store';
import { GetAppInfo } from '$wails/go/main/App';
import { getLatestReleaseInfo } from './util.svelte';

export const version = writable('');
export const releaseUrl = writable();
export const vaultPath = writable();
export const items = writable([]);
export const search = writable('');
export const order = writable('');
export const selectedGroupUuid = writable(null);
export const groupsMap = writable(new Map());
export const showSettings = writable(false);

async function loadInfo() {
    const resp = await GetAppInfo();
    version.set(resp.data.productVersion);

    const releaseInfo = await getLatestReleaseInfo();
    if(!releaseInfo || get(version) === releaseInfo.tag_name.replace('v', '')) return;

    releaseUrl.set(releaseInfo.html_url);
}

loadInfo();
