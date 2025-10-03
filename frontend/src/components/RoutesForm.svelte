<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import AddRoutesForm from './AddRoutesForm.svelte';
  import FormSection from './FormSection.svelte';
  import RouterInfo from './RouterInfo.svelte';

  export let routerConfig: RouterConfig;
  export let routeConfig: RouteConfig;
  export let isProcessing: boolean = false;

  const dispatch = createEventDispatcher<{
    'add-routes': void;
    'delete-routes': void;
    'back': void;
  }>();

  let currentView: 'menu' | 'add-routes' | 'delete-routes' = 'menu';
  let showConfirmDialog = false;

  function showAddRoutes() {
    currentView = 'add-routes';
  }

  function showDeleteRoutes() {
    currentView = 'delete-routes';
  }

  function showMenu() {
    currentView = 'menu';
  }

  function handleAddRoutes() {
    dispatch('add-routes');
  }

  function handleDeleteRoutes() {
    if (!routeConfig.interfaceId.trim()) {
      return; // Просто не делаем ничего если поле пустое
    }
    
    showConfirmDialog = true;
  }

  function confirmDelete() {
    showConfirmDialog = false;
    dispatch('delete-routes');
  }

  function cancelDelete() {
    showConfirmDialog = false;
  }

  function handleBack() {
    if (currentView === 'menu') {
      dispatch('back');
    } else {
      showMenu();
    }
  }
</script>

