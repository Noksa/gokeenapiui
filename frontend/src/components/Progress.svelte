<script lang="ts">
  export let message: string = '';
</script>

<div class="progress-container">
  <div class="spinner-icon">
    <div class="spinner"></div>
  </div>
  <h2>Обработка...</h2>
  <div class="progress-bar">
    <div class="progress-bar-fill"></div>
    <div class="progress-glow"></div>
  </div>
  <p class="progress-message">{message}</p>
</div>

<style>
  .progress-container {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 50px;
    box-shadow: 
      0 20px 40px rgba(0, 0, 0, 0.3),
      0 0 0 1px rgba(255, 255, 255, 0.2);
    max-width: 600px;
    width: 100%;
    text-align: center;
    position: relative;
    overflow: hidden;
  }

  .progress-container::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: linear-gradient(90deg, #00d4ff, #5b86e5, #36d1dc);
    animation: pulse 2s ease-in-out infinite alternate;
  }

  .spinner-icon {
    margin-bottom: 20px;
  }

  .spinner {
    width: 50px;
    height: 50px;
    border: 4px solid rgba(42, 82, 152, 0.2);
    border-top: 4px solid #2a5298;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto;
  }

  h2 {
    color: #1e3c72;
    margin-bottom: 30px;
    font-size: 1.8em;
    font-weight: 600;
  }

  .progress-bar {
    width: 100%;
    height: 12px;
    background: rgba(42, 82, 152, 0.1);
    border-radius: 6px;
    overflow: hidden;
    margin: 30px 0;
    position: relative;
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .progress-bar-fill {
    height: 100%;
    background: linear-gradient(90deg, #2a5298, #5b86e5, #36d1dc);
    border-radius: 6px;
    animation: progress 3s ease-in-out infinite;
    position: relative;
  }

  .progress-glow {
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
    animation: glow 2s ease-in-out infinite;
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }

  @keyframes progress {
    0% { width: 0%; }
    50% { width: 70%; }
    100% { width: 100%; }
  }

  @keyframes glow {
    0% { left: -100%; }
    100% { left: 100%; }
  }

  @keyframes pulse {
    0% { opacity: 0.6; }
    100% { opacity: 1; }
  }

  .progress-message {
    color: #666;
    font-style: italic;
    font-size: 1.1em;
    margin-top: 20px;
  }
</style>
