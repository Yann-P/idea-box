<script>
import { entities } from './store'

const TOP_ELEMENTS_TO_DISPLAY = 10;
let top = null;
entities.subscribe(v => { 
    top = v.sort((x, y) => y.support - x.support).slice(0, TOP_ELEMENTS_TO_DISPLAY)
})
</script>

{#if top != null && top.length > 0}
    <div class="ui container">
        <div class="ui cards">
            {#each top as t}
                <div class="card">
                    <div class="content">
                        <div class="header">{t.name}</div>
                        <div class="meta">{t.support} votes</div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
{/if}
