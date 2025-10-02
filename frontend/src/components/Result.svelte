<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let type: 'success' | 'error';
  export let message: string;

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
  <div class="result-icon">
    {#if isSuccess}
      <div class="success-icon">✓</div>
    {:else}
      <div class="error-icon">✕</div>
    {/if}
  </div>
  
  <h2>{title}</h2>
  
  {#if isSuccess}
    <p class="result-message">{message}</p>
    <div class="button-group">
      <button class="btn primary glow" on:click={() => dispatch('open-router')}>
        <span class="btn-icon">🌐</span>
        Открыть веб-интерфейс роутера
      </button>
      <button class="btn secondary" on:click={() => dispatch('quit')}>
        Назад
      </button>
    </div>
  {:else}
    <p class="error-message">{message}</p>
    <div class="button-group">
      <button class="btn primary pulse" on:click={() => dispatch('retry')}>
        <span class="btn-icon">🔄</span>
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
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 50px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.3),
      0 0 0 1px rgba(255, 255, 255, 0.2);
    max-width: 600px;
    width: 100%;
    text-align: center;
    position: relative;
    overflow: hidden;
    animation: slideIn 0.5s ease-out;
  }

  .result-container::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    animation: pulse 2s ease-in-out infinite alternate;
  }

  .result-container.success::before {
    background: linear-gradient(90deg, #28a745, #20c997, #17a2b8);
  }

  .result-container.error::before {
    background: linear-gradient(90deg, #dc3545, #fd7e14, #ffc107);
  }

  .result-icon {
    margin-bottom: 20px;
    animation: bounceIn 0.8s ease-out;
  }

  .success-icon {
    width: 80px;
    height: 80px;
    background: linear-gradient(135deg, #28a745, #20c997);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 40px;
    color: white;
    margin: 0 auto;
    box-shadow: 0 10px 30px rgba(40, 167, 69, 0.3);
    animation: successPulse 2s ease-in-out infinite;
  }

  .error-icon {
    width: 80px;
    height: 80px;
    background: linear-gradient(135deg, #dc3545, #fd7e14);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 40px;
    color: white;
    margin: 0 auto;
    box-shadow: 0 10px 30px rgba(220, 53, 69, 0.3);
    animation: shake 0.5s ease-in-out;
  }

  h2 {
    color: #1e3c72;
    margin-bottom: 30px;
    font-size: 2em;
    font-weight: 600;
  }

  .result-message {
    color: #333;
    line-height: 1.6;
    white-space: pre-line;
    font-size: 1.1em;
    margin-bottom: 30px;
  }

  .error-message {
    color: #dc3545;
    background: linear-gradient(135deg, rgba(220, 53, 69, 0.1), rgba(253, 126, 20, 0.05));
    padding: 20px;
    border-radius: 12px;
    border-left: 4px solid #dc3545;
    font-family: monospace;
    font-size: 0.9em;
    word-break: break-word;
    margin-bottom: 30px;
    animation: errorGlow 2s ease-in-out infinite alternate;
  }

  .button-group {
    display: flex;
    gap: 15px;
    justify-content: center;
    flex-wrap: wrap;
  }

  .btn {
    padding: 12px 24px;
    border: none;
    border-radius: 12px;
    font-size: 1em;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 8px;
    position: relative;
    overflow: hidden;
  }

  .btn-icon {
    font-size: 1.2em;
  }

  .btn.primary {
    background: linear-gradient(135deg, #2a5298, #1e3c72);
    color: white;
    box-shadow: 0 8px 20px rgba(42, 82, 152, 0.3);
  }

  .btn.primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 12px 30px rgba(42, 82, 152, 0.4);
  }

  .btn.secondary {
    background: rgba(108, 117, 125, 0.1);
    color: #6c757d;
    border: 2px solid rgba(108, 117, 125, 0.2);
  }

  .btn.secondary:hover {
    background: rgba(108, 117, 125, 0.2);
    transform: translateY(-1px);
  }

  .btn.glow {
    animation: glow 2s ease-in-out infinite alternate;
  }

  .btn.pulse {
    animation: buttonPulse 1.5s ease-in-out infinite;
  }

  @keyframes slideIn {
    from {
      opacity: 0;
      transform: translateY(30px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes bounceIn {
    0% {
      opacity: 0;
      transform: scale(0.3);
    }
    50% {
      transform: scale(1.1);
    }
    100% {
      opacity: 1;
      transform: scale(1);
    }
  }

  @keyframes successPulse {
    0%, 100% {
      box-shadow: 0 10px 30px rgba(40, 167, 69, 0.3);
    }
    50% {
      box-shadow: 0 15px 40px rgba(40, 167, 69, 0.5);
    }
  }

  @keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-5px); }
    75% { transform: translateX(5px); }
  }

  @keyframes errorGlow {
    0% {
      box-shadow: 0 0 5px rgba(220, 53, 69, 0.2);
    }
    100% {
      box-shadow: 0 0 20px rgba(220, 53, 69, 0.4);
    }
  }

  @keyframes pulse {
    0% { opacity: 0.6; }
    100% { opacity: 1; }
  }

  @keyframes glow {
    0% {
      box-shadow: 0 8px 20px rgba(42, 82, 152, 0.3);
    }
    100% {
      box-shadow: 0 8px 30px rgba(42, 82, 152, 0.6);
    }
  }

  @keyframes buttonPulse {
    0%, 100% {
      transform: scale(1);
    }
    50% {
      transform: scale(1.05);
    }
  }
</style>
