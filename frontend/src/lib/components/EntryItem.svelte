<script>
    import blankImgIc from '$assets/images/entry-blank-icon.svg?raw';
    import copyIc from '$assets/images/copy.svg?raw';
    import { formatCode } from '$lib/util';
    import { ClipboardSetText } from '$wails/runtime/runtime';

    /**
     * @typedef {Object} Props
     * @property {Object} item
     */

    /** @type {Props} */
    let { item } = $props();

    let { entry } = item;

    let copyMsg = $state('');

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
        copyMsg = '';
    }
</script>

<article class="entry">
    <div class="info">
        {#if entry.icon}
            <img src={`data:${entry.icon_mime};base64,${entry.icon}`} alt="">
        {:else}
            {@html blankImgIc}
        {/if}

        <div class="codeInfo">
            <p><strong>{entry.issuer}</strong> ({entry.name})</p>
            <h2>{formatCode(item.code)}</h2>
        </div>
    </div>

    {#if copyMsg}
        <button class="secondary" onclick={copyCode} onmouseleave={resetToolTip} data-tooltip={copyMsg}>
            {@html copyIc}
        </button>
    {:else}
        <button class="secondary" onclick={copyCode}>
            {@html copyIc}
        </button>
    {/if}
</article>

<style>
    article {
        width: 25rem;
        margin: 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .info {
        flex-grow: 1;
        display: flex;
        align-items: center;
        gap: calc(var(--pico-spacing) * 0.75);
    }

    .codeInfo > * {
        margin: 0;
    }

    button {
        padding: calc(var(--pico-form-element-spacing-vertical) * 0.3) calc(var(--pico-form-element-spacing-horizontal) * 0.3);
    }
</style>
