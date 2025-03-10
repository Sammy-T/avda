<script>
    import EntryItem from './EntryItem.svelte';
    import { items, order, search, selectedGroupUuid } from '$lib/stores';
    import { closeFile } from '$lib/util';
    import { EventsOn } from '$wails/runtime/runtime';
    import { getContext } from 'svelte';

    /** @type {import('svelte/store').Writable<boolean>} */
    const displaySearch = getContext('displaySearch');

    let sortedItems = $derived.by(() => {
        if($order === 'custom') {
            return $items;
        }

        return [...$items].sort((a, b) => {
            switch($order) {
                case 'account-asc':
                    return a.entry.name.localeCompare(b.entry.name);

                case 'account-desc':
                    return b.entry.name.localeCompare(a.entry.name);

                case 'issuer-asc':
                    return a.entry.issuer.localeCompare(b.entry.issuer);

                case 'issuer-desc':
                    return b.entry.issuer.localeCompare(a.entry.issuer);
            }
        });
    });

    let filteredItems = $derived(sortedItems.filter((item) => {
        // Filter by search
        const searchRE = new RegExp($search, 'i');
        const { issuer, name } = item.entry;
        const matchesSearch = searchRE.test(issuer) || searchRE.test(name);
        
        // Filter by group
        const matchesGroup = $selectedGroupUuid === null || 
            (item.entry.groups && item.entry.groups.includes($selectedGroupUuid));
        
        return matchesSearch && matchesGroup;
    }));

    /**
     * @param event {KeyboardEvent}
     */
    function onKeyEvent(event) {
        if(filteredItems.length === 0 || !event.ctrlKey) return;

        switch(event.key) {
            case 'f':
                $displaySearch = !$displaySearch;
                break;
            
            case 'w':
                closeFile();
                break;
        }
    }

    EventsOn('onCodesUpdated', (entryCodes) => $items = entryCodes);
</script>

<svelte:document on:keyup={onKeyEvent} />

{#if filteredItems.length === 0}
    <div id="empty">
        <p>No entries.</p>
        <button onclick={closeFile}>Open Vault</button>
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
        padding: calc(var(--pico-spacing) * 1.5);
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
        gap: calc(var(--pico-spacing) * 2);
    }

    #empty {
        height: 100%;
        text-align: center;
        align-self: center;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
</style>
