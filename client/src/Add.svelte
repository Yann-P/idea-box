<script>

async function postNew(name) {
  const data = new FormData;
  data.append('name', name)
  const res = await fetch('api/add', {
    body: data,
    method: 'post'
  })
}

function addNew(event) {
  event.preventDefault();
  const elt = this.querySelector('input[type=text]')
  postNew(elt.value).then(async () => {
    const old = elt.style.background;	
    elt.style.background = 'rgba(0, 255, 0, 0.6)';
    elt.style.transition = 'all .3s linear';
    await new Promise(r => setTimeout(r, 300));
    elt.style.transition = '';
    elt.style.background = old;
    elt.value = '';
  }).catch(console.error);
  
}

</script>

<div class="ui container">
  <div class="ui segments">
    <form id="add" on:submit={addNew}>    
      <div class="ui action input fluid">
        <input type="text" placeholder="Nouveau">
        <button class="ui button" type="submit">Add</button>
      </div>
    </form>
  </div>
</div>
