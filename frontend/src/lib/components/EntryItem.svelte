<script>
    import blankImgIc from '$assets/images/entry-blank-icon.svg?raw';
    import copyIc from '$assets/images/copy.svg?raw';
    import { formatCode } from '$lib/util.svelte';
    import { ClipboardSetText } from '$wails/runtime/runtime';
    import { selectedGroupUuid, groupsMap } from '$lib/stores.svelte';

    /**
     * @typedef {Object} Props
     * @property {Object} item
     */

    /** @type {Props} */
    let { item } = $props();

    let { entry } = item;

    let copyMsg = $state(null);
    
    let groupColor = $derived(
        entry.groups && entry.groups.length > 0
            ? $groupsMap.get(entry.groups[0])?.color
            : null
    );
    
    // Highlight the selected group's color if it matches
    let highlightColor = $derived(
        entry.groups?.find(g => g === $selectedGroupUuid) 
            ? $groupsMap.get($selectedGroupUuid)?.color
            : groupColor
    );

    // Get the first letter of issuer for the fallback icon
    let issuerInitial = entry.issuer.charAt(0).toUpperCase();

    async function copyCode() {
        const success = await ClipboardSetText(item.code);

        if(!success) {
            console.error('Copy to clipboard failed.');
            copyMsg = 'copy failed';
        } else {
            copyMsg = 'copied';
        }
    }

    function resetToolTip() {
        copyMsg = null;
    }
</script>

<article class="entry" style={`--group-color: ${highlightColor ?? 'transparent'};`}>
    <div class="icon" style={`--icon-color: ${highlightColor ?? '#333333'};`}>
        {#if entry.icon}
            <img src={`data:${entry.icon_mime};base64,${entry.icon}`} alt="">
        {:else}
            {@html blankImgIc}

            <div class="initial">
                {issuerInitial}
            </div>
        {/if}
    </div>

    <div class="codeInfo">
        <p class="entry-title"><strong>{entry.issuer}</strong> <span class="entry-name">({entry.name})</span></p>
        <h2>{formatCode(item.code)}</h2>
    </div>

    <button class="secondary" onclick={copyCode} onmouseleave={resetToolTip} data-tooltip={copyMsg}>
        {@html copyIc}
    </button>
</article>

<style>
    article {
        margin: 0;
        display: grid;
        grid-template-columns: 11.5% 1fr min-content;
        align-items: center;
        gap: calc(var(--pico-spacing) * 0.75);
        border-left: 4px solid var(--group-color);
    }

    .icon {
        position: relative;

        & :global svg {
            overflow: visible;
            color: var(--icon-color);
        }
    }

    .codeInfo {
        min-width: 0;
        overflow: hidden;
    }

    .codeInfo > * {
        margin: 0;
    }
    
    .entry-title {
        overflow: hidden;
        text-overflow: ellipsis;

        & > * {
            white-space: nowrap;
        }
    }
    
    .entry-name {
        opacity: 0.8;
    }

    button {
        padding: calc(var(--pico-form-element-spacing-vertical) * 0.3) calc(var(--pico-form-element-spacing-horizontal) * 0.3);
    }

    .initial {
        color: var(--pico-primary-inverse);
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: bold;
        font-size: calc(1.15vw + 12px);
    }
</style>
