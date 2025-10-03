<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { RouterConfig, RouteConfig } from '../types';
  import { OpenFileDialog } from '../../wailsjs/go/main/App.js';
  import RouterAccessSection from './RouterAccessSection.svelte';
  import FormSection from './FormSection.svelte';

  export let routerConfig: RouterConfig;
  export let routeConfig: RouteConfig;
  export let isProcessing: boolean = false;

  const dispatch = createEventDispatcher<{
    'submit': void;
    'back': void;
  }>();

  let fieldErrors = {
    url: false,
    login: false,
    password: false,
    interfaceId: false,
    batFiles: false
  };

  let newUrl = '';

  async function selectBatFiles() {
    try {
      const result = await OpenFileDialog();
      
      if (result && !routeConfig.batFiles.includes(result)) {
        routeConfig.batFiles = [...routeConfig.batFiles, result];
        clearFieldError('batFiles');
      }
    } catch (error) {
      console.error('Error selecting file:', error);
    }
  }

  function removeBatFile(index: number) {
    routeConfig.batFiles = routeConfig.batFiles.filter((_, i) => i !== index);
  }

  function addUrl() {
    if (newUrl.trim() && !routeConfig.batUrls.includes(newUrl.trim())) {
      routeConfig.batUrls = [...routeConfig.batUrls, newUrl.trim()];
      newUrl = '';
    }
  }

  function removeUrl(index: number) {
    routeConfig.batUrls = routeConfig.batUrls.filter((_, i) => i !== index);
  }

  function clearFieldError(field: keyof typeof fieldErrors) {
    fieldErrors[field] = false;
  }

  function handleSubmit() {
    // Сброс ошибок
    fieldErrors = {
      url: false,
      login: false,
      password: false,
      interfaceId: false,
      batFilePath: false
    };
    
    // Проверка обязательных полей
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
    if (!routeConfig.interfaceId.trim()) {
      fieldErrors.interfaceId = true;
      hasErrors = true;
    }
    if (!routeConfig.batFiles.length && !routeConfig.batUrls.length) {
      fieldErrors.batFiles = true;
      hasErrors = true;
    }
    
    if (hasErrors) return;
    
    dispatch('submit');
  }
</script>

