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
    <div class="router-icon-container">
      <span class="router-icon">🌐</span>
      <div class="pulse-ring"></div>
      <div class="pulse-ring delay"></div>
    </div>
    
    {#if showTooltip}
      <div class="tooltip">
        <div class="tooltip-header">
          <span class="status-dot"></span>
          Подключено к роутеру
        </div>
        <div class="tooltip-url">{routerUrl}</div>
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
    z-index: 50;
  }

  .router-icon-container {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 50px;
    height: 50px;
  }

  .router-icon {
    font-size: 1.8em;
    z-index: 1;
    position: relative;
    filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
    animation: float 3s ease-in-out infinite;
  }

  .pulse-ring {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 40px;
    height: 40px;
    border: 2px solid rgba(42, 82, 152, 0.3);
    border-radius: 50%;
    animation: pulse 2s ease-out infinite;
  }

  .pulse-ring.delay {
    animation-delay: 1s;
    border-color: rgba(34, 197, 94, 0.3);
  }

  .router-info:hover .router-icon {
    animation: bounce 0.6s ease;
    transform: scale(1.1);
  }

  .router-info:hover .pulse-ring {
    animation-play-state: paused;
    opacity: 0;
  }

  .tooltip {
    position: absolute;
    bottom: 120%;
    left: 0;
    background: linear-gradient(135deg, rgba(0, 0, 0, 0.9), rgba(30, 30, 30, 0.95));
    color: white;
    padding: 12px 16px;
    border-radius: 12px;
    font-size: 0.9em;
    white-space: nowrap;
    margin-bottom: 8px;
    animation: tooltipIn 0.3s ease;
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    min-width: 200px;
  }

  .tooltip-header {
    display: flex;
    align-items: center;
    gap: 8px;
    font-weight: 600;
    margin-bottom: 4px;
    color: #22c55e;
  }

  .status-dot {
    width: 8px;
    height: 8px;
    background: #22c55e;
    border-radius: 50%;
    animation: statusPulse 2s ease-in-out infinite;
  }

  .tooltip-url {
    color: #94a3b8;
    font-family: monospace;
    font-size: 0.85em;
  }

  .tooltip::after {
    content: '';
    position: absolute;
    top: 100%;
    left: 20px;
    border: 6px solid transparent;
    border-top-color: rgba(0, 0, 0, 0.9);
  }

  @keyframes float {
    0%, 100% { transform: translateY(0px); }
    50% { transform: translateY(-3px); }
  }

  @keyframes pulse {
    0% {
      transform: translate(-50%, -50%) scale(0.8);
      opacity: 1;
    }
    100% {
      transform: translate(-50%, -50%) scale(1.4);
      opacity: 0;
    }
  }

  @keyframes bounce {
    0%, 20%, 50%, 80%, 100% { transform: translateY(0); }
    40% { transform: translateY(-8px); }
    60% { transform: translateY(-4px); }
  }

  @keyframes tooltipIn {
    from { 
      opacity: 0; 
      transform: translateY(10px) scale(0.9); 
    }
    to { 
      opacity: 1; 
      transform: translateY(0) scale(1); 
    }
  }

  @keyframes statusPulse {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.5; }
  }
</style>
