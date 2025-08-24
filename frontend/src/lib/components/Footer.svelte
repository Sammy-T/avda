<script>
    import settingsIc from '$assets/images/settings.svg?raw';
    import githubIc from '$assets/images/brand-github.svg?raw';
    import infoIc from '$assets/images/info-circle.svg?raw';
    import { releaseUrl, showSettings, version } from '$lib/stores.svelte';
    import { openExtUrl } from '$lib/util.svelte';

    /**
     * @param {Event} event
     */
    function openSettings(event) {
        event.preventDefault();

        $showSettings = true;
    }
</script>

<footer>
    <nav>
        <ul>
            <li>
                <a href="##" class="contrast" data-tooltip="Settings" onclick={openSettings}>
                    {@html settingsIc}
                </a>
            </li>
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
        --pico-tooltip-background-color: var(--pico-background-color);
        --pico-tooltip-color: var(--pico-contrast);

        border-top: 3px solid var(--nav-fg-color);

        & nav {
            & ul {
                margin: 0;
            }

            & li {
                padding: calc(var(--pico-nav-element-spacing-vertical) * 0.5) calc(var(--pico-nav-element-spacing-horizontal) * 0.5);

                & small {
                    color: var(--nav-fg-color);
                }
            }
        }

        & a {
            color: var(--nav-fg-color);
            text-decoration: none;
        }
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
