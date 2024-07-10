<script>
    import blankImgIc from '$assets/images/entry-blank-icon.svg?raw';
    import copyIc from '$assets/images/copy.svg?raw';

    export let item;

    let entry = item.entry;
</script>

<article class="entry">
    {@html blankImgIc}

    <div class="info">
        <p><strong>{entry.issuer}</strong> ({entry.name})</p>
        <h2>{item.code}</h2>
    </div>

    <a href="##">
        {@html copyIc}
    </a>
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
</style>
