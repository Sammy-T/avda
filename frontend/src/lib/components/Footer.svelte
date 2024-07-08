<script>
    import githubIc from '$assets/images/brand-github.svg?raw';
</script>

<footer>
    <nav>
        <ul>
            <li>
                <a href="##" class="secondary" data-tooltip="GitHub">
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
        border-top: 3px solid var(--pico-secondary);
    }

    footer nav li {
        padding: calc(var(--pico-nav-element-spacing-vertical) * 0.5) calc(var(--pico-nav-element-spacing-horizontal) * 0.5);
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
