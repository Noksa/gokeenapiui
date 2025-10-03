<script lang="ts">
  import type { RouterConfig } from '../types';

  export let routerConfig: RouterConfig;
  export let fieldErrors: { url: boolean; login: boolean; password: boolean };
  export let clearFieldError: ((field: 'url' | 'login' | 'password') => void) | undefined = undefined;

  function handleFieldError(field: 'url' | 'login' | 'password') {
    if (clearFieldError) {
      clearFieldError(field);
    } else {
      fieldErrors[field] = false;
    }
  }
</script>

<h3>
  <div class="section-header">
    <div class="section-icon">
      <span class="icon">🔗</span>
    </div>
    <div class="section-title">
      <span>Доступ к роутеру</span>
      <div class="info-tooltip">
        <span class="info-icon">💡</span>
        <div class="tooltip-content">
          <strong>Примеры подключения:</strong><br>
          Внутренний адрес: <code>http://192.168.1.1</code><br>
          Внешний адрес: <code>https://super-keenetic.keenetic.pro</code><br><br>
          Если роутер в локальной сети - используйте HTTP и внутренний IP.<br>
          Если через интернет - используйте HTTPS и KeenDNS имя.
        </div>
      </div>
    </div>
  </div>
</h3>

<div class="form-group">
  <label for="router-url">
    <span class="label-icon">🔗</span>
    URL роутера
  </label>
  <input 
    id="router-url"
    type="text" 
    bind:value={routerConfig.url}
    on:input={() => handleFieldError('url')}
    placeholder="IP или DNS имя (протокол http/https) роутера"
    class:error={fieldErrors.url}
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
      on:input={() => handleFieldError('login')}
      placeholder="Логин администратора"
      class:error={fieldErrors.login}
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
      on:input={() => handleFieldError('password')}
      placeholder="Пароль администратора"
      class:error={fieldErrors.password}
    />
  </div>
</div>

<style>
  h3 {
    color: #2a5298;
    font-size: 1.1em;
    margin: 0 0 15px 0;
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .section-header {
    display: flex;
    align-items: center;
    gap: 12px;
    width: 100%;
  }

  .section-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 40px;
    height: 40px;
    background: linear-gradient(135deg, #2a5298, #4a90e2);
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(42, 82, 152, 0.3);
  }

  .section-icon .icon {
    font-size: 1.2em;
    filter: drop-shadow(0 1px 2px rgba(0, 0, 0, 0.3));
  }

  .section-title {
    display: flex;
    align-items: center;
    gap: 10px;
    flex: 1;
  }

  .info-tooltip {
    position: relative;
    display: inline-block;
    cursor: help;
    margin-left: 8px;
    background: rgba(255, 193, 7, 0.2);
    border-radius: 50%;
    padding: 4px;
    box-shadow: 0 0 10px rgba(255, 193, 7, 0.3);
    animation: glow 2s ease-in-out infinite alternate;
  }

  .info-icon {
    font-size: 1.1em;
    display: inline-block;
    transition: all 0.3s ease;
    filter: drop-shadow(0 0 3px rgba(255, 193, 7, 0.8));
  }

  .info-icon:hover {
    transform: scale(1.3);
    filter: drop-shadow(0 0 8px rgba(255, 193, 7, 1));
  }

  .info-tooltip:hover {
    animation: none;
    box-shadow: 0 0 15px rgba(255, 193, 7, 0.6);
    background: rgba(255, 193, 7, 0.3);
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

  .tooltip-content {
    position: absolute;
    top: 100%;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(42, 82, 152, 0.95);
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
    max-width: 280px;
    width: max-content;
    line-height: 1.4;
  }

  @media (max-width: 500px) {
    .tooltip-content {
      left: 0;
      right: 0;
      transform: none;
      max-width: none;
      width: auto;
      margin: 8px 10px 0 10px;
    }
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

  .info-tooltip:hover .tooltip-content {
    visibility: visible;
    opacity: 1;
  }

  .tooltip-content code {
    background: rgba(255, 255, 255, 0.2);
    padding: 2px 6px;
    border-radius: 4px;
    font-family: 'Monaco', 'Menlo', monospace;
    font-size: 0.85em;
  }

  .form-group {
    margin-bottom: 15px;
  }

  .form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 15px;
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

  input.error {
    border-color: #dc3545;
    box-shadow: 0 0 0 3px rgba(220, 53, 69, 0.2);
  }

  @media (max-width: 600px) {
    .form-row {
      grid-template-columns: 1fr;
    }
  }
</style>
