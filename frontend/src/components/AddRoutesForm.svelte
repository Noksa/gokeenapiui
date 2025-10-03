<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { RouterConfig, RouteConfig } from '../types';
  import { OpenBatFileDialog } from '../../wailsjs/go/main/App.js';
  import { BrowserOpenURL } from '../../wailsjs/runtime/runtime.js';
  import FormSection from './FormSection.svelte';
  import RouterInfo from './RouterInfo.svelte';

  export let routerConfig: RouterConfig;
  export let routeConfig: RouteConfig;
  export let isProcessing: boolean = false;
  export let interfaces: any[] = [];
  export let isLoadingInterfaces: boolean = false;

  const dispatch = createEventDispatcher<{
    'submit': void;
    'back': void;
  }>();

  let fieldErrors = {
    url: false,
    login: false,
    password: false,
    interfaceId: false,
    routes: false
  };

  let newUrl = '';

  async function selectBatFiles() {
    try {
      const result = await OpenBatFileDialog();
      
      if (result && !routeConfig.batFiles.includes(result)) {
        routeConfig.batFiles = [...routeConfig.batFiles, result];
        clearFieldError('routes');
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
      clearFieldError('routes');
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
      routes: false
    };
    
    // Проверка обязательных полей
    let hasErrors = false;
    
    // ID интерфейса обязателен
    if (!routeConfig.interfaceId.trim()) {
      fieldErrors.interfaceId = true;
      hasErrors = true;
    }
    
    // Хотя бы один BAT файл ИЛИ одна URL ссылка должны быть указаны
    if (!routeConfig.batFiles.length && !routeConfig.batUrls.length) {
      fieldErrors.routes = true;
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
      
      <div class="info-block">
        <div class="info-header">
          <span class="info-icon">💡</span>
          <span class="info-title">Зачем нужны маршруты?</span>
        </div>
        <div class="info-content">
            Маршруты определяют, какой трафик будет проходить через VPN соединение.
            BAT файлы и ссылки содержат списки IP-адресов или подсетей сайтов и сервисов,
            которые должны быть доступны через туннель.
        </div>
          <div class="info-content">
              Добавить маршруты можно через BAT-файлы на компьютере или через ссылки, по которым можно скачать BAT-файлы.
          </div>
      </div>
    </div>
    
    <FormSection title="Настройка маршрутов" icon="🛣️">
      <div class="interface-section">
        <div class="section-header">
          <h4>🔌 ID интерфейса</h4>
          <div class="interface-input-container">
            <select 
              id="interface-id"
              bind:value={routeConfig.interfaceId}
              on:change={() => clearFieldError('interfaceId')}
              class="interface-select"
              class:error={fieldErrors.interfaceId}
              disabled={isLoadingInterfaces}
            >
              <option value="">
                {isLoadingInterfaces ? 'Загрузка интерфейсов...' : 'Выберите интерфейс'}
              </option>
              {#each interfaces as iface}
                <option value={iface.Id}>
                  {iface.Id} {iface.Description ? `- ${iface.Description}` : ''}
                </option>
              {/each}
            </select>
          </div>
        </div>
      </div>

      <div class="routes-config">
        <div class="config-section">
          <div class="section-header">
            <h4>
              📄 BAT файлы
              <div class="info-tooltip">
                <span class="tooltip-icon">ℹ️</span>
                <div class="tooltip-content">
                    Локальные файлы с расширением .bat, содержащие маршруты для добавления в роутер<br>
                    Примеры BAT файлов можно найти по ссылке:<br>
                    <button type="button" class="link-button" on:click={() => BrowserOpenURL('https://github.com/Noksa/gokeenapi/tree/main/batfiles')}>
                      https://github.com/Noksa/gokeenapi/tree/main/batfiles
                    </button>
                </div>
              </div>
            </h4>
            <button type="button" class="btn primary compact" on:click={selectBatFiles}>
              <span class="btn-icon">📁</span>
              Добавить файл
              {#if routeConfig.batFiles.length > 0}
                <span class="count-badge">{routeConfig.batFiles.length}</span>
              {/if}
            </button>
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
          <div class="section-header">
            <h4>
              🔗 BAT ссылки
              <div class="info-tooltip">
                <span class="tooltip-icon">ℹ️</span>
                <div class="tooltip-content">
                    URL-адреса для загрузки BAT файлов из интернета. Файлы будут скачаны и загружены в роутер автоматически<br>
                    Пример ссылки:<br>
                    <button type="button" class="link-button" on:click={() => BrowserOpenURL('https://iplist.opencck.org/?format=bat&data=cidr4&site=youtube.com')}>
                      https://iplist.opencck.org/?format=bat&data=cidr4&site=youtube.com
                    </button>
              </div>
            </h4>
            <div class="url-input-container">
              <input 
                type="url" 
                bind:value={newUrl}
                placeholder="Введите URL..."
                class="url-input"
                on:keydown={(e) => e.key === 'Enter' && (e.preventDefault(), addUrl())}
              />
              <button type="button" class="btn secondary compact add-btn" on:click={addUrl}>
                <span class="btn-icon">➕</span>
              </button>
            </div>
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

      {#if !routeConfig.batFiles.length && !routeConfig.batUrls.length && fieldErrors.routes}
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
  
  <RouterInfo routerUrl={routerConfig.url} />
</div>

<style>
  .form-container {
    position: relative;
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
    margin: 0 0 20px 0;
    font-size: 1.4em;
    font-weight: 600;
  }

  .info-block {
    background: linear-gradient(135deg, rgba(59, 130, 246, 0.05), rgba(16, 185, 129, 0.05));
    border: 1px solid rgba(59, 130, 246, 0.2);
    border-radius: 12px;
    padding: 16px;
    margin-bottom: 24px;
  }

  .info-header {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;
  }

  .info-icon {
    font-size: 1.2em;
  }

  .info-title {
    font-weight: 600;
    color: #2a5298;
    font-size: 0.95em;
  }

  .info-content {
    font-size: 0.9em;
    color: #64748b;
    line-height: 1.5;
    margin-left: 28px;
  }

  .form-content {
    text-align: center;
    margin-bottom: 20px;
  }

  .routes-config {
    display: flex;
    flex-direction: column;
    gap: 20px;
    margin-top: 15px;
  }

  .interface-section {
    margin-bottom: 10px;
  }

  .interface-input-container {
    max-width: 250px;
  }

  .interface-select {
    width: 100%;
    padding: 8px 12px;
    border: 2px solid rgba(42, 82, 152, 0.2);
    border-radius: 8px;
    font-size: 0.9em;
    color: #2a5298;
    transition: all 0.3s ease;
    background: rgba(255, 255, 255, 0.9);
    box-sizing: border-box;
    cursor: pointer;
  }

  .interface-select:focus {
    outline: none;
    border-color: #2a5298;
    box-shadow: 0 0 0 3px rgba(42, 82, 152, 0.1);
    background: white;
  }

  .interface-select:disabled {
    background: #f9fafb;
    color: #9ca3af;
    cursor: not-allowed;
  }

  .interface-select.error {
    border-color: #dc3545;
    box-shadow: 0 0 0 3px rgba(220, 53, 69, 0.2);
  }

  .config-section {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 15px;
  }

  .config-section h4 {
    margin: 0;
    font-size: 0.95em;
    color: #2a5298;
    font-weight: 600;
    flex-shrink: 0;
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .info-tooltip {
    position: relative;
    display: inline-block;
  }

  .tooltip-icon {
    font-size: 0.9em;
    cursor: help;
    opacity: 0.6;
    transition: opacity 0.3s ease;
  }

  .tooltip-icon:hover {
    opacity: 1;
  }

  .tooltip-content {
    position: absolute;
    top: 100%;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(42, 82, 152, 1);
    color: white;
    padding: 15px;
    border-radius: 10px;
    font-size: 0.85em;
    visibility: hidden;
    opacity: 0;
    z-index: 1000;
    margin-top: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
    transition: opacity 0.3s, visibility 0.3s;
    white-space: normal;
    max-width: 500px;
    width: max-content;
    line-height: 1.4;
  }

  .tooltip-content::before {
    content: '';
    position: absolute;
    top: -6px;
    left: 50%;
    transform: translateX(-50%);
    border: 6px solid transparent;
    border-bottom-color: rgba(42, 82, 152, 0.95);
  }

  .tooltip-content .link-button {
    background: none;
    border: none;
    color: #87ceeb;
    text-decoration: underline;
    cursor: pointer;
    padding: 0;
    font-size: inherit;
    font-family: inherit;
  }

  .tooltip-content .link-button:hover {
    color: #add8e6;
    text-decoration: none;
  }

  .info-tooltip:hover .tooltip-content {
    visibility: visible;
    opacity: 1;
  }

  .tooltip-content::after {
    content: '';
    position: absolute;
    top: 100%;
    left: 20px;
    border: 6px solid transparent;
    border-top-color: rgba(0, 0, 0, 0.9);
  }

  .url-input-container {
    display: flex;
    gap: 8px;
    align-items: center;
    flex: 1;
    max-width: 300px;
  }

  .url-input {
    flex: 1;
    padding: 8px 12px;
    border: 2px solid rgba(42, 82, 152, 0.2);
    border-radius: 8px;
    font-size: 0.9em;
    transition: all 0.3s ease;
    background: rgba(255, 255, 255, 0.9);
    min-width: 0;
  }

  .url-input:focus {
    outline: none;
    border-color: #2a5298;
    box-shadow: 0 0 0 3px rgba(42, 82, 152, 0.1);
    background: white;
  }

  .btn.compact {
    padding: 8px 16px;
    font-size: 0.9em;
    position: relative;
  }

  .add-btn {
    padding: 8px 12px;
    border-radius: 8px;
    min-width: 40px;
  }

  .count-badge {
    position: absolute;
    top: -6px;
    right: -6px;
    background: #059669;
    color: white;
    border-radius: 10px;
    padding: 2px 6px;
    font-size: 0.7em;
    font-weight: 600;
    min-width: 18px;
    text-align: center;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
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
    
    .button-group {
      flex-direction: column;
    }
  }
</style>
