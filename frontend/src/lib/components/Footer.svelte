<script>
    import githubIc from '$assets/images/brand-github.svg?raw';
    import infoIc from '$assets/images/info-circle.svg?raw';
    import { releaseUrl, version } from '$lib/stores';
    import { openExtUrl } from '$lib/util';
</script>

<footer data-theme="dark">
    <nav>
        <ul>
            <li>
                <a href="https://github.com/Sammy-T/avda" class="contrast" data-tooltip="GitHub" 
                    onclick={openExtUrl}>
                    {@html githubIc}
                </a>
            </li>
        </ul>

        <ul>
            {#if $releaseUrl}
                <li>
                    <small>
                        <a href={$releaseUrl} class="contrast" data-tooltip="Update Available" 
                            onclick={openExtUrl}>
                            {@html infoIc} v{$version}
                        </a>
                    </small>
                </li>
            {:else}
                <li><small>v{$version}</small></li>
            {/if}
        </ul>
    </nav>
</footer>

<style>
    footer {
        border-top: 3px solid var(--pico-contrast);
    }

    footer nav ul {
        margin: 0;
    }

    footer nav li {
        padding: calc(var(--pico-nav-element-spacing-vertical) * 0.5) calc(var(--pico-nav-element-spacing-horizontal) * 0.5);
    }

    footer nav li small {
        color: var(--pico-contrast);
    }

    footer a {
        text-decoration: none;
    }

    nav {
        /* 
        Why I have to account for the intentional overflow is beyond me.
        see: https://picocss.com/docs/nav#overflow
        see: https://github.com/picocss/pico/issues/549
        */
        padding: 0 calc(var(--pico-nav-element-spacing-horizontal) * 4);
    }
</style>
