export interface RouterConfig {
  url: string;
  login: string;
  password: string;
}

export interface AWGConfig {
  name: string;
  filePath: string;
}

export interface RouteConfig {
  interfaceId: string;
  batFiles: string[];
  batUrls: string[];
}

export type ViewType = 'welcome' | 'create-awg' | 'add-routes' | 'progress' | 'success' | 'error';

export interface AppState {
  currentView: ViewType;
  routerConfig: RouterConfig;
  awgConfig: AWGConfig;
  routeConfig: RouteConfig;
  progressMessage: string;
  errorMessage: string;
  successMessage: string;
  isProcessing: boolean;
  lastAction: 'create-awg' | 'add-routes' | null;
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

export const DEFAULT_ROUTE_CONFIG: RouteConfig = {
  interfaceId: '',
  batFiles: [],
  batUrls: []
};
