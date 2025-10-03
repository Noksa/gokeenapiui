<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { RouterConfig } from '../types';
  import RouterAccessSection from './RouterAccessSection.svelte';
  import FormSection from './FormSection.svelte';
  import { TestRouterConnection } from '../../wailsjs/go/main/App.js';
  import { handleError } from '../utils/errorHandler';

  export let routerConfig: RouterConfig;
  export let actionTitle: string = 'Выполнить действие';
  export let actionIcon: string = '⚡';

  const dispatch = createEventDispatcher<{
    'proceed': void;
    'cancel': void;
  }>();

  let fieldErrors = {
    url: false,
    login: false,
    password: false
  };

  let isChecking = false;
  let connectionStatus: 'idle' | 'success' | 'error' = 'idle';
  let statusMessage = '';

  function clearFieldError(field: 'url' | 'login' | 'password') {
    fieldErrors[field] = false;
    // Сбрасываем статус при изменении данных
    if (connectionStatus !== 'idle') {
      connectionStatus = 'idle';
      statusMessage = '';
    }
  }

  async function handleProceed() {
    let hasErrors = false;

    if (!routerConfig.url.trim()) {
      fieldErrors.url = true;
      hasErrors = true;
    }
    if (!routerConfig.login.trim()) {
      fieldErrors.login = true;
      hasErrors = true;
    }
    if (!routerConfig.password.trim()) {
      fieldErrors.password = true;
      hasErrors = true;
    }

    if (hasErrors) return;

    isChecking = true;
    connectionStatus = 'idle';
    statusMessage = '';

    try {
      console.log('Проверяем соединение с роутером:', routerConfig);
      await TestRouterConnection(routerConfig);
      connectionStatus = 'success';
      statusMessage = 'Соединение установлено успешно!';
      
      // Автоматически продолжаем через 1 секунду
      setTimeout(() => {
        dispatch('proceed');
      }, 1000);
      
    } catch (error) {
      console.error('Ошибка соединения:', error);
      connectionStatus = 'error';
      statusMessage = handleError(error);
    } finally {
      isChecking = false;
    }
  }

  function handleCancel() {
    dispatch('cancel');
  }
</script>

<div class="modal-overlay" on:click={handleCancel} on:keydown={(e) => e.key === 'Escape' && handleCancel()}>
  <div class="modal-content" on:click|stopPropagation on:keydown|stopPropagation>
    <div class="modal-header">
      <div class="header-icon">{actionIcon}</div>
      <h2>Установка связи с роутером</h2>
      <p class="header-description">
        Для выполнения действия "{actionTitle}" необходимо подключение к роутеру
      </p>
    </div>

    <FormSection>
      <RouterAccessSection 
        bind:routerConfig={routerConfig}
        bind:fieldErrors={fieldErrors}
        {clearFieldError}
      />
    </FormSection>

    {#if connectionStatus !== 'idle'}
      <div class="status-section">
        {#if connectionStatus === 'success'}
          <div class="status success">
            <span class="status-icon">✅</span>
            <span class="status-text">{statusMessage}</span>
          </div>
        {:else if connectionStatus === 'error'}
          <div class="status error">
            <span class="status-icon">❌</span>
            <span class="status-text">{statusMessage}</span>
          </div>
        {/if}
      </div>
    {/if}

    <div class="modal-actions">
      <button type="button" class="btn secondary" on:click={handleCancel}>
        <span class="btn-icon">❌</span>
        Отмена
      </button>
      <button 
        type="button" 
        class="btn primary" 
        on:click={handleProceed}
        disabled={isChecking || connectionStatus === 'success'}
      >
        {#if isChecking}
          <span class="btn-icon spinner">⏳</span>
          Проверка...
        {:else if connectionStatus === 'success'}
          <span class="btn-icon">✅</span>
          Подключено
        {:else}
          <span class="btn-icon">{actionIcon}</span>
          Подключиться
        {/if}
      </button>
    </div>
  </div>
</div>

<style>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    backdrop-filter: blur(5px);
    animation: fadeIn 0.3s ease-out;
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  .modal-content {
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.9));
    border-radius: 20px;
    padding: 30px;
    max-width: 500px;
    width: 90%;
    max-height: 90vh;
    overflow-y: auto;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.2),
      0 0 0 1px rgba(255, 255, 255, 0.5);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
    animation: slideIn 0.4s ease-out;
  }

  @keyframes slideIn {
    from {
      opacity: 0;
      transform: translateY(30px) scale(0.95);
    }
    to {
      opacity: 1;
      transform: translateY(0) scale(1);
    }
  }

  .modal-header {
    text-align: center;
    margin-bottom: 25px;
  }

  .header-icon {
    font-size: 3em;
    margin-bottom: 10px;
    animation: float 3s ease-in-out infinite;
  }

  @keyframes float {
    0%, 100% { 
      transform: translateY(0px);
    }
    50% { 
      transform: translateY(-8px);
    }
  }

  h2 {
    color: #2a5298;
    margin: 0 0 10px 0;
    font-size: 1.4em;
    font-weight: 600;
  }

  .header-description {
    color: #6b7280;
    margin: 0;
    font-size: 0.95em;
    line-height: 1.5;
  }

  .status-section {
    margin: 20px 0;
  }

  .status {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 12px 16px;
    border-radius: 10px;
    font-size: 0.95em;
    font-weight: 500;
    animation: slideIn 0.3s ease-out;
  }

  .status.success {
    background: rgba(5, 150, 105, 0.1);
    color: #059669;
    border: 1px solid rgba(5, 150, 105, 0.3);
  }

  .status.error {
    background: rgba(220, 53, 69, 0.1);
    color: #dc3545;
    border: 1px solid rgba(220, 53, 69, 0.3);
  }

  .status-icon {
    font-size: 1.2em;
  }

  .status-text {
    flex: 1;
  }

  .modal-actions {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 25px;
  }

  .btn {
    padding: 12px 24px;
    border: none;
    border-radius: 10px;
    font-size: 1em;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 8px;
    min-width: 120px;
    justify-content: center;
  }

  .btn.primary {
    background: linear-gradient(135deg, #2a5298, #4a90e2);
    color: white;
    box-shadow: 0 4px 15px rgba(42, 82, 152, 0.4);
  }

  .btn.primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(42, 82, 152, 0.5);
  }

  .btn.secondary {
    background: linear-gradient(135deg, #6b7280, #9ca3af);
    color: white;
    box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3);
  }

  .btn.secondary:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(107, 114, 128, 0.4);
    background: linear-gradient(135deg, #4b5563, #6b7280);
    color: white;
  }

  .btn:disabled {
    opacity: 0.7;
    cursor: not-allowed;
    transform: none;
  }

  .btn:disabled:hover {
    transform: none;
    box-shadow: 0 4px 15px rgba(42, 82, 152, 0.4);
  }

  .spinner {
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    from { transform: rotate(0deg); }
    to { transform: rotate(360deg); }
  }

  .btn-icon {
    font-size: 1.1em;
  }
</style>
