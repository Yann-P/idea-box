<script>
  import Vote from './Vote.svelte';
  import Feed from './Feed.svelte';
	import Top from './Top.svelte';
	import Add from './Add.svelte';
  import {entities} from './store'

  async function getAll() {
    const res = await fetch('api/list', {
      method: 'get'
    })
    return await res.json()
  }

  (async() => {
    entities.set(await getAll());

    new EventSource("./api/sse").addEventListener('message', (message) => {
      const data = JSON.parse(message.data);
      if (data.Type === 'add') {
        $entities.push(data.Data)
        entities.set($entities)
      }
      if(data.Type === 'vote') {
        Object.assign($entities.find(e => e.id === data.Data.id), data.Data)
        entities.set($entities)
      }
    })
  })().catch(console.error)
</script>

<style>
  :global(body) {
    overflow: hidden;
  }
</style>


<h1 class="title ui center aligned header">Poll example app</h1>
<Vote/>
<h3 class="title ui center aligned header">Feed</h3>
<Feed/>
<h3 class="title ui center aligned header">Add</h3>
<Add/>
<h3 class="title ui center aligned header">Top</h3>
<Top/>
