<script lang="ts">
  export let routerUrl: string;
  
  let showTooltip = false;
</script>

{#if routerUrl}
  <div 
    class="router-info"
    on:mouseenter={() => showTooltip = true}
    on:mouseleave={() => showTooltip = false}
  >
    <span class="router-icon">🌐</span>
    
    {#if showTooltip}
      <div class="tooltip">
        Подключено к: {routerUrl}
      </div>
    {/if}
  </div>
{/if}

<style>
  .router-info {
    position: absolute;
    bottom: 20px;
    left: 20px;
    cursor: help;
    z-index: 10;
  }

  .router-icon {
    font-size: 1.5em;
    opacity: 0.6;
    transition: opacity 0.3s ease;
  }

  .router-info:hover .router-icon {
    opacity: 1;
  }

  .tooltip {
    position: absolute;
    bottom: 100%;
    left: 0;
    background: rgba(0, 0, 0, 0.8);
    color: white;
    padding: 8px 12px;
    border-radius: 6px;
    font-size: 0.9em;
    white-space: nowrap;
    margin-bottom: 8px;
    animation: fadeIn 0.2s ease;
  }

  .tooltip::after {
    content: '';
    position: absolute;
    top: 100%;
    left: 12px;
    border: 4px solid transparent;
    border-top-color: rgba(0, 0, 0, 0.8);
  }

  @keyframes fadeIn {
    from { opacity: 0; transform: translateY(5px); }
    to { opacity: 1; transform: translateY(0); }
  }
</style>
