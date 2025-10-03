import type { AppState, ViewType } from '../types';
import { DEFAULT_ROUTER_CONFIG, DEFAULT_AWG_CONFIG, DEFAULT_ROUTE_CONFIG } from '../types';

export function createInitialState(): AppState {
  return {
    currentView: 'welcome',
    routerConfig: { ...DEFAULT_ROUTER_CONFIG },
    awgConfig: { ...DEFAULT_AWG_CONFIG },
    routeConfig: { ...DEFAULT_ROUTE_CONFIG },
    progressMessage: '',
    errorMessage: '',
    successMessage: '',
    isProcessing: false
  };
}

export function updateView(state: AppState, view: ViewType): AppState {
  return {
    ...state,
    currentView: view
  };
}

export function setProcessing(state: AppState, isProcessing: boolean): AppState {
  return {
    ...state,
    isProcessing
  };
}

export function setProgress(state: AppState, message: string): AppState {
  return {
    ...state,
    progressMessage: message,
    currentView: 'progress'
  };
}

export function setSuccess(state: AppState, message: string): AppState {
  return {
    ...state,
    successMessage: message,
    currentView: 'success',
    isProcessing: false
  };
}

export function setError(state: AppState, message: string): AppState {
  return {
    ...state,
    errorMessage: message,
    currentView: 'error',
    isProcessing: false
  };
}

export function resetState(): AppState {
  return createInitialState();
}