<div class="form-container">
  <form on:submit|preventDefault={handleSubmit}>
    <div class="form-content">
      <h2>Добавление маршрутов</h2>
    </div>
    
    <FormSection>
      <RouterAccessSection 
        bind:routerConfig={routerConfig}
        bind:fieldErrors={fieldErrors}
      />
    </FormSection>

    <FormSection title="Настройка маршрутов" icon="🛣️">
      <div class="form-group">
        <label for="interface-id">
          <span class="label-icon">🔌</span>
          ID интерфейса
        </label>
        <input 
          id="interface-id"
          type="text" 
          bind:value={routeConfig.interfaceId}
          on:input={() => clearFieldError('interfaceId')}
          placeholder="Например: Wireguard0"
          class:error={fieldErrors.interfaceId}
        />
      </div>

      <div class="routes-config">
        <div class="config-section">
          <h4>📄 BAT файлы</h4>
          <div class="input-row">
            <button type="button" class="btn primary" on:click={selectBatFiles}>
              <span class="btn-icon">📁</span>
              Добавить файл
            </button>
            {#if routeConfig.batFiles.length > 0}
              <span class="count-badge">{routeConfig.batFiles.length}</span>
            {/if}
          </div>
          
          {#if routeConfig.batFiles.length > 0}
            <div class="items-list">
              {#each routeConfig.batFiles as file, index}
                <div class="item">
                  <span class="item-icon">📄</span>
                  <span class="item-text">{file.split('/').pop()}</span>
                  <button type="button" class="remove-btn" on:click={() => removeBatFile(index)}>✕</button>
                </div>
              {/each}
            </div>
          {/if}
        </div>

        <div class="config-section">
          <h4>🔗 URL ссылки</h4>
          <div class="input-row">
            <input 
              type="url" 
              bind:value={newUrl}
              placeholder="https://example.com/routes.bat"
              on:keydown={(e) => e.key === 'Enter' && (e.preventDefault(), addUrl())}
            />
            <button type="button" class="btn secondary" on:click={addUrl}>
              <span class="btn-icon">➕</span>
            </button>
          </div>
          
          {#if routeConfig.batUrls.length > 0}
            <div class="items-list">
              {#each routeConfig.batUrls as url, index}
                <div class="item">
                  <span class="item-icon">🔗</span>
                  <span class="item-text">{url}</span>
                  <button type="button" class="remove-btn" on:click={() => removeUrl(index)}>✕</button>
                </div>
              {/each}
            </div>
          {/if}
        </div>
      </div>

      {#if !routeConfig.batFiles.length && !routeConfig.batUrls.length && fieldErrors.batFiles}
        <div class="error-message">
          Добавьте хотя бы один BAT файл или URL ссылку
        </div>
      {/if}
    </FormSection>
    
    <div class="button-group">
      <button type="submit" class="btn primary pulse" disabled={isProcessing}>
        <span class="btn-icon">⚡</span>
        {isProcessing ? 'Добавление...' : 'Добавить маршруты'}
        <div class="btn-glow"></div>
      </button>
      <button type="button" class="btn secondary" on:click={() => dispatch('back')}>
        <span class="btn-icon">↩️</span>
        Назад
      </button>
    </div>
  </form>
</div>

<style>
  .form-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    padding: 0;
  }

  form {
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.9));
    border-radius: 20px;
    padding: 40px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.1),
      0 0 0 1px rgba(255, 255, 255, 0.5);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
  }

  h2 {
    text-align: center;
    color: #2a5298;
    margin: 0;
    font-size: 1.4em;
    font-weight: 600;
  }

  .form-content {
    text-align: center;
    margin-bottom: 20px;
  }

  .routes-config {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
    margin-top: 15px;
  }

  .config-section {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .config-section h4 {
    margin: 0;
    font-size: 0.9em;
    color: #2a5298;
    font-weight: 600;
  }

  .input-row {
    display: flex;
    gap: 10px;
    align-items: center;
  }

  .input-row input {
    flex: 1;
  }

  .input-row button {
    padding: 10px 16px;
    font-size: 0.9em;
  }

  .count-badge {
    background: #2a5298;
    color: white;
    border-radius: 12px;
    padding: 4px 8px;
    font-size: 0.8em;
    font-weight: 600;
    min-width: 20px;
    text-align: center;
  }

  .items-list {
    display: flex;
    flex-direction: column;
    gap: 6px;
    max-height: 150px;
    overflow-y: auto;
  }

  .item {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 10px;
    background: rgba(42, 82, 152, 0.05);
    border-radius: 6px;
    border: 1px solid rgba(42, 82, 152, 0.15);
    font-size: 0.85em;
  }

  .item-icon {
    font-size: 1em;
    opacity: 0.7;
  }

  .item-text {
    flex: 1;
    color: #2a5298;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .remove-btn {
    background: #dc3545;
    color: white;
    border: none;
    border-radius: 50%;
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    font-size: 0.7em;
    transition: all 0.3s ease;
    flex-shrink: 0;
  }

  .remove-btn:hover {
    background: #c82333;
    transform: scale(1.1);
  }

  .error-message {
    background: rgba(220, 53, 69, 0.1);
    color: #dc3545;
    padding: 10px;
    border-radius: 6px;
    font-size: 0.9em;
    text-align: center;
    margin-top: 15px;
  }

  @media (max-width: 600px) {
    .routes-config {
      grid-template-columns: 1fr;
      gap: 15px;
    }
  }

  .file-btn-container {
    display: flex;
    justify-content: flex-start;
  }

  .file-btn {
    width: auto;
    justify-content: center;
  }

  .file-btn.small {
    font-size: 0.85em;
    padding: 8px 16px;
    opacity: 0.8;
  }

  .button-group {
    display: flex;
    gap: 20px;
    justify-content: center;
    margin-top: 10px;
  }

  .btn {
    position: relative;
    overflow: hidden;
    padding: 12px 24px;
    border: none;
    border-radius: 12px;
    font-size: 1em;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 8px;
    text-decoration: none;
    min-width: 140px;
    justify-content: center;
  }

  .btn-icon {
    font-size: 1.1em;
  }

  .btn.primary {
    background: linear-gradient(135deg, #2a5298, #4a90e2);
    color: white;
    box-shadow: 0 4px 15px rgba(42, 82, 152, 0.4);
  }

  .btn.primary:hover:not(:disabled) {
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
  }

  .btn.network {
    background: linear-gradient(135deg, #059669, #10b981);
    color: white;
    box-shadow: 0 4px 15px rgba(5, 150, 105, 0.4);
  }

  .btn.network:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(5, 150, 105, 0.5);
  }

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none !important;
  }

  .btn.pulse {
    animation: pulse 2s ease-in-out infinite;
  }

  @keyframes pulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.05); }
  }

  .btn-glow {
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
    transition: left 0.5s;
  }

  .btn:hover .btn-glow {
    left: 100%;
  }

  @media (max-width: 600px) {
    .form-row {
      grid-template-columns: 1fr;
    }
    
    .button-group {
      flex-direction: column;
    }
  }
</style>
