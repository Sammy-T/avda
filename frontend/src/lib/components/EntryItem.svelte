<script>
    import blankImgIc from '$assets/images/entry-blank-icon.svg?raw';
    import copyIc from '$assets/images/copy.svg?raw';
    import { ClipboardSetText } from '$wails/runtime/runtime';

    export let item;

    let entry = item.entry;

    async function copyCode() {
        const success = await ClipboardSetText(item.code);
        if(!success) console.error('Copy to clipboard failed.');
    }
</script>

<article class="entry">
    {@html blankImgIc}

    <div class="info">
        <p><strong>{entry.issuer}</strong> ({entry.name})</p>
        <h2>{item.code}</h2>
    </div>

    <button class="secondary" on:click={copyCode}>
        {@html copyIc}
    </button>
</article>

<style>
    article {
        width: 18rem;
        margin: 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .info {
        min-width: 9rem;
    }

    .info > * {
        margin: 0;
    }

    button {
        padding: calc(var(--pico-form-element-spacing-vertical) * 0.3) calc(var(--pico-form-element-spacing-horizontal) * 0.3);
    }
</style>
