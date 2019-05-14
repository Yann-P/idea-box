<script>
    import {entities} from './store'

    const MAX_FEED_ELEMENTS = 3

    $: feed = [];

    new EventSource("./api/sse").addEventListener('message', (message) => {
        const data = JSON.parse(message.data);
        if (data.Type === 'add') {
            feed.unshift(`"${data.Data.name}" was added`)
        }
        if(data.Type === 'vote') {
            const entity = $entities.find(e => e.id === data.Data.id)
            if (Math.abs(data.Data.against - data.Data.support) === 1) {
                feed.unshift(`${entity.name} ${data.Data.vote ? 'has now over 50% support' : 'has dropped below 50% support'}`) 
            }
        }
        feed = feed.slice(0, MAX_FEED_ELEMENTS)  
    })
</script>

<style>
    ul {
        list-style: none;
        margin: 0;
        padding: 0;
    }
</style>

<div class="ui container">
    <ul>
        {#each feed as f}
            <li>{f}</li>
        {/each}
    </ul>
</div>
