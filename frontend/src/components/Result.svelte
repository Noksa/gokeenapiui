<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let type: 'success' | 'error';
  export let message: string;
  export let routerUrl: string = '';

  const dispatch = createEventDispatcher<{
    'retry': void;
    'back': void;
    'open-router': void;
    'quit': void;
  }>();

  $: isSuccess = type === 'success';
  $: title = isSuccess ? '✅ Успех!' : '❌ Ошибка';
</script>

<div class="result-container {type}">
  <h2>{title}</h2>
  
  {#if isSuccess}
    <p class="result-message">{message}</p>
    <div class="button-group">
      <button class="btn primary" on:click={() => dispatch('open-router')}>
        Открыть веб-интерфейс роутера
      </button>
      <button class="btn secondary" on:click={() => dispatch('quit')}>
        Выход
      </button>
    </div>
  {:else}
    <p class="error-message">{message}</p>
    <div class="button-group">
      <button class="btn primary" on:click={() => dispatch('retry')}>
        Попробовать снова
      </button>
      <button class="btn secondary" on:click={() => dispatch('back')}>
        Назад
      </button>
    </div>
  {/if}
</div>

<style>
  .result-container {
    background: white;
    border-radius: 12px;
    padding: 40px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    max-width: 600px;
    width: 100%;
  }

  .result-container.success {
    border-left: 4px solid #28a745;
  }

  .result-container.error {
    border-left: 4px solid #dc3545;
  }

  h2 {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
    font-size: 1.8em;
  }

  .result-message {
    color: #333;
    line-height: 1.6;
    text-align: center;
    white-space: pre-line;
  }

  .error-message {
    color: #dc3545;
    background: #f8d7da;
    padding: 15px;
    border-radius: 8px;
    font-family: monospace;
    font-size: 0.9em;
    word-break: break-word;
  }

  .button-group {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 30px;
  }
</style>