<div class="form-container">
  {#if currentView === 'menu'}
    <div class="routes-menu">
      <div class="header-section">
        <div class="animated-icon">🛣️</div>
        <h2>Управление маршрутами</h2>
      </div>
      <p class="description">
        Управление статическими маршрутами позволяет настроить направление трафика через определенные интерфейсы роутера. 
        Это полезно для VPN соединений, когда нужно направить трафик к определенным IP адресам через AWG интерфейс.
      </p>
      
      <div class="button-group">
        <button class="btn primary route-btn" on:click={showAddRoutes}>
          <span class="btn-icon">➕</span>
          Добавить маршруты
          <div class="btn-glow"></div>
        </button>
        <button class="btn secondary delete-btn" on:click={showDeleteRoutes}>
          <span class="btn-icon">🗑️</span>
          Удалить маршруты
          <div class="btn-glow"></div>
        </button>
        <button class="btn secondary" on:click={handleBack}>
          <span class="btn-icon">↩️</span>
          Назад
        </button>
      </div>
    </div>

  {:else if currentView === 'add-routes'}
    <AddRoutesForm 
      bind:routerConfig={routerConfig}
      bind:routeConfig={routeConfig}
      {isProcessing}
      on:submit={handleAddRoutes}
      on:back={handleBack}
    />

  {:else if currentView === 'delete-routes'}
    <div class="form-container">
      <form on:submit|preventDefault={handleDeleteRoutes}>
        <div class="form-content">
          <h2>Удаление маршрутов</h2>
        </div>
        
        <FormSection title="Удаление маршрутов" icon="🗑️">
          <div class="interface-section">
            <div class="section-header">
              <h4>🔌 ID интерфейса</h4>
              <div class="interface-input-container">
                <input 
                  id="interface-id"
                  type="text" 
                  bind:value={routeConfig.interfaceId}
                  placeholder="Например: Wireguard0"
                  class="interface-input"
                  required
                />
              </div>
            </div>
          </div>
          
          <div class="warning-section">
            <div class="warning-box">
              <span class="warning-icon">⚠️</span>
              <div class="warning-text">
                <strong>Внимание!</strong><br>
                Будут удалены все пользовательские маршруты для указанного интерфейса.<br>
                Это действие нельзя отменить.
              </div>
            </div>
          </div>
        </FormSection>

        <div class="button-group">
          <button type="submit" class="btn danger" disabled={isProcessing}>
            <span class="btn-icon">🗑️</span>
            Удалить маршруты
          </button>
          <button type="button" class="btn secondary" on:click={handleBack}>
            <span class="btn-icon">↩️</span>
            Назад
          </button>
        </div>
      </form>
      
      <RouterInfo routerUrl={routerConfig.url} />
    </div>
  {/if}
</div>

{#if showConfirmDialog}
  <div class="modal-overlay" on:click={cancelDelete} on:keydown={(e) => e.key === 'Escape' && cancelDelete()}>
    <div class="confirm-dialog" on:click|stopPropagation on:keydown={() => {}}>
      <div class="dialog-header">
        <span class="dialog-icon">⚠️</span>
        <h3>Подтверждение удаления</h3>
      </div>
      
      <div class="dialog-content">
        <p>Вы уверены, что хотите удалить все пользовательские маршруты для интерфейса <strong>"{routeConfig.interfaceId}"</strong>?</p>
        <p class="warning-text">Это действие нельзя отменить.</p>
      </div>
      
      <div class="dialog-buttons">
        <button class="btn danger" on:click={confirmDelete}>
          <span class="btn-icon">🗑️</span>
          Удалить
        </button>
        <button class="btn secondary" on:click={cancelDelete}>
          <span class="btn-icon">❌</span>
          Отмена
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  .form-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    padding: 0;
  }

  .routes-menu {
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.9));
    border-radius: 20px;
    padding: 40px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.1),
      0 0 0 1px rgba(255, 255, 255, 0.5);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
    text-align: center;
    animation: slideIn 0.6s ease-out;
  }

  .header-section {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 15px;
    margin-bottom: 20px;
  }

  .animated-icon {
    font-size: 3em;
    animation: float 3s ease-in-out infinite;
    filter: drop-shadow(0 4px 8px rgba(42, 82, 152, 0.3));
  }

  @keyframes float {
    0%, 100% { 
      transform: translateY(0px) rotate(0deg);
    }
    50% { 
      transform: translateY(-10px) rotate(2deg);
    }
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

  h2 {
    color: #2a5298;
    margin: 0;
    font-size: 1.4em;
    font-weight: 600;
    text-shadow: 0 2px 4px rgba(42, 82, 152, 0.1);
  }

  .description {
    color: #6b7280;
    margin-bottom: 35px;
    line-height: 1.6;
    font-size: 1.05em;
    max-width: 500px;
    margin-left: auto;
    margin-right: auto;
  }

  .button-group {
    display: flex;
    flex-direction: column;
    gap: 15px;
    align-items: center;
  }

  .btn {
    position: relative;
    overflow: hidden;
    padding: 15px 30px;
    border: none;
    border-radius: 12px;
    font-size: 1em;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 10px;
    text-decoration: none;
    min-width: 220px;
    justify-content: center;
    animation: fadeInUp 0.6s ease-out;
  }

  .btn:nth-child(1) { animation-delay: 0.1s; }
  .btn:nth-child(2) { animation-delay: 0.2s; }
  .btn:nth-child(3) { animation-delay: 0.3s; }

  @keyframes fadeInUp {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .btn-icon {
    font-size: 1.2em;
    filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.2));
  }

  .btn.primary {
    background: linear-gradient(135deg, #2a5298, #4a90e2);
    color: white;
    box-shadow: 0 6px 20px rgba(42, 82, 152, 0.4);
  }

  .btn.primary:hover {
    transform: translateY(-3px);
    box-shadow: 0 10px 30px rgba(42, 82, 152, 0.6);
  }

  .btn.secondary {
    background: linear-gradient(135deg, #6b7280, #9ca3af);
    color: white;
    box-shadow: 0 6px 20px rgba(107, 114, 128, 0.3);
  }

  .btn.secondary:hover {
    transform: translateY(-3px);
    box-shadow: 0 10px 30px rgba(107, 114, 128, 0.5);
  }

  .route-btn {
    background: linear-gradient(135deg, #059669, #10b981);
    box-shadow: 0 6px 20px rgba(5, 150, 105, 0.4);
    animation: pulse 3s ease-in-out infinite;
  }

  .route-btn:hover {
    transform: translateY(-3px);
    box-shadow: 0 10px 30px rgba(5, 150, 105, 0.6);
  }

  .delete-btn {
    background: linear-gradient(135deg, #dc2626, #ef4444);
    box-shadow: 0 6px 20px rgba(220, 38, 38, 0.4);
  }

  .delete-btn:hover {
    transform: translateY(-3px);
    box-shadow: 0 10px 30px rgba(220, 38, 38, 0.6);
  }

  @keyframes pulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.02); }
  }

  .btn-glow {
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
    transition: left 0.6s;
  }

  .btn:hover .btn-glow {
    left: 100%;
  }

  @media (max-width: 600px) {
    .button-group {
      gap: 12px;
    }
    
    .btn {
      min-width: 180px;
    }
  }
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
    margin-bottom: 30px;
    font-size: 2em;
    font-weight: 700;
    text-shadow: 0 2px 4px rgba(42, 82, 152, 0.1);
  }

  .interface-section {
    margin-bottom: 25px;
  }

  .section-header h4 {
    margin-bottom: 12px;
    color: #2a5298;
    font-weight: 600;
    font-size: 1.1em;
  }

  .interface-input-container {
    position: relative;
  }

  .interface-input {
    width: 100%;
    padding: 16px 20px;
    border: 2px solid #e5e7eb;
    border-radius: 12px;
    font-size: 1em;
    transition: all 0.3s ease;
    background: rgba(255, 255, 255, 0.8);
    backdrop-filter: blur(5px);
  }

  .interface-input:focus {
    outline: none;
    border-color: #2a5298;
    box-shadow: 0 0 0 4px rgba(42, 82, 152, 0.1);
    background: rgba(255, 255, 255, 0.95);
  }

  .button-group {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 30px;
  }

  .btn {
    position: relative;
    padding: 16px 32px;
    border: none;
    border-radius: 12px;
    font-size: 1em;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    gap: 8px;
    text-decoration: none;
    overflow: hidden;
  }

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .btn-icon {
    font-size: 1.2em;
  }

  .btn.secondary {
    background: linear-gradient(135deg, #6b7280, #9ca3af);
    color: white;
    box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3);
  }

  .btn.secondary:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(107, 114, 128, 0.4);
  }

  .btn.danger {
    background: linear-gradient(135deg, #dc2626, #ef4444);
    color: white;
    box-shadow: 0 4px 15px rgba(220, 38, 38, 0.4);
  }

  .btn.danger:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(220, 38, 38, 0.5);
  }

  .warning-section {
    margin-top: 25px;
  }

  .warning-box {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    padding: 20px;
    background: linear-gradient(135deg, #fef3c7, #fde68a);
    border: 2px solid #f59e0b;
    border-radius: 12px;
    margin-bottom: 25px;
  }

  .warning-icon {
    font-size: 1.5em;
    flex-shrink: 0;
  }

  .warning-text {
    color: #92400e;
    line-height: 1.6;
    font-weight: 500;
  }

  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }

  .confirm-dialog {
    background: white;
    border-radius: 16px;
    padding: 30px;
    max-width: 500px;
    width: 90%;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  }

  .dialog-header {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 20px;
  }

  .dialog-icon {
    font-size: 2em;
  }

  .dialog-header h3 {
    margin: 0;
    color: #2a5298;
    font-size: 1.3em;
  }

  .dialog-content {
    margin-bottom: 25px;
    line-height: 1.6;
  }

  .dialog-content .warning-text {
    color: #dc2626;
    font-weight: 600;
    margin-top: 10px;
  }

  .dialog-buttons {
    display: flex;
    gap: 12px;
    justify-content: center;
    flex-wrap: wrap;
  }

  .dialog-buttons .btn {
    min-width: 120px;
    padding: 12px 20px;
  }
</style>
