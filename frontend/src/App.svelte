<script lang="ts">
  import { 
    ValidateRouterConfig, 
    ValidateAWGConfig,
    CreateAWGInterface, 
    ConfigureAWGInterface, 
    ActivateAWGInterface, 
    GetRouterWebURL,
    AddRoutes,
    DeleteRoutes
  } from '../wailsjs/go/main/App.js';
  import { BrowserOpenURL, Quit } from '../wailsjs/runtime';
  
  import Welcome from './components/Welcome.svelte';
  import CreateAWGForm from './components/CreateAWGForm.svelte';
  import RoutesForm from './components/RoutesForm.svelte';
  import WgInterfacesList from './components/WgInterfacesList.svelte';
  import RouterAccessModal from './components/RouterAccessModal.svelte';
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
  let interfaces: any[] = [];
  let isLoadingInterfaces = false;
  let interfacesError = '';

  async function loadInterfaces() {
    if (!appState.routerConfig.url || !appState.routerConfig.login || !appState.routerConfig.password) return;
    if (isLoadingInterfaces) return; // Предотвращаем множественные вызовы
    
    isLoadingInterfaces = true;
    interfacesError = '';
    try {
      const { ShowWgInterfaces } = await import('../wailsjs/go/main/App.js');
      interfaces = await ShowWgInterfaces();
    } catch (err) {
      console.error('Ошибка загрузки интерфейсов:', err);
      interfacesError = err instanceof Error ? err.message : 'Ошибка загрузки интерфейсов';
      interfaces = [];
    } finally {
      isLoadingInterfaces = false;
    }
  }

  // Navigation handlers
  let pendingAction: 'create-awg' | 'manage-routes' | 'wg-interfaces' | 'reconnect' | null = null;

  function showCreateAWG() {
    if (appState.isRouterConnected) {
      appState = updateView(appState, 'create-awg');
    } else {
      pendingAction = 'create-awg';
    }
  }

  function retryCreateAWG() {
    appState = updateView(appState, 'create-awg');
  }

  function retryAddRoutes() {
    appState = updateView(appState, 'add-routes');
  }

  function handleRetry() {
    // Определяем, какое действие нужно повторить на основе последнего действия
    if (appState.currentView === 'error') {
      // Если мы пришли из формы маршрутов, возвращаемся туда
      if (appState.routeConfig && (appState.routeConfig.batFiles?.length > 0 || appState.routeConfig.batUrls?.length > 0)) {
        retryAddRoutes();
      } else {
        retryCreateAWG();
      }
    }
  }

  function showAddRoutes() {
    if (appState.isRouterConnected) {
      appState = updateView(appState, 'add-routes');
    } else {
      pendingAction = 'manage-routes';
    }
  }

  function showWgInterfaces() {
    if (appState.isRouterConnected) {
      appState = updateView(appState, 'wg-interfaces');
    } else {
      pendingAction = 'wg-interfaces';
    }
  }

  function handleRouterAccessProceed() {
    appState = { ...appState, isRouterConnected: true };
    
    // Загружаем интерфейсы после успешного подключения
    loadInterfaces();
    
    if (pendingAction === 'create-awg') {
      appState = updateView(appState, 'create-awg');
    } else if (pendingAction === 'manage-routes') {
      appState = updateView(appState, 'add-routes');
    } else if (pendingAction === 'wg-interfaces') {
      appState = updateView(appState, 'wg-interfaces');
    }
    pendingAction = null;
  }

  function handleRouterAccessCancel() {
    pendingAction = null;
  }

  function showRouterModal() {
    pendingAction = 'reconnect';
  }

  function showWelcome() {
    appState = { ...appState, currentView: 'welcome' };
  }

  // Main business logic
  async function createConnection() {
    if (appState.isProcessing) return;
    
    try {
      appState = { ...appState, lastAction: 'create-awg' };
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
      setTimeout(() => {
            // nothing
        }, 1000);
      appState = setProgress(appState, 'Настраиваем соединение...');
      
      // Configure AWG interface
      await ConfigureAWGInterface(appState.awgConfig, interfaceName);
      
      appState = setProgress(appState, 'Активируем соединение...');
      setTimeout(() => {
          // nothing
      }, 1000);

      // Activate AWG interface
      await ActivateAWGInterface(interfaceName);
      
      appState = setProgress(appState, 'Обновляем список интерфейсов...');
      
      // Update interfaces list
      await loadInterfaces();
        
      // Success
      appState = setSuccess(
        appState, 
        `Соединение успешно создано и включено!\nТеперь можно приступить к настройке политик подключения или маршрутизации и наслаждаться VPN!\nУдачи! 🌐`
      );
      
    } catch (error) {
      appState = setError(appState, handleError(error));
    }
  }

  // Function for adding routes
  async function addRoutes() {
    if (appState.isProcessing) return;
    
    try {
      appState = { ...appState, lastAction: 'add-routes' };
      appState = setProcessing(appState, true);
      
      // Validate form - проверяем что есть хотя бы один файл ИЛИ URL
      if (!appState.routeConfig.batFiles.length && !appState.routeConfig.batUrls.length) {
        throw new Error('Укажите хотя бы один BAT файл или URL ссылку');
      }
      
      if (!appState.routeConfig.interfaceId) {
        throw new Error('Укажите ID интерфейса');
      }
      
      appState = setProgress(appState, 'Проверяем подключение к роутеру...');
      
      appState = setProgress(appState, 'Добавляем маршруты...\nЭто может занять некоторое время если маршрутов очень много.');
      
      // Add routes using the new function
      await AddRoutes(
        appState.routeConfig.interfaceId,
        appState.routeConfig.batFiles,
        appState.routeConfig.batUrls
      );
      
      // Success
      appState = setSuccess(
        appState, 
        `Маршруты успешно добавлены для интерфейса "${appState.routeConfig.interfaceId}"!\nТеперь трафик для указанных адресов/подсетей будет направляться через этот интерфейс.`
      );
      
    } catch (error) {
      appState = setError(appState, handleError(error));
    }
  }

  // Function for deleting routes
  async function deleteRoutes() {
    if (appState.isProcessing) return;
    
    try {
      appState = { ...appState, lastAction: 'add-routes' };
      appState = setProcessing(appState, true);
      
      // Validate interface ID
      if (!appState.routeConfig.interfaceId) {
        throw new Error('Укажите ID интерфейса');
      }
      
      appState = setProgress(appState, 'Проверяем подключение к роутеру...');

      appState = setProgress(appState, 'Удаляем маршруты...');
      
      // Delete routes using the new function
      await DeleteRoutes(appState.routeConfig.interfaceId);
      
      // Success
      appState = setSuccess(
        appState, 
        `Маршруты успешно удалены для интерфейса "${appState.routeConfig.interfaceId}"!`
      );
      
    } catch (error) {
      appState = setError(appState, handleError(error));
    }
  }

  async function openRouterInterface(path: string = 'otherConnections') {
    try {
      const url = await GetRouterWebURL(appState.routerConfig.url, path);
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
      isRouterConnected={appState.isRouterConnected}
      on:create-awg={showCreateAWG}
      on:add-routes={showAddRoutes}
      on:wg-interfaces={showWgInterfaces}
      on:reconnect-router={showRouterModal}
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

  {:else if appState.currentView === 'add-routes'}
    <RoutesForm 
      bind:routerConfig={appState.routerConfig}
      bind:routeConfig={appState.routeConfig}
      isProcessing={appState.isProcessing}
      {interfaces}
      {isLoadingInterfaces}
      on:add-routes={addRoutes}
      on:delete-routes={deleteRoutes}
      on:back={showWelcome}
    />

  {:else if appState.currentView === 'wg-interfaces'}
    <WgInterfacesList 
      bind:routerConfig={appState.routerConfig}
      isLoading={isLoadingInterfaces}
      {interfaces}
      error={interfacesError}
      on:back={showWelcome}
      on:refresh={loadInterfaces}
    />
  
  {:else if appState.currentView === 'progress'}
    <Progress message={appState.progressMessage} />
  
  {:else if appState.currentView === 'success'}
    <Result 
      type="success"
      message={appState.successMessage}
      routerPath={appState.lastAction === 'add-routes' ? 'staticRoutes' : 'otherConnections'}
      on:open-router={(e) => openRouterInterface(e.detail.path)}
      on:home={showWelcome}
    />
  
  {:else if appState.currentView === 'error'}
    <Result 
      type="error"
      message={appState.errorMessage}
      on:retry={handleRetry}
      on:home={showWelcome}
    />
  {/if}

  {#if pendingAction}
    <RouterAccessModal
      bind:routerConfig={appState.routerConfig}
      actionTitle={pendingAction === 'create-awg' ? 'Создание AWG соединения' : pendingAction === 'manage-routes' ? 'Управление маршрутами' : 'Переподключение к роутеру'}
      actionIcon={pendingAction === 'create-awg' ? '🛜' : pendingAction === 'manage-routes' ? '🛣️' : '🔌'}
      on:proceed={handleRouterAccessProceed}
      on:cancel={handleRouterAccessCancel}
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
