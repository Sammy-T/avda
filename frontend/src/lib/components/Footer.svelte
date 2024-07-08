<script>
    import githubIc from '$assets/images/brand-github.svg?raw';
</script>

<footer data-theme="dark">
    <nav>
        <ul>
            <li>
                <a href="##" class="contrast" data-tooltip="GitHub">
                    {@html githubIc}
                </a>
            </li>
        </ul>

        <ul>
            <li><small>v0.1.0</small></li>
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
