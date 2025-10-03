<script lang="ts">
  export let title: string = '';
  export let icon: string = '';
</script>

<div class="form-section">
  {#if title}
    <h3>
      {#if icon}
        <span class="section-icon">{icon}</span>
      {/if}
      {title}
    </h3>
  {/if}
  
  <slot />
</div>

<style>
  .form-section {
    margin-bottom: 25px;
    padding: 20px;
    background: rgba(255, 255, 255, 0.5);
    border-radius: 15px;
    border: 1px solid rgba(42, 82, 152, 0.1);
  }

  h3 {
    color: #2a5298;
    font-size: 1.1em;
    margin: 0 0 15px 0;
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .section-icon {
    font-size: 1.2em;
  }
</style>
