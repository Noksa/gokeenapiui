<script lang="ts">
    import {createEventDispatcher} from 'svelte';
    import type {AWGConfig, RouterConfig} from '../types';
    import {OpenFileDialog} from '../../wailsjs/go/main/App.js';
    import FormSection from './FormSection.svelte';
    import RouterInfo from './RouterInfo.svelte';

    export let routerConfig: RouterConfig;
  export let awgConfig: AWGConfig;
  export let isProcessing: boolean = false;

  const dispatch = createEventDispatcher<{
    'submit': void;
    'back': void;
  }>();

  let nameManuallySet = false;
  let fieldErrors = {
    url: false,
    login: false,
    password: false,
    filePath: false
  };

  async function selectAWGFile() {
    try {
      const result = await OpenFileDialog();
      
      if (result) {
        awgConfig.filePath = result;
        clearFieldError('filePath');
        // Extract filename for connection name only if not manually set
        if (!nameManuallySet) {
            awgConfig.name = result.split('/').pop()?.replace('.conf', '') || '';
        }
      }
    } catch (error) {
      console.error('Error selecting file:', error);
    }
  }

  function handleNameInput() {
    nameManuallySet = true;
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
      filePath: false
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
    if (!awgConfig.filePath.trim()) {
      fieldErrors.filePath = true;
      hasErrors = true;
    }
    
    if (hasErrors) return;
    
    dispatch('submit');
  }
</script>

<div class="form-container">
  <div class="form-header">
    <div class="router-icon">
      <svg width="40" height="40" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="3" y="4" width="18" height="16" rx="2" stroke="currentColor" stroke-width="2"/>
        <circle cx="8" cy="8" r="1" fill="currentColor"/>
        <circle cx="12" cy="8" r="1" fill="currentColor"/>
        <circle cx="16" cy="8" r="1" fill="currentColor"/>
        <path d="M7 12h10" stroke="currentColor" stroke-width="2"/>
        <path d="M7 16h6" stroke="currentColor" stroke-width="2"/>
      </svg>
    </div>
    <h2>🛜 Создание AWG соединения</h2>
  </div>
  
  <form on:submit|preventDefault={handleSubmit}>
    <FormSection title="Настройка AWG" icon="🔒">
      
      <div class="form-group">
        <label for="connection-name">
          <span class="label-icon">📝</span>
          Имя соединения
        </label>
        <input 
          id="connection-name"
          type="text" 
          bind:value={awgConfig.name}
          on:input={handleNameInput}
          placeholder="Имя соединения (по умолчанию - имя файла)"
        />
      </div>
      
      <div class="form-group" class:error={fieldErrors.filePath}>
        <label for="awg-file-btn">
          <span class="label-icon">📄</span>
          AWG конфигурационный файл
        </label>
        {#if awgConfig.filePath}
          <div class="file-selected">
            <div class="file-info">
              <span class="file-icon">📄</span>
              <span class="file-path">{awgConfig.filePath}</span>
            </div>
            <button type="button" id="awg-file-btn" class="btn network file-btn small" on:click={selectAWGFile}>
              <span class="btn-icon">📁</span>
              Выбрать другой файл
            </button>
          </div>
        {:else}
          <div class="file-btn-container">
            <button type="button" id="awg-file-btn" class="btn network file-btn" on:click={selectAWGFile}>
              <span class="btn-icon">📁</span>
              Выбрать файл
            </button>
          </div>
        {/if}
      </div>
    </FormSection>
    
    <div class="button-group">
      <button type="submit" class="btn primary pulse" disabled={isProcessing}>
        <span class="btn-icon">⚡</span>
        {isProcessing ? 'Создание...' : 'Создать'}
        <div class="btn-glow"></div>
      </button>
      <button type="button" class="btn secondary" on:click={() => dispatch('back')}>
        <span class="btn-icon">↩️</span>
        Назад
      </button>
    </div>
  </form>
  
  <RouterInfo routerUrl={routerConfig.url} />
</div>

<style>
  .form-container {
    position: relative;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 30px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.3),
      0 0 0 1px rgba(255, 255, 255, 0.2);
    max-width: 700px;
    width: 100%;
    position: relative;
    overflow: hidden;
  }

  .form-container::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: linear-gradient(90deg, #00d4ff, #5b86e5, #36d1dc);
    animation: pulse 2s ease-in-out infinite alternate;
  }

  .form-header {
    text-align: center;
    margin-bottom: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    position: relative;
  }


  @keyframes glow {
    0% { 
      box-shadow: 0 0 5px rgba(255, 193, 7, 0.3);
      background: rgba(255, 193, 7, 0.1);
    }
    100% { 
      box-shadow: 0 0 15px rgba(255, 193, 7, 0.5);
      background: rgba(255, 193, 7, 0.25);
    }
  }
  
  .router-icon {
    color: #2a5298;
    margin-bottom: 15px;
    animation: float 3s ease-in-out infinite;
  }

  h2 {
    color: #1e3c72;
    margin: 0;
    font-size: 1.4em;
    font-weight: 600;
  }

  .form-group {
    margin-bottom: 15px;
  }

  label {
    display: flex;
    align-items: center;
    gap: 6px;
    margin-bottom: 6px;
    font-weight: 600;
    color: #2a5298;
    font-size: 0.9em;
  }

  .label-icon {
    font-size: 1.1em;
  }

  input {
    width: 100%;
    padding: 10px 12px;
    border: 2px solid rgba(42, 82, 152, 0.2);
    border-radius: 8px;
    font-size: 14px;
    transition: all 0.3s;
    box-sizing: border-box;
    background: rgba(255, 255, 255, 0.9);
  }

  input:focus {
    outline: none;
    border-color: #2a5298;
    box-shadow: 0 0 0 3px rgba(42, 82, 152, 0.1);
    background: white;
  }

  .file-selected {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 15px;
    background: rgba(42, 82, 152, 0.05);
    border-radius: 10px;
    border: 2px solid rgba(42, 82, 152, 0.2);
  }

  .file-selected .btn {
    align-self: flex-start;
  }

  .file-info {
    display: flex;
    align-items: center;
    gap: 10px;
    flex: 1;
  }

  .file-icon {
    font-size: 1.2em;
  }

  .file-path {
    font-family: 'Monaco', 'Menlo', monospace;
    font-size: 0.75em;
    color: #2a5298;
    word-break: break-all;
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


  .form-group.error {
    border: 2px solid #dc3545;
    border-radius: 8px;
    padding: 10px;
    background: rgba(220, 53, 69, 0.05);
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

  .btn:hover .btn-glow {
    left: 100%;
  }

  @keyframes float {
    0%, 100% { transform: translateY(0px); }
    50% { transform: translateY(-8px); }
  }

  @keyframes pulse {
    0% { opacity: 0.6; }
    100% { opacity: 1; }
  }

  @media (max-width: 600px) {
    
    .button-group {
      flex-direction: column;
    }
  }
</style>
