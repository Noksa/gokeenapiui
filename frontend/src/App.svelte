<script lang="ts">
  import { CreateAWGConnection, ValidateRouterConfig, GetRouterWebURL } from '../wailsjs/go/main/App.js';
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
      
      // Validate router config
      await ValidateRouterConfig(appState.routerConfig);
      
      appState = setProgress(appState, 'Создаём соединение...');
      
      // Create AWG connection
      await CreateAWGConnection(appState.routerConfig, appState.awgConfig);
      
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

<main>
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
      routerUrl={appState.routerConfig.url}
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
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
  }

  main {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
  }
</style>
