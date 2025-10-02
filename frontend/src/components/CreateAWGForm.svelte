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
        // Extract filename for default connection name if not set
        if (!awgConfig.name) {
          const filename = result.split('/').pop()?.replace('.conf', '') || '';
          awgConfig.name = filename;
        }
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
  <h2>🛜 Создание AWG соединения</h2>
  <p class="form-description">
    Для добавления и настройки нового AWG соединения в keenetic роутере введите данные ниже.<br>
    Если Ваш keenetic роутер находится в локальной сети (Вы подключены к его Wi-Fi сети),
    то к нему можно обратиться по внутреннему IP адресу и HTTP протоколу.<br>
    Если роутер доступен через интернет, к нему можно обратиться через KeenDNS имя и HTTPS протокол<br><br>
    Примеры:<br>
    Внутренний адрес: http://192.168.1.1<br>
    Внешний адрес: https://super-keenetic.keenetic.pro
  </p>
  
  <form on:submit|preventDefault={handleSubmit}>
    <div class="form-group">
      <label for="router-url">URL роутера</label>
      <input 
        id="router-url"
        type="text" 
        bind:value={routerConfig.url}
        placeholder="IP или DNS имя (протокол http/https) роутера"
        required
      />
    </div>
    
    <div class="form-group">
      <label for="router-login">Логин</label>
      <input 
        id="router-login"
        type="text" 
        bind:value={routerConfig.login}
        placeholder="Логин администратора"
        required
      />
    </div>
    
    <div class="form-group">
      <label for="router-password">Пароль</label>
      <input 
        id="router-password"
        type="password" 
        bind:value={routerConfig.password}
        placeholder="Пароль администратора"
        required
      />
    </div>
    
    <div class="form-group">
      <label for="connection-name">Имя соед.</label>
      <input 
        id="connection-name"
        type="text" 
        bind:value={awgConfig.name}
        placeholder="Имя создаваемого соединения. Опционально, по умолчанию равно имени файла"
      />
    </div>
    
    <div class="form-group">
      <label for="awg-file">AWG конф. файл</label>
      {#if awgConfig.filePath}
        <div class="file-selected">
          <span class="file-path">{awgConfig.filePath}</span>
          <button type="button" class="btn secondary small" on:click={selectAWGFile}>
            Выбрать другой файл
          </button>
        </div>
      {:else}
        <button type="button" class="btn secondary" on:click={selectAWGFile}>
          📁 Выбрать файл
        </button>
      {/if}
    </div>
    
    <div class="button-group">
      <button type="submit" class="btn primary" disabled={isProcessing}>
        Создать
      </button>
      <button type="button" class="btn secondary" on:click={() => dispatch('back')}>
        Выход
      </button>
    </div>
  </form>
</div>

<style>
  .form-container {
    background: white;
    border-radius: 12px;
    padding: 40px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    max-width: 600px;
    width: 100%;
  }

  h2 {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
    font-size: 1.8em;
  }

  .form-description {
    text-align: left;
    color: #666;
    line-height: 1.6;
    margin-bottom: 30px;
    font-size: 0.9em;
    background: #f8f9fa;
    padding: 15px;
    border-radius: 8px;
    border-left: 4px solid #667eea;
  }

  .form-group {
    margin-bottom: 20px;
  }

  label {
    display: block;
    margin-bottom: 5px;
    font-weight: 600;
    color: #333;
  }

  input {
    width: 100%;
    padding: 12px;
    border: 2px solid #e1e5e9;
    border-radius: 8px;
    font-size: 16px;
    transition: border-color 0.3s;
    box-sizing: border-box;
  }

  input:focus {
    outline: none;
    border-color: #667eea;
  }

  .file-selected {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 10px;
    background: #f8f9fa;
    border-radius: 8px;
    border: 2px solid #e1e5e9;
  }

  .file-path {
    flex: 1;
    font-family: monospace;
    font-size: 0.9em;
    color: #333;
    word-break: break-all;
  }

  .button-group {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 30px;
  }
</style>
