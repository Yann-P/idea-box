<script>
  import { entities } from './store'

  async function vote(id, vote) {
    const data = new FormData
    data.append('id', id)
    data.append('vote', vote)
    const res = await fetch('api/vote', {
      body: data,
      method: 'post'
    })
    voted = true
  }

  function showNextRandom(e) {
    console.log('nextc')
    e && e.preventDefault();
    index = ~~(Math.random() * $entities.length)
    voted = false
  }

  function getScore(c) {
    return c ? Math.round((c.support / (c.against + c.support)) * 100) : NaN
  }
  
  let index = -1
  let voted = false

  $: current = $entities.sort((e1, e2) => e1.id - e2.id)[index]
  $: score = getScore(current)

  entities.subscribe(() => {
    if(index === -1 && $entities.length !== 0)
      showNextRandom();
  })

</script>
 
{#if current != null}
  <div class="ui container">
    <div class="ui raised segment">
      <center><h1>{current.name}</h1></center>
      {#if !voted}
        <br>
        <div class="ui buttons center fluid">
          <button class="ui positive button" on:click={vote.bind(null, current.id, true)}>For</button>
          <div class="or"></div>
          <button class="ui negative button" on:click={vote.bind(null, current.id, false)}>Against</button>
        </div>
      {:else}
        <div class="ui steps center fluid">
          <div class="step">
            <div class="content center">
                <h3>{score}%</h3>
            </div>
          </div>      
        </div>
        <div class="ui button blue center fluid" on:click={showNextRandom}>
            Random
        </div>
      {/if}
    </div>
  </div>
{/if}