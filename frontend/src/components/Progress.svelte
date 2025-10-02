<script lang="ts">
  export let message: string = '';
</script>

<div class="progress-container">
  <h2>Обработка...</h2>
  <div class="progress-bar">
    <div class="progress-bar-fill"></div>
  </div>
  <p class="progress-message">{message}</p>
</div>

<style>
  .progress-container {
    background: white;
    border-radius: 12px;
    padding: 40px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    max-width: 600px;
    width: 100%;
    text-align: center;
  }

  h2 {
    text-align: center;
    color: #333;
    margin-bottom: 20px;
    font-size: 1.8em;
  }

  .progress-bar {
    width: 100%;
    height: 8px;
    background: #e1e5e9;
    border-radius: 4px;
    overflow: hidden;
    margin: 20px 0;
  }

  .progress-bar-fill {
    height: 100%;
    background: linear-gradient(90deg, #667eea, #764ba2);
    animation: progress 2s ease-in-out infinite;
  }

  @keyframes progress {
    0% { width: 0%; }
    50% { width: 70%; }
    100% { width: 100%; }
  }

  .progress-message {
    color: #666;
    font-style: italic;
  }
</style>
