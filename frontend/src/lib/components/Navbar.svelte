<script>
    import avdaIc from '$assets/images/avda.svg?raw';
    import searchIc from '$assets/images/search.svg?raw';
    import closeFileIc from '$assets/images/file-minus.svg?raw';
    import { vaultPath } from '$lib/stores';
    import { closeFile } from '$lib/util';
    import { getContext } from 'svelte';

    const displaySearch = getContext('displaySearch');

    /**
     * @param {Event} event
     */
    function toggleSearch(event) {
        event.preventDefault();

        $displaySearch = !$displaySearch;
    }
</script>

<nav data-theme="dark">
    <ul>
        <li>
            {@html avdaIc}
            <strong>avda</strong>
        </li>
    </ul>

    <ul class="filename">
        <li>
            <a href="##" class="contrast" onclick={closeFile}>
                {$vaultPath?.split(/[\\/]/).at(-1) ?? 'filename'}
            </a>
        </li>
    </ul>

    <ul>
        <li>
            <a href="##" class="contrast" data-tooltip="Search" data-placement="bottom" onclick={toggleSearch}>
                {@html searchIc}
            </a>
        </li>
        <li>
            <a href="##" class="contrast" data-tooltip="Close Vault" data-placement="bottom" onclick={closeFile}>
                {@html closeFileIc}
            </a>
        </li>
    </ul>
</nav>

<style>
    nav {
        /* 
        Why I have to account for the intentional overflow is beyond me.
        see: https://picocss.com/docs/nav#overflow
        see: https://github.com/picocss/pico/issues/549
        */
        padding: 0 calc(var(--pico-nav-element-spacing-horizontal) * 4);
        overflow-x: clip;
    }

    li strong {
        color: var(--pico-contrast);
    }

    a:hover {
        color: var(--pico-primary-background);
    }

    .filename a:hover {
        text-decoration: none;
    }
</style>
