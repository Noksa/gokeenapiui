export interface RouterConfig {
  url: string;
  login: string;
  password: string;
}

export interface AWGConfig {
  name: string;
  filePath: string;
}

export type ViewType = 'welcome' | 'create-awg' | 'progress' | 'success' | 'error';

export interface AppState {
  currentView: ViewType;
  routerConfig: RouterConfig;
  awgConfig: AWGConfig;
  progressMessage: string;
  errorMessage: string;
  successMessage: string;
  isProcessing: boolean;
}

export const DEFAULT_ROUTER_CONFIG: RouterConfig = {
  url: 'http://192.168.1.1',
  login: 'admin',
  password: ''
};

export const DEFAULT_AWG_CONFIG: AWGConfig = {
  name: '',
  filePath: ''
};
