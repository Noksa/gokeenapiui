<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { RouterConfig, AWGConfig } from '../types';
  import { OpenFileDialog } from '../../wailsjs/go/main/App.js';

  export let routerConfig: RouterConfig;
  export let awgConfig: AWGConfig;
  export let isProcessing: boolean = false;

  const dispatch = createEventDispatcher<{
    'submit': void;
    'back': void;
  }>();

  async function selectAWGFile() {
    try {
      const result = await OpenFileDialog();
      
      if (result) {
        awgConfig.filePath = result;
        // Extract filename for connection name
        const filename = result.split('/').pop()?.replace('.conf', '') || '';
        awgConfig.name = filename;
      }
    } catch (error) {
      console.error('Error selecting file:', error);
    }
  }

  function handleSubmit() {
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
  
  <p class="form-description">
    Для добавления и настройки нового AWG соединения в keenetic роутере введите данные ниже.<br>
    Если Ваш keenetic роутер находится в локальной сети (Вы подключены к его Wi-Fi сети),
    то к нему можно обратиться по внутреннему IP адресу и HTTP протоколу.<br>
    Если роутер доступен через интернет, к нему можно обратиться через KeenDNS имя и HTTPS протокол<br><br>
    <strong>Примеры:</strong><br>
    Внутренний адрес: <code>http://192.168.1.1</code><br>
    Внешний адрес: <code>https://super-keenetic.keenetic.pro</code>
  </p>
  
  <form on:submit|preventDefault={handleSubmit}>
    <div class="form-section">
      <h3>🌐 Доступ к роутеру</h3>
      
      <div class="form-group">
        <label for="router-url">
          <span class="label-icon">🔗</span>
          URL роутера
        </label>
        <input 
          id="router-url"
          type="text" 
          bind:value={routerConfig.url}
          placeholder="IP или DNS имя (протокол http/https) роутера"
          required
        />
      </div>
      
      <div class="form-row">
        <div class="form-group">
          <label for="router-login">
            <span class="label-icon">👤</span>
            Логин
          </label>
          <input 
            id="router-login"
            type="text" 
            bind:value={routerConfig.login}
            placeholder="Логин администратора"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="router-password">
            <span class="label-icon">🔐</span>
            Пароль
          </label>
          <input 
            id="router-password"
            type="password" 
            bind:value={routerConfig.password}
            placeholder="Пароль администратора"
            required
          />
        </div>
      </div>
    </div>

    <div class="form-section">
      <h3>🔒 Настройка VPN</h3>
      
      <div class="form-group">
        <label for="connection-name">
          <span class="label-icon">📝</span>
          Имя соединения
        </label>
        <input 
          id="connection-name"
          type="text" 
          bind:value={awgConfig.name}
          placeholder="Имя создаваемого соединения. Опционально, по умолчанию равно имени файла"
        />
      </div>
      
      <div class="form-group">
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
            <button type="button" id="awg-file-btn" class="btn secondary small" on:click={selectAWGFile}>
              Выбрать другой файл
            </button>
          </div>
        {:else}
          <button type="button" id="awg-file-btn" class="btn network file-btn" on:click={selectAWGFile}>
            <span class="btn-icon">📁</span>
            Выбрать файл
          </button>
        {/if}
      </div>
    </div>
    
    <div class="button-group">
      <button type="submit" class="btn primary pulse" disabled={isProcessing}>
        <span class="btn-icon">⚡</span>
        {isProcessing ? 'Создание...' : 'Создать'}
        <div class="btn-glow"></div>
      </button>
      <button type="button" class="btn secondary" on:click={() => dispatch('back')}>
        <span class="btn-icon">↩️</span>
        Выход
      </button>
    </div>
  </form>
</div>

<style>
  .form-container {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 50px;
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
    margin-bottom: 30px;
  }

  .router-icon {
    color: #2a5298;
    margin-bottom: 15px;
    animation: float 3s ease-in-out infinite;
  }

  h2 {
    color: #1e3c72;
    margin: 0;
    font-size: 2.2em;
    font-weight: 700;
  }

  h3 {
    color: #2a5298;
    font-size: 1.3em;
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .form-description {
    background: linear-gradient(135deg, rgba(42, 82, 152, 0.1), rgba(30, 60, 114, 0.05));
    padding: 20px;
    border-radius: 12px;
    border-left: 4px solid #2a5298;
    margin-bottom: 30px;
    line-height: 1.6;
    font-size: 0.8em;
  }

  code {
    background: rgba(42, 82, 152, 0.1);
    padding: 2px 6px;
    border-radius: 4px;
    font-family: 'Monaco', 'Menlo', monospace;
    font-size: 0.9em;
  }

  .form-section {
    margin-bottom: 35px;
    padding: 25px;
    background: rgba(255, 255, 255, 0.5);
    border-radius: 15px;
    border: 1px solid rgba(42, 82, 152, 0.1);
  }

  .form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 20px;
  }

  .form-group {
    margin-bottom: 20px;
  }

  label {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;
    font-weight: 600;
    color: #2a5298;
    font-size: 1.05em;
  }

  .label-icon {
    font-size: 1.1em;
  }

  input {
    width: 100%;
    padding: 14px 16px;
    border: 2px solid rgba(42, 82, 152, 0.2);
    border-radius: 10px;
    font-size: 16px;
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
    align-items: center;
    justify-content: space-between;
    padding: 15px;
    background: rgba(42, 82, 152, 0.05);
    border-radius: 10px;
    border: 2px solid rgba(42, 82, 152, 0.2);
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
    font-size: 0.9em;
    color: #2a5298;
    word-break: break-all;
  }

  .file-btn {
    width: 100%;
    justify-content: center;
  }

  .button-group {
    display: flex;
    gap: 20px;
    justify-content: center;
    margin-top: 40px;
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
    .form-row {
      grid-template-columns: 1fr;
    }
    
    .button-group {
      flex-direction: column;
    }
  }
</style>
