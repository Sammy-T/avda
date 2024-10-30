<script>
    import githubIc from '$assets/images/brand-github.svg?raw';
    import { GetAppInfo } from '$wails/go/main/App';
    import { BrowserOpenURL } from '$wails/runtime/runtime';
    import { onMount } from 'svelte';

    let version = '';

    /**
     * Captures a link click event and opens the external url
     * in the default browser.
     * @param event {Event}
     */
    function openExtUrl(event) {
        event.preventDefault();

        // @ts-ignore
        const url = event.currentTarget.href;
        BrowserOpenURL(url);
    }

    async function loadInfo() {
        const resp = await GetAppInfo();
        version = resp.data.productVersion;
    }

    onMount(() => {
        loadInfo();
    });
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
            <li><small>v{version}</small></li>
        </ul>
    </nav>
</footer>

<style>
    footer {
        border-top: 3px solid var(--pico-contrast);
    }

    footer nav li {
        padding: calc(var(--pico-nav-element-spacing-vertical) * 0.5) calc(var(--pico-nav-element-spacing-horizontal) * 0.5);
    }

    footer nav li small {
        color: var(--pico-contrast);
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
