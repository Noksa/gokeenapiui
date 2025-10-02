<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher<{
    'create-awg': void;
    'quit': void;
  }>();
</script>

<div class="welcome-container">
  <h1>Gokeenapi</h1>
  <p class="description">
    Утилита для работы с роутерами Keenetic через REST API<br><br>
    (!) Данная версия является GUI версией и содержит не весь функционал<br><br>
    Если Вам нужен весь функционал, воспользуйтесь 
    <a href="https://github.com/Noksa/gokeenapi" target="_blank">CLI версией</a>
  </p>
  
  <div class="button-group">
    <button class="btn primary" on:click={() => dispatch('create-awg')}>
      Создать AWG соединение
    </button>
    <button class="btn secondary" on:click={() => dispatch('quit')}>
      Выход
    </button>
  </div>
</div>

<style>
  .welcome-container {
    background: white;
    border-radius: 12px;
    padding: 40px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    max-width: 600px;
    width: 100%;
  }

  h1 {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
    font-size: 2.5em;
  }

  .description {
    text-align: center;
    color: #666;
    line-height: 1.6;
    margin-bottom: 30px;
  }

  .button-group {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 30px;
  }

  a {
    color: #667eea;
    text-decoration: none;
  }

  a:hover {
    text-decoration: underline;
  }
</style>
