<script lang="ts">
    import {createEventDispatcher} from 'svelte';
    import type {RouterConfig} from '../types';
    import FormSection from './FormSection.svelte';
    import RouterInfo from './RouterInfo.svelte';

    export let routerConfig: RouterConfig;
  export let isLoading: boolean = false;
  export let interfaces: any[] = [];
  export let error: string = '';

  const dispatch = createEventDispatcher();

  function handleBack() {
    dispatch('back');
  }

  function refreshInterfaces() {
    dispatch('refresh');
  }
</script>

<div class="wg-interfaces-container">
  
  <form on:submit|preventDefault>
    <div class="form-content">
      <h2>Список интерфейсов</h2>
      
      <div class="info-block">
        <div class="info-header">
          <span class="info-icon">📡</span>
          <span class="info-title">Что здесь видно</span>
        </div>
        <div class="info-content">
            На этой странице представлены все WireGuard интерфейсы, настроенные на роутере.<br>
            Так же в этом списке присутствует основной интерфейс, через который происходит соединение с провайдером.<br>
            Здесь отображаются ID интерфейсов и их описания.
        </div>
      </div>
    </div>
    
    <FormSection title={`Интерфейсы (${interfaces.length})`} icon="🔧">
      {#if isLoading}
        <div class="loading-state">
          <span class="loading-icon">⏳</span>
          Загрузка интерфейсов...
        </div>
      {:else if error}
        <div class="error-state">
          <span class="error-icon">❌</span>
          {error}
        </div>
      {:else if interfaces.length === 0}
        <div class="empty-state">
          <span class="empty-icon">📭</span>
          WireGuard интерфейсы не найдены
        </div>
      {:else}
        <div class="interfaces-list">
          {#each interfaces as iface}
            <div class="interface-item">
              <div class="interface-header">
                <div class="interface-icon">🔧</div>
                <div class="interface-details">
                  <div class="interface-field">
                    <span class="field-label">ID:</span>
                    <span class="interface-name">{iface.Id}</span>
                  </div>
                  <div class="interface-field">
                    <span class="field-label">Описание:</span>
                    <span class="interface-description">{iface.Description || 'Без описания'}</span>
                  </div>
                    <div class="interface-field">
                        <span class="field-label">Тип:</span>
                        <span class="interface-description">{iface.Type || 'Без типа'}</span>
                    </div>
                </div>
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </FormSection>

    <div class="form-actions">
      <button type="button" class="btn secondary" on:click={handleBack}>
        <span class="btn-icon">←</span>
        Назад
      </button>
      <button type="button" class="btn primary" on:click={refreshInterfaces} disabled={isLoading}>
        <span class="btn-icon">🔄</span>
        Обновить
      </button>
    </div>
      <RouterInfo routerUrl={routerConfig.url} />
  </form>
</div>

<style>
  .wg-interfaces-container {
    min-height: 100vh;
    padding: 20px;
    position: relative;
  }

  form {
    max-width: 600px;
    margin: 0 auto;
    background: rgba(255, 255, 255, 0.95);
    border-radius: 20px;
    padding: 30px;
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

  .loading-state, .error-state, .empty-state {
    text-align: center;
    padding: 40px 20px;
    color: #64748b;
  }

  .loading-icon, .error-icon, .empty-icon {
    font-size: 2em;
    display: block;
    margin-bottom: 10px;
  }

  .interfaces-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 12px;
    max-height: 400px;
    overflow-y: auto;
    padding: 8px;
  }

  .interfaces-list::-webkit-scrollbar {
    width: 8px;
  }

  .interfaces-list::-webkit-scrollbar-track {
    background: rgba(0, 0, 0, 0.1);
    border-radius: 4px;
  }

  .interfaces-list::-webkit-scrollbar-thumb {
    background: linear-gradient(135deg, #3b82f6, #10b981);
    border-radius: 4px;
  }

  .interfaces-list::-webkit-scrollbar-thumb:hover {
    background: linear-gradient(135deg, #2563eb, #059669);
  }

  .interface-item {
    background: linear-gradient(135deg, rgba(59, 130, 246, 0.08), rgba(16, 185, 129, 0.05));
    border: 1px solid rgba(59, 130, 246, 0.2);
    border-radius: 12px;
    padding: 12px;
    transition: all 0.3s ease;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }

  .interface-header {
    display: flex;
    align-items: flex-start;
    gap: 10px;
  }

  .interface-icon {
    font-size: 1.2em;
    background: linear-gradient(135deg, #3b82f6, #10b981);
    border-radius: 6px;
    padding: 6px;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 2px 6px rgba(59, 130, 246, 0.3);
    flex-shrink: 0;
  }

  .interface-details {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
  }

  .interface-field {
    display: flex;
    flex-direction: column;
    gap: 1px;
  }

  .field-label {
    font-weight: 600;
    color: #64748b;
    font-size: 0.7em;
    text-transform: uppercase;
    letter-spacing: 0.3px;
  }

  .interface-name {
    font-weight: 700;
    color: #2a5298;
    font-size: 0.9em;
    word-break: break-all;
  }

  .interface-description {
    font-size: 0.8em;
    color: #475569;
    font-style: italic;
    word-break: break-word;
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
  }

  .form-actions {
    display: flex;
    gap: 15px;
    justify-content: center;
    margin-top: 30px;
  }

  .btn {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px 24px;
    border: none;
    border-radius: 12px;
    font-weight: 600;
    font-size: 0.95em;
    cursor: pointer;
    transition: all 0.3s ease;
    text-decoration: none;
  }

  .btn.primary {
    background: linear-gradient(135deg, #2a5298, #3b82f6);
    color: white;
    box-shadow: 0 4px 15px rgba(42, 82, 152, 0.3);
  }

  .btn.primary:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(42, 82, 152, 0.4);
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
    opacity: 0.6;
    cursor: not-allowed;
    transform: none !important;
  }

  .btn-icon {
    font-size: 1.1em;
  }
</style>
