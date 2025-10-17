export class AppError extends Error {
  constructor(
    message: string,
    public code?: string,
    public details?: unknown
  ) {
    super(message);
    this.name = 'AppError';
  }
}

export function handleError(error: unknown): string {
  if (error instanceof AppError) {
    return error.message;
  }
  
  if (error instanceof Error) {
    // Переводим конкретные ошибки на русский
    if (error.message.includes('router is not in router mode')) {
      return 'Роутер не находится в режиме "Роутер". Поддерживается только режим "Роутер", а текущий режим: "Усилитель". Измените режим работы в настройках роутера или выберите другой роутер.';
    }
    
    return error.message;
  }
  
  if (typeof error === 'string') {
    // Переводим конкретные ошибки на русский
    if (error.includes('router is not in router mode')) {
      return 'Роутер не находится в режиме "Роутер". Поддерживается только режим "Роутер", а текущий режим: "Усилитель". Измените режим работы в настройках роутера или выберите другой роутер.';
    }
    
    return error;
  }
  
  return 'Произошла неизвестная ошибка';
}
