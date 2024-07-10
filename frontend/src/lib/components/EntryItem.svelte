<script>
    import blankImgIc from '$assets/images/entry-blank-icon.svg?raw';
    import copyIc from '$assets/images/copy.svg?raw';
    import { ClipboardSetText } from '$wails/runtime/runtime';

    export let item;

    let entry = item.entry;

    let copyMsg = '';

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
        {@html blankImgIc}

        <div class="codeInfo">
            <p><strong>{entry.issuer}</strong> ({entry.name})</p>
            <h2>{item.code}</h2>
        </div>
    </div>

    {#if copyMsg}
        <button class="secondary" on:click={copyCode} on:mouseleave={resetToolTip} data-tooltip={copyMsg}>
            {@html copyIc}
        </button>
    {:else}
        <button class="secondary" on:click={copyCode}>
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
        display: flex;
        align-items: center;
        gap: calc(var(--pico-spacing) * 0.75);
    }

    .codeInfo {
        min-width: 9rem;
    }

    .codeInfo > * {
        margin: 0;
    }

    button {
        padding: calc(var(--pico-form-element-spacing-vertical) * 0.3) calc(var(--pico-form-element-spacing-horizontal) * 0.3);
    }
</style>
