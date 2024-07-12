<script>
    import EntryItem from './EntryItem.svelte';
    import { closeFile } from '$lib/util';
    import { EventsOn } from '$wails/runtime/runtime';
    import { getContext } from 'svelte';

    let items = [];

    const search = getContext('search');

    $: filteredItems = items.filter(item => {
        const searchRE = new RegExp($search, 'i');
        const { issuer, name } = item.entry;
        return searchRE.test(issuer) || searchRE.test(name);
    });

    EventsOn('onCodesUpdated', (entryCodes) => items = entryCodes);
</script>

{#if filteredItems.length === 0}
    <div id="empty">
        <p>No entries.</p>
        <button on:click={closeFile}>Open Vault</button>
    </div>
{:else}
    <div class="list">
        {#each filteredItems as item (item.entry.uuid) }
            <EntryItem {item} />
        {/each}
    </div>
{/if}

<style>
    .list {
        flex-grow: 1;
        padding: calc(var(--pico-spacing) * 1.5);
        overflow-y: auto;
        display: flex;
        align-content: flex-start;
        align-items: flex-start;
        flex-wrap: wrap;
        gap: calc(var(--pico-spacing) * 2);
    }

    #empty {
        height: 100%;
        text-align: center;
        align-self: center;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }
</style>
