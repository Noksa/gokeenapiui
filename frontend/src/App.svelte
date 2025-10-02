<script lang="ts">
  import { 
    ValidateRouterConfig, 
    ValidateAWGConfig, 
    CreateAWGInterface, 
    ConfigureAWGInterface, 
    ActivateAWGInterface, 
    GetRouterWebURL 
  } from '../wailsjs/go/main/App.js';
  import { BrowserOpenURL, Quit } from '../wailsjs/runtime';
  
  import Welcome from './components/Welcome.svelte';
  import CreateAWGForm from './components/CreateAWGForm.svelte';
  import Progress from './components/Progress.svelte';
  import Result from './components/Result.svelte';
  
  import type { AppState } from './types';
  import { handleError } from './utils/errorHandler';
  import { 
    createInitialState, 
    updateView, 
    setProgress, 
    setSuccess, 
    setError, 
    resetState,
    setProcessing 
  } from './utils/stateManager';
  
  import './styles/buttons.css';

  // Centralized state management using AppState interface and state manager
  let appState: AppState = createInitialState();

  // Navigation handlers
  function showCreateAWG() {
    appState = updateView(appState, 'create-awg');
  }

  function showWelcome() {
    appState = resetState();
  }

  // Main business logic
  async function createConnection() {
    if (appState.isProcessing) return;
    
    try {
      appState = setProcessing(appState, true);
      
      // Validate form
      if (!appState.awgConfig.filePath) {
        throw new Error('Выберите AWG конфиг файл');
      }
      
      appState = setProgress(appState, 'Проверяем подключение к роутеру...');
      
      // Validate router config
      await ValidateRouterConfig(appState.routerConfig);
      
      appState = setProgress(appState, 'Проверяем AWG конфигурацию...');
      
      // Validate AWG config and authenticate
      await ValidateAWGConfig(appState.routerConfig, appState.awgConfig);
      
      appState = setProgress(appState, 'Создаём AWG интерфейс...');
      
      // Create AWG interface
      const interfaceName = await CreateAWGInterface(appState.awgConfig);
      
      appState = setProgress(appState, 'Настраиваем соединение...');
      
      // Configure AWG interface
      await ConfigureAWGInterface(appState.awgConfig, interfaceName);
      
      appState = setProgress(appState, 'Активируем соединение...');
      
      // Activate AWG interface
      await ActivateAWGInterface(interfaceName);
        
      // Success
      appState = setSuccess(
        appState, 
        `Соединение успешно создано и включено!\nТеперь можно приступить к настройке политик подключения или маршрутизации и наслаждаться VPN!\nУдачи! 🌐`
      );
      
    } catch (error) {
      appState = setError(appState, handleError(error));
    }
  }

  async function openRouterInterface() {
    try {
      const url = await GetRouterWebURL(appState.routerConfig.url);
      BrowserOpenURL(url);
    } catch (error) {
      console.error('Error opening router interface:', error);
    }
  }

  function quit() {
    Quit();
  }
</script>

<main class:center-content={appState.currentView === 'progress' || appState.currentView === 'success' || appState.currentView === 'error'}>
  {#if appState.currentView === 'welcome'}
    <Welcome 
      on:create-awg={showCreateAWG}
      on:quit={quit}
    />
  
  {:else if appState.currentView === 'create-awg'}
    <CreateAWGForm 
      bind:routerConfig={appState.routerConfig}
      bind:awgConfig={appState.awgConfig}
      isProcessing={appState.isProcessing}
      on:submit={createConnection}
      on:back={showWelcome}
    />
  
  {:else if appState.currentView === 'progress'}
    <Progress message={appState.progressMessage} />
  
  {:else if appState.currentView === 'success'}
    <Result 
      type="success"
      message={appState.successMessage}
      on:open-router={openRouterInterface}
      on:quit={quit}
    />
  
  {:else if appState.currentView === 'error'}
    <Result 
      type="error"
      message={appState.errorMessage}
      on:retry={showCreateAWG}
      on:back={showWelcome}
    />
  {/if}
</main>

<style>
  :global(html, body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    background: linear-gradient(135deg, #1e3c72 0%, #2a5298 50%, #0f2027 100%);
    height: 100vh;
    overflow: hidden;
    position: fixed;
    width: 100%;
  }

  :global(body::before) {
    content: '';
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-image: 
      radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
      radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.15) 0%, transparent 50%),
      radial-gradient(circle at 40% 40%, rgba(120, 219, 255, 0.1) 0%, transparent 50%);
    pointer-events: none;
    z-index: -1;
  }

  main {
    height: 100vh;
    display: flex;
    align-items: flex-start;
    justify-content: center;
    padding: 20px;
    overflow-y: auto;
    box-sizing: border-box;
  }

  main.center-content {
    align-items: center;
  }
</style>
