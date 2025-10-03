<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { BrowserOpenURL } from '../../wailsjs/runtime/runtime';
  
  export let isRouterConnected: boolean = false;
  
  const dispatch = createEventDispatcher<{
    'create-awg': void;
    'add-routes': void;
    'reconnect-router': void;
    'quit': void;
  }>();

  function openGithubLink() {
    BrowserOpenURL('https://github.com/Noksa/gokeenapi');
  }
</script>

<div class="welcome-container">
  <div class="network-icon">
    <svg width="80" height="80" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M12 2L2 7L12 12L22 7L12 2Z" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
      <path d="M2 17L12 22L22 17" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
      <path d="M2 12L12 17L22 12" stroke="currentColor" stroke-width="2" stroke-linejoin="round"/>
    </svg>
  </div>
  
  <h1>Gokeenapi</h1>
  <p class="description">
    Утилита для работы с роутерами Keenetic через REST API<br><br>
    (!) Данная версия является GUI версией и содержит не весь функционал<br><br>
    Если Вам нужен весь функционал, воспользуйтесь 
    <button type="button" class="link-button" on:click={openGithubLink}>CLI версией</button>
  </p>
  
  <div class="button-group">
    <button class="btn primary network-btn" on:click={() => dispatch('create-awg')}>
      <span class="btn-icon">🔗</span>
      Создать AWG соединение
      <div class="btn-glow"></div>
    </button>
    <button class="btn primary route-btn" on:click={() => dispatch('add-routes')}>
      <span class="btn-icon">🛣️</span>
      Управление маршрутами
      <div class="btn-glow"></div>
    </button>
    
    {#if isRouterConnected}
      <button class="btn secondary reconnect-btn" on:click={() => dispatch('reconnect-router')}>
        <span class="btn-icon">🔌</span>
        Переподключиться к роутеру
      </button>
    {/if}
    
    <button class="btn secondary" on:click={() => dispatch('quit')}>
      <span class="btn-icon">🚪</span>
      Выход
    </button>
  </div>
</div>

<style>
  .welcome-container {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 50px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.3),
      0 0 0 1px rgba(255, 255, 255, 0.2);
    max-width: 650px;
    width: 100%;
    position: relative;
    overflow: hidden;
  }

  .welcome-container::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: linear-gradient(90deg, #00d4ff, #5b86e5, #36d1dc);
    animation: pulse 2s ease-in-out infinite alternate;
  }

  .network-icon {
    text-align: center;
    color: #2a5298;
    margin-bottom: 20px;
    animation: float 3s ease-in-out infinite;
  }

  h1 {
    text-align: center;
    color: #1e3c72;
    margin-bottom: 30px;
    font-size: 2.8em;
    font-weight: 700;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .description {
    text-align: center;
    color: #4a5568;
    line-height: 1.8;
    margin-bottom: 40px;
    font-size: 1.1em;
  }

  .button-group {
    display: flex;
    flex-direction: column;
    gap: 20px;
    align-items: center;
    margin-top: 40px;
  }

  .btn {
    position: relative;
    overflow: hidden;
    display: flex;
    align-items: center;
    gap: 12px;
    min-width: 280px;
    justify-content: center;
  }

  .btn-icon {
    font-size: 1.2em;
  }

  .network-btn {
    position: relative;
  }

  .route-btn {
    position: relative;
    background: linear-gradient(135deg, #059669, #10b981);
  }

  .btn-glow {
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
    transition: left 0.6s;
  }

  .network-btn:hover .btn-glow {
    left: 100%;
  }

  .route-btn:hover .btn-glow {
    left: 100%;
  }

  @keyframes float {
    0%, 100% { transform: translateY(0px); }
    50% { transform: translateY(-10px); }
  }

  @keyframes pulse {
    0% { opacity: 0.6; }
    100% { opacity: 1; }
  }

  .link-button {
    background: none;
    border: none;
    color: #007bff;
    text-decoration: underline;
    cursor: pointer;
    font: inherit;
    padding: 0;
  }

  .link-button:hover {
    color: #0056b3;
  }
</style>
